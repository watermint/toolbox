package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnScopedIndividual(name string) dbx_conn.ConnScopedIndividual {
	return &connScopedIndividual{
		name:   name,
		scopes: []string{},
		ctx:    nil,
	}
}

type connScopedIndividual struct {
	name   string
	scopes []string
	ctx    dbx_context.Context
}

func (z *connScopedIndividual) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.Scopes(), z.name, ctl, dbx_auth.NewScopedIndividual(ctl))
	return err
}

func (z *connScopedIndividual) PeerName() string {
	return z.name
}

func (z *connScopedIndividual) SetPeerName(name string) {
	z.name = name
}

func (z *connScopedIndividual) ScopeLabel() string {
	return api_auth.DropboxScopedIndividual
}

func (z *connScopedIndividual) ServiceName() string {
	return api_conn.ServiceDropbox
}

func (z *connScopedIndividual) Context() dbx_context.Context {
	return z.ctx
}

func (z *connScopedIndividual) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connScopedIndividual) Scopes() []string {
	return z.scopes
}

func (z *connScopedIndividual) IsIndividual() bool {
	return true
}
