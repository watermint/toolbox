package nw_client

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
)

type Rest interface {
	Call(ctx api_context.Context, req api_request.Request) (res api_response.Response, err error)
}

type Http interface {
	Call()
}
