package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Delete struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnScopedIndividual
	Path           mo_path.DropboxPath
	ProgressDelete app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()

	return ig_file.DeleteRecursively(z.Peer.Client(), z.Path, func(path mo_path.DropboxPath) {
		ui.Progress(z.ProgressDelete.With("Path", path.Path()))
	})
}

func (z *Delete) Test(c app_control.Control) error {
	// replay test
	{
		err := rc_exec.ExecReplay(c, &Delete{}, "recipe-file-delete.json.gz", func(r rc_recipe.Recipe) {
			m := r.(*Delete)
			m.Path = mo_path.NewDropboxPath("target")
		})
		if err != nil {
			return err
		}
	}

	err := rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("delete")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}
	return qt_errors.ErrorScenarioTest
}
