package gh_context

import "github.com/watermint/toolbox/infra/api/api_context"

type Context interface {
	api_context.Context
	api_context.PostContext
	api_context.GetContext
	api_context.UploadContext
}
