package api_list

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
)

type List interface {
	Param(param interface{}) List
	Continue(endpoint string) List
	UseHasMore(use bool) List
	ResultTag(tag string) List
	OnFailure(func(err error) error) List
	OnResponse(func(res api_rpc.Response) error) List
	OnEntry(func(entry ListEntry) error) List
	Call() (err error)
}

type ListEntry interface {
	Json() (res gjson.Result, err error)
	Model(v interface{}) error
	ModelWithPath(v interface{}, path string) error
}
