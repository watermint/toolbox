package mo_content

import "encoding/json"

type Content struct {
	Raw  json.RawMessage
	Name string `path:"content.name"`
	Path string `path:"content.path"`
}
