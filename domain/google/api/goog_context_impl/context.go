package goog_context_impl

import (
	"context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response_impl"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/api/goog_request"
	"github.com/watermint/toolbox/domain/google/api/goog_response_impl"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"net/http"
)

const (
	ApiEndpoint = "https://www.googleapis.com/"
)

func NewMock(ctl app_control.Control) goog_context.Context {
	client := nw_rest.New(
		nw_rest.Mock())
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: goog_request.NewBuilder(ctl, nil),
	}
}

func NewReplayMock(ctl app_control.Control, rr []nw_replay.Response) goog_context.Context {
	client := nw_rest.New(
		nw_rest.Assert(dbx_response_impl.AssertResponse),
		nw_rest.ReplayMock(rr))
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: goog_request.NewBuilder(ctl, nil),
	}
}

func New(ctl app_control.Control, token api_auth.Context) goog_context.Context {
	client := nw_rest.New(
		nw_rest.Client(token.Config().Client(context.Background(), token.Token())),
	)
	return &ctxImpl{
		client:  client,
		ctl:     ctl,
		builder: goog_request.NewBuilder(ctl, token),
	}
}

type ctxImpl struct {
	builder goog_request.Builder
	client  nw_client.Rest
	ctl     app_control.Control
}

func (z ctxImpl) UI() app_ui.UI {
	return z.ctl.UI()
}

func (z ctxImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodGet,
		ApiEndpoint+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		ApiEndpoint+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPut,
		ApiEndpoint+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodDelete,
		ApiEndpoint+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
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
