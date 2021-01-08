package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/strings/es_mailaddr"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type DeleteRecord struct {
	TeamFolderName         string `json:"team_folder_name"`
	Path                   string `json:"path"`
	GroupNameOrMemberEmail string `json:"group_name_or_member_email"`
}

func (z DeleteRecord) DropboxPath() mo_path.DropboxPath {
	if z.Path == "" {
		return mo_path.NewDropboxPath("")
	} else {
		return mo_path.NewDropboxPath("/" + z.Path)
	}
}

type Delete struct {
	rc_recipe.RemarkIrreversible
	Peer               dbx_conn.ConnScopedTeam
	File               fd_file.RowFeed
	OperationLog       rp_model.TransactionReport
	AdminGroupName     string
	SkipFolderNotFound app_msg.Message
	SkipNotAMember     app_msg.Message
}

func (z *Delete) Preset() {
	z.AdminGroupName = uc_teamfolder.DefaultAdminWorkGroupName
	z.OperationLog.SetModel(&DeleteRecord{}, nil)
	z.File.SetModel(&DeleteRecord{})
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Delete) delete(r *DeleteRecord, c app_control.Control, tc uc_teamfolder.TeamContent, sg sv_group.Group) error {
	l := c.Log()
	l.Debug("Remove", esl.Any("record", r))

	tf, err := tc.GetTeamFolder(r.TeamFolderName)
	switch err {
	case nil:
		l.Debug("The folder found")
		// continue

	case uc_teamfolder.ErrorUnableToIdentifyFolder:
		l.Debug("The folder not found")
		z.OperationLog.Skip(z.SkipFolderNotFound, r)
		return nil

	default:
		l.Debug("Unable to resolve the team folder", esl.Error(err))
		z.OperationLog.Failure(err, r)
		return err
	}

	handleResult := func(err error) error {
		switch err {
		case uc_teamfolder.ErrorUnableToIdentifyFolder:
			l.Debug("The folder not found", esl.Error(err))
			z.OperationLog.Skip(z.SkipFolderNotFound, r)
			return nil

		case uc_teamfolder.ErrorNotAMember:
			l.Debug("The member does not have an access", esl.Error(err))
			z.OperationLog.Skip(z.SkipNotAMember, r)
			return nil

		default:
			l.Debug("Successfully removed")
			z.OperationLog.Success(r, nil)
			return nil
		}
	}

	if _, err := sg.ResolveByName(r.GroupNameOrMemberEmail); err != nil {
		// assume the field is email
		if !es_mailaddr.IsEmailAddr(r.GroupNameOrMemberEmail) {
			l.Debug("The field look like not an email address")
			z.OperationLog.Failure(errors.New("group not found"), r)
			return errors.New("group not found")
		}

		return handleResult(tf.MemberRemoveUser(r.DropboxPath(), r.GroupNameOrMemberEmail))
	} else {
		// adding the group
		return handleResult(tf.MemberRemoveGroup(r.DropboxPath(), r.GroupNameOrMemberEmail))
	}
}

func (z *Delete) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	tc, err := uc_teamfolder.New(z.Peer.Context(), z.AdminGroupName)
	if err != nil {
		return err
	}
	sg := sv_group.NewCached(z.Peer.Context())

	var lastErr, fileErr error
	queueId := "delete"
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueId, z.delete, c, tc, sg)
		q := s.Get(queueId)

		fileErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	if fileErr != nil {
		l.Debug("Error on read file", esl.Error(fileErr))
		return fileErr
	}
	if lastErr != nil {
		l.Debug("Error on the process", esl.Error(lastErr))
		return lastErr

	}
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("member", "Tokyo Sales,,Sales\nTokyo Sales,Sales Report,Audit\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.File.SetFilePath(f)
	})
}
