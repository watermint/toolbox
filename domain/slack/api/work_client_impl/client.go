package work_client_impl

import (
	"github.com/watermint/toolbox/domain/slack/api/work_auth"
	"github.com/watermint/toolbox/domain/slack/api/work_client"
	"github.com/watermint/toolbox/domain/slack/api/work_request"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/api/api_response"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

func NewMock(name string, ctl app_control.Control) work_client.Client {
	client := nw_rest_factory.New(nw_rest_factory.Mock())
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: work_request.New(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func New(name string, ctl app_control.Control, entity api_auth.OAuthEntity) work_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.OAuthEntity(work_auth.Slack, func(appKey string) (clientId, clientSecret string) {
			return app_apikey.Resolve(ctl, appKey)
		}, entity),
		nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
			return nw_auth.NewOAuthRestClient(entity, ctl.AuthRepository(), client)
		}),
		nw_rest_factory.Assert(api_response.AssertResponse),
	)

	return &clientImpl{
		name:    name,
		client:  nw_retry.NewRetry(nw_retry.NewRatelimit(client)),
		ctl:     ctl,
		builder: work_request.New(ctl, entity),
	}
}

const (
	Endpoint = "https://slack.com/api/"
)

type clientImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder work_request.Builder
}

func (z clientImpl) Name() string {
	return z.name
}

func (z clientImpl) ClientHash() string {
	return z.builder.ClientHash()
}

func (z clientImpl) Log() esl.Logger {
	return z.builder.Log()
}

func (z clientImpl) Capture() esl.Logger {
	return z.ctl.Capture()
}

func (z clientImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodGet, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPost, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}
