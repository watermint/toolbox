package artifact

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Up struct {
	LocalPath   mo_path.FileSystemPath
	DropboxPath mo_path.DropboxPath
}

func (z *Up) Preset() {
}

func (z *Up) Exec(c app_control.Control) error {
	return nil
}

func (z *Up) Test(c app_control.Control) error {
	tp, err := ioutil.TempDir("", "up")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(tp, "artifact.txt"), []byte(time.Now().String()), 0644); err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(tp)
	}()

	return rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = mo_path.NewFileSystemPath(tp)
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("dev-ci-artifact", time.Now().Format(time.RFC3339))
	})
}
