package dbx_async_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"testing"
)

func TestMockAsync_Call(t *testing.T) {
	var m dbx_async.Async
	m = &MockAsync{}
	m = m.Status("")
	m = m.Param(map[string]string{"test": "test"})
	m = m.PollInterval(1)
	_, err := m.Call()
	if err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
