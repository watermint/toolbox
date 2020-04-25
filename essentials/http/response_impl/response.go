package response_impl

import (
	"github.com/watermint/toolbox/essentials/http/context"
	"github.com/watermint/toolbox/essentials/http/response"
	"net/http"
	"strings"
)

func New(ctx context.Context, res *http.Response) response.Response {
	l := ctx.Log()
	l.Debug("Read response body")
	body, err := Read(ctx, res.Body)
	if err != nil {
		l.Debug("Error on transport")
		return newErrorResponse(err, res)
	}

	headers := createHeader(res)
	switch res.StatusCode / 100 {
	case response.Code2xxSuccess:
		l.Debug("Success response")
		return newSuccessResponse(res.StatusCode, headers, body)

	default:
		l.Debug("Alternative response")
		return newAltResponse(res.StatusCode, headers, body)
	}
}

func newSuccessResponse(code int, headers map[string]string, body response.Body) response.Response {
	return &resImpl{
		code:         code,
		headers:      headers,
		headersLower: createHeaderLower(headers),
		success:      body,
		alt:          newEmptyBody(),
		isSuccess:    true,
	}
}

func newAltResponse(code int, headers map[string]string, body response.Body) response.Response {
	return &resImpl{
		code:         code,
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

type resImpl struct {
	code         int
	headers      map[string]string
	headersLower map[string]string
	success      response.Body
	alt          response.Body
	isSuccess    bool
}

func (z resImpl) Error() error {
	return nil
}

func (z resImpl) IsSuccess() bool {
	return z.isSuccess
}

func (z resImpl) Code() int {
	return z.code
}

func (z resImpl) CodeCategory() response.CodeCategory {
	return response.CodeCategory(z.code / 100)
}

func (z resImpl) Headers() map[string]string {
	return z.headers
}

func (z resImpl) Header(header string) string {
	return z.headersLower[strings.ToLower(header)]
}

func (z resImpl) Success() response.Body {
	return z.success
}

func (z resImpl) Alt() response.Body {
	return z.alt
}
