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

func NewConnGoogleSheets(name string) goog_conn.ConnGoogleSheets {
	return &connSheets{
		name:   name,
		scopes: []string{},
	}
}

type connSheets struct {
	name   string
	scopes []string
	ctx    goog_client.Client
}

func (z *connSheets) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_auth.Sheets, goog_client_impl.EndpointGoogleSheets, z.scopes, z.name, ctl)
	return
}

func (z *connSheets) PeerName() string {
	return z.name
}

func (z *connSheets) SetPeerName(name string) {
	z.name = name
}

func (z *connSheets) ScopeLabel() string {
	return app.ServiceGoogleSheets
}

func (z *connSheets) ServiceName() string {
	return api_conn.ServiceGoogleSheets
}

func (z *connSheets) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connSheets) Scopes() []string {
	return z.scopes
}

func (z *connSheets) Client() goog_client.Client {
	return z.ctx
}

func (z *connSheets) IsSheets() bool {
	return true
}
