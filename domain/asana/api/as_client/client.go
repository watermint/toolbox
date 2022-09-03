package as_client

import (
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

type Client interface {
	api_client.Client
	api_client.Get
	api_client.Post
	api_client.Delete
	api_client.Put

	// Pagination request
	GetWithPagination(endpoint string, offset string, limit int, d ...api_request.RequestDatum) (res es_response.Response)
}
