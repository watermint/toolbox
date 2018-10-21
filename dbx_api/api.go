package dbx_api

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"time"
)

var (
	ReqHeaderSelectUser    = "Dropbox-API-Select-User"
	ReqHeaderSelectAdmin   = "Dropbox-API-Select-Admin"
	ResHeaderRetryAfter    = "Retry-After"
	ResJsonDotTag          = "\\.tag"
	DefaultClientTimeout   = time.Duration(60) * time.Second
	DateTimeFormat         = "2006-01-02T15:04:05Z"
	ErrorBadInputParam     = 400
	ErrorBadOrExpiredToken = 401
	ErrorAccessError       = 403
	ErrorEndpointSpecific  = 409
	ErrorRateLimit         = 429
	ErrorSuccess           = 0
	ErrorTransport         = 1000
	ErrorUnknown           = 1001
	ErrorServerError       = 1500
)

func ParseApiError(responseBody string) (ae ApiError) {
	ae.ErrorTag = gjson.Get(responseBody, "error."+ResJsonDotTag).String()
	ae.ErrorSummary = gjson.Get(responseBody, "error_summary").String()
	ae.UserMessageLocale = gjson.Get(responseBody, "user_message.locale").String()
	ae.UserMessage = gjson.Get(responseBody, "user_message.text").String()
	ae.ErrorBody = json.RawMessage(gjson.Get(responseBody, "error").Raw)
	ae.UserMessageBody = json.RawMessage(gjson.Get(responseBody, "user_message").Raw)

	return
}

func ParseAccessError(responseBody string) (ae AccessError) {
	ae.PaperAccessDenied = gjson.Get(responseBody, "invalid_account_type.\\.tag").String()
	ae.InvalidAccountType = gjson.Get(responseBody, "paper_access_denied.\\.tag").String()
	ae.ErrorBody = json.RawMessage(responseBody)

	return
}

type ErrorAnnotation struct {
	ErrorType int
	Error     error
}

func (e ErrorAnnotation) IsSuccess() bool {
	return e.ErrorType == ErrorSuccess
}
func (e ErrorAnnotation) IsFailure() bool {
	return e.ErrorType != ErrorSuccess
}
func (e ErrorAnnotation) ApiError() *ApiError {
	switch ae := e.Error.(type) {
	case ApiError:
		return &ae
	}
	return nil
}
func (e ErrorAnnotation) AccessError() *AccessError {
	switch ae := e.Error.(type) {
	case AccessError:
		return &ae
	}
	return nil
}
func (e ErrorAnnotation) ErrorTypeLabel() string {
	switch e.ErrorType {
	case ErrorBadInputParam:
		return "bad_input_param"
	case ErrorBadOrExpiredToken:
		return "bad_or_expired_token"
	case ErrorAccessError:
		return "access_error"
	case ErrorEndpointSpecific:
		return "endpoint_specific"
	case ErrorRateLimit:
		return "rate_limit"
	case ErrorSuccess:
		return "success"
	case ErrorTransport:
		return "transport_error"
	case ErrorUnknown:
		return "unknown"
	case ErrorServerError:
		return "server_error"
	}
	return "unknown"
}
func (e ErrorAnnotation) UserMessage() string {
	if e.Error == nil {
		return ""
	}
	if ae := e.ApiError(); ae != nil {
		if ae.UserMessage != "" {
			return ae.UserMessage
		} else {
			return ae.ErrorSummary
		}
	}
	if ae := e.AccessError(); ae != nil {
		return ae.Error()
	}
	return e.Error.Error()
}

type ArgAsyncJobId struct {
	AsyncJobId string `json:"async_job_id"`
}

type ServerError struct {
	StatusCode int
}

func (e ServerError) Error() string {
	return fmt.Sprintf("An error occurred on the Dropbox servers (%d). Check status.dropbox.com for announcements about Dropbox service issues.", e.StatusCode)
}

type ApiError struct {
	ErrorTag          string          `json:"error,omitempty"`
	ErrorSummary      string          `json:"error_summary,omitempty"`
	ErrorBody         json.RawMessage `json:"error,omitempty"`
	UserMessageLocale string          `json:"user_message_lang,omitempty"`
	UserMessage       string          `json:"user_message,omitempty"`
	UserMessageBody   json.RawMessage `json:"user_message,omitempty"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Endpoint specific error[%s] %s", e.ErrorTag, e.ErrorSummary)
}

type AccessError struct {
	InvalidAccountType string          `json:"invalid_account_type,omitempty"`
	PaperAccessDenied  string          `json:"paper_access_denied,omitempty"`
	ErrorBody          json.RawMessage `json:"error,omitempty"`
}

func (a AccessError) Error() string {
	if a.InvalidAccountType != "" {
		return a.InvalidAccountType
	}
	if a.PaperAccessDenied != "" {
		return a.PaperAccessDenied
	}
	return "The user or team account doesn't have access to the endpoint or feature"
}

type ApiErrorRateLimit struct {
	RetryAfter int
}

func (e ApiErrorRateLimit) Error() string {
	return fmt.Sprintf("API Rate limit (retry after %d sec)", e.RetryAfter)
}

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

type Context struct {
	Token      string
	Client     *http.Client
	RetryAfter time.Time
}

func NewContext(token string) *Context {
	return &Context{
		Token:  token,
		Client: &http.Client{Timeout: DefaultClientTimeout},
	}
}

//func (a *Context) CallRpc(route string, arg interface{}) (apiRes *ApiRpcResponse, err error) {
//	req := ApiRpcRequest{
//		Param:      arg,
//		Endpoint:      route,
//		AuthHeader: true,
//		Context:    a,
//	}
//	return req.Call()
//}
//
//func (a *Context) CallRpcAsMemberId(route, memberId string, arg interface{}) (apiRes *ApiRpcResponse, err error) {
//	req := ApiRpcRequest{
//		Param:      arg,
//		Endpoint:      route,
//		AuthHeader: true,
//		Context:    a,
//		AsMemberId: memberId,
//	}
//	return req.Call()
//}
//
//func (a *Context) CallRpcAsAdminId(route, adminId string, arg interface{}) (apiRes *ApiRpcResponse, err error) {
//	req := ApiRpcRequest{
//		Param:      arg,
//		Endpoint:      route,
//		AuthHeader: true,
//		Context:    a,
//		AsAdminId:  adminId,
//	}
//	return req.Call()
//}
//
//func (a *Context) NewApiRpcRequest(route string, arg interface{}) *ApiRpcRequest {
//	return &ApiRpcRequest{
//		Param:      arg,
//		Endpoint:      route,
//		AuthHeader: true,
//		Context:    a,
//	}
//}
