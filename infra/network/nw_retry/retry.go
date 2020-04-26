package nw_retry

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"time"
)

func NewRetry(client nw_client.Rest) nw_client.Rest {
	return &Retry{
		client: client,
	}
}

type Retry struct {
	client nw_client.Rest
}

func (z *Retry) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res response.Response) {
	l := ctx.Log().With(
		zap.String("Url", req.Endpoint()),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	res = z.client.Call(ctx, req)

	// return when on success
	if res.IsSuccess() {
		return res
	}

	switch er := res.TransportError().(type) {
	case *ErrorRateLimit:
		l.Debug("Rate limit, waiting for reset",
			zap.Int("limit", er.Limit),
			zap.Int("remaining", er.Remaining),
			zap.String("reset", er.Reset.Format(time.RFC3339)))
		nw_ratelimit.UpdateRetryAfter(ctx.ClientHash(), req.Endpoint(), er.Reset)
		return z.Call(ctx, req)

	default:
		abort := nw_ratelimit.AddError(ctx.ClientHash(), req.Endpoint(), er)
		if abort {
			l.Debug("Abort retry due to retries exceeds retry limit")
			return res
		}
		l.Debug("Retrying", zap.Error(er))
		return z.Call(ctx, req)
	}
}
