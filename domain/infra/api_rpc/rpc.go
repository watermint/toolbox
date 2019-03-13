package api_rpc

import "github.com/tidwall/gjson"

type Request interface {
	Param(param interface{}) Request
	OnSuccess(func(res Response) error) Request
	OnFailure(func(err error) error) Request
	Call() (res Response, err error)
}

type Response interface {
	Error() error
	StatusCode() int
	Body() (body string, err error)
	Json() (res gjson.Result, err error)
	Model(v interface{}) error
	ModelWithPath(v interface{}, path string) error
}
