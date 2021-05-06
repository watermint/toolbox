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

// Deprecated: ConnBusinessAudit
type ConnBusinessAudit interface {
	ConnLegacyDropboxApi
	IsBusinessAudit()
}

// Deprecated: ConnBusinessInfo
type ConnBusinessInfo interface {
	ConnLegacyDropboxApi
	IsBusinessInfo()
}

// Deprecated: ConnBusinessMgmt
type ConnBusinessMgmt interface {
	ConnLegacyDropboxApi
	IsBusinessMgmt()
}

// Deprecated: ConnBusinessFile
type ConnBusinessFile interface {
	ConnLegacyDropboxApi
	IsBusinessFile()
}
