package api_list

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
)

type List interface {
	Param(param interface{}) List
	Continue(endpoint string) List
	UseHasMore(use bool) List
	ResultTag(tag string) List
	OnResponse(response func(res response.Response) error) List
	OnEntry(entry func(entry tjson.Json) error) List
	OnLastCursor(f func(cursor string)) List
	Call() (err error)
}
