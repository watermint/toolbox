package api_auth

import (
	"golang.org/x/oauth2"
)

func NewContext(token *oauth2.Token, peerName, scope string) Context {
	return &contextImpl{
		token:    token,
		scope:    scope,
		peerName: peerName,
	}
}

func NewContextWithAttr(c Context, desc, suppl string) Context {
	return &contextImpl{
		token:    c.Token(),
		peerName: c.PeerName(),
		scope:    c.Scope(),
		desc:     desc,
		suppl:    suppl,
	}
}

// Auth context
type Context interface {
	Token() *oauth2.Token
	Scope() string
	PeerName() string
	Description() string
	Supplemental() string
	IsNoAuth() bool
}

func NewNoAuth() Context {
	return &noAuthContext{}
}

type noAuthContext struct {
}

func (z *noAuthContext) Scope() string {
	return ""
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
	token    *oauth2.Token
	peerName string
	scope    string
	desc     string
	suppl    string
}

func (z *contextImpl) Scope() string {
	return z.scope
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
