package mo_thread

import "encoding/json"

type Thread struct {
	Raw     json.RawMessage
	Id      string `json:"id" path:"id"`
	Snippet string `json:"snippet" path:"snippet"`
}
