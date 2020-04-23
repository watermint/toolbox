package api_request

import (
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
)

type MockRequest struct {
}

func (z MockRequest) Header(key, value string) Request {
	return &z
}

func (z MockRequest) ParamString() string {
	return ""
}

func (z MockRequest) Param(p interface{}) Request {
	return &z
}

func (z MockRequest) Call() (res api_response.Response, err error) {
	return nil, qt_errors.ErrorMock
}

func (z MockRequest) Endpoint() string {
	return ""
}

func (z MockRequest) Url() string {
	return ""
}

func (z MockRequest) Headers() map[string]string {
	return make(map[string]string)
}

func (z MockRequest) Method() string {
	return ""
}

func (z MockRequest) ContentLength() int64 {
	return 0
}

func (z MockRequest) Make() (req *http.Request, err error) {
	return nil, qt_errors.ErrorMock
}
