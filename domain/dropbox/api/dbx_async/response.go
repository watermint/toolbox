package dbx_async

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
)

type Response interface {
	es_response.Response

	// True when the async job completed.
	IsCompleted() bool

	// Completed body. Returns nil if the operation is not yet completed.
	Complete() es_json.Json
}

func NewCompleted(res es_response.Response, complete es_json.Json) Response {
	return &resImpl{
		Proxy:     es_response.NewProxy(res),
		completed: true,
		complete:  complete,
	}
}

func NewIncomplete(res es_response.Response) Response {
	return &resImpl{
		Proxy:     es_response.NewProxy(res),
		completed: false,
		complete:  nil,
	}
}

type resImpl struct {
	es_response.Proxy
	completed bool
	complete  es_json.Json
}

func (z resImpl) IsAuthInvalidToken() bool {
	return false
}

func (z resImpl) IsTextContentType() bool {
	return es_response_impl.IsTextContentType(z)
}

func (z resImpl) IsCompleted() bool {
	return z.completed
}

func (z resImpl) Complete() es_json.Json {
	return z.complete
}
