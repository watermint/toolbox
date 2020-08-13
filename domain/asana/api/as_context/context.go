package as_context

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Context interface {
	api_context.Context
	api_context.Get
	api_context.Post
	api_context.Delete
	api_context.Put

	// Pagination request
	GetWithPagination(endpoint string, offset string, limit int, d ...api_request.RequestDatum) (res es_response.Response)
}
