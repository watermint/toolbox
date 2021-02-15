package sv_thread

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMsgImpl_List(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := goog_context_impl.NewMock(goog_context_impl.EndpointGoogleApis, "mock", ctl)
		sv := New(mc, "me")
		_, err := sv.List()
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMsgImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := goog_context_impl.NewMock(goog_context_impl.EndpointGoogleApis, "mock", ctl)
		sv := New(mc, "me")
		_, err := sv.Resolve("123", ResolveFormat("full"))
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
