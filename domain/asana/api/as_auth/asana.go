package as_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
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
