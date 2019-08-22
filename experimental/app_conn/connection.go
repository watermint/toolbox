package app_conn

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/experimental/app_control"
)

type ConnDropboxApi interface {
	Connect(control app_control.Control) (ctx api_context.Context, err error)
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
