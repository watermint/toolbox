package dbx_list

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type MockList struct {
}

func (z MockList) Param(param interface{}) api_list.List {
	return &z
}

func (z MockList) Continue(endpoint string) api_list.List {
	return &z
}

func (z MockList) UseHasMore(use bool) api_list.List {
	return &z
}

func (z MockList) ResultTag(tag string) api_list.List {
	return &z
}

func (z MockList) OnResponse(response func(res response.Response) error) api_list.List {
	return &z
}

func (z MockList) OnEntry(entry func(entry tjson.Json) error) api_list.List {
	return &z
}

func (z MockList) OnLastCursor(f func(cursor string)) api_list.List {
	return &z
}

func (z MockList) Call() (err error) {
	return qt_errors.ErrorMock
}
