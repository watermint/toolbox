package dbx_auth

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/dbx_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"strings"
)

type CcAuth struct {
	control  app_control.Control
	app      api_auth.App
	peerName string
}

func (z *CcAuth) Auth(tokenType string) (ctx api_context.Context, err error) {
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

func (z *CcAuth) wrapToken(tokenType, token string, cause error) (ctx api_context.Context, err error) {
	ui := z.control.UI()
	if cause != nil {
		ui.Error(MCcAuth.FailedOrCancelled.With("Cause", cause.Error()))
		return nil, cause
	}
	tc := api_auth.TokenContainer{
		Token:     token,
		TokenType: tokenType,
		PeerName:  z.peerName,
	}
	ctx = dbx_context.New(z.control, tc)

	_, _, err = VerifyToken(tokenType, ctx)
	if err != nil {
		z.control.Log().Debug("failed verify token", zap.Error(err))
		ui.Error(MCcAuth.VerifyFailed)
		return nil, err
	}
	return ctx, nil
}

func (z *CcAuth) init() {
	z.app = NewApp(z.control)
}

func (z *CcAuth) generatedTokenInstruction(tokenType string) {
	ui := z.control.UI()
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

	ui.Info(MCcAuth.GeneratedToken1.With("API", api).With("TypeOfAccess", toa))
}

func (z *CcAuth) generatedToken(tokenType string) (string, error) {
	ui := z.control.UI()
	z.generatedTokenInstruction(tokenType)
	for {
		code, cancel := ui.AskSecure(MCcAuth.GeneratedToken2)
		if cancel {
			return "", errors.New("user cancelled")
		}
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim, nil
		}
	}
}

func (z *CcAuth) authGenerated(tokenType string) (string, error) {
	z.control.Log().Debug("No appKey/appSecret found. Try asking 'Generate Token'")
	tok, err := z.generatedToken(tokenType)
	return tok, err
}

func (z *CcAuth) oauthStart(tokenType string) (string, error) {
	l := z.control.Log()
	l.Debug("Start OAuth sequence")
	state, err := sc_random.GenerateRandomString(8)
	if err != nil {
		l.Error("Unable to generate `state`", zap.Error(err))
		return "", err
	}

	tok, err := z.oauthAskCode(tokenType, state)
	if err != nil {
		l.Debug("Authentication failed due to the error", zap.Error(err))
		return "", err
	}
	return tok.AccessToken, nil
}

func (z *CcAuth) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *CcAuth) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *CcAuth) oauthCode(state string) string {
	ui := z.control.UI()
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

func (z *CcAuth) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
	ui := z.control.UI()
	cfg := z.app.Config(tokenType)
	url := z.oauthUrl(cfg, state)

	ui.Info(MCcAuth.OauthSeq1.With("Url", url))

	code := z.oauthCode(state)
	if code == "" {
		return nil, errors.New("user might cancelled auth sequence, or quiet mode (require pre-authentication)")
	}

	return z.oauthExchange(cfg, code)
}