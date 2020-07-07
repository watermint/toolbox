package nw_http

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_throttle"
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
	l := esl.Default().With(
		esl.String("Endpoint", endpoint),
		esl.String("Routine", es_goroutine.GetGoRoutineName()),
	)

	l.Debug("Call")
	var callStart, callEnd time.Time
	nw_throttle.Throttle(hash, endpoint, func() {
		callStart = time.Now()
		res, err = z.client.Do(req)
		callEnd = time.Now()
	})

	latency = callEnd.Sub(callStart)

	if err != nil {
		return nil, latency, err
	}
	return res, latency, nil
}
