package dbx_recovery

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

var (
	ErrorBadContentResponse = errors.New("bad response from server: res_code 400 with html body")
)

func New(client nw_client.Rest) nw_client.Rest {
	return &retryImpl{client: client}
}

type retryImpl struct {
	client nw_client.Rest
}

func (z *retryImpl) Call(ctx api_context.Context, req api_request.Request) (res response.Response, err error) {
	l := ctx.Log().With(
		zap.String("Endpoint", req.Endpoint()),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	res, err = z.client.Call(ctx, req)
	if err != nil {
		return nil, err
	}

	ll := l.With(zap.Int("statusCode", res.Code()))
	if res.Code() == http.StatusOK {
		return res, nil
	}
	bodyString := res.Success().BodyString()

	// Handle API error
	switch res.Code() {
	case dbx_context.DropboxApiErrorBadInputParam: // Bad input param
		// In case of the server returned unexpected HTML response
		// Response body should be plain text
		if strings.HasPrefix(bodyString, "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", zap.String("response", bodyString))
			return nil, ErrorBadContentResponse
		}
		ll.Debug("Bad input param", zap.String("Error", bodyString))
		return nil, dbx_error.ParseApiError(bodyString)

	case dbx_context.DropboxApiErrorBadOrExpiredToken: // Bad or expired token
		ll.Debug("Bad or expired token", zap.String("Error", bodyString))
		return nil, dbx_error.ParseApiError(bodyString)

	case dbx_context.DropboxApiErrorAccessError: // Access Error
		ll.Debug("Access Error", zap.String("Error", bodyString))
		return nil, dbx_error.ParseAccessError(bodyString)

	case dbx_context.DropboxApiErrorEndpointSpecific: // Endpoint specific
		ll.Debug("Endpoint specific error", zap.String("Error", bodyString))
		return nil, dbx_error.ParseApiError(bodyString)

	case dbx_context.DropboxApiErrorNoPermission: // No permission
		ll.Debug("No Permission", zap.String("Error", bodyString))
		return nil, dbx_error.ParseAccessError(bodyString)

	case dbx_context.DropboxApiErrorRateLimit: // Rate limit
		return nil, nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers())
	}

	if res.CodeCategory() == response.Code5xxServerErrors {
		ll.Debug("server error", zap.String("body", bodyString))
		return nil, err
	}

	ll.Warn("Unknown or server error", zap.String("Result", bodyString))
	return nil, err
}
