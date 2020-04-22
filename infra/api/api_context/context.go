package api_context

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_feature"
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

type AsyncContext interface {
	Async(endpoint string) api_async.Async
}
type ListContext interface {
	List(endpoint string) api_list.List
}
type NotifyContext interface {
	Notify(endpoint string) api_request.Request
}
type UploadContext interface {
	Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request
}
type DownloadContext interface {
	Download(endpoint string) api_request.Request
}
type PostContext interface {
	Post(endpoint string) api_request.Request
}
type GetContext interface {
	Get(endpoint string) api_request.Request
}

type Context interface {
	ClientGroup

	Log() *zap.Logger
	Capture() *zap.Logger
	Feature() app_feature.Feature

	NoRetryOnError() Context
	IsNoRetry() bool
	MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error)
}
