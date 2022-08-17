package dbx_context_impl

import (
	"context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest"
	"github.com/watermint/toolbox/essentials/network/nw_simulator"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"net/http"
)

func NewMock(name string, ctl app_control.Control) dbx_context.Context {
	client := nw_rest.New(
		nw_rest.Assert(dbx_response_impl.AssertResponse),
		nw_rest.Mock())
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, nil),
	}
}

func NewSeqReplayMock(name string, ctl app_control.Control, rr []nw_replay.Response) dbx_context.Context {
	client := nw_rest.New(
		nw_rest.Assert(dbx_response_impl.AssertResponse),
		nw_rest.ReplayMock(rr))
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, nil),
	}
}

func NewReplayMock(name string, ctl app_control.Control, replay kv_storage.Storage) dbx_context.Context {
	client := nw_replay.NewHashReplay(replay)
	return &ctxImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, nil),
	}
}

func newClientOpts(feature app_feature.Feature, l esl.Logger) (opts []nw_rest.ClientOpt) {
	opts = make([]nw_rest.ClientOpt, 0)
	opts = append(opts, nw_rest.Assert(dbx_response_impl.AssertResponse))

	// too many requests error simulator
	if feature.Experiment(app.ExperimentDbxClientConditionerNarrow20) {
		l.Debug("Experiment: Network conditioner enabled: 20%")
		opts = append(opts, nw_rest.RateLimitSimulator(20, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	} else if feature.Experiment(app.ExperimentDbxClientConditionerNarrow40) {
		l.Debug("Experiment: Network conditioner enabled: 40%")
		opts = append(opts, nw_rest.RateLimitSimulator(40, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	} else if feature.Experiment(app.ExperimentDbxClientConditionerNarrow100) {
		l.Debug("Experiment: Network conditioner enabled: 100%")
		opts = append(opts, nw_rest.RateLimitSimulator(100, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	}

	// server error simulator
	if feature.Experiment(app.ExperimentDbxClientConditionerError20) {
		l.Debug("Experiment: Network conditioner enabled: 20%")
		opts = append(opts, nw_rest.ServerErrorSimulator(20, http.StatusInternalServerError, decorateServerError))
	} else if feature.Experiment(app.ExperimentDbxClientConditionerError40) {
		l.Debug("Experiment: Network conditioner enabled: 40%")
		opts = append(opts, nw_rest.ServerErrorSimulator(40, http.StatusInternalServerError, decorateServerError))
	} else if feature.Experiment(app.ExperimentDbxClientConditionerError100) {
		l.Debug("Experiment: Network conditioner enabled: 100%")
		opts = append(opts, nw_rest.ServerErrorSimulator(100, http.StatusInternalServerError, decorateServerError))
	}

	return opts
}

func newClientWithToken(feature app_feature.Feature, l esl.Logger, token api_auth.OAuthContext) nw_client.Rest {
	opts := newClientOpts(feature, l)
	opts = append(opts, nw_rest.Client(token.Config().Client(context.Background(), token.Token())))
	opts = append(opts)
	return nw_rest.New(opts...)
}

func newClientNoAuth(feature app_feature.Feature, l esl.Logger) nw_client.Rest {
	opts := newClientOpts(feature, l)
	opts = append(opts, nw_rest.Client(&http.Client{}))
	return nw_rest.New(opts...)
}

func New(name string, ctl app_control.Control, token api_auth.OAuthContext) dbx_context.Context {
	return &ctxImpl{
		name:    name,
		client:  newClientWithToken(ctl.Feature(), ctl.Log(), token),
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, token),
	}
}

func decorateRateLimit(endpoint string, res *http.Response) {
}

func decorateServerError(endpoint string, res *http.Response) {
}

type ctxImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder dbx_request.Builder
	noRetry bool
}

func (z ctxImpl) Name() string {
	return z.name
}

func (z ctxImpl) Feature() app_feature.Feature {
	return z.ctl.Feature()
}

func (z ctxImpl) NoRetryOnError() bool {
	return z.noRetry
}

func (z ctxImpl) NoRetry() dbx_context.Context {
	z.noRetry = true
	return z
}

func (z ctxImpl) UI() app_ui.UI {
	return z.ctl.UI()
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

func (z ctxImpl) Async(endpoint string, d ...api_request.RequestDatum) dbx_async.Async {
	return dbx_async_impl.New(&z, endpoint, d)
}

func (z ctxImpl) List(endpoint string, d ...api_request.RequestDatum) dbx_list.List {
	return dbx_list_impl.New(&z, endpoint, d)
}

func (z ctxImpl) Post(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		RpcRequestUrl(RpcEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Upload(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		ContentRequestUrl(endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Download(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		ContentRequestUrl(endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) Notify(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		RpcRequestUrl(NotifyEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) ContentHead(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodHead,
		RpcRequestUrl(ContentEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z ctxImpl) AsMemberId(teamMemberId string) dbx_context.Context {
	z.builder = z.builder.AsMemberId(teamMemberId)
	return z
}

func (z ctxImpl) AsAdminId(teamMemberId string) dbx_context.Context {
	z.builder = z.builder.AsAdminId(teamMemberId)
	return z
}

func (z ctxImpl) WithPath(pathRoot dbx_context.PathRoot) dbx_context.Context {
	z.builder = z.builder.WithPath(pathRoot)
	return z
}

func (z ctxImpl) NoAuth() dbx_context.Context {
	z.builder = z.builder.NoAuth()
	z.client = newClientNoAuth(z.ctl.Feature(), z.ctl.Log())
	return z
}
