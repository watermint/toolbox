package api_async_impl

import (
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type MockAsync struct {
}

func (z *MockAsync) Param(p interface{}) api_async.Async {
	return z
}

func (z *MockAsync) Status(endpoint string) api_async.Async {
	return z
}

func (z *MockAsync) PollInterval(second int) api_async.Async {
	return z
}

func (z *MockAsync) OnSuccess(func(res api_async.Response) error) api_async.Async {
	return z
}

func (z *MockAsync) OnFailure(func(err error) error) api_async.Async {
	return z
}

func (z *MockAsync) Call() (res api_async.Response, err error) {
	return nil, qt_errors.ErrorMock
}
