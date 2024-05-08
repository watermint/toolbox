package mo_list

type ListInfo struct {
	NumPages   int `json:"num_pages" path:"num_pages"`
	NumResults int `json:"num_results" path:"num_results"`
	Page       int `json:"page" path:"page"`
	PageSize   int `json:"page_size" path:"page_size"`
}
