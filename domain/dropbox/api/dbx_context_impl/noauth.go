package dbx_context_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var (
	noAuthImplSeed = atomic.Int64{}
)

func NewNoAuth(feature app_feature.Feature) dbx_context.NoAuthContext {
	return &noAuthImpl{seed: noAuthImplSeed.Add(1), noRetry: false, feature: feature}
}

type noAuthImpl struct {
	seed    int64
	noRetry bool
	feature app_feature.Feature
}

func (z *noAuthImpl) Feature() app_feature.Feature {
	return z.feature
}

func (z *noAuthImpl) ClientHash() string {
	return ClientHash([]string{
		"s", strconv.Itoa(int(z.seed)),
		"n", strconv.FormatBool(z.noRetry),
	})
}

func (z *noAuthImpl) Capture() *zap.Logger {
	return app_root.Capture()
}

func (z *noAuthImpl) Log() *zap.Logger {
	return app_root.Log()
}

func (z *noAuthImpl) Post(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		z,
		endpoint,
		"",
		"",
		nil,
		api_auth.NewNoAuth(),
		dbx_request.NotifyEndpoint,
	)
}

func (z *noAuthImpl) Notify(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		z,
		endpoint,
		"",
		"",
		nil,
		api_auth.NewNoAuth(),
		dbx_request.NotifyEndpoint,
	)
}

func (z *noAuthImpl) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return dbx_request.NewUploadRequest(
		z,
		endpoint,
		content,
		"",
		"",
		nil,
		api_auth.NewNoAuth(),
	)
}

func (z *noAuthImpl) Download(endpoint string) api_request.Request {
	return dbx_request.NewDownloadRequest(
		z,
		endpoint,
		"",
		"",
		nil,
		api_auth.NewNoAuth(),
	)
}

func (z *noAuthImpl) NoRetryOnError() api_context.Context {
	return &noAuthImpl{seed: noAuthImplSeed.Add(1), noRetry: true}
}

func (z *noAuthImpl) IsNoRetry() bool {
	return z.noRetry
}

func (z *noAuthImpl) MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error) {
	return NewResponse(z, req, res)
}
