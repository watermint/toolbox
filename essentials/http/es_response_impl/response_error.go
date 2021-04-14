package es_response_impl

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"net/http"
)

func NewNoResponse(err error) es_response.Response {
	return NewTransportErrorHttpResponse(err, nil)
}

func NewTransportErrorResponse(err error, res es_response.Response) es_response.Response {
	return &errorResponse{
		err:         err,
		proto:       res.Proto(),
		code:        res.Code(),
		headers:     res.Headers(),
		headerLower: createHeaderLower(res.Headers()),
	}
}

func NewTransportErrorHttpResponse(err error, res *http.Response) es_response.Response {
	if res != nil {
		headers := createHeader(res)
		headersLower := createHeaderLower(headers)
		return &errorResponse{
			err:         err,
			proto:       res.Proto,
			code:        res.StatusCode,
			headers:     headers,
			headerLower: headersLower,
		}
	} else {
		return &errorResponse{
			err:         err,
			code:        -1,
			headers:     make(map[string]string),
			headerLower: make(map[string]string),
		}
	}
}

type errorResponse struct {
	headers     map[string]string
	headerLower map[string]string
	code        int
	proto       string
	err         error
}

func (z errorResponse) IsTextContentType() bool {
	return IsTextContentType(z)
}

func (z errorResponse) Proto() string {
	return z.proto
}

func (z errorResponse) Failure() (error, bool) {
	return z.err, true
}

func (z errorResponse) Code() int {
	return z.code
}

func (z errorResponse) CodeCategory() es_response.CodeCategory {
	return es_response.CodeCategory(z.code / 100)
}

func (z errorResponse) Headers() map[string]string {
	return z.headers
}

func (z errorResponse) Header(header string) string {
	return z.headerLower[header]
}

func (z errorResponse) IsSuccess() bool {
	return false
}

func (z errorResponse) Success() es_response.Body {
	return newEmptyBody()
}

func (z errorResponse) Alt() es_response.Body {
	return newEmptyBody()
}

func (z errorResponse) TransportError() error {
	return z.err
}
