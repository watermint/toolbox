package api_rpc_impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_capture"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func New(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_rpc.Caller {

	ri := CallerImpl{
		ctx:        ctx,
		endpoint:   endpoint,
		asMemberId: asMemberId,
		asAdminId:  asAdminId,
		base:       base,
		token:      token,
	}
	return &ri
}

type CallerImpl struct {
	ctx        api_context.Context
	asMemberId string
	asAdminId  string
	base       api_context.PathRoot
	param      interface{}
	token      api_auth.TokenContainer
	endpoint   string
	success    func(res api_rpc.Response) error
	failure    func(err error) error
}

func (z *CallerImpl) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, z.endpoint)
}

func (z *CallerImpl) createRequest() (req api_rpc.Request, err error) {
	url := z.requestUrl()
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))

	headers := make(map[string]string)

	headers[api_rpc.ReqHeaderContentType] = "application/json"
	if z.token.TokenType != api_auth.DropboxTokenNoAuth {
		headers[api_rpc.ReqHeaderAuthorization] = "Bearer " + z.token.Token
	}
	if z.asMemberId != "" {
		headers[api_rpc.ReqHeaderSelectUser] = z.asMemberId
	}
	if z.asAdminId != "" {
		headers[api_rpc.ReqHeaderSelectAdmin] = z.asAdminId
	}
	if z.base != nil {
		pr, err := json.Marshal(z.base)
		if err != nil {
			log.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		headers[api_rpc.ReqHeaderPathRoot] = string(pr)
	}

	return newPostRequest(z.ctx, url, z.param, headers)
}

func (z *CallerImpl) ensureRetryOnError(lastErr error) (res api_rpc.Response, err error) {
	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		if rc.IsNoRetry() {
			z.ctx.Log().Debug("Abort retry due to NoRetryOnError", zap.Error(lastErr))
			return nil, lastErr
		}

		sameErrorCount := 0
		rc.AddError(lastErr)
		for _, e := range rc.LastErrors() {
			if e.Error() == lastErr.Error() {
				sameErrorCount++
			}
		}

		if sameErrorCount >= SameErrorRetryCount {
			z.ctx.Log().Debug(
				"Abort retry due to `same_error_count` exceed threshold",
				zap.Int("same_error_count", sameErrorCount),
				zap.Error(lastErr),
			)
			return nil, lastErr
		}

		after := time.Now().Add(SameErrorRetryWait)
		rc.UpdateRetryAfter(after)
		z.ctx.Log().Debug("Retry after", zap.Error(err), zap.String("retry_after", after.String()))

		return z.Call()

	default:
		z.ctx.Log().Debug("No retry context")
		return z.Call()
	}
}

func (z *CallerImpl) Param(param interface{}) api_rpc.Caller {
	z.param = param
	return z
}

func (z *CallerImpl) OnSuccess(success func(res api_rpc.Response) error) api_rpc.Caller {
	z.success = success
	return z
}

func (z *CallerImpl) OnFailure(failure func(err error) error) api_rpc.Caller {
	z.failure = failure
	return z
}

func (z *CallerImpl) waitForRetryIfRequired() {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))

	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		now := time.Now()
		if !rc.RetryAfter().IsZero() && now.Before(rc.RetryAfter()) {
			log.Debug("Sleep until", zap.String("retry_after", rc.RetryAfter().String()))
			time.Sleep(rc.RetryAfter().Sub(now))
		}
	}
}

func (z *CallerImpl) handleRetryAfterResponse(retryAfterSec int) bool {
	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		after := time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		z.ctx.Log().Debug("Retry after", zap.Int("RetryAfterSec", retryAfterSec))
		rc.UpdateRetryAfter(after)
		z.ctx.Log().Debug("Precaution wait for rate limit", zap.Duration("wait", PrecautionRateLimitWait))
		time.Sleep(PrecautionRateLimitWait)

		return true

	default:
		// do not retry
		return false
	}
}

func (z *CallerImpl) handleResponse(apiResImpl *ResponseImpl) (apiRes api_rpc.Response, err error) {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))
	//if app.Root().IsDebug() {
	//	log.Debug("Response", zap.Int("code", apiResImpl.resStatusCode), zap.String("body", apiResImpl.resBodyString))
	//}

	switch apiResImpl.resStatusCode {
	case http.StatusOK:
		return apiResImpl, nil

	case ErrorBadInputParam: // Bad input param
		log.Debug("Bad input param", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorBadOrExpiredToken: // Bad or expired token
		log.Debug("Bad or expired token", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorAccessError: // Access Error
		log.Debug("Access Error", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseAccessError(apiResImpl.resBodyString)

	case ErrorEndpointSpecific: // Endpoint specific
		log.Debug("Endpoint specific error", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorNoPermission: // No permission
		log.Debug("No Permission", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseAccessError(apiResImpl.resBodyString)

	case ErrorRateLimit: // Rate limit
		retryAfter := apiResImpl.resHeader.Get(api_rpc.ResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			log.Debug("Unable to parse header for RateLimit", zap.String("header", retryAfter), zap.Error(err))
			return nil, errors.New("unknown retry param")
		}

		log.Debug("Rate limit", zap.Int("retryAfterSec", retryAfterSec))
		if z.handleRetryAfterResponse(retryAfterSec) {
			// Retry
			return z.Call()
		} else {
			return nil, api_rpc.ApiErrorRateLimit{RetryAfter: retryAfterSec}
		}
	}

	if int(apiResImpl.resStatusCode/100) == 5 {
		log.Debug("server error", zap.Int("status_code", apiResImpl.resStatusCode), zap.String("body", apiResImpl.resBodyString))
		return z.ensureRetryOnError(api_rpc.ServerError{StatusCode: apiResImpl.resStatusCode})
	}

	log.Warn("Unknown or server error", zap.Int("Code", apiResImpl.resStatusCode), zap.String("Body", apiResImpl.resBodyString))
	return nil, err
}

func (z *CallerImpl) Call() (apiRes api_rpc.Response, err error) {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))
	req, err := z.createRequest()
	if err != nil {
		log.Warn("Unable to prepare HTTP Caller", zap.Error(err))
		return nil, errors.New(fmt.Sprintf("unable to prepare request for [%s]", z.endpoint))
	}

	z.waitForRetryIfRequired()

	switch cc := z.ctx.(type) {
	case api_context.ClientContext:
		log.Debug("Caller", zap.Any("param", z.param), zap.Any("root", z.base))
		apiResImpl := &ResponseImpl{}
		apiResImpl.resStatusCode, apiResImpl.resHeader, apiResImpl.resBody, err = cc.DoRequest(req)
		if apiResImpl.resBody != nil {
			apiResImpl.resBodyString = string(apiResImpl.resBody)
		}

		var cp api_capture.Capture
		if cac, ok := cc.(api_context.CaptureContext); ok {
			cp = api_capture.NewCapture(cac.Capture())
		} else {
			cp = api_capture.Current()
		}
		cp.Rpc(req, apiResImpl, err)

		if err != nil {
			log.Debug("Transport error", zap.Error(err))
			return z.ensureRetryOnError(err)
		}

		return z.handleResponse(apiResImpl)
	}

	log.Warn("No client available")
	return nil, errors.New("no client available")
}
