package api_context

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/essentials/http/context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/util/ut_io"
)

type AsyncContext interface {
	Async(endpoint string) dbx_async.Async
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
	context.Context

	Feature() app_feature.Feature

	NoRetryOnError() Context
	IsNoRetry() bool
}
