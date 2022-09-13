package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/api/goog_client_impl"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
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
	ctx    goog_client.Client
}

func (z *connGoogleCalendar) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_auth.Calendar, goog_client_impl.EndpointGoogleCalendar, z.scopes, z.name, ctl)
	return
}

func (z *connGoogleCalendar) PeerName() string {
	return z.name
}

func (z *connGoogleCalendar) SetPeerName(name string) {
	z.name = name
}

func (z *connGoogleCalendar) ScopeLabel() string {
	return app.ServiceGoogleCalendar
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

func (z *connGoogleCalendar) Client() goog_client.Client {
	return z.ctx
}

func (z *connGoogleCalendar) IsCalendar() bool {
	return true
}
