package es_response_impl

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"net/http"
)

func NewNoResponse(err error) es_response.Response {
	return NewTransportErrorHttpResponse(err, nil)
}

func NewTransportErrorResponse(err error, res es_response.Response) es_response.Response {
	return newErrorResponse(
		ResErrTypeTransport,
		err,
		res.Proto(),
		res.Code(),
		res.Headers(),
	)
}

func NewAuthErrorResponse(err error, res es_response.Response) es_response.Response {
	return newErrorResponse(
		ResErrTypeAuth,
		err,
		res.Proto(),
		res.Code(),
		res.Headers(),
	)
}

func NewTransportErrorHttpResponse(err error, res *http.Response) es_response.Response {
	if res != nil {
		headers := createHeader(res)
		return newErrorResponse(
			ResErrTypeTransport,
			err,
			res.Proto,
			res.StatusCode,
			headers,
		)
	} else {
		return newErrorResponse(
			ResErrTypeTransport,
			err,
			"",
			-1,
			make(map[string]string),
		)
	}
}

func newErrorResponse(
	errType int,
	err error,
	proto string,
	code int,
	headers map[string]string,
) es_response.Response {
	lower := make(map[string]string)
	if 0 < len(headers) {
		lower = createHeaderLower(headers)
	}
	return &errorResponse{
		headers:     headers,
		headerLower: lower,
		code:        code,
		proto:       proto,
		err:         err,
		errType:     errType,
	}
}

const (
	ResErrTypeGeneral = iota
	ResErrTypeTransport
	ResErrTypeAuth
)

type errorResponse struct {
	headers     map[string]string
	headerLower map[string]string
	code        int
	proto       string
	err         error
	errType     int
}

func (z errorResponse) IsAuthInvalidToken() bool {
	return z.errType == ResErrTypeAuth
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
