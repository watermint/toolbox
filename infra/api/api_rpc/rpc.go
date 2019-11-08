package api_rpc

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const (
	ReqHeaderContentType   = "Content-Type"
	ReqHeaderAuthorization = "Authorization"
	ReqHeaderSelectUser    = "Dropbox-API-Select-User"
	ReqHeaderSelectAdmin   = "Dropbox-API-Select-Admin"
	ReqHeaderPathRoot      = "Dropbox-API-Path-Root"
	ReqHeaderArg           = "Dropbox-API-Arg"
	ResHeaderRetryAfter    = "Retry-After"
)

type Caller interface {
	Param(param interface{}) Caller
	OnSuccess(func(res Response) error) Caller
	OnFailure(func(err error) error) Caller
	Call() (res Response, err error)
	Upload(r io.Reader) (res Response, err error)
}

type Request interface {
	// Request param.
	Param() string

	// Request url.
	Url() string

	// Headers
	Headers() map[string]string

	// Method
	Method() string

	// Request
	Request() (req *http.Request, err error)

	// Upload request
	Upload(r io.Reader) (req *http.Request, err error)
}

type Response interface {
	// Response code. Returns -1 if a response does not contain status code.
	StatusCode() int

	// Returns body string. Returns empty & error if a response does not contain body.
	Body() (body string, err error)

	// Returns JSON result. Returns empty & error if a response is not a JSON document.
	Json() (res gjson.Result, err error)

	// Returns first element of the array.
	// Returns empty & error if a response is not an array of JSON
	JsonArrayFirst() (res gjson.Result, err error)

	// Parse model.
	Model(v interface{}) error

	// Parse model with given JSON path.
	ModelWithPath(v interface{}, path string) error

	// Parse model for a first element of the array of JSON.
	ModelArrayFirst(v interface{}) error
}
