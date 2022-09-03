package nw_assert

import (
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/network/nw_client"
)

// AssertResponse asserts broken response or rate limit for retry
type AssertResponse func(res es_response.Response) es_response.Response

type AssertClient struct {
	assert AssertResponse
	client nw_client.Rest
}

func (z AssertClient) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	res = z.client.Call(ctx, req)
	if !res.IsSuccess() && z.assert != nil {
		return z.assert(res)
	}
	return res
}

func NewAssert(assert AssertResponse, client nw_client.Rest) nw_client.Rest {
	return &AssertClient{
		assert: assert,
		client: client,
	}
}
