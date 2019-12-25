package rc_conn

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
)

// Deprecated: use ConnDropboxApi
type OldConnDropboxApi interface {
	Connect(control app_control.Control) (ctx api_context.Context, err error)
	Name() string
}

// Deprecated: use ConnUserFile
type OldConnUserFile interface {
	OldConnDropboxApi
	IsUserFile()
}

// Deprecated: use ConnBusinessAudit
type OldConnBusinessAudit interface {
	OldConnDropboxApi
	IsBusinessAudit()
}

// Deprecated: use ConnBusinessInfo
type OldConnBusinessInfo interface {
	OldConnDropboxApi
	IsBusinessInfo()
}

// Deprecated: use ConnBusinessMgmt
type OldConnBusinessMgmt interface {
	OldConnDropboxApi
	IsBusinessMgmt()
}

// Deprecated: use ConnBusinessFile
type OldConnBusinessFile interface {
	OldConnDropboxApi
	IsBusinessFile()
}
