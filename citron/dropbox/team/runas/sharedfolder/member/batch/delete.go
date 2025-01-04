package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Delete struct {
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	LeaveCopy    bool
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&DeleteMember{})
	z.OperationLog.SetModel(&DeleteMember{}, nil)
}

func (z *Delete) deleteMember(m *DeleteMember, resolver uc_sharedfolder.Resolver, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(m.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	rm, err := resolver.GroupOrEmailForRemoveMember(m.GroupOrEmail)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	cm := z.Peer.Client().AsMemberId(member.TeamMemberId).WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	sf, err := uc_sharedfolder.NewResolver(cm).Resolve(mo_path.NewDropboxPath(m.Path))
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	opts := make([]sv_sharedfolder_member.RemoveOption, 0)
	if z.LeaveCopy {
		opts = append(opts, sv_sharedfolder_member.LeaveACopy())
	}
	err = sv_sharedfolder_member.New(cm, sf).Remove(rm, opts...)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}
	z.OperationLog.Success(m, nil)
	return nil
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr, listErr error
	sfr := uc_sharedfolder.NewResolver(z.Peer.Client())
	svm := sv_member.NewCached(z.Peer.Client())

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_member", z.deleteMember, sfr, svm, c)
		q := s.Get("delete_member")
		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return es_lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
	//f, err := qt_file.MakeTestFile("share", "john@example.com,/shared,emma@examplel.com\nemma@example.com,/project,john@example.com")
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	_ = os.Remove(f)
	//}()
	//
	//return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
	//	m := r.(*Delete)
	//	m.File.SetFilePath(f)
	//})
}
