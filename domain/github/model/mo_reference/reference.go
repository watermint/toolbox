package mo_reference

import "encoding/json"

type Reference struct {
	Raw        json.RawMessage
	Ref        string `json:"ref" path:"ref"`
	ObjectType string `json:"object_type" path:"object.type"`
	ObjectSha  string `json:"object_sha" path:"object.sha"`
	ObjectUrl  string `json:"object_url" path:"object.url"`
}
