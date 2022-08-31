package api_auth

import "golang.org/x/oauth2"

// Deprecated: OAuthAppLegacy OAuth Application key/secret manager
type OAuthAppLegacy interface {
	// Config OAuth2 config
	Config(scope []string) *oauth2.Config

	// UsePKCE Use PKCE on authentication
	UsePKCE() bool
}

// OAuthConsole OAuth interface for console UI
type OAuthConsole interface {
	Auth

	Start(scope []string) (token OAuthContext, err error)
}
