package gh_context_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
	"net/http"
)

func NewMock(ctl app_control.Control) gh_context.Context {
	return &Mock{}
}

type Mock struct {
	ctl app_control.Control
}

func (z Mock) Feature() app_feature.Feature {
	return z.ctl.Feature()
}

func (z Mock) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return &api_request.MockRequest{}
}

func (z Mock) ClientHash() string {
	return ""
}

func (z Mock) Log() *zap.Logger {
	return app_root.Log()
}

func (z Mock) Capture() *zap.Logger {
	return app_root.Capture()
}

func (z Mock) NoRetryOnError() api_context.Context {
	return &z
}

func (z Mock) IsNoRetry() bool {
	return false
}

func (z Mock) MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error) {
	return nil, qt_errors.ErrorMock
}

func (z Mock) Post(endpoint string) api_request.Request {
	return &api_request.MockRequest{}
}

func (z Mock) Get(endpoint string) api_request.Request {
	return &api_request.MockRequest{}
}
