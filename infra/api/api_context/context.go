package api_context

import (
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Notify interface {
	Notify(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Upload interface {
	Upload(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Download interface {
	Download(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Post interface {
	Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Get interface {
	Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}

type Context interface {
	es_context.Context
}
