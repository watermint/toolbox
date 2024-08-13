package bo_spreadsheet

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_sheet"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Spreadsheet struct {
	Raw            json.RawMessage
	SpreadsheetId  string `path:"spreadsheetId" json:"spreadsheet_id"`
	SpreadsheetUrl string `path:"spreadsheetUrl" json:"spreadsheet_url"`
	Title          string `path:"properties.title" json:"title"`
}

func (z Spreadsheet) Sheets() (sheets []bo_sheet.Sheet, err error) {
	sheets = make([]bo_sheet.Sheet, 0)
	m, err := es_json.Parse(z.Raw)
	if err != nil {
		return nil, err
	}

	sheetsData, found := m.FindArray("sheets")
	if !found {
		return
	}
	for _, sheetData := range sheetsData {
		sheet := bo_sheet.Sheet{}
		if err := sheetData.Model(&sheet); err != nil {
			return nil, err
		}
		sheets = append(sheets, sheet)
	}
	return sheets, nil
}
