package response

import (
	"github.com/watermint/toolbox/essentials/http/context"
	"net/http"
	"strings"
)

const (
	Code1xxInformational = 1
	Code2xxSuccess       = 2
	Code3xxRedirection   = 3
	Code4xxClientErrors  = 4
	Code5xxServerErrors  = 5
)

type CodeCategory int

type Response interface {
	// Status code.
	Code() int

	// Status code category.
	CodeCategory() CodeCategory

	// Response headers.
	Headers() map[string]string

	// Get header value. Ignore cases.
	// Returns empty string, if no header found in the response.
	Header(header string) string

	// Response body.
	Body() Body
}

func New(ctx context.Context, res *http.Response) Response {
	body := Read(ctx, res.Body)
	return &resImpl{
		code:         res.StatusCode,
		headers:      createHeader(res),
		headersLower: createHeaderLower(res),
		body:         body,
	}
}

func createHeader(res *http.Response) map[string]string {
	h := make(map[string]string)
	for k, v := range res.Header {
		h[k] = v[0]
	}
	return h
}

func createHeaderLower(res *http.Response) map[string]string {
	h := make(map[string]string)
	for k, v := range res.Header {
		h[strings.ToLower(k)] = v[0]
	}
	return h
}

type resImpl struct {
	code         int
	headers      map[string]string
	headersLower map[string]string
	body         Body
}

func (z resImpl) Code() int {
	return z.code
}

func (z resImpl) CodeCategory() CodeCategory {
	return CodeCategory(z.code / 100)
}

func (z resImpl) Headers() map[string]string {
	return z.headers
}

func (z resImpl) Header(header string) string {
	return z.headersLower[strings.ToLower(header)]
}

func (z resImpl) Body() Body {
	return z.body
}
