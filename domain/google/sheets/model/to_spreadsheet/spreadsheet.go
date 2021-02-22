package to_spreadsheet

import (
	"errors"
	"regexp"
)

type SpreadsheetProperties struct {
	Title    string `json:"title,omitempty"`
	Locale   string `json:"locale,omitempty"`
	TimeZone string `json:"timeZone,omitempty"`
}

type Spreadsheet struct {
	SpreadsheetId  string                 `json:"spreadsheetId,omitempty"`
	SpreadsheetUrl string                 `json:"spreadsheetUrl,omitempty"`
	Properties     *SpreadsheetProperties `json:"properties,omitempty"`
}

var (
	// https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id
	SpreadsheetIdPattern = regexp.MustCompile(`^([a-zA-Z0-9-_]+)$`)

	ErrorInvalidSpreadsheetId = errors.New("invalid spreadsheet id")
)

func IsValidSpreadsheetId(id string) bool {
	return SpreadsheetIdPattern.MatchString(id)
}
