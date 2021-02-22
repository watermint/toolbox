package xlsx

import (
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Create struct {
	rc_recipe.RemarkTransient
	File  mo_path.FileSystemPath
	Sheet string
}

func (z *Create) Preset() {
}

func (z *Create) Exec(c app_control.Control) error {
	f := xlsx.NewFile()
	_, err := f.AddSheet(z.Sheet)
	if err != nil {
		return err
	}
	return f.Save(z.File.Path())
}

func (z *Create) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("xlsx", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()

	return rc_exec.Exec(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.File = mo_path.NewFileSystemPath(filepath.Join(p, "test.xlsx"))
		m.Sheet = "Sheet1"
	})
}
