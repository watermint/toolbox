package dbx_auth_attr

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
)

func NewConsole(c app_control.Control, peerName string) api_auth.Console {
	l := c.Log().With(zap.String("peerName", peerName))
	var oa api_auth.Console

	// Make redirect impl. hidden for while
	if f, found := c.Feature().OptInGet(&api_auth_impl.FeatureRedirect{}); found && f.OptInIsEnabled() {
		oa = api_auth_impl.NewConsoleRedirect(c, peerName, dbx_auth.NewApp(c))
	} else {
		oa = api_auth_impl.NewConsoleOAuth(c, peerName)
	}
	aa := NewConsoleAttr(c, oa)
	if c.Feature().IsSecure() {
		l.Debug("Skip caching")
		return aa
	}
	l.Debug("Token cache enabled")
	ca := api_auth_impl.NewConsoleCache(c, aa)
	return ca
}

func NewConsoleAttr(c app_control.Control, auth api_auth.Console) api_auth.Console {
	return &Attr{
		ctl:  c,
		auth: auth,
	}
}
