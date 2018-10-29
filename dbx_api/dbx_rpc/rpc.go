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
	RpcEndpoint = "api.dropboxapi.com"
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

	log.Debug(
		"request params",
		zap.String("params", string(requestParam)),
	)

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

	log.Debug("do request")
	res, err := c.Client.Do(req)

	if err != nil {
		log.Debug("transport error", zap.Error(err))
		return annotate(nil, dbx_api.ErrorTransport, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
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
		return annotate(nil, dbx_api.ErrorServerError,
			dbx_api.ServerError{
				StatusCode: res.StatusCode,
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
