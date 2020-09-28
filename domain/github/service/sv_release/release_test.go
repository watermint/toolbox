package sv_release

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestReleaseImpl_List(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock("mock", ctl)
		sv := New(mc, "watermint", "toolbox_sandbox")
		if _, err := sv.List(); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})

}

func TestReleaseImpl_CreateDraft(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock("mock", ctl)
		sv := New(mc, "watermint", "toolbox_sandbox")
		if _, err := sv.CreateDraft("0.0.0", "test", "test body", "master"); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestReleaseImpl_Get(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_context_impl.NewMock("mock", ctl)
		sv := New(mc, "watermint", "toolbox_sandbox")
		if _, err := sv.Get("0.0.2"); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
