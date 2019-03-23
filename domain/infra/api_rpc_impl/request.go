package api_rpc_impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func New(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_rpc.Request {

	ri := RequestImpl{
		ctx:        ctx,
		endpoint:   endpoint,
		asMemberId: asMemberId,
		asAdminId:  asAdminId,
		base:       base,
		token:      token,
	}
	return &ri
}

type RequestImpl struct {
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

func (z *RequestImpl) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, z.endpoint)
}

func (z *RequestImpl) httpRequest() (req *http.Request, err error) {
	url := z.requestUrl()
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))

	// param
	requestParam, err := json.Marshal(z.param)
	if err != nil {
		log.Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}
	req, err = http.NewRequest("POST", url, bytes.NewReader(requestParam))
	if err != nil {
		log.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if z.token.TokenType != api_auth.DropboxTokenNoAuth {
		req.Header.Add("Authorization", "Bearer "+z.token.Token)
	}
	if z.asMemberId != "" {
		req.Header.Add(ReqHeaderSelectUser, z.asMemberId)
	}
	if z.asAdminId != "" {
		req.Header.Add(ReqHeaderSelectAdmin, z.asAdminId)
	}
	if z.base != nil {
		pr, err := json.Marshal(z.base)
		if err != nil {
			log.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		req.Header.Add(ReqHeaderPathRoot, string(pr))
	}
	return
}

func (z *RequestImpl) ensureRetryOnError(lastErr error) (res api_rpc.Response, err error) {
	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		sameErrorCount := 0
		rc.AddError(err)
		for _, e := range rc.LastErrors() {
			if e.Error() == lastErr.Error() {
				sameErrorCount++
			}
		}

		if sameErrorCount >= SameErrorRetryCount {
			z.ctx.Log().Debug(
				"Abort retry due to `same_error_count` exceed threshold",
				zap.Int("same_error_count", sameErrorCount),
				zap.Error(err),
			)
			return nil, err
		}

		rc.UpdateRetryAfter(time.Now().Add(SameErrorRetryWait))
		z.ctx.Log().Debug("Retry after", zap.Error(err), zap.Time("retry_after", rc.RetryAfter()))

		return z.Call()

	default:
		z.ctx.Log().Debug("No retry context")
		return z.Call()
	}
}

func (z *RequestImpl) Param(param interface{}) api_rpc.Request {
	z.param = param
	return z
}

func (z *RequestImpl) OnSuccess(success func(res api_rpc.Response) error) api_rpc.Request {
	z.success = success
	return z
}

func (z *RequestImpl) OnFailure(failure func(err error) error) api_rpc.Request {
	z.failure = failure
	return z
}

func (z *RequestImpl) waitForRetryIfRequired() {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))

	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		now := time.Now()
		if !rc.RetryAfter().IsZero() && now.Before(rc.RetryAfter()) {
			log.Debug("Sleep until", zap.Time("retry_after", rc.RetryAfter()))
			time.Sleep(rc.RetryAfter().Sub(now))
		}
	}
}

func (z *RequestImpl) handleRetryAfterResponse(retryAfterSec int) bool {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))

	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		after := time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		rc.UpdateRetryAfter(after)
		log.Debug("Retry after", zap.Int("RetryAfterSec", retryAfterSec), zap.Time("After", after))

		return true

	default:
		// do not retry
		return false
	}
}

func (z *RequestImpl) handleResponse(apiResImpl *ResponseImpl) (apiRes api_rpc.Response, err error) {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))
	if app.Root().IsDebug() {
		log.Debug("Response", zap.Int("code", apiResImpl.resStatusCode), zap.String("body", apiResImpl.resBodyString))
	}

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
		retryAfter := apiResImpl.resHeader.Get(ResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			log.Debug("Unable to parse header for RateLimit", zap.String("header", retryAfter), zap.Error(err))
			return nil, errors.New("unknown retry param")
		}

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

func (z *RequestImpl) Call() (apiRes api_rpc.Response, err error) {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint))
	req, err := z.httpRequest()
	if err != nil {
		log.Warn("Unable to prepare HTTP Request", zap.Error(err))
		return nil, errors.New(fmt.Sprintf("unable to prepare request for [%s]", z.endpoint))
	}

	z.waitForRetryIfRequired()

	switch cc := z.ctx.(type) {
	case api_context.ClientContext:
		log.Debug("Request", zap.Any("param", z.param), zap.Any("root", z.base))
		apiResImpl := &ResponseImpl{}
		apiResImpl.resStatusCode, apiResImpl.resHeader, apiResImpl.resBody, err = cc.DoRequest(req)

		if err != nil {
			log.Debug("Transport error", zap.Error(err))
			return z.ensureRetryOnError(err)
		}
		apiResImpl.resBodyString = string(apiResImpl.resBody)

		return z.handleResponse(apiResImpl)
	}

	log.Warn("No client available")
	return nil, errors.New("no client available")
}
