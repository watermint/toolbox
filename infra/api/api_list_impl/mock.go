package api_list_impl

import (
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_response"
)

type MockList struct {
}

func (z *MockList) Param(param interface{}) api_list.List {
	return z
}

func (z *MockList) Continue(endpoint string) api_list.List {
	return z
}

func (z *MockList) UseHasMore(use bool) api_list.List {
	return z
}

func (z *MockList) ResultTag(tag string) api_list.List {
	return z
}

func (z *MockList) OnFailure(failure func(err error) error) api_list.List {
	return z
}

func (z *MockList) OnResponse(response func(res api_response.Response) error) api_list.List {
	return z
}

func (z *MockList) OnEntry(entry func(entry api_list.ListEntry) error) api_list.List {
	return z
}

func (z *MockList) OnLastCursor(f func(cursor string)) api_list.List {
	return z
}

func (z *MockList) Call() (err error) {
	return nil
}
