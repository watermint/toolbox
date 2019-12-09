package nw_client

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_concurrency"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Jar:     nil,
		Timeout: 1 * time.Minute,
	}
)

// Call RPC. res will be nil on an error
func Call(hash, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error) {
	l := app_root.Log().With(
		zap.String("Endpoint", endpoint),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	l.Debug("Call")
	nw_ratelimit.WaitIfRequired(hash, endpoint)
	nw_concurrency.Start()
	callStart := time.Now()
	res, err = client.Do(req)
	callEnd := time.Now()
	nw_concurrency.End()

	latency = callEnd.Sub(callStart)

	if err != nil {
		return nil, latency, err
	}
	return res, latency, nil
}
