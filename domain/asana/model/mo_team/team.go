package mo_team

import "encoding/json"

type Team struct {
	Raw          json.RawMessage
	Gid          string `json:"gid" path:"gid"`
	Name         string `json:"name" path:"name"`
	ResourceType string `json:"resource_type" path:"resource_type"`
}
