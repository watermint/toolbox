package mo_paper

type LegacyPaper struct {
	Owner    string `path:"owner" json:"owner"`
	Title    string `path:"title" json:"title"`
	Revision int64  `path:"revision" json:"revision"`
	MimeType string `path:"mime_type" json:"mime_type"`
}

type LegacyFolder struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
