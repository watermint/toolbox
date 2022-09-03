package nw_retry

import (
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_congestion"
	"net/http"
	"time"
)

func NewRatelimit(client nw_client.Rest) nw_client.Rest {
	return &RateLimit{
		client: client,
	}
}

type RateLimit struct {
	client nw_client.Rest
}

const (
	exitReasonSuccess = iota
	exitReasonTransportError
	exitReasonRateLimit
)

func (z RateLimit) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	var errRateLimit *ErrorRateLimit
	exitReason := exitReasonTransportError
	l := esl.Default()
	defer func() {
		switch exitReason {
		case exitReasonSuccess:
			nw_congestion.EndSuccess(ctx.ClientHash(), req.Endpoint())
		case exitReasonRateLimit:
			reset := time.Now()
			if errRateLimit != nil {
				reset = errRateLimit.Reset
			}
			nw_congestion.EndRateLimit(ctx.ClientHash(), req.Endpoint(), reset)
		default:
			nw_congestion.EndTransportError(ctx.ClientHash(), req.Endpoint())
		}
	}()

	nw_congestion.Start(ctx.ClientHash(), req.Endpoint())
	res = z.client.Call(ctx, req)
	if res.IsSuccess() || res.TransportError() == nil {
		if res.IsSuccess() {
			exitReason = exitReasonSuccess
		}
		return res
	}

	switch res.Code() {
	case http.StatusTooManyRequests:
		l.Debug("Ratelimit", esl.Int("code", res.Code()))
		exitReason = exitReasonRateLimit
		errRateLimit = NewErrorRateLimitFromHeadersFallback(res.Headers())
		return es_response_impl.NewTransportErrorResponse(errRateLimit, res)
	}
	return res
}
