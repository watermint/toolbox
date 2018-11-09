package dbx_rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
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
}

func (a *RpcRequest) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, a.Endpoint)
}

func (a *RpcRequest) rpcRequest(c *dbx_api.Context) (req *http.Request, err error) {
	url := a.requestUrl()
	log := c.Log().With(zap.String("endpoint", a.Endpoint))

	// param
	requestParam, err := json.Marshal(a.Param)
	if err != nil {
		log.Debug(
			"unable to marshal params",
			zap.Error(err),
		)

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
	if !a.NoAuthHeader {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	if a.AsMemberId != "" {
		req.Header.Add(dbx_api.ReqHeaderSelectUser, a.AsMemberId)
	}
	if a.AsAdminId != "" {
		req.Header.Add(dbx_api.ReqHeaderSelectAdmin, a.AsAdminId)
	}
	return
}

func (a *RpcRequest) ensureRetryOnError(c *dbx_api.Context, annotation dbx_api.ErrorAnnotation) (apiRes *RpcResponse, ea dbx_api.ErrorAnnotation, err error) {
	sameErrorCount := 0
	if c.LastErrors == nil {
		c.LastErrors = make([]dbx_api.ErrorAnnotation, 1)
		c.LastErrors[0] = annotation
	} else {
		for _, e := range c.LastErrors {
			if e.ErrorType == annotation.ErrorType {
				sameErrorCount++
			}
		}
		c.LastErrors = append(c.LastErrors, annotation)
	}

	if sameErrorCount >= SameErrorRetryCount {
		c.Log().Debug(
			"Abort retry due to `same_error_count` exceed threshold",
			zap.Int("same_error_count", sameErrorCount),
			zap.Int("error_type", annotation.ErrorType),
			zap.Error(annotation.Error),
		)
		return nil, annotation, annotation.Error
	}

	c.RetryAfter = time.Now().Add(SameErrorRetryWait)
	c.Log().Debug(
		"retry after",
		zap.Error(annotation.Error),
		zap.Int("error_type", annotation.ErrorType),
		zap.Time("retry_after", c.RetryAfter),
	)

	return a.Call(c)
}

func (a *RpcRequest) Call(c *dbx_api.Context) (apiRes *RpcResponse, ea dbx_api.ErrorAnnotation, err error) {
	annotate := func(res *RpcResponse, et int, err error) (*RpcResponse, dbx_api.ErrorAnnotation, error) {
		return res, dbx_api.ErrorAnnotation{
			ErrorType: et,
			Error:     err,
		}, err
	}
	log := c.Log().With(zap.String("endpoint", a.Endpoint))
	req, err := a.rpcRequest(c)
	if err != nil {
		log.Debug("unable to prepare request", zap.Error(err))
		return annotate(nil, dbx_api.ErrorUnknown, errors.New(fmt.Sprintf("unable to prepare request for [%s]", a.Endpoint)))
	}

	now := time.Now()
	if !c.RetryAfter.IsZero() && now.Before(c.RetryAfter) {
		log.Debug("sleep until",
			zap.Time("retry_after", c.RetryAfter),
		)
		time.Sleep(c.RetryAfter.Sub(now))
	}

	log.Debug("do_request", zap.Any("param", a.Param))
	res, err := c.Client.Do(req)

	if err != nil {
		log.Debug("transport error", zap.Error(err))
		return a.ensureRetryOnError(
			c,
			dbx_api.ErrorAnnotation{
				ErrorType: dbx_api.ErrorTransport,
				Error:     err,
			},
		)
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

		return annotate(nil, dbx_api.ErrorEndpointSpecific, dbx_api.ParseAccessError(bodyString))

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
		return a.Call(c)
	}

	if int(res.StatusCode/100) == 5 {
		log.Debug(
			"server error",
			zap.Int("status_code", res.StatusCode),
			zap.String("body", bodyString),
		)
		return a.ensureRetryOnError(
			c,
			dbx_api.ErrorAnnotation{
				ErrorType: dbx_api.ErrorServerError,
				Error: dbx_api.ServerError{
					StatusCode: res.StatusCode,
				},
			},
		)
	}

	log.Debug(
		"unknown or server error",
		zap.Int("status_code", res.StatusCode),
		zap.String("body", bodyString),
	)
	return annotate(nil, dbx_api.ErrorUnknown, err)
}
