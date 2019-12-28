package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Delete struct {
	Peer rc_conn.ConnUserFile
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
		switch api_util.ErrorSummary(err) {
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
	return qt_endtoend.ScenarioTest()
}
