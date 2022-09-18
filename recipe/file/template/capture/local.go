package capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_template"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Local struct {
	Path mo_path.ExistingFileSystemPath
	Out  mo_path.FileSystemPath
}

func (z *Local) Preset() {
}

func (z *Local) Exec(c app_control.Control) error {
	lfs := es_filesystem_local.NewFileSystem()
	cp := es_template.NewCapture(lfs, es_template.CaptureOpts{})

	template, err := cp.Capture(es_filesystem_local.NewPath(z.Path.Path()))
	if err != nil {
		return err
	}
	tj, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(z.Out.Path(), tj, 0644)
}

func (z *Local) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("capture", true)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.Exec(c, &Local{}, func(r rc_recipe.Recipe) {
		m := r.(*Local)
		m.Path = mo_path.NewExistingFileSystemPath(f)
		m.Out = mo_path.NewFileSystemPath(filepath.Join(f, "test.json"))
	})
}
