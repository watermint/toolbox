package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnScopedDropboxApi interface {
	api_conn.Connection

	Context() dbx_context.Context

	// True when the connection type is personal
	IsPersonal() bool

	// True when the connection type is business
	IsBusiness() bool

	// Update scopes
	SetScopes(scopes []string)

	// Scopes
	Scopes() []string
}

type ConnScopedTeam interface {
	ConnScopedDropboxApi
	IsTeam() bool
}

type ConnScopedIndividual interface {
	ConnScopedDropboxApi
	IsIndividual() bool
}
