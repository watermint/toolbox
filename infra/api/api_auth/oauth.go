package api_auth

import (
	"golang.org/x/oauth2"
	"sort"
)

type OAuthEndpointStyle string

const (
	AuthStyleAutoDetect OAuthEndpointStyle = ""
	AuthStyleInParams   OAuthEndpointStyle = "param"
	AuthStyleInHeader   OAuthEndpointStyle = "header"
)

type OAuthAppData struct {
	// Name to retrieve client_id/client_secret from app key registry.
	AppKeyName string `json:"app_key_name"`

	// Auth Endpoint
	EndpointAuthUrl string `json:"endpoint_auth_url"`

	// Token Endpoint
	EndpointTokenUrl string `json:"endpoint_token_url"`

	// Endpoint parameter type
	EndpointStyle OAuthEndpointStyle `json:"endpoint_style"`

	// Use PKCE (Proof Key for Code Exchange, RFC7636) or not
	UsePKCE bool `json:"use_pkce"`

	// Redirect URL
	RedirectUrl string `json:"redirect_url"`
}

type OAuthKeyResolver func(appKey string) (clientId, clientSecret string)

func (z OAuthAppData) Config(scopes []string, resolve OAuthKeyResolver) *oauth2.Config {
	sortedScopes := make([]string, len(scopes))
	copy(sortedScopes[:], scopes[:])
	sort.Strings(sortedScopes)

	style := oauth2.AuthStyleAutoDetect
	switch z.EndpointStyle {
	case AuthStyleInParams:
		style = oauth2.AuthStyleInParams
	case AuthStyleInHeader:
		style = oauth2.AuthStyleInHeader
	}

	clientId, clientSecret := resolve(z.AppKeyName)

	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   z.EndpointAuthUrl,
			TokenURL:  z.EndpointTokenUrl,
			AuthStyle: style,
		},
		RedirectURL: z.RedirectUrl,
		Scopes:      scopes,
	}
}

type OAuthSessionData struct {
	AppData  OAuthAppData `json:"app_data"`
	PeerName string       `json:"peer_name"`
	Scopes   []string     `json:"scopes"`
}

type OAuthSession interface {
	Start(session OAuthSessionData) (entity OAuthEntity, err error)
}
