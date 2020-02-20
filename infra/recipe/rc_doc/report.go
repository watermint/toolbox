package rc_doc

type ReportColumn struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Report struct {
	Name    string          `json:"name"`
	Desc    string          `json:"desc"`
	Columns []*ReportColumn `json:"columns"`
}
