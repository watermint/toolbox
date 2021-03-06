package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConnScopedTeam(name string) dbx_conn.ConnScopedTeam {
	return &connScopedTeam{
		name:   name,
		scopes: []string{},
		ctx:    nil,
	}
}

type connScopedTeam struct {
	name   string
	scopes []string
	ctx    dbx_context.Context
}

func (z *connScopedTeam) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.Scopes(), z.name, ctl, dbx_auth.NewScopedTeam(ctl))
	return err
}

func (z *connScopedTeam) PeerName() string {
	return z.name
}

func (z *connScopedTeam) SetPeerName(name string) {
	z.name = name
}

func (z *connScopedTeam) ScopeLabel() string {
	return api_auth.DropboxScopedTeam
}

func (z *connScopedTeam) ServiceName() string {
	return api_conn.ServiceDropboxBusiness
}

func (z *connScopedTeam) Context() dbx_context.Context {
	return z.ctx
}

func (z *connScopedTeam) SetScopes(scopes ...string) {
	z.scopes = scopes
}

func (z *connScopedTeam) Scopes() []string {
	return z.scopes
}

func (z *connScopedTeam) IsTeam() bool {
	return true
}
