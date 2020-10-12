package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Permdelete struct {
	rc_recipe.RemarkIrreversible
	rc_recipe.RemarkExperimental
	Peer           dbx_conn.ConnUserFile
	Path           mo_path.DropboxPath
	ProgressDelete app_msg.Message
}

func (z *Permdelete) Preset() {
}

func (z *Permdelete) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Context()

	var del func(path mo_path.DropboxPath) error
	del = func(path mo_path.DropboxPath) error {
		ui.Progress(z.ProgressDelete.With("Path", path.Path()))
		_, err := sv_file.NewFiles(ctx).Remove(path)
		if err == nil {
			return nil
		}
		de := dbx_error.NewErrors(err)
		if de.IsTooManyFiles() {
			entries, err := sv_file.NewFiles(ctx).List(path)
			if err != nil {
				return err
			}
			for _, entry := range entries {
				if f, ok := entry.File(); ok {
					del(f.Path())
				}
				if f, ok := entry.Folder(); ok {
					del(f.Path())
				}
			}
			return del(path)
		} else {
			return err
		}
	}

	return del(z.Path)
}

func (z *Permdelete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Permdelete{}, func(r rc_recipe.Recipe) {
		m := r.(*Permdelete)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("test-perm-delete")
	})
}
