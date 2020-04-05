package dbx_context

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
	"net/http"
)

func NewMock(c app_control.Control) api_context.DropboxApiContext {
	return &Mock{l: c.Log()}
}

type Mock struct {
	l *zap.Logger
}

func (z *Mock) MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error) {
	return nil, qt_errors.ErrorMock
}

func (z *Mock) Log() *zap.Logger {
	return z.l
}

func (z *Mock) Rpc(endpoint string) api_request.Request {
	return &dbx_request.MockRequest{}
}

func (z *Mock) Notify(endpoint string) api_request.Request {
	return &dbx_request.MockRequest{}
}

func (z *Mock) List(endpoint string) api_list.List {
	return &dbx_list.MockList{}
}

func (z *Mock) Async(endpoint string) api_async.Async {
	return &dbx_async.MockAsync{}
}

func (z *Mock) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return &dbx_request.MockRequest{}
}

func (z *Mock) Download(endpoint string) api_request.Request {
	return &dbx_request.MockRequest{}
}

func (z *Mock) AsMemberId(teamMemberId string) api_context.DropboxApiContext {
	return z
}

func (z *Mock) AsAdminId(teamMemberId string) api_context.DropboxApiContext {
	return z
}

func (z *Mock) WithPath(pathRoot api_context.PathRoot) api_context.DropboxApiContext {
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
