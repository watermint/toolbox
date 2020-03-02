package api_async_impl

import (
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"testing"
)

func TestMockAsync_Call(t *testing.T) {
	var m api_async.Async
	m = &MockAsync{}
	m = m.Status("")
	m = m.Param(map[string]string{"test": "test"})
	m = m.OnFailure(func(err error) error { return nil })
	m = m.OnSuccess(func(res api_async.Response) error { return nil })
	m = m.PollInterval(1)
	_, err := m.Call()
	if err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
