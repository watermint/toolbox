package nw_diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRuntime(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		err := Runtime(ctl)
		if err != nil {
			t.Error(err)
		}
	})
}
