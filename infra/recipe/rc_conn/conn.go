package rc_conn

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ConnDropboxApi interface {
	Name() string
	Context() api_context.Context
	Connect(ctl app_control.Control) (err error)
	SetName(name string)
	ScopeLabel() string
	IsPersonal() bool
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
