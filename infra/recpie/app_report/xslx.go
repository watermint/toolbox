package app_report

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"path/filepath"
	"time"
)

const (
	xlsxThemeColor = "ff548235"
)

func xlsxHeaderStyle() *xlsx.Style {
	headerStyle := xlsx.NewStyle()
	headerStyle.ApplyFill = true
	headerStyle.Fill = xlsx.Fill{
		PatternType: "solid",
		FgColor:     xlsxThemeColor,
	}
	headerStyle.Font.Color = "ffffffff"
	return headerStyle
}

func xlsxDataStyle() *xlsx.Style {
	dataStyle := xlsx.NewStyle()
	dataStyle.ApplyBorder = true
	dataStyle.Border.Bottom = "thin"
	dataStyle.Border.BottomColor = xlsxThemeColor
	dataStyle.Border.Left = "thin"
	dataStyle.Border.LeftColor = xlsxThemeColor
	dataStyle.Border.Top = "thin"
	dataStyle.Border.TopColor = xlsxThemeColor
	dataStyle.Border.Right = "thin"
	dataStyle.Border.RightColor = xlsxThemeColor
	return dataStyle
}

func NewXlsx(name string, row interface{}, ctl app_control.Control) (r Report, err error) {
	path, err := ctl.Workspace().Descendant(reportPath)
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(path, name+".xlsx")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet(name)
	if err != nil {
		return nil, err
	}
	err = file.Save(filePath)
	if err != nil {
		return nil, err
	}
	parser := NewColumn(row, ctl)
	r = &Xlsx{
		Ctl:      ctl,
		FilePath: filePath,
		File:     file,
		Sheet:    sheet,
		Parser:   parser,
	}
	return r, nil
}

type Xlsx struct {
	Ctl      app_control.Control
	FilePath string
	File     *xlsx.File
	Sheet    *xlsx.Sheet
	Parser   Column
	Index    int
}

func (z *Xlsx) addRow(cols []interface{}, style *xlsx.Style) error {
	row := z.Sheet.AddRow()
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

func (z *Xlsx) Success(input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgSuccess.Key(), msgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgFailure.Key(), msgFailure.Params()...),
		Reason: z.Ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgSkip.Key(), msgFailure.Params()...),
		Reason: z.Ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Row(row interface{}) {
	if z.Index == 0 {
		header := make([]interface{}, 0)
		for _, h := range z.Parser.Header() {
			header = append(header, h)
		}
		z.addRow(header, xlsxHeaderStyle())
	}
	z.addRow(z.Parser.Values(row), xlsxDataStyle())
	z.Index++
}

func (z *Xlsx) Flush() {
	z.File.Save(z.FilePath)
}

func (z *Xlsx) Close() {
	z.File.Save(z.FilePath)
}
