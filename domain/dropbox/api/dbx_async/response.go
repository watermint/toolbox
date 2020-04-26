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
		Proxy:     response.NewProxy(res),
		completed: true,
		complete:  complete,
	}
}

func NewIncomplete(res response.Response) Response {
	return &resImpl{
		Proxy:     response.NewProxy(res),
		completed: false,
		complete:  nil,
	}
}

type resImpl struct {
	response.Proxy
	completed bool
	complete  tjson.Json
}

func (z resImpl) IsCompleted() bool {
	return z.completed
}

func (z resImpl) Complete() tjson.Json {
	return z.complete
}
