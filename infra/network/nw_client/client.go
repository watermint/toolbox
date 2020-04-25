package nw_client

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"net/http"
	"time"
)

type Rest interface {
	Call(ctx api_context.Context, req api_request.Request) (res response.Response, err error)
}

type Http interface {
	Call(clientHash string, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error)
}
