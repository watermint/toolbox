package dbx_auth_attr

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConsole(c app_control.Control, peerName string, app api_auth.OAuthApp) api_auth.OAuthConsole {
	l := c.Log().With(esl.String("peerName", peerName))
	var oa api_auth.OAuthConsole

	// Make redirect impl. hidden for while
	if f, found := c.Feature().OptInGet(&api_auth_oauth.OptInFeatureRedirect{}); found && f.OptInIsEnabled() {
		oa = api_auth_oauth.NewConsoleRedirect(c, peerName, app)
	} else {
		oa = api_auth_oauth.NewConsoleOAuth(c, peerName, app)
	}
	aa := NewConsoleAttr(c, oa, app)
	if c.Feature().IsSecure() {
		l.Debug("Skip caching")
		return aa
	}
	l.Debug("Token cache enabled")
	ca := api_auth_oauth.NewConsoleCache(c, aa, app)
	return ca
}

func NewConsoleAttr(c app_control.Control, auth api_auth.OAuthConsole, app api_auth.OAuthApp) api_auth.OAuthConsole {
	return &Attr{
		app:  app,
		ctl:  c,
		auth: auth,
	}
}
