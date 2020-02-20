package rc_doc

type FeedColumn struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Example string `json:"example"`
}

type Feed struct {
	Name    string        `json:"name"`
	Desc    string        `json:"desc"`
	Columns []*FeedColumn `json:"columns"`
}
