package nw_diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestNetwork(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		err := Network(ctl)
		if err != nil {
			t.Error(err)
		}
	})
}
