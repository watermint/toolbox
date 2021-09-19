package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnLegacyDropboxApi interface {
	api_conn.Connection

	Context() dbx_context.Context
}

// Deprecated: ConnUserFile
type ConnUserFile interface {
	ConnLegacyDropboxApi
	IsUserFile()
}
