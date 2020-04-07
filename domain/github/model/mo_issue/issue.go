package mo_issue

import "encoding/json"

type Issue struct {
	Raw    json.RawMessage
	Id     string `path:"id" json:"id"`
	Number string `path:"number" json:"number"`
	Url    string `path:"url" json:"url"`
	Title  string `path:"title" json:"title"`
	State  string `path:"state" json:"state"`
}
