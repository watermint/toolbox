package sv_project

import (
	"github.com/watermint/toolbox/domain/asana/api/as_context_impl"
	"github.com/watermint/toolbox/domain/asana/model/mo_team"
	"github.com/watermint/toolbox/domain/asana/model/mo_workspace"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestPrjImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := as_context_impl.NewMock(ctl)
		svc := New(ctx)
		_, err := svc.Resolve("12345")
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestPrjImpl_List(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := as_context_impl.NewMock(ctl)
		svc := New(ctx)
		_, err := svc.List(
			Workspace(&mo_workspace.Workspace{Gid: "12345"}),
			Team(&mo_team.Team{Gid: "23456"}),
		)
		if err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
