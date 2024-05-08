package mo_content

import (
	"encoding/json"
)

type Content struct {
	Raw     json.RawMessage
	Type    string `path:"type" json:"type"`
	Name    string `path:"name" json:"name"`
	Path    string `path:"path" json:"path"`
	Sha     string `path:"sha" json:"sha"`
	Size    int    `path:"size" json:"size"`
	Target  string `path:"target" json:"target"`
	Content string `path:"content" json:"-""`
}
