package api_auth

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
