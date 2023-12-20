package as_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_definitions"
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
		AppKeyName:       app_definitions.ServiceAsana,
		EndpointAuthUrl:  "https://app.asana.com/-/oauth_authorize",
		EndpointTokenUrl: "https://app.asana.com/-/oauth_token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)
