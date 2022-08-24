package sv_task

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client_impl"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestTaskImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := as_client_impl.NewMock("mock", ctl)
		svc := New(ctx)
		_, err := svc.Resolve("12345")
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTaskImpl_List(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := as_client_impl.NewMock("mock", ctl)
		svc := New(ctx)
		_, err := svc.List(Project(&mo_project.Project{Gid: "12345"}))
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
