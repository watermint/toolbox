package batch

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Unshare struct {
	Peer                dbx_conn.ConnScopedTeam
	File                fd_file.RowFeed
	OperationLog        rp_model.TransactionReport
	LeaveCopy           bool
	SkipNotSharedFolder app_msg.Message
}

func (z *Unshare) Preset() {
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

func (z *Unshare) unshare(mf *MemberFolder, svm sv_member.Member, c app_control.Control) error {
	member, err := svm.ResolveByEmail(mf.MemberEmail)
	if err != nil {
		z.OperationLog.Failure(err, mf)
		return err
	}

	cm := z.Peer.Context().AsMemberId(member.TeamMemberId)

	f1, err := sv_file.NewFiles(cm).Resolve(mo_path.NewDropboxPath(mf.Path))
	if err != nil {
		return err
	}
	f2, ok := f1.Folder()
	if !ok {
		err = errors.New("shared folder not found")
		z.OperationLog.Failure(err, mf)
		return err
	}
	if f2.EntrySharedFolderId == "" {
		z.OperationLog.Skip(z.SkipNotSharedFolder, mf)
		return nil
	}
	sf, err := sv_sharedfolder.New(cm).Resolve(f2.EntrySharedFolderId)
	if err != nil {
		z.OperationLog.Failure(err, mf)
		return err
	}

	err = sv_sharedfolder.New(cm).Remove(sf, sv_sharedfolder.LeaveACopy(z.LeaveCopy))
	z.OperationLog.Success(mf, sf)
	return nil
}

func (z *Unshare) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Context())

	var lastErr, listErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("unshare", z.unshare, svm, c)
		q := s.Get("unshare")

		listErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	})

	return lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Unshare) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("share", "john@example.com,/shared\nemma@example.com,/project")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Unshare{}, func(r rc_recipe.Recipe) {
		m := r.(*Unshare)
		m.File.SetFilePath(f)
	})
}
