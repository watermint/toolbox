package response_impl

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
)

func newEmptyBody() response.Body {
	return &emptyBody{}
}

type emptyBody struct {
}

func (z emptyBody) Json() tjson.Json {
	return tjson.Null()
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
	return "", response.ErrorNoContent
}

func (z emptyBody) AsJson() (tjson.Json, error) {
	return nil, response.ErrorNoContent
}
