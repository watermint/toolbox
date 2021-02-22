package to_cell

type CellData struct {
}

type ValueRange struct {
	Range          string          `json:"range"`
	MajorDimension string          `json:"majorDimension"`
	Values         [][]interface{} `json:"values"`
}
