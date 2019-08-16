package api_auth_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_zap"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func NewApp(kitchen app_kitchen.Kitchen) api_auth.App {
	a := &App{
		kitchen: kitchen,
		keys:    make(map[string]string),
	}
	a.loadKeys()
	return a
}

type App struct {
	kitchen  app_kitchen.Kitchen
	keys map[string]string
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
	kb, err := app_zap.Unzap(z.kitchen.Control())
	if err != nil {
		kb, err = z.kitchen.Control().Resource("toolbox.appkeys")
		if err != nil {
			z.kitchen.Log().Debug("Skip loading app keys")
			return
		}
	}
	err = json.Unmarshal(kb, &z.keys)
	if err != nil {
		z.kitchen.Log().Debug("Skip loading app keys: unable to unmarshal resource", zap.Error(err))
		return
	}
}
