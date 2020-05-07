package dc_recipe

type ReportColumn struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (z *ReportColumn) ColName() string {
	return z.Name
}

func (z *ReportColumn) ColDesc() string {
	return z.Desc
}

func (z *ReportColumn) ColExample() string {
	return ""
}

type Report struct {
	Name    string          `json:"name"`
	Desc    string          `json:"desc"`
	Columns []*ReportColumn `json:"columns"`
}

func (z *Report) RowsName() string {
	return z.Name
}

func (z *Report) RowsDesc() string {
	return z.Desc
}

func (z *Report) RowsCols() []DocColumn {
	cols := make([]DocColumn, 0)
	for _, c := range z.Columns {
		cols = append(cols, c)
	}
	return cols
}

func (z *Report) RowsHasExample() bool {
	return false
}
