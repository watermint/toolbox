package api_rpc

import "github.com/tidwall/gjson"

type Request interface {
	Param(param interface{}) Request
	OnSuccess(func(res Response) error) Request
	OnFailure(func(err error) error) Request
	Call() (res Response, err error)
}

type Response interface {
	// Returns error. Returns nil if a response succeed.
	Error() error

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
