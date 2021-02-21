package sheet

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	xlsx2 "github.com/watermint/toolbox/recipe/util/xlsx"
	"os"
	"path/filepath"
)

type Import struct {
	File     mo_path.ExistingFileSystemPath
	Data     da_griddata.GridDataInput
	Sheet    string
	Position string
}

func (z *Import) Preset() {
	z.Position = "A1"
}

func (z *Import) Exec(c app_control.Control) error {
	l := c.Log()
	f, err := xlsx.OpenFile(z.File.Path())
	if err != nil {
		l.Debug("Unable to open a file", esl.Error(err))
		return err
	}
	sheet, ok := f.Sheet[z.Sheet]
	if !ok {
		l.Debug("The sheet not found for a name. try add a sheet", esl.String("sheet", z.Sheet))
		sheet, err = f.AddSheet(z.Sheet)
		if err != nil {
			l.Debug("Unable to add a sheet", esl.Error(err))
			return err
		}
	}

	startCol, startRow, err := xlsx.GetCoordsFromCellIDString(z.Position)
	if err != nil {
		l.Debug("Unable to determine cell position", esl.Error(err))
		return err
	}

	err = z.Data.EachRow(func(cols []interface{}, rowIndex int) error {
		for c, col := range cols {
			cell := sheet.Cell(startRow+rowIndex, startCol+c)
			switch cv := col.(type) {
			case string:
				cell.SetString(cv)
			case int, int8, int16, int32, int64:
				cell.SetInt64(col.(int64))
			case float32, float64:
				cell.SetFloat(col.(float64))
			default:
				cell.SetString(fmt.Sprintf("%v", col))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return f.Save(z.File.Path())
}

func (z *Import) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("xlsx", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()
	d, err := qt_file.MakeTestCsv("import.csv")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(d)
	}()

	f := filepath.Join(p, "test.xlsx")
	err = rc_exec.Exec(c, &xlsx2.Create{}, func(r rc_recipe.Recipe) {
		m := r.(*xlsx2.Create)
		m.File = mo_path.NewFileSystemPath(f)
		m.Sheet = "Sheet1"
	})
	if err != nil {
		return err
	}

	return rc_exec.Exec(c, &Import{}, func(r rc_recipe.Recipe) {
		m := r.(*Import)
		m.File = mo_path.NewExistingFileSystemPath(f)
		m.Data.SetFilePath(d)
		m.Sheet = "Sheet1"
	})
}
