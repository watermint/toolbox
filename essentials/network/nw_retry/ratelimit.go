package nw_retry

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_congestion"
	"github.com/watermint/toolbox/infra/api/api_context"
	"net/http"
)

func NewRatelimit(client nw_client.Rest) nw_client.Rest {
	return &RateLimit{
		client: client,
	}
}

type RateLimit struct {
	client nw_client.Rest
}

func (z RateLimit) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res es_response.Response) {
	l := esl.Default()

	nw_congestion.Start(ctx.ClientHash(), req.Endpoint())
	res = z.client.Call(ctx, req)
	if res.IsSuccess() || res.TransportError() == nil {
		if res.IsSuccess() {
			nw_congestion.EndSuccess(ctx.ClientHash(), req.Endpoint())
		} else {
			nw_congestion.EndTransportError(ctx.ClientHash(), req.Endpoint())
		}
		return res
	}

	switch res.Code() {
	case http.StatusTooManyRequests:
		l.Debug("Ratelimit", esl.Int("code", res.Code()))
		nw_congestion.EndRateLimit(ctx.ClientHash(), req.Endpoint())
		return es_response_impl.NewTransportErrorResponse(NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
	}
	return res
}
