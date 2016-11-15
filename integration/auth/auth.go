package auth

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial"
	"github.com/watermint/toolbox/infra/util"
	"golang.org/x/oauth2"
	"strings"
)

type DropboxAuthenticator struct {
	AppKey    string
	AppSecret string
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

	authGeneratedToken1 = `========================================================
1. Visit the MyApp page (you mihgt have to login first):

https://www.dropbox.com/developers/apps

2. Proceed with "Create App"
3. Choose "Dropbox API"
4. Choose "Full Dropbox"
5. Enter name of your app
6. Proceed with "Create App"
7. Hit "Generate" button near "Generated access token"
8. Copy generated token:
`
	authGeneratedToken2 = `
Enter the generated token here:
`
)

func (d *DropboxAuthenticator) Authorise() (string, error) {
	seelog.Flush()

	if d.AppKey == "" || d.AppSecret == "" {
		return d.acquireToken()
	} else {
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
		return tok.AccessToken, nil
	}
}

func (d *DropboxAuthenticator) acquireToken() (string, error) {
	fmt.Println(authGeneratedToken1)

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
	return cfg.Exchange(oauth2.NoContext, code)
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
	client := dropbox.Client(token, dropbox.Options{})
	err := client.TokenRevoke()
	if err != nil {
		seelog.Warnf("Error during clean up token: %s", err)
	}
}
