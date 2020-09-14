package as_context_impl

import (
	"context"
	"github.com/watermint/toolbox/domain/asana/api/as_context"
	"github.com/watermint/toolbox/domain/asana/api/as_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

func NewMock(ctl app_control.Control) as_context.Context {
	client := nw_rest.New(nw_rest.Mock())
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, nil),
	}
}

func NewReplayMock(ctl app_control.Control, rr []nw_replay.Response) as_context.Context {
	client := nw_rest.New(nw_rest.ReplayMock(rr))
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, nil),
	}
}

func New(ctl app_control.Control, token api_auth.Context) as_context.Context {
	client := nw_rest.New(
		nw_rest.Client(token.Config().Client(context.Background(), token.Token())),
	)
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: as_request.NewBuilder(ctl, token),
	}
}

const (
	Endpoint = "https://app.asana.com/api/1.0/"
)

type ctxImpl struct {
	client  nw_client.Rest
	ctl     app_control.Control
	builder as_request.Builder
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

func (z ctxImpl) GetWithPagination(endpoint string, offset string, limit int, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodGet, Endpoint+endpoint, api_request.Combine(d)).WithOffset(limit, offset)
	return z.client.Call(&z, b)
}

func (z ctxImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPost, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodDelete, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}

func (z ctxImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(http.MethodPut, Endpoint+endpoint, api_request.Combine(d))
	return z.client.Call(&z, b)
}
