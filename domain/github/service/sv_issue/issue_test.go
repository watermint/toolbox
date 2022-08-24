package sv_issue

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRepoIssueImpl_List(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		mc := gh_client_impl.NewMock("mock", ctl)
		sv := New(mc, "watermint", "toolbox")
		if _, err := sv.List(); err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
