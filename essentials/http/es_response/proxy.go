package es_response

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
)

// Create new instance of the proxy instance.
func NewProxy(res Response) Proxy {
	if res == nil {
		es_log.Default().Error("null response")
	}
	return Proxy{
		res: res,
	}
}

// Proxy implementation of Response
type Proxy struct {
	res Response
}

func (z Proxy) Failure() (error, bool) {
	return z.res.Failure()
}

func (z Proxy) Code() int {
	return z.res.Code()
}

func (z Proxy) CodeCategory() CodeCategory {
	return z.res.CodeCategory()
}

func (z Proxy) Headers() map[string]string {
	return z.res.Headers()
}

func (z Proxy) Header(header string) string {
	return z.res.Header(header)
}

func (z Proxy) IsSuccess() bool {
	return z.res.IsSuccess()
}

func (z Proxy) Success() Body {
	return z.res.Success()
}

func (z Proxy) Alt() Body {
	return z.res.Alt()
}

func (z Proxy) TransportError() error {
	return z.res.TransportError()
}
