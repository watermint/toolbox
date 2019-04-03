package api_async_impl

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
)

type responseImpl struct {
	res            api_rpc.Response
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
		return errors.New("no result")
	}
	return api_parser.ParseModel(v, z.complete.Get(path))
}
