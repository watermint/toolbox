package nw_http

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_concurrency"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"net/http"
	"time"
)

func NewClient() nw_client.Http {
	return &Client{
		client: http.Client{
			Jar:     nil,
			Timeout: 1 * time.Minute,
		},
	}
}

type Client struct {
	client http.Client
}

// Call RPC. res will be nil on an error
func (z *Client) Call(hash, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error) {
	l := es_log.Default().With(
		es_log.String("Endpoint", endpoint),
		es_log.String("Routine", es_goroutine.GetGoRoutineName()),
	)

	l.Debug("Call")
	nw_ratelimit.WaitIfRequired(hash, endpoint)
	nw_concurrency.Start()
	callStart := time.Now()
	res, err = z.client.Do(req)
	callEnd := time.Now()
	nw_concurrency.End()

	latency = callEnd.Sub(callStart)

	if err != nil {
		return nil, latency, err
	}
	return res, latency, nil
}
