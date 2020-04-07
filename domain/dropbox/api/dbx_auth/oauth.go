package dbx_auth

import (
	"context"
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"strings"
)

func NewApp(control app_control.Control) api_auth.App {
	a := &App{
		control: control,
		keys:    make(map[string]string),
	}
	a.loadKeys()
	return a
}

func NewConsoleOAuth(c app_control.Control, peerName string) api_auth.Console {
	return &OAuth{
		ctl:      c,
		app:      NewApp(c),
		peerName: peerName,
	}
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

type OAuth struct {
	ctl      app_control.Control
	app      api_auth.App
	peerName string
}

func (z *OAuth) PeerName() string {
	return z.peerName
}

func (z *OAuth) Auth(scope string) (tc api_auth.Context, err error) {
	l := z.ctl.Log().With(zap.String("peerName", z.peerName), zap.String("scope", scope))
	ui := z.ctl.UI()

	l.Debug("Start OAuth sequence")
	t, err := z.oauthStart(scope)
	if err != nil {
		l.Debug("Authentication finished with an error", zap.Error(err))
		ui.Error(MCcAuth.FailedOrCancelled.With("Cause", err))
		return nil, err
	}
	return api_auth.NewContext(t, z.peerName, scope), nil
}

func (z *OAuth) oauthStart(scope string) (*oauth2.Token, error) {
	l := z.ctl.Log()
	l.Debug("Start OAuth sequence")
	state, err := sc_random.GenerateRandomString(8)
	if err != nil {
		l.Error("Unable to generate `state`", zap.Error(err))
		return nil, err
	}

	tok, err := z.oauthAskCode(scope, state)
	if err != nil {
		l.Debug("Authentication failed due to the error", zap.Error(err))
		return nil, err
	}
	return tok, nil
}

func (z *OAuth) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *OAuth) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *OAuth) oauthCode(state string) string {
	ui := z.ctl.UI()
	for {
		code, cancel := ui.AskSecure(MCcAuth.OauthSeq2)
		if cancel {
			return ""
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (z *OAuth) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
	ui := z.ctl.UI()
	cfg := z.app.Config(tokenType)
	url := z.oauthUrl(cfg, state)

	ui.Info(MCcAuth.OauthSeq1.With("Url", url))

	code := z.oauthCode(state)
	if code == "" {
		return nil, ErrorUserCancelled
	}

	return z.oauthExchange(cfg, code)
}
