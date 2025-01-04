package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Leave struct {
	Peer                dbx_conn.ConnScopedTeam
	File                fd_file.RowFeed
	OperationLog        rp_model.TransactionReport
	KeepCopy            bool
	SkipNotSharedFolder app_msg.Message
}

func (z *Leave) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.File.SetModel(&MemberFolder{})
	z.OperationLog.SetModel(&MemberFolder{}, &mo_sharedfolder.SharedFolder{})
}

func (z *Leave) leave(mf *MemberFolder, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(mf.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, mf)
		return err
	}

	cm := z.Peer.Client().AsMemberId(member.TeamMemberId).WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	sf, err := uc_sharedfolder.NewResolver(cm).Resolve(mo_path.NewDropboxPath(mf.Path))
	switch err {
	case nil:
		// fall through
	case uc_sharedfolder.ErrorNotSharedFolder:
		z.OperationLog.Skip(z.SkipNotSharedFolder, mf)
		return nil

	default:
		z.OperationLog.Failure(err, mf)
		return err
	}

	err = sv_sharedfolder.New(cm).Leave(sf, sv_sharedfolder.LeaveACopy(z.KeepCopy))
	if err != nil {
		z.OperationLog.Failure(err, sf)
		return err
	}
	z.OperationLog.Success(mf, sf)
	return nil
}

func (z *Leave) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Client())

	var lastErr, listErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("leave", z.leave, svm, c)
		q := s.Get("leave")

		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return es_lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Leave) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
