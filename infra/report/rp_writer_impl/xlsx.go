package rp_writer_impl

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"path/filepath"
	"sync"
	"time"
)

const (
	xlsxThemeColor      = "ff548235"
	XlsxMaxRows         = 10000
	XlsxMaxMemoryTarget = 4 * 1_048_576 // 4MB
)

type MsgXlsxWriter struct {
	UnableToOpen app_msg.Message
}

var (
	MXlsxWriter = app_msg.Apply(&MsgXlsxWriter{}).(*MsgXlsxWriter)
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

func NewXlsxWriter(name string, ctl app_control.Control) rp_writer.Writer {
	return &xlsxWriter{
		name: name,
		ctl:  ctl,
	}
}

type xlsxWriter struct {
	name           string
	nameSuffix     string
	index          int
	path           string
	mutex          sync.Mutex
	ctl            app_control.Control
	colModel       rp_column.Column
	omitError      bool
	rotateCount    int
	rotateFailed   bool
	filePath       string
	fileAvailable  bool
	file           *xlsx.File
	fileIndex      int
	sheet          *xlsx.Sheet
	estMemoryUsage int64
}

func (z *xlsxWriter) rotate() {
	l := z.ctl.Log()

	// Ignore once rotate failed
	if z.rotateFailed {
		return
	}

	l.Debug("Rotate xlsx report", es_log.Int("fileIndex", z.fileIndex))

	// rotate
	if err := z.open(); err != nil {
		if !z.omitError {
			z.ctl.UI().Error(MXlsxWriter.UnableToOpen.
				With("Path", z.filePath).
				With("Error", err.Error()))
			z.omitError = true
		}
		z.rotateFailed = true
	}
	z.rotateCount++
}

func (z *xlsxWriter) open() (err error) {
	l := z.ctl.Log()
	if z.fileAvailable {
		path := z.filePath
		if z.rotateCount == 0 {
			path = filepath.Join(z.ctl.Workspace().Report(), z.name+z.nameSuffix+"_0000.xlsx")
		}
		if err = z.file.Save(path); err != nil {
			l.Debug("Unable to save file", es_log.Error(err), es_log.String("path", path))
			return err
		}
	}

	z.fileAvailable = false

	name := z.name
	if z.fileIndex != 0 {
		name = fmt.Sprintf("%s_%04d", z.name, z.fileIndex)
	}
	l = l.With(es_log.String("name", name))
	z.filePath = filepath.Join(z.ctl.Workspace().Report(), name+z.nameSuffix+".xlsx")

	file := xlsx.NewFile()
	l.Debug("Create xlsx report", es_log.String("filePath", z.filePath))
	sheetName := z.name
	if len(sheetName) >= 31 {
		sheetName = sheetName[:30]
	}
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		l.Debug("Unable to add sheet", es_log.Error(err))
		return err
	}

	z.fileAvailable = true
	z.file = file
	z.sheet = sheet
	z.fileIndex++
	z.index = 0
	z.estMemoryUsage = 0

	return nil
}

func (z *xlsxWriter) addRow(cols []interface{}, style *xlsx.Style) error {
	rowSize := 0
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
		rowSize += len(cell.String())
	}
	z.estMemoryUsage += int64(rowSize)
	return nil
}

func (z *xlsxWriter) Name() string {
	return z.name
}

func (z *xlsxWriter) Row(r interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if !z.fileAvailable {
		return
	}

	if z.index == 0 {
		header := make([]interface{}, 0)
		for _, h := range z.colModel.Header() {
			header = append(header, h)
		}
		z.addRow(header, xlsxHeaderStyle())
	}
	z.addRow(z.colModel.Values(r), xlsxDataStyle())
	z.index++

	if z.index > XlsxMaxRows || z.estMemoryUsage > XlsxMaxMemoryTarget {
		z.rotate()
	}
}

func (z *xlsxWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	z.ctl = ctl
	z.colModel = rp_column_impl.NewModel(model, opts...)
	z.fileAvailable = false
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}
	z.nameSuffix = ro.ReportSuffix

	if err := z.open(); err != nil {
		return err
	}
	return nil
}

func (z *xlsxWriter) Close() {
	if !z.fileAvailable {
		return
	}
	if z.index > 0 {
		z.file.Save(z.filePath)
	}
}
