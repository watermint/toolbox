package _import

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
)

type ViaAppVO struct {
	Peer            app_conn.ConnUserFile
	WorkDropboxPath string
	DestDropboxPath string
	LocalPath       string
}

type ViaAppWorker struct {
	ctx api_context.Context
}
