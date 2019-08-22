package api_auth_impl

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"strings"
)

type KitchenAuth struct {
	control  app_control.Control
	app      api_auth.App
	peerName string
}

func (z *KitchenAuth) Auth(tokenType string) (ctx api_context.Context, err error) {
	if z.control.IsTest() {
		return nil, errors.New("test mode")
	}

	key, secret := z.app.AppKey(tokenType)
	if key == "" || secret == "" {
		t, err := z.authGenerated(tokenType)
		return z.wrapToken(tokenType, t, err)
	} else {
		t, err := z.oauthStart(tokenType)
		return z.wrapToken(tokenType, t, err)
	}
}

func (z *KitchenAuth) wrapToken(tokenType, token string, cause error) (ctx api_context.Context, err error) {
	if cause != nil {
		return nil, err
	}
	tc := api_auth.TokenContainer{
		Token:     token,
		TokenType: tokenType,
		PeerName:  z.peerName,
	}
	ctx = api_context_impl.NewKC(z.control, tc)

	_, _, err = VerifyToken(tokenType, ctx)
	if err != nil {
		z.control.Log().Debug("failed verify token", zap.Error(err))
		z.control.UI().Error("auth.basic.verify.failed")
		return nil, err
	}
	return ctx, nil
}

func (z *KitchenAuth) init() {
	z.app = NewApp(z.control)
}

func (z *KitchenAuth) generatedTokenInstruction(tokenType string) {
	api := ""
	toa := ""

	switch tokenType {
	case api_auth.DropboxTokenFull:
		api = "Dropbox API"
		toa = "Full Dropbox"
	case api_auth.DropboxTokenApp:
		api = "Dropbox API"
		toa = "App folder"
	case api_auth.DropboxTokenBusinessInfo:
		api = "Dropbox Business API"
		toa = "Team information"
	case api_auth.DropboxTokenBusinessAudit:
		api = "Dropbox Business API"
		toa = "Team auditing"
	case api_auth.DropboxTokenBusinessFile:
		api = "Dropbox Business API"
		toa = "Team member file access"
	case api_auth.DropboxTokenBusinessManagement:
		api = "Dropbox Business API"
		toa = "Team member management"
	default:
		z.control.Log().Fatal("Undefined token type", zap.String("type", tokenType))
	}

	z.control.UI().Info(
		"auth.basic.generated_token1",
		app_msg.P("API", api),
		app_msg.P("TypeOfAccess", toa),
	)
}

func (z *KitchenAuth) generatedToken(tokenType string) (string, error) {
	z.generatedTokenInstruction(tokenType)
	for {
		code, cancel := z.control.UI().AskSecure("auth.basic.generated_token2")
		if cancel {
			return "", errors.New("user cancelled")
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim, nil
		}
	}
}

func (z *KitchenAuth) authGenerated(tokenType string) (string, error) {
	z.control.Log().Debug("No appKey/appSecret found. Try asking 'Generate Token'")
	tok, err := z.generatedToken(tokenType)
	return tok, err
}

func (z *KitchenAuth) oauthStart(tokenType string) (string, error) {
	l := z.control.Log()
	l.Debug("Start OAuth sequence")
	state, err := sc_random.GenerateRandomString(8)
	if err != nil {
		l.Error("Unable to generate `state`", zap.Error(err))
		return "", err
	}

	tok, err := z.oauthAskCode(tokenType, state)
	if err != nil {
		l.Error("Authentication failed due to the error", zap.Error(err))
		return "", err
	}
	return tok.AccessToken, nil
}

func (z *KitchenAuth) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *KitchenAuth) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *KitchenAuth) oauthCode(state string) string {
	for {
		code, cancel := z.control.UI().AskSecure("auth.basic.oauth_seq2")
		if cancel {
			return ""
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (z *KitchenAuth) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
	cfg := z.app.Config(tokenType)
	url := z.oauthUrl(cfg, state)

	z.control.UI().Info("auth.basic.oauth_seq1", app_msg.P("Url", url))

	code := z.oauthCode(state)

	return z.oauthExchange(cfg, code)
}
