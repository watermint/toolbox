package api_context_impl

import (
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_async_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_list_impl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_request_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
)

func NewMock(c app_control.Control) api_context.Context {
	return &Mock{l: c.Log()}
}

type Mock struct {
	l *zap.Logger
}

func (z *Mock) Log() *zap.Logger {
	return z.l
}

func (z *Mock) Rpc(endpoint string) api_request.Request {
	return &api_request_impl.MockRequest{}
}

func (z *Mock) Notify(endpoint string) api_request.Request {
	return &api_request_impl.MockRequest{}
}

func (z *Mock) List(endpoint string) api_list.List {
	return &api_list_impl.MockList{}
}

func (z *Mock) Async(endpoint string) api_async.Async {
	return &api_async_impl.MockAsync{}
}

func (z *Mock) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return &api_request_impl.MockRequest{}
}

func (z *Mock) Download(endpoint string) api_request.Request {
	return &api_request_impl.MockRequest{}
}

func (z *Mock) AsMemberId(teamMemberId string) api_context.Context {
	return z
}

func (z *Mock) AsAdminId(teamMemberId string) api_context.Context {
	return z
}

func (z *Mock) WithPath(pathRoot api_context.PathRoot) api_context.Context {
	return z
}

func (z *Mock) NoRetryOnError() api_context.Context {
	return z
}

func (z *Mock) IsNoRetry() bool {
	return false
}

func (z *Mock) Hash() string {
	return ""
}

func (z *Mock) NoAuth() api_context.Context {
	return z
}
