package response_impl

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"net/http"
)

func newErrorResponse(err error, res *http.Response) response.Response {
	if res != nil {
		headers := createHeader(res)
		headersLower := createHeaderLower(headers)
		return &errorResponse{
			err:         err,
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
	err         error
}

func (z errorResponse) Code() int {
	return z.code
}

func (z errorResponse) CodeCategory() response.CodeCategory {
	return response.CodeCategory(z.code / 100)
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

func (z errorResponse) Success() response.Body {
	return newEmptyBody()
}

func (z errorResponse) Alt() response.Body {
	return newEmptyBody()
}

func (z errorResponse) Error() error {
	return z.err
}
