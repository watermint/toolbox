package dbx_response_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"strings"
)

var (
	ErrorBadContentResponse = errors.New("bad response from server: res_code 400 with html body")
)

func AssertResponse(res es_response.Response) es_response.Response {
	l := es_log.Default()

	switch res.Code() {
	case dbx_context.DropboxApiErrorBadInputParam:
		// In case of the server returned unexpected HTML response
		// Response body should be plain text
		if strings.HasPrefix(res.Alt().BodyString(), "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", es_log.String("response", res.Alt().BodyString()))
			return es_response_impl.NewTransportErrorResponse(ErrorBadContentResponse, res)
		}

	case dbx_context.DropboxApiErrorRateLimit:
		return es_response_impl.NewTransportErrorResponse(nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
	}

	return res
}
