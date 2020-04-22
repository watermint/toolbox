package gh_rest

import (
	"github.com/watermint/toolbox/domain/github/api/gh_recovery"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/network/nw_capture"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_http"
	"github.com/watermint/toolbox/infra/network/nw_retry"
)

var (
	defaultClient = nw_retry.NewRetry(
		gh_recovery.New(
			nw_capture.New(
				nw_http.NewClient())))
	testClient = nw_capture.New(nw_http.NewClient())
)

func Default(feature app_feature.Feature) nw_client.Rest {
	if feature.IsTest() {
		return testClient
	} else {
		return defaultClient
	}
}
