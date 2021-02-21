package sheet

import (
	"errors"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Export struct {
	File  mo_path.ExistingFileSystemPath
	Data  da_griddata.GridDataOutput
	Sheet string
}

func (z *Export) Preset() {
}

func (z *Export) Exec(c app_control.Control) error {
	l := c.Log()
	f, err := xlsx.OpenFile(z.File.Path())
	if err != nil {
		l.Debug("Unable to open a file", esl.Error(err))
		return err
	}
	sheet, ok := f.Sheet[z.Sheet]
	if !ok {
		l.Debug("The sheet not found")
		return errors.New("the sheet not found")
	}

	for _, sheetRow := range sheet.Rows {
		columns := make([]interface{}, 0)
		for _, cell := range sheetRow.Cells {
			fv, err := cell.FormattedValue()
			if err != nil {
				l.Debug("Unable to retrieve formatted value, fallback to Value", esl.Error(err))
				columns = append(columns, cell.Value)
			} else {
				columns = append(columns, fv)
			}
		}
		z.Data.Row(columns)
	}
	return nil
}

func (z *Export) Test(c app_control.Control) error {
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

	return rc_exec.Exec(c, &Export{}, func(r rc_recipe.Recipe) {
		m := r.(*Export)
		m.File = mo_path.NewExistingFileSystemPath(f)
		m.Sheet = "Sheet1"
	})
}
