package as_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
)

// https://developers.asana.com/docs/oauth
const (
	// Provides access to all endpoints documented in our API reference. If no scopes are requested, this scope is assumed by default.
	ScopeDefault = "default"

	// Provides access to OpenID Connect ID tokens and the OpenID Connect user info endpoint.
	ScopeOpenId = "openid"

	// Provides access to the user's email through the OpenID Connect user info endpoint.
	ScopeEmail = "email"

	// Provides access to the user's name and profile photo through the OpenID Connect user info endpoint.
	ScopeProfile = "profile"
)

var (
	Asana = api_auth.OAuthAppData{
		AppKeyName:       api_auth.Asana,
		EndpointAuthUrl:  "https://app.asana.com/-/oauth_authorize",
		EndpointTokenUrl: "https://app.asana.com/-/oauth_token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)

func New(ctl app_control.Control) api_auth.OAuthAppLegacy {
	return &App{
		ctl: ctl,
		res: api_appkey.New(ctl),
	}
}

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z App) Config(scope []string) *oauth2.Config {
	key, secret := z.res.Key(api_auth.Asana)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://app.asana.com/-/oauth_authorize",
			TokenURL: "https://app.asana.com/-/oauth_token",
		},
		Scopes: scope,
	}
}

func (z App) UsePKCE() bool {
	return true
}
