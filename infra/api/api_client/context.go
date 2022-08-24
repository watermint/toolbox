package api_client

import (
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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
type Put interface {
	Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Patch interface {
	Patch(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type Delete interface {
	Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response)
}
type UI interface {
	UI() app_ui.UI
}

type Client interface {
	es_context.Context
}

type QualityContext interface {
	NoRetryOnError() bool
}
