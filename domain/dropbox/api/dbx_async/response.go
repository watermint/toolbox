package dbx_async

import (
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
)

type Response interface {
	response.Response

	// True when the async job completed.
	IsCompleted() bool

	// Completed body. Returns nil if the operation is not yet completed.
	Complete() tjson.Json
}

func NewCompleted(res response.Response, complete tjson.Json) Response {
	return &resImpl{
		res:       res,
		completed: true,
		complete:  complete,
	}
}

func NewIncomplete(res response.Response) Response {
	return &resImpl{
		res:       res,
		completed: false,
		complete:  nil,
	}
}

type resImpl struct {
	res       response.Response
	completed bool
	complete  tjson.Json
}

func (z resImpl) IsSuccess() bool {
	return z.res.IsSuccess()
}

func (z resImpl) Alt() response.Body {
	return z.res.Alt()
}

func (z resImpl) Error() error {
	return z.res.Error()
}

func (z resImpl) Code() int {
	return z.res.Code()
}

func (z resImpl) CodeCategory() response.CodeCategory {
	return z.res.CodeCategory()
}

func (z resImpl) Headers() map[string]string {
	return z.res.Headers()
}

func (z resImpl) Header(header string) string {
	return z.res.Header(header)
}

func (z resImpl) Success() response.Body {
	return z.res.Success()
}

func (z resImpl) IsCompleted() bool {
	return z.completed
}

func (z resImpl) Complete() tjson.Json {
	return z.complete
}
