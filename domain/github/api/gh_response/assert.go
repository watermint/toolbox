package gh_response

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/infra/network/nw_retry"
)

func AssertResponse(res es_response.Response) es_response.Response {
	switch res.CodeCategory() {
	case es_response.Code4xxClientErrors:
		erl, found := nw_retry.NewErrorRateLimitFromHeaders(res.Headers())
		if found && erl.Remaining < 1 {
			return es_response_impl.NewTransportErrorResponse(erl, res)
		}
	}
	return res
}
