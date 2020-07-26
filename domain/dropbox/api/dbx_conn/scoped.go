package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnScopedDropboxApi interface {
	api_conn.ScopedConnection

	Context() dbx_context.Context
}

type ConnScopedTeam interface {
	ConnScopedDropboxApi
	IsTeam() bool
}

type ConnScopedIndividual interface {
	ConnScopedDropboxApi
	IsIndividual() bool
}
