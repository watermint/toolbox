package gh_context_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_request"
	"github.com/watermint/toolbox/domain/github/api/gh_response"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

const (
	ServerRpc    = "https://api.github.com/"
	ServerUpload = "https://uploads.github.com/"
)

func NewMock(name string, ctl app_control.Control) gh_context.Context {
	client := nw_rest.New(
		nw_rest.Assert(gh_response.AssertResponse),
		nw_rest.Mock())
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: gh_request.NewBuilder(ctl, nil),
	}
}

func New(name string, ctl app_control.Control, token api_auth.Context) gh_context.Context {
	client := nw_rest.New(
		nw_rest.Assert(gh_response.AssertResponse))
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: gh_request.NewBuilder(ctl, token),
	}
}

type ctxImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder gh_request.Builder
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

func (z ctxImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodGet,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Upload(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		ServerUpload+endpoint, // Upload endpoint
		api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPut,
		ServerRpc+endpoint,
		api_request.Combine(d))
	return z.client.Call(&z, b)
}
