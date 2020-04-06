package nw_retry

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_capture"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	retryIntervalSecOnNoRetryAfterParam = 10
)

var (
	retryInstance = retryImpl{}
)

func Call(ctx api_context.Context, req api_request.Request) (res api_response.Response, err error) {
	return retryInstance.Call(ctx, req)
}

type retryImpl struct {
}

func (z *retryImpl) Call(ctx api_context.Context, req api_request.Request) (res api_response.Response, err error) {
	l := ctx.Log().With(
		zap.String("Endpoint", req.Endpoint()),
		zap.String("Routine", ut_runtime.GetGoRoutineName()),
	)

	// Error handling
	retryOnError := func(lastErr error) (res api_response.Response, err error) {
		if ctx.IsNoRetry() {
			l.Debug("Abort retry due to NoRetryOnError", zap.Error(lastErr))
			return nil, lastErr
		}

		// Add lastErr and wait if required
		abort := nw_ratelimit.AddError(ctx.ClientHash(), req.Endpoint(), lastErr)
		if abort {
			l.Debug("Abort retry due to rateLimit", zap.Error(lastErr))
			return nil, lastErr
		}

		l.Debug("Retrying", zap.Error(lastErr))
		return z.Call(ctx, req)
	}

	res, err = nw_capture.Call(ctx, req)
	if err != nil {
		return retryOnError(err)
	}

	// Handle API error
	switch res.StatusCode() {
	case http.StatusOK:
		return res, nil

	case api_response.DropboxApiErrorBadInputParam: // Bad input param
		// In case of the server returned unexpected HTML response
		// Response body should be plain text
		if strings.HasPrefix(res.ResultString(), "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", zap.String("response", res.ResultString()))

			// add error & retry
			nw_ratelimit.AddError(ctx.ClientHash(), req.Endpoint(), errors.New("bad response from server: res_code 400 with html body"))
			return z.Call(ctx, req)
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
		retryAfter := res.Header(api_response.DropboxApiResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			l.Debug("Unable to parse header for RateLimit, retry with predefined interval",
				zap.String("header", retryAfter),
				zap.Error(err),
			)
			retryAfterSec = retryIntervalSecOnNoRetryAfterParam
			// fall through
		}

		after := time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		l.Debug("Rewind after",
			zap.Int("RetryAfterSec", retryAfterSec),
			zap.String("RetryAfter", after.String()),
			zap.Bool("NoRetry", ctx.IsNoRetry()),
		)
		nw_ratelimit.UpdateRetryAfter(ctx.ClientHash(), req.Endpoint(), after)
		return z.Call(ctx, req)
	}

	if int(res.StatusCode()/100) == 5 {
		l.Debug("server error", zap.Int("status_code", res.StatusCode()), zap.String("body", res.ResultString()))
		return retryOnError(dbx_error.ServerError{StatusCode: res.StatusCode()})
	}

	l.Warn("Unknown or server error", zap.Int("Code", res.StatusCode()), zap.String("Result", res.ResultString()))
	return nil, err
}
