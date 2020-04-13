package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
)

func DropboxOAuthEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z *App) Config(tokenType string) *oauth2.Config {
	key, secret := z.AppKey(tokenType)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     DropboxOAuthEndpoint(),
	}
}

func (z *App) AppKey(tokenType string) (key, secret string) {
	return z.res.Key(tokenType)
}

func NewApp(ctl app_control.Control) api_auth.App {
	a := &App{
		ctl: ctl,
		res: api_appkey.New(ctl),
	}
	return a
}
