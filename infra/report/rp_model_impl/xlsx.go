package rp_model_impl

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"path/filepath"
	"sync"
	"time"
)

const (
	xlsxThemeColor = "ff548235"
	XlsxMaxRows    = 10000
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

func NewXlsx(name string, row interface{}, ctl app_control.Control, opts ...rp_model.ReportOpt) (r rp_model.Report, err error) {
	parser := NewColumn(row, opts...)
	x := &Xlsx{
		fileAvailable: false,
		name:          name,
		ctl:           ctl,
		parser:        parser,
	}
	if err = x.open(); err != nil {
		return nil, err
	}
	return x, nil
}

type Xlsx struct {
	ctl           app_control.Control
	name          string
	omitError     bool
	rotateCount   int
	rotateFailed  bool
	filePath      string
	fileAvailable bool
	file          *xlsx.File
	sheet         *xlsx.Sheet
	parser        Column
	index         int
	fileIndex     int
	mutex         sync.Mutex
}

func (z *Xlsx) rotate() {
	l := z.ctl.Log()

	// Ignore once rotate failed
	if z.rotateFailed {
		return
	}

	l.Debug("Rotate xlsx report", zap.Int("fileIndex", z.fileIndex))

	// rotate
	if err := z.open(); err != nil {
		if !z.omitError {
			z.ctl.UI().Error("report.xlsx.unable_to_open", app_msg.P{
				"Path":  z.filePath,
				"Error": err.Error(),
			})
			z.omitError = true
		}
		z.rotateFailed = true
	}
	z.rotateCount++
}

func (z *Xlsx) open() (err error) {
	l := z.ctl.Log()
	if z.fileAvailable {
		path := z.filePath
		if z.rotateCount == 0 {
			path = filepath.Join(z.ctl.Workspace().Report(), z.name+"_0000.xlsx")
		}
		if err = z.file.Save(path); err != nil {
			l.Debug("Unable to save file", zap.Error(err), zap.String("path", path))
			return err
		}
	}

	z.fileAvailable = false

	name := z.name
	if z.fileIndex != 0 {
		name = fmt.Sprintf("%s_%04d", z.name, z.fileIndex)
	}
	l = l.With(zap.String("name", name))
	z.filePath = filepath.Join(z.ctl.Workspace().Report(), name+".xlsx")

	file := xlsx.NewFile()
	l.Debug("Create xlsx report", zap.String("filePath", z.filePath))
	sheet, err := file.AddSheet(z.name)
	if err != nil {
		l.Debug("Unable to add sheet", zap.Error(err))
		return err
	}

	z.fileAvailable = true
	z.file = file
	z.sheet = sheet
	z.fileIndex++
	z.index = 0

	return nil
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
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSuccess.Key(), rp_model.MsgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *Xlsx) Failure(err error, input interface{}) {
	z.Row(rowForFailure(z.ctl.UI(), err, input))
}

func (z *Xlsx) Skip(reason app_msg.Message, input interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSkip.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: nil,
	})
}

func (z *Xlsx) Row(row interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if !z.fileAvailable {
		return
	}

	if z.index == 0 {
		header := make([]interface{}, 0)
		for _, h := range z.parser.Header() {
			header = append(header, h)
		}
		z.addRow(header, xlsxHeaderStyle())
	}
	z.addRow(z.parser.Values(row), xlsxDataStyle())
	z.index++

	if z.index > XlsxMaxRows {
		z.rotate()
	}
}

func (z *Xlsx) Close() {
	if !z.fileAvailable {
		return
	}
	if z.index > 0 {
		z.file.Save(z.filePath)
	}
}
