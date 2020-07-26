package mo_label

import "encoding/json"

type Label struct {
	Raw  json.RawMessage
	Id   string `json:"id" path:"id"`
	Name string `json:"name" path:"name"`
	Type string `json:"type" path:"type"`
}
