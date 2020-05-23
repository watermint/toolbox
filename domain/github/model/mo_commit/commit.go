package mo_commit

import "encoding/json"

type Commit struct {
	Raw json.RawMessage
	Sha string `path:"sha" json:"sha"`
	Url string `path:"url" json:"url"`
}
