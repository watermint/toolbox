package api_auth_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func NewApp(control app_control.Control) api_auth.App {
	a := &App{
		control: control,
		keys:    make(map[string]string),
	}
	a.loadKeys()
	return a
}

type App struct {
	control app_control.Control
	keys    map[string]string
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
	var e bool
	if key, e = z.keys[tokenType+".key"]; !e {
		return "", ""
	}
	if secret, e = z.keys[tokenType+".secret"]; !e {
		return "", ""
	}
	return
}

func (z *App) loadKeys() {
	kb, err := sc_zap.Unzap(z.control)
	if err != nil {
		kb, err = z.control.Resource("toolbox.appkeys")
		if err != nil {
			z.control.Log().Debug("Skip loading app keys")
			return
		}
	}
	err = json.Unmarshal(kb, &z.keys)
	if err != nil {
		z.control.Log().Debug("Skip loading app keys: unable to unmarshal resource", zap.Error(err))
		return
	}
}
