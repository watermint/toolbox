package es_response_impl

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

func newEmptyBody() es_response.Body {
	return &emptyBody{}
}

type emptyBody struct {
}

func (z emptyBody) Json() es_json.Json {
	return es_json.Null()
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
	return "", es_response.ErrorNoContent
}

func (z emptyBody) AsJson() (es_json.Json, error) {
	return nil, es_response.ErrorNoContent
}
