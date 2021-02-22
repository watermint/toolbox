package sheet

import (
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Sheet struct {
	Name   string `json:"name"`
	Rows   int    `json:"rows"`
	Cols   int    `json:"cols"`
	Hidden bool   `json:"hidden"`
}

type List struct {
	File   mo_path.ExistingFileSystemPath
	Sheets rp_model.RowReport
}

func (z *List) Preset() {
	z.Sheets.SetModel(&Sheet{})
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	f, err := xlsx.OpenFile(z.File.Path())
	if err != nil {
		l.Debug("Unable to open a file", esl.Error(err))
		return err
	}
	if err := z.Sheets.Open(); err != nil {
		return err
	}

	for _, s := range f.Sheets {

		z.Sheets.Row(&Sheet{
			Name:   s.Name,
			Rows:   s.MaxRow,
			Cols:   s.MaxCol,
			Hidden: s.Hidden,
		})
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("xlsx", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()
	f := filepath.Join(p, "test.xlsx")
	d, err := qt_file.MakeTestCsv("import.csv")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(d)
	}()

	err = rc_exec.Exec(c, &Import{}, func(r rc_recipe.Recipe) {
		m := r.(*Import)
		m.File = mo_path.NewExistingFileSystemPath(f)
		m.Data.SetFilePath(d)
		m.Sheet = "Sheet1"
	})
	if err != nil {
		return err
	}

	return rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.File = mo_path.NewExistingFileSystemPath(f)
	})
}
