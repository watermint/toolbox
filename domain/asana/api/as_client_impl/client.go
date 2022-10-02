package as_client_impl

import (
	"github.com/watermint/toolbox/domain/asana/api/as_auth"
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/domain/asana/api/as_request"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

func NewMock(name string, ctl app_control.Control) as_client.Client {
	client := nw_rest_factory.New(nw_rest_factory.Mock())
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func NewReplayMock(name string, ctl app_control.Control, rr []nw_replay.Response) as_client.Client {
	client := nw_rest_factory.New(nw_rest_factory.ReplayMock(rr))
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func New(name string, ctl app_control.Control, entity api_auth.OAuthEntity) as_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.OAuthEntity(as_auth.Asana, func(appKey string) (clientId, clientSecret string) {
			return app_apikey.Resolve(ctl, appKey)
		}, entity),
		nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
			return nw_auth.NewOAuthRestClient(entity, ctl.AuthRepository(), client)
		}),
	)
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, entity),
	}
}

const (
	Endpoint = "https://app.asana.com/api/1.0/"
)

type clientImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder as_request.Builder
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

func (z clientImpl) GetWithPagination(endpoint string, offset string, limit int, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodGet, Endpoint+endpoint, api_request.Combine(d)).WithOffset(limit, offset)
	return z.client.Call(&z, b)
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPost, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodDelete, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPut, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}
