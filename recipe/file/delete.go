package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Delete struct {
	Peer dbx_conn.ConnUserFile
	Path mo_path.DropboxPath
}

func (z *Delete) Preset() {
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Context()

	var delete func(path mo_path.DropboxPath) error
	delete = func(path mo_path.DropboxPath) error {
		ui.InfoK("recipe.file.delete.progress.deleting", app_msg.P{
			"Path": path.Path(),
		})
		_, err := sv_file.NewFiles(ctx).Remove(path)
		if err == nil {
			return nil
		}
		switch dbx_util.ErrorSummary(err) {
		case "too_many_files":
			entries, err := sv_file.NewFiles(ctx).List(path)
			if err != nil {
				return err
			}
			for _, entry := range entries {
				if f, ok := entry.File(); ok {
					delete(f.Path())
				}
				if f, ok := entry.Folder(); ok {
					delete(f.Path())
				}
			}
			return delete(path)

		default:
			return err
		}
	}

	return delete(z.Path)
}

func (z *Delete) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Path = qt_recipe.NewTestDropboxFolderPath("delete")
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}
	return qt_errors.ErrorScenarioTest
}
