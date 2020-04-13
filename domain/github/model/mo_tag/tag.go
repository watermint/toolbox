package mo_tag

import "encoding/json"

type Tag struct {
	Raw     json.RawMessage
	TagName string `path:"tag" json:"tag"`
	Sha     string `path:"sha" json:"sha"`
	Message string `path:"message" json:"message"`
	Url     string `path:"url" json:"url"`
}
