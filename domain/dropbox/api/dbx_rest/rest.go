package dbx_rest

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_recovery"
	"github.com/watermint/toolbox/infra/network/nw_capture"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_http"
)

var (
	defaultClient = dbx_recovery.New(nw_capture.New(nw_http.NewClient()))
)

func Default() nw_client.Rest {
	return defaultClient
}
