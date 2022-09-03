package api_auth

import (
	"golang.org/x/oauth2"
)

// Deprecated: NewContext
func NewContext(token *oauth2.Token, cfg *oauth2.Config, peerName string, scopes []string) OAuthContext {
	return &contextImpl{
		cfg:      cfg,
		token:    token,
		scopes:   scopes,
		peerName: peerName,
	}
}

// Deprecated: OAuthContext of OAuth
type OAuthContext interface {
	Config() *oauth2.Config
	Token() *oauth2.Token
	Scopes() []string
	PeerName() string
	Description() string
	Supplemental() string
	IsNoAuth() bool
}

type contextImpl struct {
	cfg      *oauth2.Config
	token    *oauth2.Token
	peerName string
	scopes   []string
	desc     string
	suppl    string
}

func (z *contextImpl) Config() *oauth2.Config {
	return z.cfg
}

func (z *contextImpl) Scopes() []string {
	return z.scopes
}

func (z *contextImpl) Token() *oauth2.Token {
	return z.token
}

func (z *contextImpl) PeerName() string {
	return z.peerName
}

func (z *contextImpl) Description() string {
	return z.desc
}

func (z *contextImpl) Supplemental() string {
	return z.suppl
}

func (z *contextImpl) IsNoAuth() bool {
	return false
}
