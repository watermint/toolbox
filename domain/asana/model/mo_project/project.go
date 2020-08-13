package mo_project

import "encoding/json"

type Project struct {
	Raw          json.RawMessage
	Gid          string `json:"gid" path:"gid"`
	ResourceType string `json:"resource_type" path:"resource_type"`
	Name         string `json:"name" path:"name"`
}
