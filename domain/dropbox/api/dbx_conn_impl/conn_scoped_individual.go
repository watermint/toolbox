package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
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
	currentScope := z.Scopes()
	if ctl.Feature().Experiment(app.ExperimentDbxAuthCourseGrainedScope) {
		currentScope = []string{}
	}
	z.ctx, err = connect(currentScope, z.name, ctl, dbx_auth.DropboxIndividual)
	return err
}

func (z *connScopedIndividual) PeerName() string {
	return z.name
}

func (z *connScopedIndividual) SetPeerName(name string) {
	z.name = name
}

func (z *connScopedIndividual) ScopeLabel() string {
	return app.ServiceDropboxIndividual
}

func (z *connScopedIndividual) ServiceName() string {
	return api_conn.ServiceDropbox
}

func (z *connScopedIndividual) Client() dbx_client.Client {
	return z.ctx
}

func (z *connScopedIndividual) SetScopes(scopes ...string) {
	ss := make([]string, len(scopes))
	copy(ss[:], scopes[:])
	if !dbx_auth.HasAccountInfoRead(scopes) {
		ss = append(ss, dbx_auth.ScopeAccountInfoRead)
	}
	sort.Strings(ss)
	z.scopes = ss
}

func (z *connScopedIndividual) Scopes() []string {
	return z.scopes
}

func (z *connScopedIndividual) IsIndividual() bool {
	return true
}
