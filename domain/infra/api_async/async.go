package api_async

import "github.com/tidwall/gjson"

type Async interface {
	Param(p interface{}) Async
	Status(endpoint string) Async
	PollInterval(second int) Async
	OnSuccess(func(res Response) error) Async
	OnFailure(func(err error) error) Async
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
