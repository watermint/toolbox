package api_context

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
)

var (
	ErrorIncompatibleContextType = errors.New("incompatible context type")
)

type ClientGroup interface {
	ClientHash() string
}

type Context interface {
	ClientGroup

	Log() *zap.Logger

	Rpc(endpoint string) api_request.Request
	Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request
	Download(endpoint string) api_request.Request

	NoRetryOnError() Context
	IsNoRetry() bool
	NoAuth() Context
	MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error)
}

type CaptureContext interface {
	Capture() *zap.Logger
}
