package nw_retry

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
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

func (z *Retry) Call(ctx api_context.Context, req api_request.Request) (res response.Response, err error) {
	// path through when no retry enabled.
	if ctx.IsNoRetry() {
		return z.client.Call(ctx, req)
	}

	l := ctx.Log().With(
		zap.String("Url", req.Url()),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	res, err = z.client.Call(ctx, req)

	// return when on success
	if err == nil {
		return res, nil
	}

	switch er := err.(type) {
	case *ErrorRateLimit:
		l.Debug("Rate limit, waiting for reset",
			zap.Int("limit", er.Limit),
			zap.Int("remaining", er.Remaining),
			zap.String("reset", er.Reset.Format(time.RFC3339)))
		nw_ratelimit.UpdateRetryAfter(ctx.ClientHash(), req.Endpoint(), er.Reset)
		return z.Call(ctx, req)

	default:
		abort := nw_ratelimit.AddError(ctx.ClientHash(), req.Endpoint(), err)
		if abort {
			l.Debug("Abort retry due to retries exceeds retry limit")
			return nil, err
		}
		l.Debug("Retrying", zap.Error(err))
		return z.Call(ctx, req)
	}
}
