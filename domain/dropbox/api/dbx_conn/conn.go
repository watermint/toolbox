package dbx_conn

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

type ConnDropboxApi interface {
	rc_conn.Connection

	Name() string

	Context() dbx_context.Context

	// Update peer name
	SetPeerName(name string)

	// Scope label
	ScopeLabel() string

	// True when the connection required to verify before the operation
	IsPreVerify() bool

	// True when the connection type is personal
	IsPersonal() bool

	// True when the connection type is business
	IsBusiness() bool
}

type ConnUserFile interface {
	ConnDropboxApi
	IsUserFile()
}

type ConnBusinessAudit interface {
	ConnDropboxApi
	IsBusinessAudit()
}

type ConnBusinessInfo interface {
	ConnDropboxApi
	IsBusinessInfo()
}

type ConnBusinessMgmt interface {
	ConnDropboxApi
	IsBusinessMgmt()
}

type ConnBusinessFile interface {
	ConnDropboxApi
	IsBusinessFile()
}
