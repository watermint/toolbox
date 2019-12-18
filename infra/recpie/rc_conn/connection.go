package rc_conn

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ConnDropboxApi interface {
	Connect(control app_control.Control) (ctx api_context.Context, err error)
	Name() string
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
