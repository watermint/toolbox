package sv_profile

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"testing"
)

func TestCurrentImpl_User(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc)
	if _, err := sv.User(); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
