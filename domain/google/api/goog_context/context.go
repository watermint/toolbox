package goog_context

import "github.com/watermint/toolbox/infra/api/api_context"

type Context interface {
	api_context.Context
	api_context.Get
	api_context.Post
	api_context.Put
	api_context.Delete
}
