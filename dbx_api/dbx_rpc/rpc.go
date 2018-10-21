package dbx_rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
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

	// param
	requestParam, err := json.Marshal(a.Param)
	if err != nil {
		seelog.Debugf("Endpoint[%s] Unable to marshal params. error[%s]", a.Endpoint, err)
		return nil, err
	}
	seelog.Debugf("Request Params[%s]", string(requestParam))

	req, err = http.NewRequest("POST", url, bytes.NewReader(requestParam))
	if err != nil {
		seelog.Debugf("Endpoint[%s] Unable create request. error[%s]", a.Endpoint, err)
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

	req, err := a.rpcRequest(c)
	if err != nil {
		seelog.Tracef("Endpoint[%s] Unable to prepare request : error[%s]", a.Endpoint, err)
		return annotate(nil, dbx_api.ErrorUnknown, errors.New(fmt.Sprintf("unable to prepare request for [%s]", a.Endpoint)))
	}

	now := time.Now()
	if c.RetryAfter.Before(now) {

		time.Sleep(c.RetryAfter.Sub(now))
	}

	seelog.Tracef("Endpoint[%s]", a.Endpoint)
	res, err := c.Client.Do(req)

	if err != nil {
		seelog.Debugf("Endpoint[%s] Transport error[%s]", a.Endpoint, err)
		return annotate(nil, dbx_api.ErrorTransport, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		seelog.Debugf("Endpoint[%s] Unable to read body. error[%s]", a.Endpoint, err)
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
		seelog.Debugf("Endpoint[%s] Bad input param. error[%s]", a.Endpoint, err)
		return annotate(nil, dbx_api.ErrorBadInputParam, dbx_api.ParseApiError(bodyString))

	case dbx_api.ErrorBadOrExpiredToken: // Bad or expired token
		seelog.Debugf("Endpoint[%s] Bad or expired token.", a.Endpoint)
		return annotate(nil, dbx_api.ErrorBadOrExpiredToken, dbx_api.ParseApiError(bodyString))

	case dbx_api.ErrorAccessError: // Access Error
		seelog.Debugf("Endpoint[%s] Access Error.", a.Endpoint)
		return annotate(nil, dbx_api.ErrorAccessError, dbx_api.ParseAccessError(bodyString))

	case dbx_api.ErrorEndpointSpecific: // Endpoint specific
		seelog.Debugf("Endpoint[%s] Endpoint specific Error.", a.Endpoint)
		return annotate(nil, dbx_api.ErrorEndpointSpecific, dbx_api.ParseAccessError(bodyString))

	case dbx_api.ErrorRateLimit: // Rate limit
		retryAfter := res.Header.Get(dbx_api.ResHeaderRetryAfter)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			seelog.Debugf("Endpoint[%s] Unable to parse '%s' header. HeaderContent[%s] error[%s]", a.Endpoint, retryAfter, err)
			return annotate(nil, dbx_api.ErrorRateLimit, errors.New("unknown retry param"))
		}

		c.RetryAfter = time.Now().Add(time.Duration(retryAfterSec+1) * time.Second)
		seelog.Debugf("Endpoint[%s] Retry after (%d sec, after %s)", retryAfterSec, c.RetryAfter.Format(dbx_api.DateTimeFormat))

		// Retry
		return a.Call(c)
	}

	if int(res.StatusCode/100) == 5 {
		seelog.Debugf("Endpoint[%s] Server error", a.Endpoint)
		return annotate(nil, dbx_api.ErrorServerError,
			dbx_api.ServerError{
				StatusCode: res.StatusCode,
			},
		)
	}

	seelog.Debugf("Endpoint[%s] unknown or server error[%s]", a.Endpoint, err)
	return annotate(nil, dbx_api.ErrorUnknown, err)
}
