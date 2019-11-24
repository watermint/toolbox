package api_rpc_impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_capture"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

// Stateful object
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
	upload     []byte
}

func (z *CallerImpl) rpcRequestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, z.endpoint)
}

func (z *CallerImpl) contentRequestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", ContentEndpoint, z.endpoint)
}

func (z *CallerImpl) createRequest(url string, contentType string, arg interface{}) (req api_rpc.Request, err error) {
	log := z.ctx.Log().With(zap.String("endpoint", z.endpoint), zap.String("Routine", ut_runtime.GetGoRoutineName()))

	headers := make(map[string]string)

	headers[api_rpc.ReqHeaderContentType] = contentType
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
	if arg != nil {
		p, err := json.Marshal(arg)
		if err != nil {
			z.ctx.Log().Debug("Unable to marshal params", zap.Error(err))
			return nil, err
		}
		headers[api_rpc.ReqHeaderArg] = string(p)
	}

	return newPostRequest(z.ctx, url, z.param, headers)
}

func (z *CallerImpl) ensureRetryOnError(lastErr error) (res api_rpc.Response, err error) {
	l := z.ctx.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))
	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		if rc.IsNoRetry() {
			l.Debug("Abort retry due to NoRetryOnError", zap.Error(lastErr))
			return nil, lastErr
		}

		// Add lastErr and wait if required
		abort := nw_ratelimit.AddError(z.ctx.Hash(), z.endpoint, lastErr)
		if abort {
			l.Debug("Abort retry due to rateLimit", zap.Error(lastErr))
			return nil, lastErr
		}

		return z.Call()

	default:
		l.Debug("No retry context")
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

func (z *CallerImpl) handleRetryAfterResponse(retryAfterSec int) bool {
	switch rc := z.ctx.(type) {
	case api_context.RetryContext:
		after := time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		z.ctx.Log().Debug("Retry after", zap.Int("RetryAfterSec", retryAfterSec), zap.Bool("noRetry", rc.IsNoRetry()))
		nw_ratelimit.UpdateRetryAfter(z.ctx.Hash(), z.endpoint, after)

		return true

	default:
		// do not retry
		return false
	}
}

func (z *CallerImpl) handleResponse(apiResImpl *ResponseImpl) (apiRes api_rpc.Response, err error) {
	l := z.ctx.Log().With(zap.String("endpoint", z.endpoint), zap.String("Routine", ut_runtime.GetGoRoutineName()))
	//if app.Root().IsDebug() {
	//	log.Debug("Response", zap.Int("code", apiResImpl.resStatusCode), zap.String("body", apiResImpl.resBodyString))
	//}

	switch apiResImpl.resStatusCode {
	case http.StatusOK:
		return apiResImpl, nil

	case ErrorBadInputParam: // Bad input param
		// In case of the server returned unexpected HTML response
		// Response body should be plain text
		if strings.HasPrefix(apiResImpl.resBodyString, "<!DOCTYPE html>") {
			l.Debug("Bad response from server, assume that can retry", zap.String("response", apiResImpl.resBodyString))

			// add error & retry
			nw_ratelimit.AddError(z.ctx.Hash(), z.endpoint, errors.New("bad response from server: res_code 400 with html body"))
			return z.Call()
		}
		l.Debug("Bad input param", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorBadOrExpiredToken: // Bad or expired token
		l.Debug("Bad or expired token", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorAccessError: // Access Error
		l.Debug("Access Error", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseAccessError(apiResImpl.resBodyString)

	case ErrorEndpointSpecific: // Endpoint specific
		l.Debug("Endpoint specific error", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseApiError(apiResImpl.resBodyString)

	case ErrorNoPermission: // No permission
		l.Debug("No Permission", zap.String("Error", apiResImpl.resBodyString))
		return nil, api_rpc.ParseAccessError(apiResImpl.resBodyString)

	case ErrorRateLimit: // Rate limit
		retryAfter := apiResImpl.resHeader.Get(api_rpc.ResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			l.Debug("Unable to parse header for RateLimit", zap.String("header", retryAfter), zap.Error(err))
			return nil, errors.New("unknown retry param")
		}

		l.Debug("Rate limit", zap.Int("retryAfterSec", retryAfterSec))
		if z.handleRetryAfterResponse(retryAfterSec) {
			// Retry
			return z.Call()
		} else {
			return nil, api_rpc.ApiErrorRateLimit{RetryAfter: retryAfterSec}
		}
	}

	if int(apiResImpl.resStatusCode/100) == 5 {
		l.Debug("server error", zap.Int("status_code", apiResImpl.resStatusCode), zap.String("body", apiResImpl.resBodyString))
		return z.ensureRetryOnError(api_rpc.ServerError{StatusCode: apiResImpl.resStatusCode})
	}

	l.Warn("Unknown or server error", zap.Int("Code", apiResImpl.resStatusCode), zap.String("Body", apiResImpl.resBodyString))
	return nil, err
}

func (z *CallerImpl) rpcCall() (apiRes api_rpc.Response, err error) {
	return z.doCall(func() (api_rpc.Request, *http.Request, error) {
		r, err := z.createRequest(z.rpcRequestUrl(), "application/json", nil)
		if err != nil {
			return nil, nil, err
		}
		hr, err := r.Request()
		if err != nil {
			return nil, nil, err
		}
		return r, hr, nil
	})
}

func (z *CallerImpl) Call() (apiRes api_rpc.Response, err error) {
	if z.upload != nil {
		return z.uploadCall()
	} else {
		return z.rpcCall()
	}
}

func (z *CallerImpl) uploadCall() (res api_rpc.Response, err error) {
	return z.doCall(func() (api_rpc.Request, *http.Request, error) {
		rq, err := z.createRequest(z.contentRequestUrl(), "application/octet-stream", z.param)
		if err != nil {
			return nil, nil, err
		}
		hr, err := rq.Upload(nw_bandwidth.WrapReader(bytes.NewReader(z.upload)))
		if err != nil {
			return nil, nil, err
		}
		return rq, hr, nil
	})
}

func (z *CallerImpl) Upload(r io.Reader) (res api_rpc.Response, err error) {
	l := z.ctx.Log()
	z.upload, err = ioutil.ReadAll(r)
	if err != nil {
		l.Debug("Unable to read", zap.Error(err))
		return nil, err
	}

	return z.uploadCall()
}

func (z *CallerImpl) doCall(mkReq func() (api_rpc.Request, *http.Request, error)) (apiRes api_rpc.Response, err error) {
	l := z.ctx.Log().With(zap.String("endpoint", z.endpoint), zap.String("Routine", ut_runtime.GetGoRoutineName()))
	l.Debug("doCall: Make request")
	req, httpReq, err := mkReq()
	if err != nil {
		l.Warn("Unable to prepare HTTP Caller", zap.Error(err))
		return nil, errors.New(fmt.Sprintf("unable to prepare request for [%s]", z.endpoint))
	}

	nw_ratelimit.WaitIfRequired(z.ctx.Hash(), z.endpoint)

	switch cc := z.ctx.(type) {
	case api_context.ClientContext:
		l.Debug("Caller", zap.String("url", req.Url()), zap.Any("param", z.param), zap.Any("root", z.base))
		callStart := time.Now()
		apiResImpl := &ResponseImpl{}
		apiResImpl.resStatusCode, apiResImpl.resHeader, apiResImpl.resBody, err = cc.DoRequest(httpReq)
		callEnd := time.Now()
		if apiResImpl.resBody != nil {
			apiResImpl.resBodyString = string(apiResImpl.resBody)
		}

		var cp api_capture.Capture
		if cac, ok := cc.(api_context.CaptureContext); ok {
			cp = api_capture.NewCapture(cac.Capture())
		} else {
			cp = api_capture.Current()
		}
		cp.Rpc(req, apiResImpl, err, callEnd.UnixNano()-callStart.UnixNano())

		if err != nil {
			l.Debug("Transport error", zap.Error(err))
			return z.ensureRetryOnError(err)
		}

		return z.handleResponse(apiResImpl)
	}

	l.Warn("No client available")
	return nil, errors.New("no client available")
}
