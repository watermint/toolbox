package rc_conn

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ConnDropboxApi interface {
	Name() string

	Context() api_context.Context

	// Connect to api
	Connect(ctl app_control.Control) (err error)

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
