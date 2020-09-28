package sv_content

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCtsImpl_Get(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock("mock", ctl)
		sv := New(mc, "watermint", "toolbox")
		if _, err := sv.Get("/README.md"); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
