package dbx_context_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
)

func NewMock(c app_control.Control) dbx_context.Context {
	return &Mock{l: c.Log(), c: c.Capture(), feature: c.Feature()}
}

type Mock struct {
	l       *zap.Logger
	c       *zap.Logger
	feature app_feature.Feature
}

func (z *Mock) Feature() app_feature.Feature {
	return z.feature
}

func (z *Mock) Capture() *zap.Logger {
	return z.c
}

func (z *Mock) Log() *zap.Logger {
	return z.l
}

func (z *Mock) Post(endpoint string) api_request.Request {
	return &api_request.MockRequest{}
}

func (z *Mock) Notify(endpoint string) api_request.Request {
	return &api_request.MockRequest{}
}

func (z *Mock) List(endpoint string) api_list.List {
	return &dbx_list.MockList{}
}

func (z *Mock) Async(endpoint string) dbx_async.Async {
	return &dbx_async_impl.MockAsync{}
}

func (z *Mock) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return &api_request.MockRequest{}
}

func (z *Mock) Download(endpoint string) api_request.Request {
	return &api_request.MockRequest{}
}

func (z *Mock) AsMemberId(teamMemberId string) dbx_context.Context {
	return z
}

func (z *Mock) AsAdminId(teamMemberId string) dbx_context.Context {
	return z
}

func (z *Mock) WithPath(pathRoot dbx_context.PathRoot) dbx_context.Context {
	return z
}

func (z *Mock) NoRetryOnError() api_context.Context {
	return z
}

func (z *Mock) IsNoRetry() bool {
	return false
}

func (z *Mock) ClientHash() string {
	return ""
}

func (z *Mock) NoAuth() api_context.Context {
	return z
}
