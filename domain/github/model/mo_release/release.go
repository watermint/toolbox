package mo_release

import (
	"encoding/json"
)

type Release struct {
	Raw     json.RawMessage
	Id      string `path:"id" json:"id"`
	TagName string `path:"tag_name" json:"tag_name"`
	Name    string `path:"name" json:"name"`
	Draft   bool   `path:"draft" json:"draft"`
	Url     string `path:"html_url" json:"url"`
}
