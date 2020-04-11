package mo_installation

import "encoding/json"

type Installation struct {
	Raw        json.RawMessage
	Id         string `path:"id" json:"id"`
	TargetType string `path:"target_type" json:"target_type"`
}
