package oper_cli

import (
	"errors"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/poc/oper"
	"github.com/watermint/toolbox/poc/oper/oper_api"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type BasicAuthenticator struct {
	Ctx oper.Context
}

func (z *BasicAuthenticator) Auth(t oper_api.DropboxApiToken) (oper_api.DropboxApiToken, error) {
	if t.ApiKey() == "" {
		return z.generatedToken(t)
	} else {
		return z.oauthSeqStart(t)
	}
}

func (z *BasicAuthenticator) generatedToken(t oper_api.DropboxApiToken) (rt oper_api.DropboxApiToken, err error) {
	z.Ctx.UI().Tell(
		z.Ctx.Message("auth.basic.generated_token1").WithTmplData(struct {
			API          string
			TypeOfAccess string
		}{
			API:          t.AppTypeLabel(),
			TypeOfAccess: t.AppAccessLabel(),
		}),
	)

	token := z.Ctx.UI().AskText(z.Ctx.Message("auth.basic.generated_token2"))
	if token == "" {
		z.Ctx.UI().TellFailure(z.Ctx.Message("auth.basic.user_cancelled"))
		return nil, errors.New("user cancelled operation")
	} else {
		return t.WithApi(dbx_api.NewContext(token, z.Ctx.Log())), nil
	}
}

func (z *BasicAuthenticator) oauthSeqStart(t oper_api.DropboxApiToken) (rt oper_api.DropboxApiToken, err error) {
	state, err := app_util.GenerateRandomString(8)
	if err != nil {
		z.Ctx.Log().Error("unable to generate `state`", zap.Error(err))
		return nil, err
	}
	cfg := &oauth2.Config{
		ClientID:     t.ApiKey(),
		ClientSecret: t.ApiSecret(),
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
	}
	url := cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)

	z.Ctx.UI().Tell(
		z.Ctx.Message("auth.basic.oauth_seq1").WithTmplData(struct {
			Url string
		}{
			Url: url,
		}),
	)

	code := z.Ctx.UI().AskText(z.Ctx.Message("auth.basic.oauth_seq2"))
	if code == "" {
		z.Ctx.UI().TellFailure(z.Ctx.Message("auth.basic.user_cancelled"))
		return nil, errors.New("user cancelled operation")
	}

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		z.Ctx.UI().TellError(z.Ctx.Message("auth.basic.failed"))
		z.Ctx.Log().Error("Authentication failed", zap.Error(err))
		return nil, err
	}
	return t.WithApi(dbx_api.NewContext(token.AccessToken, z.Ctx.Log())), nil
}
