package dbx_async_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type MockAsync struct {
}

func (z *MockAsync) Param(p interface{}) dbx_async.Async {
	return z
}

func (z *MockAsync) Status(endpoint string) dbx_async.Async {
	return z
}

func (z *MockAsync) PollInterval(second int) dbx_async.Async {
	return z
}

func (z *MockAsync) OnSuccess(func(res dbx_async.Response) error) dbx_async.Async {
	return z
}

func (z *MockAsync) OnFailure(func(err error) error) dbx_async.Async {
	return z
}

func (z *MockAsync) Call() (res dbx_async.Response, err error) {
	return nil, qt_errors.ErrorMock
}
