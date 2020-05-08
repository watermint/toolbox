package nw_rest

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_http"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/infra/api/api_context"
)

// Assert broken response or rate limit for retry
type AssertResponse func(res es_response.Response) es_response.Response

type ClientOpts struct {
	Assert     AssertResponse
	Mock       bool
	ReplayMock []nw_replay.Response
}

func (z ClientOpts) Apply(opts ...ClientOpt) ClientOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		x, y := opts[0], opts[1:]
		return x(z).Apply(y...)
	}
}

type ClientOpt func(o ClientOpts) ClientOpts

func Mock() ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.Mock = true
		return o
	}
}

func ReplayMock(rm []nw_replay.Response) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.ReplayMock = rm
		return o
	}
}

func Assert(ar AssertResponse) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.Assert = ar
		return o
	}
}

func New(opts ...ClientOpt) nw_client.Rest {
	co := ClientOpts{}.Apply(opts...)
	var hc nw_client.Http
	switch {
	case co.Mock:
		hc = nw_http.Mock{}
	case len(co.ReplayMock) > 0:
		hc = nw_replay.NewReplay(co.ReplayMock)
	default:
		hc = nw_http.NewClient()
	}

	c0 := NewAssert(co.Assert, nw_capture.New(hc))
	return nw_retry.NewRetry(nw_retry.NewRatelimit(c0))
}

func NewAssert(assert AssertResponse, client nw_client.Rest) nw_client.Rest {
	return &AssertClient{
		assert: assert,
		client: client,
	}
}

type AssertClient struct {
	assert AssertResponse
	client nw_client.Rest
}

func (z AssertClient) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res es_response.Response) {
	res = z.client.Call(ctx, req)
	if !res.IsSuccess() && z.assert != nil {
		return z.assert(res)
	}
	return res
}
