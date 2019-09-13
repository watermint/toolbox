package app_report

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"path/filepath"
	"sync"
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
		ctl:      ctl,
		filePath: filePath,
		file:     file,
		sheet:    sheet,
		parser:   parser,
	}
	return r, nil
}

type Xlsx struct {
	ctl      app_control.Control
	filePath string
	file     *xlsx.File
	sheet    *xlsx.Sheet
	parser   Column
	index    int
	mutex    sync.Mutex
}

func (z *Xlsx) addRow(cols []interface{}, style *xlsx.Style) error {
	row := z.sheet.AddRow()
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
		Status: z.ctl.UI().Text(msgSuccess.Key(), msgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.ctl.UI().Text(msgFailure.Key(), msgFailure.Params()...),
		Reason: z.ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.ctl.UI().Text(msgSkip.Key(), msgFailure.Params()...),
		Reason: z.ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Row(row interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.index == 0 {
		header := make([]interface{}, 0)
		for _, h := range z.parser.Header() {
			header = append(header, h)
		}
		z.addRow(header, xlsxHeaderStyle())
	}
	z.addRow(z.parser.Values(row), xlsxDataStyle())
	z.index++
}

func (z *Xlsx) Flush() {
	z.file.Save(z.filePath)
}

func (z *Xlsx) Close() {
	z.file.Save(z.filePath)
}
