package gh_client_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/api/gh_request"
	"github.com/watermint/toolbox/domain/github/api/gh_response"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

const (
	ServerRpc    = "https://api.github.com/"
	ServerUpload = "https://uploads.github.com/"
)

func NewMock(name string, ctl app_control.Control) gh_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Assert(gh_response.AssertResponse),
		nw_rest_factory.Mock())
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: gh_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func New(name string, ctl app_control.Control, entity api_auth.OAuthEntity) gh_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Assert(gh_response.AssertResponse))
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: gh_request.NewBuilder(ctl, entity),
	}
}

type clientImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder gh_request.Builder
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

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Patch(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPatch,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodGet,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Upload(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		ServerUpload+endpoint, // Upload endpoint
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z clientImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPut,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}
