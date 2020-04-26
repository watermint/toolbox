package gh_response

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/essentials/http/response_impl"
	"github.com/watermint/toolbox/infra/network/nw_retry"
)

func AssertResponse(res response.Response) response.Response {
	switch res.CodeCategory() {
	case response.Code4xxClientErrors:
		erl, found := nw_retry.NewErrorRateLimitFromHeaders(res.Headers())
		if found && erl.Remaining < 1 {
			return response_impl.NewTransportErrorResponse(erl, res)
		}
	}
	return res
}
