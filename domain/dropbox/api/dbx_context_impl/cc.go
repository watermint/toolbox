package dbx_context_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"strconv"
)

func New(ctl app_control.Control, token api_auth.Context) dbx_context.Context {
	c := &ccImpl{
		ctl:            ctl,
		token:          token,
		noRetryOnError: false,
	}
	return c
}

type ccImpl struct {
	ctl            app_control.Control
	token          api_auth.Context
	asMemberId     string
	asAdminId      string
	basePath       dbx_context.PathRoot
	noRetryOnError bool
	hashComputed   string
}

func (z ccImpl) Feature() app_feature.Feature {
	return z.ctl.Feature()
}

func (z ccImpl) Capture() *zap.Logger {
	return z.ctl.Capture()
}

func (z ccImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z ccImpl) Log() *zap.Logger {
	return z.ctl.Log()
}

func (z ccImpl) Post(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		&z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
		dbx_request.RpcEndpoint,
	)
}

func (z ccImpl) Notify(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		&z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
		dbx_request.NotifyEndpoint,
	)
}

func (z ccImpl) List(endpoint string) api_list.List {
	return dbx_list.New(&z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z ccImpl) Async(endpoint string) dbx_async.Async {
	return dbx_async_impl.New(&z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z ccImpl) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return dbx_request.NewUploadRequest(
		&z,
		endpoint,
		content,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
	)
}

func (z ccImpl) Download(endpoint string) api_request.Request {
	return dbx_request.NewDownloadRequest(
		&z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
	)
}

func (z ccImpl) AsMemberId(teamMemberId string) dbx_context.Context {
	z.asMemberId = teamMemberId
	z.asAdminId = ""
	return &z
}

func (z ccImpl) AsAdminId(teamMemberId string) dbx_context.Context {
	z.asMemberId = ""
	z.asAdminId = teamMemberId
	return &z
}

func (z ccImpl) WithPath(pathRoot dbx_context.PathRoot) dbx_context.Context {
	z.basePath = pathRoot
	return &z
}

func (z ccImpl) NoRetryOnError() api_context.Context {
	z.noRetryOnError = true
	return &z
}

func (z ccImpl) ClientHash() string {
	seeds := []string{
		"m", z.asMemberId,
		"a", z.asAdminId,
		"p", z.token.PeerName(),
		"t", z.token.Token().AccessToken,
		"y", z.token.Scope(),
		"n", strconv.FormatBool(z.noRetryOnError),
	}
	if z.basePath != nil {
		seeds = append(seeds, z.basePath.Header())
	}
	return ClientHash(seeds)
}
