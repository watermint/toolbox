package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func ConnectByRedirect(session api_auth.OAuthSessionData, ctl app_control.Control) (entity api_auth.OAuthEntity, useMock bool, err error) {
	l := ctl.Log()
	ui := ctl.UI()

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return api_auth.NewNoAuthOAuthEntity(), true, nil
	}
	if ctl.Feature().IsTest() && qt_endtoend.IsSkipEndToEndTest() {
		l.Debug("Skip end to end test")
		return api_auth.NewNoAuthOAuthEntity(), true, nil
	}
	if !ui.IsConsole() {
		l.Debug("non console UI is not supported")
		return api_auth.NewNoAuthOAuthEntity(), false, qt_errors.ErrorUnsupportedUI
	}
	a := api_auth_oauth.NewSessionRepository(api_auth_oauth.NewSessionRedirect(ctl), ctl.AuthRepository())
	l.Debug("Start auth sequence", esl.Strings("scopes", session.Scopes))
	entity, err = a.Start(session)
	return entity, false, err
}
