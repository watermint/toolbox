package dbx_response_impl

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"strings"
)

var (
	ErrorBadContentResponse  = errors.New("bad response from server: res_code 400 with html body")
	ErrorInternalServerError = errors.New("internal server error")
)

type ErrorMissingScope struct {
	ErrorSummary  string `json:"error_summary" path:"error_summary"`
	RequiredScope string `json:"required_scope" path:"error.required_scope"`
}

func (z ErrorMissingScope) Error() string {
	return fmt.Sprintf("missing scope [%s]", z.RequiredScope)
}

func AssertResponse(res es_response.Response) es_response.Response {
	l := esl.Default()

	switch res.Code() {
	case dbx_context.DropboxApiErrorBadInputParam:
		// In case of the server returned unexpected HTML response;
		// Response body should be plain text
		if strings.HasPrefix(res.Alt().BodyString(), "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", esl.String("response", res.Alt().BodyString()))
			return es_response_impl.NewTransportErrorResponse(ErrorBadContentResponse, res)
		}

	case dbx_context.DropboxApiErrorBadOrExpiredToken:
		errMissingScope := ErrorMissingScope{}
		if err := res.Alt().Json().Model(&errMissingScope); err != nil {
			l.Debug("The response is not a JSON form. fall back to transport error", esl.Error(err))
			return es_response_impl.NewTransportErrorResponse(ErrorBadContentResponse, res)
		}
		if errMissingScope.RequiredScope != "" {
			l.Warn("Missing scope", esl.String("missingScope", errMissingScope.RequiredScope))
			return es_response_impl.NewTransportErrorResponse(errMissingScope, res)
		}

	case dbx_context.DropboxApiErrorEndpointSpecific:
		if j, err := res.Alt().AsJson(); err != nil {
			dbxErr := &dbx_error.ErrorInfo{}
			if err = j.Model(dbxErr); err != nil {
				dbxErrs := dbx_error.NewErrors(dbxErr)
				switch {
				case dbxErrs.Path().IsTooManyWriteOperations(), dbxErrs.IsTooManyWriteOperations():
					l.Debug("Too many write operations")
					return es_response_impl.NewTransportErrorResponse(nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
				}
			}
		}

	case dbx_context.DropboxApiErrorRateLimit:
		return es_response_impl.NewTransportErrorResponse(nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers()), res)
	}

	// Internal server error
	if res.Code()/100 == 5 {
		return es_response_impl.NewTransportErrorResponse(ErrorInternalServerError, res)
	}

	return res
}
