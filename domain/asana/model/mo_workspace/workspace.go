package mo_workspace

import "encoding/json"

type Workspace struct {
	Raw            json.RawMessage
	Gid            string `json:"gid" path:"gid"`
	ResourceType   string `json:"resource_type" path:"resource_type"`
	Name           string `json:"name" path:"name"`
	IsOrganization bool   `json:"is_organization" path:"is_organization"`
}
