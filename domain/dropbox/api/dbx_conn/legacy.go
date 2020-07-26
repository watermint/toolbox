package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/api/api_conn"
)

type ConnLegacyDropboxApi interface {
	api_conn.Connection

	Context() dbx_context.Context
}

type ConnUserFile interface {
	ConnLegacyDropboxApi
	IsUserFile()
}

type ConnBusinessAudit interface {
	ConnLegacyDropboxApi
	IsBusinessAudit()
}

type ConnBusinessInfo interface {
	ConnLegacyDropboxApi
	IsBusinessInfo()
}

type ConnBusinessMgmt interface {
	ConnLegacyDropboxApi
	IsBusinessMgmt()
}

type ConnBusinessFile interface {
	ConnLegacyDropboxApi
	IsBusinessFile()
}
