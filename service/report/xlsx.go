package report

import (
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/tealeg/xlsx"
	"github.com/watermint/toolbox/infra/util"
	"sync"
)

const (
	XLSX_THEME_COLOR = "FF548235"
)

func writeXlsxHeader(sheet *xlsx.Sheet, header ReportHeader) {
	row := sheet.AddRow()

	hs := xlsx.NewStyle()
	hs.ApplyFill = true
	hs.Fill = xlsx.Fill{
		PatternType: "solid",
		FgColor:     XLSX_THEME_COLOR,
	}
	hs.Font.Color = "FFFFFFFF"

	for _, h := range header.Headers {
		cell := row.AddCell()
		cell.Value = h
		cell.SetStyle(hs)
	}
}

func writeXlsxData(sheet *xlsx.Sheet, data ReportData) {
	row := sheet.AddRow()

	ds := xlsx.NewStyle()
	ds.ApplyBorder = true
	ds.Border.Bottom = "thin"
	ds.Border.BottomColor = XLSX_THEME_COLOR
	ds.Border.Left = "thin"
	ds.Border.LeftColor = XLSX_THEME_COLOR
	ds.Border.Top = "thin"
	ds.Border.TopColor = XLSX_THEME_COLOR
	ds.Border.Right = "thin"
	ds.Border.RightColor = XLSX_THEME_COLOR

	for _, d := range data.Data {
		cell := row.AddCell()
		cell.SetStyle(ds)
		switch a := d.(type) {
		case uint64:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case uint32:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case uint16:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case uint8:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case uint:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case int64:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case int32:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case int16:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case int8:
			cell.SetFormula(fmt.Sprintf("%d", a))
		case int:
			cell.SetFormula(fmt.Sprintf("%d", a))
		default:
			cell.SetValue(d)
		}
	}
}

func WriteXlsx(path, sheetName string, report chan ReportRow, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	f := xlsx.NewFile()
	sheet, err := f.AddSheet(sheetName)
	if err != nil {
		return err
	}
	defer f.Save(path)

	for r := range report {
		switch row := r.(type) {
		case ReportHeader:
			seelog.Tracef("Header(%s)", util.MarshalObjectToString(row.Headers))
			writeXlsxHeader(sheet, row)

		case ReportData:
			seelog.Tracef("Data(%s)", util.MarshalObjectToString(row.Data))
			writeXlsxData(sheet, row)

		case ReportEOF:
			seelog.Trace("EOF")
			return nil

		default:
			seelog.Warnf("Unexpected row")
			return errors.New("Unexpected row detected")
		}
	}
	return nil
}
