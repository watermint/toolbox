package response

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
)

type emptyBody struct {
}

func (z emptyBody) ContentLength() int64 {
	return 0
}

func (z emptyBody) Body() []byte {
	return []byte{}
}

func (z emptyBody) BodyString() string {
	return ""
}

func (z emptyBody) File() string {
	return ""
}

func (z emptyBody) IsFile() bool {
	return false
}

func (z emptyBody) AsFile() (string, error) {
	return "", ErrorNoContent
}

func (z emptyBody) AsJson() (tjson.Json, error) {
	return nil, ErrorNoContent
}
