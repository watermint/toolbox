package api_response

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
)

func AssertResponse(res es_response.Response) es_response.Response {
	l := esl.Default()

	switch res.Code() {
	case dbx_context.DropboxApiErrorRateLimit:
		l.Debug("Rate limit", esl.Int("code", res.Code()))
		return es_response_impl.NewTransportErrorResponse(nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
	}
	return res
}
