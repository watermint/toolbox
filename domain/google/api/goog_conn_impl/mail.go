package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnGoogleMail(name string) goog_conn.ConnGoogleMail {
	return &connGoogleMail{
		name:   name,
		scopes: []string{},
	}
}

type connGoogleMail struct {
	name   string
	scopes []string
	ctx    goog_context.Context
}

func (z *connGoogleMail) IsGmail() bool {
	return true
}

func (z *connGoogleMail) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_context_impl.EndpointGoogleApis, z.scopes, z.name, ctl)
	return
}

func (z *connGoogleMail) PeerName() string {
	return z.name
}

func (z *connGoogleMail) SetPeerName(name string) {
	z.name = name
}

func (z *connGoogleMail) ScopeLabel() string {
	return api_auth.GoogleMail
}

func (z *connGoogleMail) ServiceName() string {
	return api_conn.ServiceGoogleMail
}

func (z *connGoogleMail) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connGoogleMail) Scopes() []string {
	return z.scopes
}

func (z *connGoogleMail) Context() goog_context.Context {
	return z.ctx
}
