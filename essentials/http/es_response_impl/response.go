package es_response_impl

import (
	"errors"
	"github.com/watermint/toolbox/essentials/http/es_context"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"net/http"
	"strconv"
	"strings"
)

func New(ctx es_context.Context, res *http.Response) es_response.Response {
	l := ctx.Log()
	l.Debug("Read response body")
	body, err := Read(ctx, res.Body)
	if err != nil {
		l.Debug("Error on transport")
		return NewTransportErrorHttpResponse(err, res)
	}

	headers := createHeader(res)
	switch res.StatusCode / 100 {
	case es_response.Code2xxSuccess:
		l.Debug("Success response")
		return newSuccessResponse(res.StatusCode, res.Proto, headers, body)

	default:
		l.Debug("Alternative response")
		return newAltResponse(res.StatusCode, res.Proto, headers, body)
	}
}

func newSuccessResponse(code int, proto string, headers map[string]string, body es_response.Body) es_response.Response {
	return &resImpl{
		code:         code,
		proto:        proto,
		headers:      headers,
		headersLower: createHeaderLower(headers),
		success:      body,
		alt:          newEmptyBody(),
		isSuccess:    true,
	}
}

func newAltResponse(code int, proto string, headers map[string]string, body es_response.Body) es_response.Response {
	return &resImpl{
		code:         code,
		proto:        proto,
		headers:      headers,
		headersLower: createHeaderLower(headers),
		success:      newEmptyBody(),
		alt:          body,
		isSuccess:    false,
	}
}

func createHeader(res *http.Response) map[string]string {
	h := make(map[string]string)
	for k, v := range res.Header {
		h[k] = v[0]
	}
	return h
}

func createHeaderLower(headers map[string]string) map[string]string {
	h := make(map[string]string)
	for k, v := range headers {
		h[strings.ToLower(k)] = v
	}
	return h
}

func IsTextContentType(res es_response.Response) bool {
	contentType := res.Header("Content-Type")
	if strings.HasPrefix(contentType, "text") {
		return true
	}
	switch contentType {
	case "application/json",
		"application/xml",
		"application/xhtml+xml":
		return true

	default:
		return false
	}
}

type resImpl struct {
	code         int
	proto        string
	headers      map[string]string
	headersLower map[string]string
	success      es_response.Body
	alt          es_response.Body
	isSuccess    bool
}

func (z resImpl) IsTextContentType() bool {
	return IsTextContentType(z)
}

func (z resImpl) Proto() string {
	return z.proto
}

func (z resImpl) Failure() (error, bool) {
	if z.isSuccess {
		return nil, false
	}
	st := http.StatusText(z.code)
	if st != "" {
		return errors.New(st), true
	}
	return errors.New("status code " + strconv.FormatInt(int64(z.code), 10)), true
}

func (z resImpl) TransportError() error {
	return nil
}

func (z resImpl) IsSuccess() bool {
	return z.isSuccess
}

func (z resImpl) Code() int {
	return z.code
}

func (z resImpl) CodeCategory() es_response.CodeCategory {
	return es_response.CodeCategory(z.code / 100)
}

func (z resImpl) Headers() map[string]string {
	return z.headers
}

func (z resImpl) Header(header string) string {
	return z.headersLower[strings.ToLower(header)]
}

func (z resImpl) Success() es_response.Body {
	return z.success
}

func (z resImpl) Alt() es_response.Body {
	return z.alt
}
