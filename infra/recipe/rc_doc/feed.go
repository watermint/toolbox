package rc_doc

type FeedColumn struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Example string `json:"example"`
}

func (z FeedColumn) ColName() string {
	return z.Name
}

func (z FeedColumn) ColDesc() string {
	return z.Desc
}

func (z FeedColumn) ColExample() string {
	return z.Example
}

type Feed struct {
	Name    string        `json:"name"`
	Desc    string        `json:"desc"`
	Columns []*FeedColumn `json:"columns"`
}

func (z Feed) RowsName() string {
	return z.Name
}

func (z Feed) RowsDesc() string {
	return z.Desc
}

func (z Feed) RowsCols() []DocColumn {
	cols := make([]DocColumn, 0)
	for _, c := range z.Columns {
		cols = append(cols, c)
	}
	return cols
}

func (z Feed) RowsHasExample() bool {
	return true
}
