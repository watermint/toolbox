package work_context_impl

import (
	"context"
	"github.com/watermint/toolbox/domain/slack/api/work_context"
	"github.com/watermint/toolbox/domain/slack/api/work_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

func NewMock(name string, ctl app_control.Control) work_context.Context {
	client := nw_rest.New(nw_rest.Mock())
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: work_request.New(ctl, nil),
	}
}

func New(name string, ctl app_control.Control, token api_auth.OAuthContext) work_context.Context {
	client := nw_rest.New(
		nw_rest.Client(token.Config().Client(context.Background(), token.Token())),
		nw_rest.Assert(api_response.AssertResponse),
	)

	return &ctxImpl{
		name:    name,
		client:  nw_retry.NewRetry(nw_retry.NewRatelimit(client)),
		ctl:     ctl,
		builder: work_request.New(ctl, token),
	}
}

const (
	Endpoint = "https://slack.com/api/"
)

type ctxImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder work_request.Builder
}

func (z ctxImpl) Name() string {
	return z.name
}

func (z ctxImpl) ClientHash() string {
	return z.builder.ClientHash()
}

func (z ctxImpl) Log() esl.Logger {
	return z.builder.Log()
}

func (z ctxImpl) Capture() esl.Logger {
	return z.ctl.Capture()
}

func (z ctxImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodGet, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPost, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}
