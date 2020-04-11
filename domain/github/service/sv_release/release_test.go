package sv_release

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"testing"
)

func TestReleaseImpl_List(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc, "watermint", "toolbox")
	if _, err := sv.List(); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
