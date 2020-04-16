package api_request

import (
	"github.com/watermint/toolbox/infra/api/api_response"
	"net/http"
)

const (
	ReqHeaderContentType           = "Content-Type"
	ReqHeaderAccept                = "Accept"
	ReqHeaderContentLength         = "Content-Length"
	ReqHeaderAuthorization         = "Authorization"
	ReqHeaderUserAgent             = "User-Agent"
	ReqHeaderDropboxApiSelectUser  = "Dropbox-API-Select-User"
	ReqHeaderDropboxApiSelectAdmin = "Dropbox-API-Select-Admin"
	ReqHeaderDropboxApiPathRoot    = "Dropbox-API-Path-Root"
	ReqHeaderDropboxApiArg         = "Dropbox-API-Arg"
)

type Request interface {
	// Request param as string. Might be empty string until Make call.
	ParamString() string

	// With a param
	Param(p interface{}) Request

	// With a header
	Header(key, value string) Request

	// Call request
	Call() (res api_response.Response, err error)

	// Endpoint.
	Endpoint() string

	// Request url. Might be empty string until Make call.
	Url() string

	// Headers. Might be empty map until Make call.
	Headers() map[string]string

	// Method. Might be empty string until Make call.
	Method() string

	// Content length. Might be zero until Make call.
	ContentLength() int64

	// Make request
	Make() (req *http.Request, err error)
}
