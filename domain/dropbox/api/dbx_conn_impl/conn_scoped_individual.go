package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"sort"
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
	ctx    dbx_client.Client
}

func (z *connScopedIndividual) Connect(ctl app_control.Control) (err error) {
	z.ctx, err = connect(z.Scopes(), z.name, ctl, dbx_auth.DropboxIndividual)
	return err
}

func (z *connScopedIndividual) PeerName() string {
	return z.name
}

func (z *connScopedIndividual) SetPeerName(name string) {
	z.name = name
}

func (z *connScopedIndividual) ScopeLabel() string {
	return api_auth.DropboxIndividual
}

func (z *connScopedIndividual) ServiceName() string {
	return api_conn.ServiceDropbox
}

func (z *connScopedIndividual) Context() dbx_client.Client {
	return z.ctx
}

func (z *connScopedIndividual) SetScopes(scopes ...string) {
	ss := make([]string, len(scopes))
	copy(ss[:], scopes[:])
	sort.Strings(ss)
	z.scopes = ss
}

func (z *connScopedIndividual) Scopes() []string {
	return z.scopes
}

func (z *connScopedIndividual) IsIndividual() bool {
	return true
}
