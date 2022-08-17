package api_auth

import (
	"golang.org/x/oauth2"
)

func NewContext(token *oauth2.Token, cfg *oauth2.Config, peerName string, scopes []string) OAuthContext {
	return &contextImpl{
		cfg:      cfg,
		token:    token,
		scopes:   scopes,
		peerName: peerName,
	}
}

func NewContextWithAttr(c OAuthContext, cfg *oauth2.Config, desc, suppl string) OAuthContext {
	return &contextImpl{
		cfg:      cfg,
		token:    c.Token(),
		peerName: c.PeerName(),
		scopes:   c.Scopes(),
		desc:     desc,
		suppl:    suppl,
	}
}

// OAuthContext of OAuth
type OAuthContext interface {
	Config() *oauth2.Config
	Token() *oauth2.Token
	Scopes() []string
	PeerName() string
	Description() string
	Supplemental() string
	IsNoAuth() bool
}

func NewNoAuth() OAuthContext {
	return &noAuthContext{}
}

type noAuthContext struct {
}

func (z *noAuthContext) Config() *oauth2.Config {
	return &oauth2.Config{}
}

func (z *noAuthContext) Scopes() []string {
	return []string{}
}

func (z *noAuthContext) Token() *oauth2.Token {
	return &oauth2.Token{}
}

func (z *noAuthContext) PeerName() string {
	return ""
}

func (z *noAuthContext) Description() string {
	return ""
}

func (z *noAuthContext) Supplemental() string {
	return ""
}

func (z *noAuthContext) IsNoAuth() bool {
	return true
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
