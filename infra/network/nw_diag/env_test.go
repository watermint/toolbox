package nw_diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRuntime(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		err := Runtime(ctl)
		if err != nil {
			t.Error(err)
		}
	})
}
