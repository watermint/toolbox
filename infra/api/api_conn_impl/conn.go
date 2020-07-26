package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func Connect(scopes []string, peerName string, app api_auth.App, ctl app_control.Control) (ctx api_auth.Context, useMock bool, err error) {
	l := ctl.Log()
	ui := ctl.UI()

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return api_auth.NewNoAuth(), true, nil
	}
	if ctl.Feature().IsTest() && qt_endtoend.IsSkipEndToEndTest() {
		l.Debug("Skip end to end test")
		return api_auth.NewNoAuth(), true, nil
	}
	if !ui.IsConsole() {
		l.Debug("non console UI is not supported")
		return nil, false, qt_errors.ErrorUnsupportedUI
	}
	a := api_auth_impl.NewConsoleRedirect(ctl, peerName, app)
	if !ctl.Feature().IsSecure() {
		l.Debug("Enable cache")
		a = api_auth_impl.NewConsoleCache(ctl, a, app)
	}
	l.Debug("Start auth sequence", esl.Strings("scopes", scopes))
	ctx, err = a.Auth(scopes)
	return ctx, false, err
}
