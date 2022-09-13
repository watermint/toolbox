package api_conn_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func OAuthConnectByRedirect(session api_auth.OAuthSessionData, ctl app_control.Control) (entity api_auth.OAuthEntity, useMock bool, err error) {
	l := ctl.Log()

	if isTest, mock, err := isTestMode(ctl); isTest {
		return api_auth.NewNoAuthOAuthEntity(), mock, err
	}

	a := api_auth_oauth.NewSessionRepository(api_auth_oauth.NewSessionRedirect(ctl), ctl.AuthRepository())
	l.Debug("Start auth sequence", esl.Strings("scopes", session.Scopes))
	entity, err = a.Start(session)
	return entity, false, err
}
