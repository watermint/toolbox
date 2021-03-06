package goog_conn_impl

import (
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
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
	ctx    goog_context.Context
}

func (z *connSheets) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(goog_context_impl.EndpointGoogleSheets, z.scopes, z.name, ctl)
	return
}

func (z *connSheets) PeerName() string {
	return z.name
}

func (z *connSheets) SetPeerName(name string) {
	z.name = name
}

func (z *connSheets) ScopeLabel() string {
	return api_auth.GoogleSheets
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

func (z *connSheets) Context() goog_context.Context {
	return z.ctx
}

func (z *connSheets) IsSheets() bool {
	return true
}
