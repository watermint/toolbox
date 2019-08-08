package api_auth_impl

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/app/app_zap"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_context_impl"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"strings"
)

// ExecContext auth context
type ECAuth struct {
	ec       *app.ExecContext
	peerName string
	keys     map[string]string
}

func (z *ECAuth) Auth(tokenType string) (ctx api_context.Context, err error) {
	if z.ec.IsTest() {
		return nil, errors.New("test mode")
	}

	key, secret := z.appKeys(tokenType)
	if key == "" || secret == "" {
		t, err := z.authGenerated(tokenType)
		return z.wrapToken(tokenType, t, err)
	} else {
		t, err := z.oauthStart(tokenType)
		return z.wrapToken(tokenType, t, err)
	}
}

func (z *ECAuth) appKeys(tokenType string) (key, secret string) {
	var e bool
	if key, e = z.keys[tokenType+".key"]; !e {
		return "", ""
	}
	if secret, e = z.keys[tokenType+".secret"]; !e {
		return "", ""
	}
	return
}

func (z *ECAuth) verifyToken(tokenType string, ctx api_context.Context) error {
	switch tokenType {
	case api_auth.DropboxTokenFull, api_auth.DropboxTokenApp:
		_, err := ctx.Request("users/get_current_account").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return err
		}
		ctx.Log().Debug("Token Verified")

		return nil

	case api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessAudit:
		_, err := ctx.Request("team/token/get_authenticated_admin").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return err
		}
		ctx.Log().Debug("Token Verified")

		return nil

	default:
		return nil
	}
}

func (z *ECAuth) wrapToken(tokenType, token string, cause error) (ctx api_context.Context, err error) {
	if err != nil {
		return nil, err
	}
	tc := api_auth.TokenContainer{
		Token:     token,
		TokenType: tokenType,
		PeerName:  z.peerName,
	}
	ctx = api_context_impl.New(z.ec, tc)

	err = z.verifyToken(tokenType, ctx)
	if err != nil {
		z.ec.Log().Debug("failed verify token", zap.Error(err))
		z.ec.Msg("auth.basic.verify.failed").TellError()
		return nil, err
	}
	return ctx, nil
}

func (z *ECAuth) loadKeys() {
	kb, err := app_zap.Unzap(z.ec)
	if err != nil {
		kb, err = z.ec.ResourceBytes("toolbox.appkeys")
		if err != nil {
			z.ec.Log().Debug("Skip loading app keys")
		}
		return
	}
	err = json.Unmarshal(kb, &z.keys)
	if err != nil {
		z.ec.Log().Debug("Skip loading app keys: unable to unmarshal resource", zap.Error(err))
		return
	}
}

func (z *ECAuth) init() {
	z.keys = make(map[string]string)
	z.loadKeys()
}

func (z *ECAuth) generatedTokenInstruction(tokenType string) {
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
		z.ec.Log().Fatal("Undefined token type", zap.String("type", tokenType))
	}

	z.ec.Msg("auth.basic.generated_token1").WithData(struct {
		API          string
		TypeOfAccess string
	}{
		API:          api,
		TypeOfAccess: toa,
	}).Tell()
}

func (z *ECAuth) generatedToken(tokenType string) (string, error) {
	z.generatedTokenInstruction(tokenType)
	for {
		code := z.ec.Msg("auth.basic.generated_token2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim, nil
		}
	}
}

func (z *ECAuth) authGenerated(tokenType string) (string, error) {
	z.ec.Log().Debug(
		"No appKey/appSecret found. Try asking 'Generate Token'",
	)
	tok, err := z.generatedToken(tokenType)
	return tok, err
}

func (z *ECAuth) oauthStart(tokenType string) (string, error) {
	log := z.ec.Log()
	log.Debug("Start OAuth sequence")
	state, err := app_util.GenerateRandomString(8)
	if err != nil {
		log.Error("Unable to generate `state`",
			zap.Error(err),
		)
		return "", err
	}

	tok, err := z.oauthAskCode(tokenType, state)
	if err != nil {
		log.Error("Authentication failed due to the error",
			zap.Error(err),
		)
		return "", err
	}
	return tok.AccessToken, nil
}

func (z *ECAuth) oauthEndpoint() *oauth2.Endpoint {
	return &oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}

func (z *ECAuth) oauthConfig(tokenType string) *oauth2.Config {
	key, secret := z.appKeys(tokenType)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     *z.oauthEndpoint(),
	}
}

func (z *ECAuth) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *ECAuth) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *ECAuth) oauthCode(state string) string {
	for {
		code := z.ec.Msg("auth.basic.oauth_seq2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (z *ECAuth) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
	cfg := z.oauthConfig(tokenType)
	url := z.oauthUrl(cfg, state)

	z.ec.Msg("auth.basic.oauth_seq1").WithData(struct {
		Url string
	}{
		Url: url,
	}).Tell()

	code := z.oauthCode(state)

	return z.oauthExchange(cfg, code)
}
