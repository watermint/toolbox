package response

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
)

func newErrorBody(err error) Body {
	return &errorBody{err: err}
}

type errorBody struct {
	err error
}

func (z errorBody) Json() tjson.Json {
	return tjson.Null()
}

func (z errorBody) Error() error {
	return z.err
}

func (z errorBody) ContentLength() int64 {
	return 0
}

func (z errorBody) Body() []byte {
	return []byte{}
}

func (z errorBody) BodyString() string {
	return ""
}

func (z errorBody) File() string {
	return ""
}

func (z errorBody) IsFile() bool {
	return false
}

func (z errorBody) AsFile() (string, error) {
	return "", ErrorNoContent
}

func (z errorBody) AsJson() (tjson.Json, error) {
	return nil, ErrorNoContent
}
