package sync

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"os"
)

type Up struct {
	Peer        rc_conn.ConnUserFile
	LocalPath   mo_path.FileSystemPath
	DropboxPath mo_path.DropboxPath
	ChunkSizeKb int
	Upload      *file.Upload
}

func (z *Up) Preset() {
	z.ChunkSizeKb = 150 * 1024
}

func (z *Up) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*file.Upload)
		ru.EstimateOnly = false
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = true
		ru.CreateFolder = true
		ru.Context = z.Peer.Context()
		if z.ChunkSizeKb > 0 {
			ru.ChunkSizeKb = z.ChunkSizeKb
		}
	})
}

func (z *Up) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Up{}, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = mo_path.NewFileSystemPath(os.TempDir())
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("up")
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}
