package api_conn_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func IsTestMode(ctl app_control.Control) (isTest, mock bool, err error) {
	l := ctl.Log()
	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return true, true, nil
	}
	if ctl.Feature().IsTest() && qt_endtoend.IsSkipEndToEndTest() {
		l.Debug("Skip end to end test")
		return true, true, nil
	}
	if !ctl.UI().IsConsole() {
		l.Debug("non console UI is not supported")
		return true, false, qt_errors.ErrorUnsupportedUI
	}
	return false, false, nil
}
