package mo_paper

type PaperUpdate struct {
	Url           string `json:"url" path:"url"`
	ResultPath    string `json:"result_path" path:"result_path"`
	FileId        string `json:"file_id" path:"file_id"`
	PaperRevision string `json:"paper_revision" path:"paper_revision"`
}
