package dbx_rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	RpcEndpoint         = "api.dropboxapi.com"
	SameErrorRetryCount = 5
	SameErrorRetryWait  = time.Duration(60) * time.Second
)

type RpcResponse struct {
	StatusCode int
	Tag        string
	Body       string
	Error      error
}

type RpcRequest struct {
	Param        interface{}
	Endpoint     string
	NoAuthHeader bool
	AsMemberId   string
	AsAdminId    string

	// Dropbox-API-Path-Root header. See https://www.dropbox.com/developers/reference/namespace-guide
	PathRoot interface{}
}

func (z *RpcRequest) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, z.Endpoint)
}

func (z *RpcRequest) rpcRequest(c *dbx_api.Context) (req *http.Request, err error) {
	url := z.requestUrl()
	log := c.Log().With(zap.String("endpoint", z.Endpoint))

	// param
	requestParam, err := json.Marshal(z.Param)
	if err != nil {
		log.Debug("unable to marshal params", zap.Error(err))
		return nil, err
	}
	req, err = http.NewRequest("POST", url, bytes.NewReader(requestParam))
	if err != nil {
		log.Debug(
			"unable create request",
			zap.Error(err),
		)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if !z.NoAuthHeader {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	if z.AsMemberId != "" {
		req.Header.Add(dbx_api.ReqHeaderSelectUser, z.AsMemberId)
	}
	if z.AsAdminId != "" {
		req.Header.Add(dbx_api.ReqHeaderSelectAdmin, z.AsAdminId)
	}
	if z.PathRoot != nil {
		pr, err := json.Marshal(z.PathRoot)
		if err != nil {
			log.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		req.Header.Add(dbx_api.ReqHeaderPathRoot, string(pr))
	}
	return
}

func (z *RpcRequest) ensureRetryOnError(c *dbx_api.Context, lastErr error) (apiRes *RpcResponse, err error) {
	sameErrorCount := 0
	if c.LastErrors == nil {
		c.LastErrors = make([]error, 1)
		c.LastErrors[0] = lastErr
	} else {
		for _, e := range c.LastErrors {
			if e.Error() == lastErr.Error() {
				sameErrorCount++
			}
		}
		c.LastErrors = append(c.LastErrors, lastErr)
	}

	if sameErrorCount >= SameErrorRetryCount {
		c.Log().Debug(
			"Abort retry due to `same_error_count` exceed threshold",
			zap.Int("same_error_count", sameErrorCount),
			zap.Error(err),
		)
		return nil, err
	}

	c.RetryAfter = time.Now().Add(SameErrorRetryWait)
	c.Log().Debug(
		"retry after",
		zap.Error(err),
		zap.Time("retry_after", c.RetryAfter),
	)

	return z.Call(c)
}

func (z *RpcRequest) Call(c *dbx_api.Context) (apiRes *RpcResponse, err error) {
	annotate := func(res *RpcResponse, et int, err error) (*RpcResponse, error) {
		return res, err
	}
	log := c.Log().With(zap.String("endpoint", z.Endpoint))
	req, err := z.rpcRequest(c)
	if err != nil {
		log.Debug("unable to prepare request", zap.Error(err))
		return annotate(nil, dbx_api.ErrorUnknown, errors.New(fmt.Sprintf("unable to prepare request for [%s]", z.Endpoint)))
	}

	now := time.Now()
	if !c.RetryAfter.IsZero() && now.Before(c.RetryAfter) {
		log.Debug("sleep until",
			zap.Time("retry_after", c.RetryAfter),
		)
		time.Sleep(c.RetryAfter.Sub(now))
	}

	log.Debug("do_request", zap.Any("param", z.Param), zap.Any("root", z.Param))
	res, err := c.Client.Do(req)

	if err != nil {
		log.Debug("transport error", zap.Error(err))
		return z.ensureRetryOnError(c, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Do not retry
		log.Debug("unable to read boy", zap.Error(err))
		return annotate(nil, dbx_api.ErrorTransport, err)
	}
	res.Body.Close()

	bodyString := string(body)

	if res.StatusCode == http.StatusOK {
		jsonBody := bodyString
		tag := gjson.Get(jsonBody, dbx_api.ResJsonDotTag)
		responseTag := ""
		if tag.Exists() {
			responseTag = tag.String()
		}

		return annotate(
			&RpcResponse{
				StatusCode: res.StatusCode,
				Body:       jsonBody,
				Tag:        responseTag,
				Error:      nil,
			},
			dbx_api.ErrorSuccess,
			nil,
		)
	}

	switch res.StatusCode {
	case dbx_api.ErrorBadInputParam: // Bad input param
		log.Debug("bad input param",
			zap.String("error_body", bodyString),
		)
		return annotate(nil, dbx_api.ErrorBadInputParam, dbx_api.ParseApiError(bodyString))

	case dbx_api.ErrorBadOrExpiredToken: // Bad or expired token
		log.Debug(
			"bad or expired token",
			zap.String("error_body", bodyString),
		)
		return annotate(nil, dbx_api.ErrorBadOrExpiredToken, dbx_api.ParseApiError(bodyString))

	case dbx_api.ErrorAccessError: // Access Error
		log.Debug(
			"access error",
			zap.String("error_body", bodyString),
		)
		return annotate(nil, dbx_api.ErrorAccessError, dbx_api.ParseAccessError(bodyString))

	case dbx_api.ErrorEndpointSpecific: // Endpoint specific
		log.Debug(
			"endpoint specific error",
			zap.String("error_body", bodyString),
		)

		return annotate(nil, dbx_api.ErrorEndpointSpecific, dbx_api.ParseApiError(bodyString))

	case dbx_api.ErrorNoPermission: // No permission
		log.Debug(
			"access error",
			zap.String("error_body", bodyString),
		)
		return annotate(nil, dbx_api.ErrorNoPermission, dbx_api.ParseAccessError(bodyString))

	case dbx_api.ErrorRateLimit: // Rate limit
		retryAfter := res.Header.Get(dbx_api.ResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			log.Debug(
				"unable to parse header for RateLimit",
				zap.String("header", retryAfter),
				zap.Error(err),
			)
			return annotate(nil, dbx_api.ErrorRateLimit, errors.New("unknown retry param"))
		}

		c.RetryAfter = time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		log.Debug(
			"retry after",
			zap.Int("retry_after_second", retryAfterSec),
			zap.Time("retry_after", c.RetryAfter),
		)

		// Retry
		return z.Call(c)
	}

	if int(res.StatusCode/100) == 5 {
		log.Debug(
			"server error",
			zap.Int("status_code", res.StatusCode),
			zap.String("body", bodyString),
		)
		return z.ensureRetryOnError(c, dbx_api.ServerError{StatusCode: res.StatusCode})
	}

	log.Debug(
		"unknown or server error",
		zap.Int("status_code", res.StatusCode),
		zap.String("body", bodyString),
	)
	return annotate(nil, dbx_api.ErrorUnknown, err)
}
