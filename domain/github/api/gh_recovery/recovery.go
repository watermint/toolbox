package gh_recovery

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"go.uber.org/zap"
)

var (
	ErrorNoResponse               = errors.New("no response")
	ErrorGeneralApiError          = errors.New("general api error")
	ErrorUnexpectedResponseFormat = errors.New("unexpected response format")
)

func New(client nw_client.Rest) nw_client.Rest {
	return &recoveryImpl{
		client: client,
	}
}

type recoveryImpl struct {
	client nw_client.Rest
}

func (z recoveryImpl) Call(ctx api_context.Context, req api_request.Request) (res api_response.Response, err error) {
	l := ctx.Log()
	res, err = z.client.Call(ctx, req)
	if err != nil {
		l.Debug("nil response")
		return nil, err
	}

	if res == nil {
		return nil, ErrorNoResponse
	}

	switch {
	case res.StatusCode()/100 == 2:
		return res, nil

	default:
		// Rate limit
		erl, found := nw_retry.NewErrorRateLimitFromHeaders(res.Headers())
		if found && erl.Remaining < 1 {
			return nil, erl
		}

		// General errors
		apiErr := &gh_context.ApiError{}
		if err := res.Model(apiErr); err != nil {
			l.Debug("unable to parse the error", zap.Error(err))
			return nil, ErrorUnexpectedResponseFormat
		}
		return nil, apiErr
	}
}
