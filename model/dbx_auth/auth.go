package dbx_auth

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"strings"
)

type DropboxTokenType int

const (
	DropboxTokenFull               = "user_full"
	DropboxTokenApp                = "user_app"
	DropboxTokenBusinessInfo       = "business_info"
	DropboxTokenBusinessAudit      = "business_audit"
	DropboxTokenBusinessFile       = "business_file"
	DropboxTokenBusinessManagement = "business_management"
)

func NewDefaultAuth(ec *app.ExecContext) Authenticator {
	return NewAuth(ec, "default")
}

func NewAuth(ec *app.ExecContext, peerName string) Authenticator {
	ua := &UIAuthenticator{
		ec: ec,
	}
	ua.init()
	ca := &CachedAuthenticator{
		peerName: peerName,
		ec:       ec,
		auth:     ua,
	}
	ca.init()
	return ca
}

func IsCacheAvailable(ec *app.ExecContext, peerName string) bool {
	ca := &CachedAuthenticator{
		peerName: peerName,
		ec:       ec,
	}
	return ca.loadResource() != nil
}

type Authenticator interface {
	Auth(tokenType string) (*dbx_api.Context, error)
}

type CachedAuthenticator struct {
	peerName string
	tokens   map[string]string
	ec       *app.ExecContext
	auth     Authenticator
}

func (z *CachedAuthenticator) init() {
	z.tokens = make(map[string]string)

	if z.loadFile() == nil {
		return // return on success
	}
	if z.loadResource() == nil {
		return // return on success
	}
}

func (z *CachedAuthenticator) cacheFile() string {
	px := sha256.Sum224([]byte(z.peerName))
	pn := fmt.Sprintf("%x.tokens", px)
	return z.ec.FileOnWorkPath(pn)
}

func (z *CachedAuthenticator) loadFile() error {
	tf := z.cacheFile()
	_, err := os.Stat(tf)
	if os.IsNotExist(err) {
		z.ec.Log().Debug("token file not found", zap.String("path", tf))
		return err
	}
	tb, err := ioutil.ReadFile(tf)
	if err != nil {
		z.ec.Log().Debug("unable to read tokens file", zap.String("path", tf), zap.Error(err))
		return err
	}
	err = json.Unmarshal(tb, &z.tokens)
	if err != nil {
		z.ec.Log().Debug("unable to unmarshal tokens file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CachedAuthenticator) loadResource() error {
	tb, err := z.ec.ResourceBytes(z.peerName + ".tokens")
	if err != nil {
		z.ec.Log().Debug("unable to load tokens file", zap.Error(err))
		return err
	}

	err = json.Unmarshal(tb, &z.tokens)
	if err != nil {
		z.ec.Log().Debug("unable to unmarshal tokens file", zap.Error(err))
		return err
	}
	return nil
}

func (z *CachedAuthenticator) updateCache(tokenType, token string) {
	z.tokens[tokenType] = token
	tb, err := json.Marshal(z.tokens)
	if err != nil {
		z.ec.Log().Debug("unable to marshal tokens", zap.Error(err))
		return
	}
	tf := z.cacheFile()
	err = ioutil.WriteFile(tf, tb, 0600)
	if err != nil {
		z.ec.Log().Debug("unable to write tokens into file", zap.Error(err))
		return
	}
}

func (z *CachedAuthenticator) Auth(tokenType string) (*dbx_api.Context, error) {
	if t, e := z.tokens[tokenType]; e {
		return dbx_api.NewContext(
			z.ec,
			tokenType,
			t,
		), nil
	}

	if t, err := z.auth.Auth(tokenType); err != nil {
		return nil, err
	} else {
		z.updateCache(tokenType, t.Token)
		return t, nil
	}
}

type UIAuthenticator struct {
	ec   *app.ExecContext
	keys map[string]string
}

func (z *UIAuthenticator) appKeys(tokenType string) (key, secret string) {
	var e bool
	if key, e = z.keys[tokenType+".key"]; !e {
		return "", ""
	}
	if secret, e = z.keys[tokenType+".secret"]; !e {
		return "", ""
	}
	return
}

func (z *UIAuthenticator) wrapToken(tokenType, token string, err error) (*dbx_api.Context, error) {
	if err != nil {
		return nil, err
	}
	return dbx_api.NewContext(
		z.ec,
		tokenType,
		token,
	), nil
}

func (z *UIAuthenticator) Auth(tokenType string) (*dbx_api.Context, error) {
	key, secret := z.appKeys(tokenType)
	if key == "" || secret == "" {
		t, err := z.authGenerated(tokenType)
		return z.wrapToken(tokenType, t, err)
	} else {
		t, err := z.oauthStart(tokenType)
		return z.wrapToken(tokenType, t, err)
	}
}

func (z *UIAuthenticator) loadKeys() {
	kb, err := z.ec.ResourceBytes("toolbox.appkeys")
	if err != nil {
		z.ec.Log().Debug("unable to load resource `toolbox.appkeys`", zap.Error(err))
		return
	}
	err = json.Unmarshal(kb, &z.keys)
	if err != nil {
		z.ec.Log().Debug("unable to unmarshal resource `toolbox.appkeys`", zap.Error(err))
		return
	}
}

func (z *UIAuthenticator) init() {
	z.keys = make(map[string]string)
	z.loadKeys()
}

func (z *UIAuthenticator) generatedTokenInstruction(tokenType string) {
	api := ""
	toa := ""

	switch tokenType {
	case DropboxTokenFull:
		api = "Dropbox API"
		toa = "Full Dropbox"
	case DropboxTokenApp:
		api = "Dropbox API"
		toa = "App folder"
	case DropboxTokenBusinessInfo:
		api = "Dropbox Business API"
		toa = "Team information"
	case DropboxTokenBusinessAudit:
		api = "Dropbox Business API"
		toa = "Team auditing"
	case DropboxTokenBusinessFile:
		api = "Dropbox Business API"
		toa = "Team member file access"
	case DropboxTokenBusinessManagement:
		api = "Dropbox Business API"
		toa = "Team member management"
	default:
		z.ec.Log().Fatal(
			"Undefined token type",
			zap.String("type", tokenType),
		)
	}

	z.ec.Msg("auth.basic.generated_token1").WithData(struct {
		API          string
		TypeOfAccess string
	}{
		API:          api,
		TypeOfAccess: toa,
	}).Tell()
}

func (z *UIAuthenticator) generatedToken(tokenType string) (string, error) {
	z.generatedTokenInstruction(tokenType)
	for {
		code := z.ec.Msg("auth.basic.generated_token2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim, nil
		}
	}
}

func (z *UIAuthenticator) authGenerated(tokenType string) (string, error) {
	z.ec.Log().Debug(
		"No appKey/appSecret found. Try asking 'Generate Token'",
	)
	tok, err := z.generatedToken(tokenType)
	return tok, err
}

func (z *UIAuthenticator) oauthStart(tokenType string) (string, error) {
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

func (z *UIAuthenticator) oauthEndpoint() *oauth2.Endpoint {
	return &oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}

func (z *UIAuthenticator) oauthConfig(tokenType string) *oauth2.Config {
	key, secret := z.appKeys(tokenType)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     *z.oauthEndpoint(),
	}
}

func (z *UIAuthenticator) oauthUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (z *UIAuthenticator) oauthExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (z *UIAuthenticator) oauthCode(state string) string {
	for {
		code := z.ec.Msg("auth.basic.oauth_seq2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (z *UIAuthenticator) oauthAskCode(tokenType, state string) (*oauth2.Token, error) {
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
