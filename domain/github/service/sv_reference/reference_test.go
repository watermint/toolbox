package sv_reference

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestReferenceImpl_Create(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock(ctl)
		sv := New(mc, "watermint", "toolbox")
		_, err := sv.Create(
			"refs/tags/63.4.129",
			"273cb137be80ece8b4a2324e4f2f9bf1eabede36",
		)
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
