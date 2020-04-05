package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"golang.org/x/oauth2"
)

func NewContext(token *oauth2.Token, peerName, scope string) api_auth.Context {
	return &Context{
		token:    token,
		scope:    scope,
		peerName: peerName,
	}
}

func NewContextWithAttr(c api_auth.Context, desc, suppl string) api_auth.Context {
	return &Context{
		token:    c.Token(),
		peerName: c.PeerName(),
		scope:    c.Scope(),
		desc:     desc,
		suppl:    suppl,
	}
}

func NewNoAuth() api_auth.Context {
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

type Context struct {
	token    *oauth2.Token
	peerName string
	scope    string
	desc     string
	suppl    string
}

func (z *Context) Scope() string {
	return z.scope
}

func (z *Context) Token() *oauth2.Token {
	return z.token
}

func (z *Context) PeerName() string {
	return z.peerName
}

func (z *Context) Description() string {
	return z.desc
}

func (z *Context) Supplemental() string {
	return z.suppl
}

func (z *Context) IsNoAuth() bool {
	return false
}
