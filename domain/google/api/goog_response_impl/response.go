package goog_response_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/google/api/goog_error"
	"github.com/watermint/toolbox/domain/google/api/goog_response"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
)

func New(res es_response.Response) goog_response.Response {
	return &resImpl{
		Proxy: es_response.NewProxy(res),
	}
}

type resImpl struct {
	es_response.Proxy
}

func (z resImpl) IsTextContentType() bool {
	return es_response_impl.IsTextContentType(z)
}

func (z resImpl) GoogleError() (err goog_error.GoogleError) {
	if z.IsSuccess() {
		return
	}
	_ = z.Alt().Json().Model(err)
	return
}

func (z resImpl) Failure() (error, bool) {
	if z.IsSuccess() {
		return z.Proxy.Failure()
	}
	ge := &goog_error.GoogleError{}
	if err := json.Unmarshal(z.Alt().Body(), ge); err != nil {
		return z.Proxy.Failure()
	}
	return ge, true
}
