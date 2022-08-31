package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
)

type Legacy struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z *Legacy) UsePKCE() bool {
	return false
}

func (z *Legacy) Config(scopes []string) *oauth2.Config {
	if len(scopes) != 1 {
		panic("Unsupported scope type")
	}
	scope := scopes[0]
	key, secret := z.res.Key(scope)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     DropboxOAuthEndpoint(),
	}
}

func NewLegacyApp(ctl app_control.Control) api_auth.OAuthAppLegacy {
	a := &Legacy{
		ctl: ctl,
		res: api_appkey.New(ctl),
	}
	return a
}
