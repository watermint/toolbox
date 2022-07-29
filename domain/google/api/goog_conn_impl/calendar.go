package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnGoogleCalendar(name string) goog_conn.ConnGoogleCalendar {
	return &connGoogleCalendar{
		name:   name,
		scopes: []string{},
	}
}

type connGoogleCalendar struct {
	name   string
	scopes []string
	ctx    goog_context.Context
}

func (z *connGoogleCalendar) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_context_impl.EndpointGoogleCalendar, z.scopes, z.name, ctl)
	return
}

func (z *connGoogleCalendar) PeerName() string {
	return z.name
}

func (z *connGoogleCalendar) SetPeerName(name string) {
	z.name = name
}

func (z *connGoogleCalendar) ScopeLabel() string {
	return api_auth.GoogleCalendar
}

func (z *connGoogleCalendar) ServiceName() string {
	return api_conn.ServiceGoogleCalendar
}

func (z *connGoogleCalendar) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connGoogleCalendar) Scopes() []string {
	return z.scopes
}

func (z *connGoogleCalendar) Context() goog_context.Context {
	return z.ctx
}

func (z *connGoogleCalendar) IsCalendar() bool {
	return true
}
