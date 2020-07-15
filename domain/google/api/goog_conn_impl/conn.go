package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/api/api_conn_impl"
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
	ac, useMock, err := api_conn_impl.Connect(z.scopes, z.name, goog_auth.NewApp(ctl), ctl)
	if useMock {
		z.ctx = goog_context_impl.NewMock(ctl)
		return nil
	}
	if ac != nil {
		z.ctx = goog_context_impl.New(ctl, ac)
		return nil
	}
	return err
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
