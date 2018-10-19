package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
	"golang.org/x/oauth2"
	"io/ioutil"
	"strings"
)

type DropboxTokenType int

const (
	DropboxTokenFull DropboxTokenType = iota
	DropboxTokenApp
	DropboxTokenBusinessInfo
	DropboxTokenBusinessAudit
	DropboxTokenBusinessFile
	DropboxTokenBusinessManagement
)

type DropboxAuthenticator struct {
	AuthFile  string
	AppKey    string
	AppSecret string
	TokenType DropboxTokenType
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
		seelog.Errorf("Undefined token type: %d", d.TokenType)
		return errors.New(fmt.Sprintf("Undefined token type"))
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
		seelog.Debugf("Unable to compile template: [[[%s]]]", authGeneratedToken1Tmpl)
		return errors.New("unable to generate instruction")
	}
	fmt.Println(instr)
	return nil
}

func (d *DropboxAuthenticator) TokenFileLoadMap() (map[string]string, error) {
	seelog.Tracef("Loading token from file: [%s]", d.AuthFile)
	f, err := ioutil.ReadFile(d.AuthFile)
	if err != nil {
		seelog.Tracef("Unable to load file: file[%s], err[%s]", d.AuthFile, err)
		return nil, err
	}
	m := make(map[string]string)
	err = json.Unmarshal(f, &m)
	if err != nil {
		seelog.Tracef("Unable to unmarshal: error[%v]", err)
		return nil, err
	}
	return m, nil
}

func (d *DropboxAuthenticator) TokenFileLoad() (string, error) {
	m, err := d.TokenFileLoadMap()
	if err != nil {
		return "", err
	}
	if t, ex := m[d.AppKey]; !ex {
		seelog.Tracef("Appkey[%s] not found in loaded token map", d.AppKey)
		return "", errors.New("app key not found in loaded token map")
	} else {
		seelog.Tracef("Token for App key[%s] found in map", d.AppKey)
		return t, nil
	}
}

func (d *DropboxAuthenticator) TokenFileSave(token string) error {
	seelog.Infof("Saving token to the file: [%s]", d.AuthFile)

	// TODO: check file exist or not
	m, err := d.TokenFileLoadMap()
	if err != nil {
		m = make(map[string]string)
	}

	// overwrite token for AppKey
	m[d.AppKey] = token

	f, err := json.Marshal(m)
	if err != nil {
		seelog.Error("Unable to marshal auth tokens. Failed to save token to the file")
		return err
	}

	err = ioutil.WriteFile(d.AuthFile, f, 0600)

	if err != nil {
		seelog.Errorf("Unable to write authentication token to file: %s", d.AuthFile)
		return err
	}

	return nil
}

func (d *DropboxAuthenticator) LoadOrAuth(business bool, storeToken bool) (string, error) {
	t, err := d.TokenFileLoad()
	if err != nil {
		return d.Authorise(storeToken)
	}

	return t, nil
}

func (d *DropboxAuthenticator) Authorise(storeToken bool) (string, error) {
	seelog.Debugf("Authorize(storeToken:%t)", storeToken)
	seelog.Flush()

	if d.AppKey == "" || d.AppSecret == "" {
		seelog.Tracef("No AppKey/AppSecret found. Try asking 'Generate Token'")
		tok, err := d.acquireToken()
		if err == nil && storeToken {
			d.TokenFileSave(tok)
		}
		return tok, err
	} else {
		seelog.Tracef("Start auth sequence for AppKey[%s]", d.AppKey)
		state, err := util.GenerateRandomString(8)
		if err != nil {
			seelog.Errorf("Unable to generate `state` [%s]", err)
			return "", err
		}

		tok, err := d.auth(state)
		if err != nil {
			seelog.Errorf("Authentication failed due to the error [%s]", err)
			return "", err
		}
		if storeToken {
			d.TokenFileSave(tok.AccessToken)
		}
		return tok.AccessToken, nil
	}
}

func (d *DropboxAuthenticator) acquireToken() (string, error) {
	d.generateTokenInstruction()

	var code string

	for {
		fmt.Println(authGeneratedToken2)
		if _, err := fmt.Scan(&code); err != nil {
			seelog.Errorf("Input error (%s), try again.", err)
			continue
		}
		trim := strings.TrimSpace(code)
		if len(trim) < 1 {
			seelog.Errorf("Input error, try again.")
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

func RevokeToken(token string) {
	//config := dropbox.Config{Token: token}
	//client := auth.New(config)
	//err := client.TokenRevoke()
	//if err != nil {
	//	seelog.Warnf("Error during clean up token: %s", err)
	//} else {
	//	seelog.Info("Token successfully revoked")
	//}
}
