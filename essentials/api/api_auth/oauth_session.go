package api_auth

type OAuthSessionData struct {
	AppData           OAuthAppData `json:"app_data"`
	PeerName          string       `json:"peer_name"`
	Scopes            []string     `json:"scopes"`
	UseSecureRedirect bool         `json:"use_secure_redirect"`
}

type OAuthSession interface {
	Start(session OAuthSessionData) (entity OAuthEntity, err error)
}
