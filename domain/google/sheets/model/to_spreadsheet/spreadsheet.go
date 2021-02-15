package to_spreadsheet

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
