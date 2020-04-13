package sync

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Up struct {
	Peer        dbx_conn.ConnUserFile
	LocalPath   mo_path2.ExistingFileSystemPath
	DropboxPath mo_path.DropboxPath
	ChunkSizeKb mo_int.RangeInt
	Upload      *file.Upload
}

func (z *Up) Preset() {
	z.ChunkSizeKb.SetRange(1, 150*1024, 150*1024)
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
		ru.ChunkSizeKb = z.ChunkSizeKb.Value()
	})
}

func (z *Up) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Up{}, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = qt_recipe.NewTestExistingFileSystemFolderPath(c, "up")
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("up")
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}
