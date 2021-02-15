package to_sheet

import "github.com/watermint/toolbox/domain/google/sheets/model/to_cell"

type SheetProperties struct {
	SheetId   string `json:"sheetId,omitempty"`
	SheetType string `json:"sheetType,omitempty"`
	Title     string `json:"title,omitempty"`
	Index     *int   `json:"index,omitempty"`
	Hidden    *bool  `json:"hidden,omitempty"`
}

type Sheet struct {
	Properties *SheetProperties `json:"properties"`
}

type GridData struct {
	StartRow       *int                   `json:"startRow,omitempty"`
	StartColumn    *int                   `json:"startColumn,omitempty"`
	RowData        []*RowData             `json:"rowData,omitempty"`
	RowMetadata    []*DimensionProperties `json:"rowMetadata,omitempty"`
	ColumnMetadata []*DimensionProperties `json:"columnMetadata,omitempty"`
}

type DimensionProperties struct {
	HiddenByFilter *bool `json:"hiddenByFilter,omitempty"`
	HiddenByUser   *bool `json:"hiddenByUser,omitempty"`
	PixelSize      *int  `json:"pixelSize,omitempty"`
}

type RowData struct {
	Values []*to_cell.CellData `json:"values,omitempty"`
}
