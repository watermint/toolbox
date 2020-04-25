package gh_recovery

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
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

func (z recoveryImpl) Call(ctx api_context.Context, req api_request.Request) (res response.Response, err error) {
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
	case res.CodeCategory() == response.Code2xxSuccess:
		return res, nil

	default:
		// Rate limit
		erl, found := nw_retry.NewErrorRateLimitFromHeaders(res.Headers())
		if found && erl.Remaining < 1 {
			return nil, erl
		}
		// General errors
		apiErr := &gh_context.ApiError{}

		if j, err := res.Success().AsJson(); err != nil {
			l.Debug("response body is not a json", zap.Error(err))
			return nil, ErrorUnexpectedResponseFormat
		} else if _, err := j.Model(apiErr); err != nil {
			l.Debug("unexpected error json format", zap.Error(err))
			return nil, ErrorUnexpectedResponseFormat
		} else {
			return nil, apiErr
		}
	}
}
