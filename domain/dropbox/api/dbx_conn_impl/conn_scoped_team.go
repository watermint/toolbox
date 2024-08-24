package dbx_conn_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"sort"
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
	ctx    dbx_client.Client
}

func (z *connScopedTeam) Connect(ctl app_control.Control) (err error) {
	currentScope := z.Scopes()
	if ctl.Feature().Experiment(app_definitions.ExperimentDbxAuthCourseGrainedScope) {
		currentScope = []string{}
	}
	z.ctx, err = connect(currentScope, z.name, ctl, dbx_auth.DropboxTeam)
	return err
}

func (z *connScopedTeam) PeerName() string {
	return z.name
}

func (z *connScopedTeam) SetPeerName(name string) {
	z.name = name
}

func (z *connScopedTeam) ScopeLabel() string {
	return app_definitions.ScopeLabelDropboxTeam
}

func (z *connScopedTeam) AppKeyName() string {
	return app_definitions.AppKeyDropboxTeam
}

func (z *connScopedTeam) Client() dbx_client.Client {
	return z.ctx
}

func (z *connScopedTeam) SetScopes(scopes ...string) {
	ss := make([]string, len(scopes))
	copy(ss[:], scopes[:])
	if !dbx_auth.HasTeamInfoRead(scopes) {
		ss = append(ss, dbx_auth.ScopeTeamInfoRead)
	}
	sort.Strings(ss)
	z.scopes = ss
}

func (z *connScopedTeam) Scopes() []string {
	return z.scopes
}

func (z *connScopedTeam) IsTeam() bool {
	return true
}
