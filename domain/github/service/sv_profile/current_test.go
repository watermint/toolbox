package sv_profile

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestCurrentImpl_User(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock(ctl)
		sv := New(mc)
		if _, err := sv.User(); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
