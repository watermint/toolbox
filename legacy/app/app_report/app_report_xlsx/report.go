package app_report_xlsx

import (
	"fmt"
	"github.com/tealeg/xlsx"
	app2 "github.com/watermint/toolbox/legacy/app"
	app_report_column2 "github.com/watermint/toolbox/legacy/app/app_report/app_report_column"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

const (
	themeColor = "ff548235"
)

type XlsxReport struct {
	ec          *app2.ExecContext
	parsers     map[string]app_report_column2.Row
	sheets      map[string]*xlsx.Sheet
	headerStyle *xlsx.Style
	dataStyle   *xlsx.Style
	file        *xlsx.File
	ReportPath  string
}

func (z *XlsxReport) log() *zap.Logger {
	return z.ec.Log().With(zap.String("path", z.ReportPath))
}

func (z *XlsxReport) Init(ec *app2.ExecContext) error {
	z.ec = ec
	if z.sheets == nil {
		z.sheets = make(map[string]*xlsx.Sheet)
	}
	if z.parsers == nil {
		z.parsers = make(map[string]app_report_column2.Row)
	}
	if z.file == nil {
		z.file = xlsx.NewFile()
	}
	if z.headerStyle == nil {
		z.headerStyle = xlsx.NewStyle()
		z.headerStyle.ApplyFill = true
		z.headerStyle.Fill = xlsx.Fill{
			PatternType: "solid",
			FgColor:     themeColor,
		}
		z.headerStyle.Font.Color = "ffffffff"
	}
	if z.dataStyle == nil {
		z.dataStyle = xlsx.NewStyle()
		z.dataStyle.ApplyBorder = true
		z.dataStyle.Border.Bottom = "thin"
		z.dataStyle.Border.BottomColor = themeColor
		z.dataStyle.Border.Left = "thin"
		z.dataStyle.Border.LeftColor = themeColor
		z.dataStyle.Border.Top = "thin"
		z.dataStyle.Border.TopColor = themeColor
		z.dataStyle.Border.Right = "thin"
		z.dataStyle.Border.RightColor = themeColor
	}
	return nil
}

func (z *XlsxReport) Close() {
	if z.file == nil {
		return
	}
	z.file.Save(filepath.Join(z.ReportPath, "report.xlsx"))
}

func (z *XlsxReport) appendRow(cols []interface{}, sheet *xlsx.Sheet, style *xlsx.Style) error {
	row := sheet.AddRow()
	for _, col := range cols {
		cell := row.AddCell()
		cell.SetStyle(style)
		if col == nil {
			continue
		}
		switch c := col.(type) {
		case string:
			cell.SetString(c)
		case int:
			cell.SetInt(c)
		case int64:
			cell.SetInt64(c)
		case int8, int16, int32, uint, uint8, uint16, uint32, uint64:
			cell.SetFormula(fmt.Sprintf("%d", c))
		case time.Time:
			cell.SetDateTime(c)
		default:
			cell.SetValue(c)
		}
	}
	return nil
}

func (z *XlsxReport) Report(row interface{}) (err error) {
	name := app_report_column2.RowName(row)

	cols, e := z.parsers[name]
	if !e {
		cols = app_report_column2.NewRow(row, z.ec)
		z.parsers[name] = cols
	}

	sheet, e := z.sheets[name]
	if !e {
		sheet, err = z.file.AddSheet(name)
		if err != nil {
			z.log().Debug("unable to add sheet", zap.String("name", name), zap.Error(err))
			return err
		}
		z.sheets[name] = sheet

		header := cols.Header()
		hi := make([]interface{}, len(header))
		for i := len(header) - 1; i >= 0; i-- {
			hi[i] = header[i]
		}
		if err := z.appendRow(hi, sheet, z.headerStyle); err != nil {
			z.log().Debug("unable to add header", zap.Strings("header", cols.Header()))
		}
	}

	return z.appendRow(cols.Values(row), sheet, z.dataStyle)
}