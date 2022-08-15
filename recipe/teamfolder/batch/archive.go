package batch

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"strings"
)

type Archive struct {
	rc_recipe.RemarkIrreversible
	ErrTeamFolderNotFound                 app_msg.Message
	ErrUnableToArchive                    app_msg.Message
	ErrUnableToRetrieveCurrentTeamFolders app_msg.Message
	File                                  fd_file.RowFeed
	OperationLog                          rp_model.TransactionReport
	Peer                                  dbx_conn.ConnScopedTeam
	ProgressArchiveFolder                 app_msg.Message
	ErrorTeamSpaceNotSupported            app_msg.Message
}

func (z *Archive) Exec(c app_control.Control) error {
	if ok, _ := teamfolder.IsTeamSpaceSupported(z.Peer.Context()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupported)
		return errors.New("team space is not supported by this command")
	}

	ui := c.UI()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	folders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		ui.Error(z.ErrUnableToRetrieveCurrentTeamFolders.With("Error", err.Error()))
		return err
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*TeamFolderName)
		ui.Info(z.ProgressArchiveFolder.With("Name", r.Name))

		var folder *mo_teamfolder.TeamFolder
		for _, tf := range folders {
			if strings.ToLower(r.Name) == strings.ToLower(tf.Name) {
				folder = tf
				break
			}
		}
		if folder == nil {
			ui.Error(z.ErrTeamFolderNotFound.With("Name", r.Name))
			z.OperationLog.Failure(errors.New("team folder not found"), r)
			return nil
		}

		archived, err := sv_teamfolder.New(z.Peer.Context()).Archive(folder)
		if err != nil {
			ui.Error(z.ErrUnableToArchive.With("Name", r.Name).With("Error", err.Error()))
			z.OperationLog.Failure(err, r)
			return nil
		}
		z.OperationLog.Success(r, archived)
		return nil
	})
}

func (z *Archive) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Archive{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("test-batch-archive", "Marketing\nSales\n")
		if err != nil {
			return
		}
		m := r.(*Archive)
		m.File.SetFilePath(f)
	})
}

func (z *Archive) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.File.SetModel(&TeamFolderName{})
	z.OperationLog.SetModel(&TeamFolderName{}, &mo_teamfolder.TeamFolder{},
		rp_model.HiddenColumns(
			"result.team_folder_id",
		),
	)
}
