package dbx_auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/app/util"
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

type DropboxAuthenticator struct {
	AuthFile  string
	AppKey    string
	AppSecret string
	TokenType string
	Logger    *zap.Logger
}

const (
	authPromptMessage1 = `=================================================
1. Visit the URL for the auth dialog:

%s

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
`
	authPromptMessage2 = `
Enter the authorisation code here: `

	authGeneratedToken1Tmpl = `========================================================
1. Visit the MyApp page (you mihgt have to login first):

https://www.dropbox.com/developers/apps

2. Proceed with "Create App"
3. Choose "{{.API}}"
4. Choose "{{.TypeOfAccess}}"
5. Enter name of your app
6. Proceed with "Create App"
7. Hit "Generate" button near "Generated access token"
8. Copy generated token:
`
	authGeneratedToken2 = `
Enter the generated token here:
`
)

func (d *DropboxAuthenticator) Log() *zap.Logger {
	return d.Logger
}

func (d *DropboxAuthenticator) generateTokenInstruction() error {
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

	data := struct {
		API          string
		TypeOfAccess string
	}{
		API:          api,
		TypeOfAccess: toa,
	}
	instr, err := util.CompileTemplate(authGeneratedToken1Tmpl, data)
	if err != nil {
		d.Log().Fatal(
			"Unable to compile template",
			zap.String("tmpl", authGeneratedToken1Tmpl),
			zap.Error(err),
		)

		return errors.New("unable to generate instruction")
	}
	fmt.Println(instr)
	return nil
}

func (d *DropboxAuthenticator) TokenFileLoadMap() (map[string]string, error) {
	log := d.Log().With(
		zap.String("file", d.AuthFile),
	)
	log.Debug(
		"Loading token from file",
	)
	f, err := ioutil.ReadFile(d.AuthFile)
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
	if t, ok := m[d.AppKey]; ok {
		d.Log().Debug(
			"Token for app key found in map",
			zap.String("appKey", d.AppKey),
		)
		return t, nil
	}
	if t, ok := m[d.TokenType]; ok {
		d.Log().Debug(
			"Token for token type dound in map",
			zap.String("tokenType", d.TokenType),
		)
		return t, nil
	}

	d.Log().Debug(
		"Token not found in loaded token map",
		zap.String("appKey", d.AppKey),
	)
	return "", errors.New("app key not found in loaded token map")
}

func (d *DropboxAuthenticator) TokenFileSave(token string) error {
	log := d.Log().With(zap.String("file", d.AuthFile))
	log.Info("Saving token to the file")

	// TODO: check file exist or not
	m, err := d.TokenFileLoadMap()
	if err != nil {
		m = make(map[string]string)
	}

	if d.AppKey == "" {
		// save as token type string
		m[d.TokenType] = token
	} else {
		// overwrite token for AppKey
		m[d.AppKey] = token
	}

	f, err := json.Marshal(m)
	if err != nil {
		log.Error(
			"Unable to marshal auth tokens. Failed to save token to the file",
			zap.Error(err),
		)
		return err
	}

	err = ioutil.WriteFile(d.AuthFile, f, 0600)

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
	if d.AppKey == "" || d.AppSecret == "" {
		d.Log().Debug(
			"No AppKey/AppSecret found. Try asking 'Generate Token'",
		)
		tok, err := d.acquireToken()
		if err == nil {
			d.TokenFileSave(tok)
		}
		return tok, err
	} else {
		log := d.Log().With(
			zap.String("appKey", d.AppKey),
		)
		log.Debug(
			"Start auth sequence for AppKey",
		)
		state, err := util.GenerateRandomString(8)
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

	var code string

	for {
		fmt.Println(authGeneratedToken2)
		if _, err := fmt.Scan(&code); err != nil {
			d.Log().Error(
				"Input error (%s), try again.",
				zap.Error(err),
			)
			continue
		}
		trim := strings.TrimSpace(code)
		if len(trim) < 1 {
			d.Log().Error("Input error, try again.")
			continue
		}
		return trim, nil
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
		ClientID:     d.AppKey,
		ClientSecret: d.AppSecret,
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
	var code string

	fmt.Println(authPromptMessage2)

	if _, err := fmt.Scan(&code); err != nil {
		fmt.Errorf("%s\n", err)
		return ""
	}
	return code
}

func (d *DropboxAuthenticator) auth(state string) (*oauth2.Token, error) {
	cfg := d.authConfig()
	url := d.authUrl(cfg, state)

	fmt.Printf(authPromptMessage1, url)

	code := d.codeDialogue(state)

	return d.authExchange(cfg, code)
}
