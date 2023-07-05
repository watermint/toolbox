package format

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Xlsx struct {
	Source   mo_path.ExistingFileSystemPath
	Template mo_path.ExistingFileSystemPath
	Dest     mo_path.FileSystemPath
	Sheet    string
	Position string
}

func (z *Xlsx) Preset() {
	z.Position = "A1"
}

func (z *Xlsx) Exec(c app_control.Control) error {
	l := c.Log()
	f, err := xlsx.OpenFile(z.Source.Path())
	if err != nil {
		return err
	}
	sheet, ok := f.Sheet[z.Sheet]
	if !ok {
		return err
	}
	startCol, startRow, err := xlsx.GetCoordsFromCellIDString(z.Position)
	if err != nil {
		l.Debug("Unable to determine cell position", esl.Error(err))
		return err
	}

	d, err := os.Create(z.Dest.Path())
	if err != nil {
		return err
	}
	defer func() {
		_ = d.Close()
	}()

	tmplData, err := os.ReadFile(z.Template.Path())
	if err != nil {
		return err
	}
	rowFmt, err := template.New("format").Parse(string(tmplData))
	if err != nil {
		return err
	}

	for r, row := range sheet.Rows {
		rowData := make(map[string]interface{})
		hasData := false
		for col, cell := range row.Cells {
			if r < startRow || col < startCol {
				continue
			}
			hasData = true
			coord := xlsx.GetCellIDStringFromCoords(col, r)
			for i := 0; i <= 9; i++ {
				coord = strings.ReplaceAll(coord, fmt.Sprintf("%d", i), "")
			}
			rowData[coord], err = cell.FormattedValue()
			if err != nil {
				l.Debug("Unable to retrieve formatted value, fallback to Value", esl.Error(err))
				rowData[coord] = cell.Value
			}

			l.Debug("Cell", esl.Any("cell", cell))
		}
		if hasData {
			err := rowFmt.Execute(d, rowData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Xlsx) Test(c app_control.Control) error {
	d, err := qt_file.MakeTestFolder("xlsx", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()
	src := xlsx.NewFile()
	sheet, err := src.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	row1 := sheet.AddRow()
	row1.AddCell().SetString("A1")
	row1.AddCell().SetString("B1")
	row2 := sheet.AddRow()
	row2.AddCell().SetString("A2")
	row2.AddCell().SetString("B2")
	srcPath := filepath.Join(d, "test.xlsx")
	err = src.Save(srcPath)
	if err != nil {
		return err
	}
	tmplPath := filepath.Join(d, "format.txt")
	tmplData := `A: {{.A}}\nB: {{.B}}`
	err = os.WriteFile(tmplPath, []byte(tmplData), 0644)
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Xlsx{}, func(r rc_recipe.Recipe) {
		m := r.(*Xlsx)
		m.Source = mo_path.NewExistingFileSystemPath(srcPath)
		m.Template = mo_path.NewExistingFileSystemPath(tmplPath)
		m.Dest = mo_path.NewFileSystemPath(d)
		m.Sheet = "Sheet1"
		m.Position = "A1"
	})
}
