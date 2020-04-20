package dbx_recovery

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	retryIntervalSecOnNoRetryAfterParam = 10
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

func (z *retryImpl) Call(ctx api_context.Context, req api_request.Request) (res api_response.Response, err error) {
	l := ctx.Log().With(
		zap.String("Endpoint", req.Endpoint()),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	res, err = z.client.Call(ctx, req)

	// Handle API error
	switch res.StatusCode() {
	case http.StatusOK:
		return res, nil

	case api_response.DropboxApiErrorBadInputParam: // Bad input param
		// In case of the server returned unexpected HTML response
		// Response body should be plain text
		if strings.HasPrefix(res.ResultString(), "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", zap.String("response", res.ResultString()))
			return nil, ErrorBadContentResponse
		}
		l.Debug("Bad input param", zap.String("Error", res.ResultString()))
		return nil, dbx_error.ParseApiError(res.ResultString())

	case api_response.DropboxApiErrorBadOrExpiredToken: // Bad or expired token
		l.Debug("Bad or expired token", zap.String("Error", res.ResultString()))
		return nil, dbx_error.ParseApiError(res.ResultString())

	case api_response.DropboxApiErrorAccessError: // Access Error
		l.Debug("Access Error", zap.String("Error", res.ResultString()))
		return nil, dbx_error.ParseAccessError(res.ResultString())

	case api_response.DropboxApiErrorEndpointSpecific: // Endpoint specific
		l.Debug("Endpoint specific error", zap.String("Error", res.ResultString()))
		return nil, dbx_error.ParseApiError(res.ResultString())

	case api_response.DropboxApiErrorNoPermission: // No permission
		l.Debug("No Permission", zap.String("Error", res.ResultString()))
		return nil, dbx_error.ParseAccessError(res.ResultString())

	case api_response.DropboxApiErrorRateLimit: // Rate limit
		return nil, nw_retry.NewErrorRateLimitFromHeadersFallback(res.Headers())
	}

	if int(res.StatusCode()/100) == 5 {
		l.Debug("server error", zap.Int("status_code", res.StatusCode()), zap.String("body", res.ResultString()))
		return nil, err
	}

	l.Warn("Unknown or server error", zap.Int("Code", res.StatusCode()), zap.String("Result", res.ResultString()))
	return nil, err
}
