package dbx_auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/app/app_util"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"io/ioutil"
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

func NewDropboxAuthenticator(
	authFile string,
	appKey string,
	appSecret string,
	tokenType string,
	msg *app_ui.UIMessageContainer,
	logger *zap.Logger,
) *DropboxAuthenticator {
	return &DropboxAuthenticator{
		authFile:  authFile,
		appKey:    appKey,
		appSecret: appSecret,
		TokenType: tokenType,
		mc:        msg,
		logger:    logger,
	}
}

type DropboxAuthenticator struct {
	authFile  string
	appKey    string
	appSecret string
	TokenType string
	mc        *app_ui.UIMessageContainer
	logger    *zap.Logger
}

func (d *DropboxAuthenticator) Log() *zap.Logger {
	return d.logger
}

func (d *DropboxAuthenticator) generateTokenInstruction() {
	api := ""
	toa := ""

	if d.TokenType == DropboxTokenFull {
		api = "Dropbox API"
		toa = "Full Dropbox"
	} else if d.TokenType == DropboxTokenApp {
		api = "Dropbox API"
		toa = "App folder"
	} else if d.TokenType == DropboxTokenBusinessInfo {
		api = "Dropbox Business API"
		toa = "Team information"
	} else if d.TokenType == DropboxTokenBusinessAudit {
		api = "Dropbox Business API"
		toa = "Team auditing"
	} else if d.TokenType == DropboxTokenBusinessFile {
		api = "Dropbox Business API"
		toa = "Team member file access"
	} else if d.TokenType == DropboxTokenBusinessManagement {
		api = "Dropbox Business API"
		toa = "Team member management"
	} else {
		d.Log().Fatal(
			"Undefined token type",
			zap.String("type", d.TokenType),
		)
	}

	d.mc.Msg("auth.basic.generated_token1").WithData(struct {
		API          string
		TypeOfAccess string
	}{
		API:          api,
		TypeOfAccess: toa,
	}).Tell()
}

func (d *DropboxAuthenticator) TokenFileLoadMap() (map[string]string, error) {
	log := d.Log().With(
		zap.String("file", d.authFile),
	)
	log.Debug(
		"Loading token from file",
	)
	f, err := ioutil.ReadFile(d.authFile)
	if err != nil {
		log.Debug(
			"Unable to read file",
			zap.Error(err),
		)
		return nil, err
	}
	m := make(map[string]string)
	err = json.Unmarshal(f, &m)
	if err != nil {
		log.Debug(
			"Unable to unmarshal file data",
			zap.Error(err),
		)
		return nil, err
	}
	return m, nil
}

func (d *DropboxAuthenticator) TokenFileLoad() (string, error) {
	m, err := d.TokenFileLoadMap()
	if err != nil {
		return "", err
	}
	if t, ok := m[d.appKey]; ok {
		d.Log().Debug(
			"Token for app key found in map",
			zap.String("appKey", d.appKey),
		)
		return t, nil
	}
	if t, ok := m[d.TokenType]; ok {
		d.Log().Debug(
			"Token for token type dound in map",
			zap.String("TokenType", d.TokenType),
		)
		return t, nil
	}

	d.Log().Debug(
		"Token not found in loaded token map",
		zap.String("appKey", d.appKey),
	)
	return "", errors.New("app key not found in loaded token map")
}

func (d *DropboxAuthenticator) TokenFileSave(token string) error {
	log := d.Log().With(zap.String("file", d.authFile))
	log.Info("Saving token to the file")

	// TODO: check file exist or not
	m, err := d.TokenFileLoadMap()
	if err != nil {
		m = make(map[string]string)
	}

	if d.appKey == "" {
		// save as token type string
		m[d.TokenType] = token
	} else {
		// overwrite token for appKey
		m[d.appKey] = token
	}

	f, err := json.Marshal(m)
	if err != nil {
		log.Error(
			"Unable to marshal auth tokens. Failed to save token to the file",
			zap.Error(err),
		)
		return err
	}

	err = ioutil.WriteFile(d.authFile, f, 0600)

	if err != nil {
		log.Error(
			"Unable to write authentication token to file",
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (d *DropboxAuthenticator) LoadOrAuth(business bool) (string, error) {
	t, err := d.TokenFileLoad()
	if err != nil {
		return d.Authorise()
	}

	return t, nil
}

func (d *DropboxAuthenticator) Authorise() (string, error) {
	if d.appKey == "" || d.appSecret == "" {
		d.Log().Debug(
			"No appKey/appSecret found. Try asking 'Generate Token'",
		)
		tok, err := d.acquireToken()
		if err == nil {
			d.TokenFileSave(tok)
		}
		return tok, err
	} else {
		log := d.Log().With(
			zap.String("appKey", d.appKey),
		)
		log.Debug(
			"Start auth sequence for appKey",
		)
		state, err := app_util.GenerateRandomString(8)
		if err != nil {
			log.Error("Unable to generate `state`",
				zap.Error(err),
			)
			return "", err
		}

		tok, err := d.auth(state)
		if err != nil {
			log.Error("Authentication failed due to the error",
				zap.Error(err),
			)
			return "", err
		}
		d.TokenFileSave(tok.AccessToken)
		return tok.AccessToken, nil
	}
}

func (d *DropboxAuthenticator) acquireToken() (string, error) {
	d.generateTokenInstruction()
	for {
		code := d.mc.Msg("auth.basic.generated_token2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim, nil
		}
	}
}

func (d *DropboxAuthenticator) authEndpoint() *oauth2.Endpoint {
	return &oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}

func (d *DropboxAuthenticator) authConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     d.appKey,
		ClientSecret: d.appSecret,
		Scopes:       []string{},
		Endpoint:     *d.authEndpoint(),
	}
}

func (d *DropboxAuthenticator) authUrl(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
	)
}

func (d *DropboxAuthenticator) authExchange(cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}

func (d *DropboxAuthenticator) codeDialogue(state string) string {
	for {
		code := d.mc.Msg("auth.basic.oauth_seq2").AskText()
		trim := strings.TrimSpace(code)
		if len(trim) > 0 {
			return trim
		}
	}
}

func (d *DropboxAuthenticator) auth(state string) (*oauth2.Token, error) {
	cfg := d.authConfig()
	url := d.authUrl(cfg, state)

	d.mc.Msg("auth.basic.oauth_seq1").WithData(struct {
		Url string
	}{
		Url: url,
	}).Tell()

	code := d.codeDialogue(state)

	return d.authExchange(cfg, code)
}
