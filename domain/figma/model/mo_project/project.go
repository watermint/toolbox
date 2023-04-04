package mo_project

import "encoding/json"

type Project struct {
	Raw  json.RawMessage
	Id   string `json:"id" path:"id"`
	Name string `json:"name" path:"name"`
}
