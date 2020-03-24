package dbx_async

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/api/api_response"
)

var (
	ErrorNoResult = errors.New("no result")
)

type responseImpl struct {
	res            api_response.Response
	complete       gjson.Result
	completeExists bool
}

func (z *responseImpl) Json() (res gjson.Result, err error) {
	return z.complete, nil
}

func (z *responseImpl) Model(v interface{}) error {
	return api_parser.ParseModel(v, z.complete)
}

func (z *responseImpl) ModelWithPath(v interface{}, path string) error {
	if !z.completeExists {
		return ErrorNoResult
	}
	return api_parser.ParseModel(v, z.complete.Get(path))
}
