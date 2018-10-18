package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

var (
	API_RPC_ENDPOINT                                  = "api.dropboxapi.com"
	API_REQ_HEADER_SELECT_USER                        = "Dropbox-API-Select-User"
	API_RES_HEADER_RETRY_AFTER                        = "Retry-After"
	API_RES_JSON_DOT_TAG                              = "\\.tag"
	API_DEFAULT_UPLOAD_CHUNKED_UPLOAD_THRESHOLD int64 = 150 * 1048576
	API_DEFAULT_UPLOAD_CHUNK_SIZE               int64 = 150 * 1048576
	API_DEFAULT_CLIENT_TIMEOUT                        = 60
)

type DropboxPath struct {
	Path string
}

type ArgAsyncJobId struct {
	AsyncJobId string `json:"async_job_id"`
}

type ApiServerError struct {
	StatusCode int
}

func (e ApiServerError) Error() string {
	return fmt.Sprintf("An error occurred on the Dropbox servers (%d). Check status.dropbox.com for announcements about Dropbox service issues.", e.StatusCode)
}

type ApiEndpointSpecificError struct {
	ErrorTag     string `json:"error,omitempty"`
	ErrorSummary string `json:"error_summary,omitempty"`
	UserMessage  string `json:"user_message,omitempty"`
}

func (e ApiEndpointSpecificError) Error() string {
	return fmt.Sprintf("Endpoint specific error[%s] %s", e.ErrorTag, e.ErrorSummary)
}

type ApiInvalidTokenError struct {
}

func (e ApiInvalidTokenError) Error() string {
	return "Bad or expired token"
}

type ApiAccessError struct {
}

func (e ApiAccessError) Error() string {
	return "The user or team account doesn't have access to the endpoint or feature"
}

type ApiBadInputParamError struct {
	ErrorSummary string `json:"error_summary"`
}

func (e ApiBadInputParamError) Error() string {
	return e.ErrorSummary
}

type ApiErrorRateLimit struct {
	RetryAfter int
}

func (e ApiErrorRateLimit) Error() string {
	return fmt.Sprintf("API Rate limit (retry after %d sec)", e.RetryAfter)
}

func NewDropboxPath(path string) *DropboxPath {
	return &DropboxPath{
		Path: path,
	}
}

func (d *DropboxPath) CleanPath() string {
	p := filepath.ToSlash(filepath.Clean(d.Path))
	if p == "/" {
		return ""
	} else {
		return p
	}
}

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

type ApiConfig struct {
	Timeout                      time.Duration
	UploadChunkedUploadThreshold int64
	UploadChunkedUploadChunkSize int64
}

func NewDefaultApiConfig() *ApiConfig {
	return &ApiConfig{
		Timeout: time.Duration(API_DEFAULT_CLIENT_TIMEOUT) * time.Second,
		UploadChunkedUploadThreshold: API_DEFAULT_UPLOAD_CHUNKED_UPLOAD_THRESHOLD,
		UploadChunkedUploadChunkSize: API_DEFAULT_UPLOAD_CHUNK_SIZE,
	}
}

type ApiContext struct {
	Token  string
	Client *http.Client
	Config *ApiConfig
}

func (a *ApiContext) CallRpc(route string, arg interface{}) (apiRes *ApiRpcResponse, err error) {
	req := ApiRpcRequest{
		Param:      arg,
		Route:      route,
		AuthHeader: true,
		Context:    a,
	}
	return req.Call()
}

func (a *ApiContext) CallRpcAsMemberId(route, memberId string, arg interface{}) (apiRes *ApiRpcResponse, err error) {
	req := ApiRpcRequest{
		Param:      arg,
		Route:      route,
		AuthHeader: true,
		Context:    a,
		AsMemberId: memberId,
	}
	return req.Call()
}

func (a *ApiContext) NewApiRpcRequest(route string, arg interface{}) *ApiRpcRequest {
	return &ApiRpcRequest{
		Param:      arg,
		Route:      route,
		AuthHeader: true,
		Context:    a,
	}
}

func NewDefaultApiContext(token string) *ApiContext {
	config := NewDefaultApiConfig()
	return &ApiContext{
		Token:  token,
		Client: &http.Client{Timeout: config.Timeout},
		Config: config,
	}
}

func (c *ApiContext) PrepareHeader(req *http.Request) *http.Request {

	return req
}

type ApiRpcResponse struct {
	StatusCode int
	Tag        string
	Body       string
	Error      error
}

type ApiRpcRequest struct {
	Param      interface{}
	AuthHeader bool
	Route      string
	AsMemberId string
	Context    *ApiContext
}

func (a *ApiRpcRequest) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", API_RPC_ENDPOINT, a.Route)
}

func (a *ApiRpcRequest) rpcRequest() (req *http.Request, err error) {
	url := a.requestUrl()

	// param
	requestParam, err := json.Marshal(a.Param)
	if err != nil {
		seelog.Debugf("Route[%s] Unable to marshal params. error[%s]", a.Route, err)
		return nil, err
	}
	seelog.Debugf("Request Params[%s]", string(requestParam))

	req, err = http.NewRequest("POST", url, bytes.NewReader(requestParam))
	if err != nil {
		seelog.Debugf("Route[%s] Unable create request. error[%s]", a.Route, err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if a.AuthHeader {
		req.Header.Add("Authorization", "Bearer "+a.Context.Token)
	}
	if a.AsMemberId != "" {
		req.Header.Add(API_REQ_HEADER_SELECT_USER, a.AsMemberId)
	}
	a.Context.PrepareHeader(req)
	return
}

func (a *ApiRpcRequest) Call() (apiRes *ApiRpcResponse, err error) {
	req, err := a.rpcRequest()
	if err != nil {
		seelog.Tracef("Route[%s] Unable to prepare request : error[%s]", a.Route, err)
		return
	}

	seelog.Tracef("Route[%s]", a.Route)
	res, err := a.Context.Client.Do(req)

	if err != nil {
		seelog.Debugf("Route[%s] Transport error[%s]", a.Route, err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		seelog.Debugf("Route[%s] Unable to read body. error[%s]", a.Route, err)
		return nil, err
	}
	res.Body.Close()

	bodyString := string(body)

	if res.StatusCode == http.StatusOK {
		jsonBody := bodyString
		tag := gjson.Get(jsonBody, API_RES_JSON_DOT_TAG)
		responseTag := ""
		if tag.Exists() {
			responseTag = tag.String()
		}

		return &ApiRpcResponse{
			StatusCode: res.StatusCode,
			Body:       jsonBody,
			Tag:        responseTag,
			Error:      nil,
		}, nil
	}

	switch res.StatusCode {
	case 400: // Bad input param
		seelog.Debugf("Route[%s] Bad input param. error[%s]", a.Route, err)
		return nil, ApiBadInputParamError{
			ErrorSummary: bodyString,
		}

	case 401: // Bad or expired token
		seelog.Debugf("Route[%s] Bad or expired token.", a.Route)
		return nil, ApiInvalidTokenError{}

	case 403: // Access Error
		seelog.Debugf("Route[%s] Access Error.", a.Route)
		return nil, ApiAccessError{}

	case 409: // Endpoint specific
		seelog.Debugf("Route[%s] Endpoint specific Error.", a.Route)
		apiErr := ApiEndpointSpecificError{}
		ume := json.Unmarshal(body, &apiErr)
		if ume != nil {
			seelog.Debugf("Route[%s] unknown or server error. response body[%s], unmarshal err[%s]", a.Route, bodyString, err)
			return nil, ApiEndpointSpecificError{
				ErrorSummary: bodyString,
			}
		}

		return nil, apiErr

	case 429: // Rate limit
		retryAfter := res.Header.Get(API_RES_HEADER_RETRY_AFTER)
		retryAfterSec, err := strconv.Atoi(retryAfter)
		if err != nil {
			seelog.Debugf("Route[%s] Unable to parse '%s' header. HeaderContent[%s] error[%s]", a.Route, retryAfter, err)
			return nil, errors.New("unknown retry param")
		}
		seelog.Debugf("Route[%s] Wait for retry [%d] seconds.", retryAfterSec)

		return nil, ApiErrorRateLimit{RetryAfter: retryAfterSec}
	}

	if int(res.StatusCode/100) == 5 {
		seelog.Debugf("Route[%s] Server error", a.Route)
		return nil, ApiServerError{
			StatusCode: res.StatusCode,
		}
	}

	seelog.Debugf("Route[%s] unknown or server error[%s]", a.Route, err)
	return nil, err
}
