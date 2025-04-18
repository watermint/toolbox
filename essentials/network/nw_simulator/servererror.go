package nw_simulator

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"

	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_throttle"
)

func NewServerError(client nw_client.Rest, rate, code int, decorator ResponseDecorator) nw_client.Rest {
	return &serverErrorClient{
		rate:      rate,
		code:      code,
		decorator: decorator,
		client:    client,
	}
}

type serverErrorClient struct {
	// too many requests error rate in percent
	rate int

	// error code
	code int

	// response decorator
	decorator ResponseDecorator

	// nested client
	client nw_client.Rest
}

func (z serverErrorClient) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	if rand.Intn(100) >= z.rate {
		return z.client.Call(ctx, req)
	} else {
		hr := &http.Response{}
		if z.code >= 500 && z.code < 600 {
			hr.StatusCode = z.code
		} else {
			hr.StatusCode = http.StatusInternalServerError
		}

		if z.decorator != nil {
			z.decorator(req.Endpoint(), hr)
		}
		if hr.Body == nil {
			hr.Body = io.NopCloser(&bytes.Buffer{})
		}

		nw_throttle.Throttle(ctx.ClientHash(), req.Endpoint(), func() {
			res = es_response_impl.New(ctx, hr)
		})
		return res
	}
}
