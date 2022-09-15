package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Add struct {
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	Silent       bool
	Message      mo_string.OptionalString
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&AddMember{})
	z.OperationLog.SetModel(&AddMember{}, nil)
}

func (z *Add) addMember(m *AddMember, resolver uc_sharedfolder.Resolver, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(m.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	am, err := resolver.GroupOrEmailForAddMember(m.GroupOrEmail, m.AccessLevel)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	cm := z.Peer.Client().AsMemberId(member.TeamMemberId)
	sf, err := uc_sharedfolder.NewResolver(cm).Resolve(mo_path.NewDropboxPath(m.Path))
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}

	opts := make([]sv_sharedfolder_member.AddOption, 0)
	if z.Silent {
		opts = append(opts, sv_sharedfolder_member.AddQuiet())
	}
	if z.Message.IsExists() {
		opts = append(opts, sv_sharedfolder_member.AddCustomMessage(z.Message.Value()))
	}
	err = sv_sharedfolder_member.New(cm, sf).Add(am, opts...)
	if err != nil {
		z.OperationLog.Failure(err, m)
		return err
	}
	z.OperationLog.Success(m, nil)
	return nil
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr, listErr error
	sfr := uc_sharedfolder.NewResolver(z.Peer.Client())
	svm := sv_member.NewCached(z.Peer.Client())

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("add_member", z.addMember, sfr, svm, c)
		q := s.Get("add_member")
		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Add) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
	//f, err := qt_file.MakeTestFile("share", "john@example.com,/shared,editor,emma@examplel.com\nemma@example.com,/project,viewer,john@example.com")
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	_ = os.Remove(f)
	//}()
	//
	//return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
	//	m := r.(*Add)
	//	m.File.SetFilePath(f)
	//})
}
