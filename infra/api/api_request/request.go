package api_request

import (
	"github.com/watermint/toolbox/infra/api/api_response"
	"net/http"
)

const (
	ReqHeaderContentType   = "Content-Type"
	ReqHeaderAuthorization = "Authorization"
	ReqHeaderSelectUser    = "Dropbox-API-Select-User"
	ReqHeaderSelectAdmin   = "Dropbox-API-Select-Admin"
	ReqHeaderPathRoot      = "Dropbox-API-Path-Root"
	ReqHeaderArg           = "Dropbox-API-Arg"
)

type Request interface {
	// Request param as string. Might be empty string until Make call.
	ParamString() string

	// Param
	Param(p interface{}) Request

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

	// Make request
	Make() (req *http.Request, err error)
}
