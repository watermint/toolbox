package dc_license

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestDetect(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		inventory, err := Detect(ctl)
		if err != nil {
			t.Error(err)
			return
		}

		l := ctl.Log()

		for _, info := range inventory {
			l.Info("Library", esl.String("name", info.Package), esl.String("type", info.LicenseType))
		}
	})
}
