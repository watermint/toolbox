package bo_sheet

import "encoding/json"

type Sheet struct {
	Raw         json.RawMessage
	SheetId     string `json:"sheet_id" path:"properties.sheetId"`
	Title       string `json:"title" path:"properties.title"`
	Index       int    `json:"index" path:"properties.index"`
	SheetType   string `json:"sheet_type" path:"properties.sheetType"`
	RowCount    int    `json:"row_count" path:"properties.gridProperties.rowCount"`
	ColumnCount int    `json:"column_count" path:"properties.gridProperties.columnCount"`
}

type ValueRange struct {
	Raw            json.RawMessage
	Range          string `path:"range" json:"range"`
	MajorDimension string `path:"majorDimension" json:"major_dimension"`
}

type ValueRows struct {
	Columns []interface{}
}
