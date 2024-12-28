package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnScopedDropboxApi interface {
	api_conn.ScopedConnection

	Client() dbx_client.Client
}

type ConnScopedTeam interface {
	ConnScopedDropboxApi
	IsTeam() bool
}

type ConnScopedIndividual interface {
	ConnScopedDropboxApi
	IsIndividual() bool
}

type FileSystemIdentifier interface {
	Version() (mo_team.TeamFileSystemType, error)
}
