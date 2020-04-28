package nw_rest

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/network/nw_capture"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/network/nw_http"
	"github.com/watermint/toolbox/infra/network/nw_retry"
)

// Assert broken response or rate limit for retry
type AssertResponse func(res es_response.Response) es_response.Response

type ClientOpts struct {
	Assert AssertResponse
	Mock   bool
}

type ClientOpt func(o ClientOpts) ClientOpts

func Mock() ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.Mock = true
		return o
	}
}
func Assert(ar AssertResponse) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.Assert = ar
		return o
	}
}

func New(feature app_feature.Feature, opts ...ClientOpt) nw_client.Rest {
	co := ClientOpts{}
	for _, o := range opts {
		co = o(co)
	}
	var hc nw_client.Http
	if co.Mock {
		hc = nw_http.Mock{}
	} else {
		hc = nw_http.NewClient()
	}

	c0 := NewAssert(co.Assert, nw_capture.New(hc))
	return nw_retry.NewRetry(c0)
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
