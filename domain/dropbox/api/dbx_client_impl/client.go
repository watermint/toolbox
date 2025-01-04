package dbx_client_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/essentials/network/nw_simulator"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"net/http"
)

func NewMock(name string, ctl app_control.Control) dbx_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Assert(dbx_response_impl.AssertResponse),
		nw_rest_factory.Mock())
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity(), dbx_filesystem.NewEmptyHelper()),
	}
}

func NewSeqReplayMock(name string, ctl app_control.Control, rr []nw_replay.Response) dbx_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Assert(dbx_response_impl.AssertResponse),
		nw_rest_factory.ReplayMock(rr))
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity(), dbx_filesystem.NewEmptyHelper()),
	}
}

func NewReplayMock(name string, ctl app_control.Control, replay kv_storage.Storage) dbx_client.Client {
	client := nw_replay.NewHashReplay(replay)
	return &clientImpl{
		name:    name,
		client:  client,
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity(), dbx_filesystem.NewEmptyHelper()),
	}
}

func newClientOpts(feature app_feature.Feature, l esl.Logger) (opts []nw_rest_factory.ClientOpt) {
	opts = make([]nw_rest_factory.ClientOpt, 0)
	opts = append(opts, nw_rest_factory.Assert(dbx_response_impl.AssertResponse))

	// too many requests error simulator
	if feature.Experiment(app_definitions.ExperimentDbxClientConditionerNarrow20) {
		l.Debug("Experiment: Network conditioner enabled: 20%")
		opts = append(opts, nw_rest_factory.RateLimitSimulator(20, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	} else if feature.Experiment(app_definitions.ExperimentDbxClientConditionerNarrow40) {
		l.Debug("Experiment: Network conditioner enabled: 40%")
		opts = append(opts, nw_rest_factory.RateLimitSimulator(40, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	} else if feature.Experiment(app_definitions.ExperimentDbxClientConditionerNarrow100) {
		l.Debug("Experiment: Network conditioner enabled: 100%")
		opts = append(opts, nw_rest_factory.RateLimitSimulator(100, nw_simulator.RetryAfterHeaderRetryAfter, decorateRateLimit))
	}

	// server error simulator
	if feature.Experiment(app_definitions.ExperimentDbxClientConditionerError20) {
		l.Debug("Experiment: Network conditioner enabled: 20%")
		opts = append(opts, nw_rest_factory.ServerErrorSimulator(20, http.StatusInternalServerError, decorateServerError))
	} else if feature.Experiment(app_definitions.ExperimentDbxClientConditionerError40) {
		l.Debug("Experiment: Network conditioner enabled: 40%")
		opts = append(opts, nw_rest_factory.ServerErrorSimulator(40, http.StatusInternalServerError, decorateServerError))
	} else if feature.Experiment(app_definitions.ExperimentDbxClientConditionerError100) {
		l.Debug("Experiment: Network conditioner enabled: 100%")
		opts = append(opts, nw_rest_factory.ServerErrorSimulator(100, http.StatusInternalServerError, decorateServerError))
	}

	return opts
}

func newClientWithToken(ctl app_control.Control, l esl.Logger, app api_auth.OAuthAppData, entity api_auth.OAuthEntity) nw_client.Rest {
	opts := newClientOpts(ctl.Feature(), l)
	opts = append(opts, nw_rest_factory.OAuthEntity(app, func(appKey string) (clientId, clientSecret string) {
		return app_apikey.Resolve(ctl, app.AppKeyName)
	}, entity))
	opts = append(opts, nw_rest_factory.Auth(func(client nw_client.Rest) nw_client.Rest {
		return nw_auth.NewOAuthRestClient(entity, ctl.AuthRepository(), client)
	}))
	opts = append(opts)
	return nw_rest_factory.New(opts...)
}

func newClientNoAuth(feature app_feature.Feature, l esl.Logger) nw_client.Rest {
	opts := newClientOpts(feature, l)
	return nw_rest_factory.New(opts...)
}

func New(ctl app_control.Control, app api_auth.OAuthAppData, entity api_auth.OAuthEntity, helper *dbx_filesystem.FileSystemBuilderHelper) dbx_client.Client {
	return &clientImpl{
		name:    entity.PeerName,
		client:  newClientWithToken(ctl, ctl.Log(), app, entity),
		ctl:     ctl,
		builder: dbx_request.NewBuilder(ctl, entity, helper),
	}
}

func decorateRateLimit(endpoint string, res *http.Response) {
}

func decorateServerError(endpoint string, res *http.Response) {
}

type clientImpl struct {
	name    string
	client  nw_client.Rest
	ctl     app_control.Control
	builder dbx_request.Builder
	noRetry bool
}

func (z clientImpl) Name() string {
	return z.name
}

func (z clientImpl) Feature() app_feature.Feature {
	return z.ctl.Feature()
}

func (z clientImpl) NoRetryOnError() bool {
	return z.noRetry
}

func (z clientImpl) NoRetry() dbx_client.Client {
	z.noRetry = true
	return z
}

func (z clientImpl) UI() app_ui.UI {
	return z.ctl.UI()
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

func (z clientImpl) Async(endpoint string, d ...api_request.RequestDatum) dbx_async.Async {
	return dbx_async_impl.New(&z, endpoint, d)
}

func (z clientImpl) List(endpoint string, d ...api_request.RequestDatum) dbx_list.List {
	return dbx_list_impl.New(&z, endpoint, d)
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		RpcRequestUrl(RpcEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Upload(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		ContentRequestUrl(endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Download(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		ContentRequestUrl(endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) DownloadRPC(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		RpcRequestUrl(RpcEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Notify(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodPost,
		RpcRequestUrl(NotifyEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) ContentHead(endpoint string, d ...api_request.RequestDatum) dbx_response.Response {
	b := z.builder.With(
		http.MethodHead,
		RpcRequestUrl(ContentEndpoint, endpoint),
		api_request.Combine(d),
	)
	return dbx_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) AsMemberId(teamMemberId string) dbx_client.Client {
	z.builder = z.builder.AsMemberId(teamMemberId)
	return z
}

func (z clientImpl) AsAdminId(teamMemberId string) dbx_client.Client {
	z.builder = z.builder.AsAdminId(teamMemberId)
	return z
}

func (z clientImpl) WithPath(pathRoot dbx_client.PathRoot) dbx_client.Client {
	z.builder = z.builder.WithPath(pathRoot)
	return z
}

func (z clientImpl) NoAuth() dbx_client.Client {
	z.builder = z.builder.NoAuth()
	z.client = newClientNoAuth(z.ctl.Feature(), z.ctl.Log())
	return z
}
