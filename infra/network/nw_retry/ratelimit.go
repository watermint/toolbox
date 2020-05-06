package nw_retry

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/network/nw_client"
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
	l := es_log.Default()

	res = z.client.Call(ctx, req)
	if res.IsSuccess() || res.TransportError() == nil {
		return res
	}

	switch res.Code() {
	case http.StatusTooManyRequests:
		l.Debug("Ratelimit", es_log.Int("code", res.Code()))
		return es_response_impl.NewTransportErrorResponse(NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
	}
	return res
}
