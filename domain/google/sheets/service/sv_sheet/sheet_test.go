package sv_sheet

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"testing"
)

func TestParseSheetFromBatchUpdateAddSheetResponse(t *testing.T) {
	resSample := `{
	"spreadsheetId": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"replies": [
		{
		"addSheet": {
			"properties": {
			"sheetId": 1234567890,
			"title": "Test",
			"index": 2,
			"sheetType": "GRID",
			"gridProperties": {
				"rowCount": 1000,
				"columnCount": 26
			}
			}
		}
		}
	]
}`

	d, err := parseSheetFromBatchUpdateAddSheetResponse(es_json.MustParseString(resSample))
	if err != nil {
		t.Error(err)
	}
	if d.SheetId != "1234567890" {
		t.Error(d.SheetId)
	}
}
