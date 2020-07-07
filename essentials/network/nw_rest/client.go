package nw_rest

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_http"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/essentials/network/nw_simulator"
	"github.com/watermint/toolbox/infra/api/api_context"
)

// Assert broken response or rate limit for retry
type AssertResponse func(res es_response.Response) es_response.Response

type ClientOpts struct {
	Assert     AssertResponse
	Mock       bool
	ReplayMock []nw_replay.Response

	// rate limit simulator
	rateLimitRate       int
	rateLimitDecorator  nw_simulator.ResponseDecorator
	rateLimitHeaderType nw_simulator.RetryAfterHeaderType
	rateLimitEnabled    bool

	// server error simulator
	serverErrorRate      int
	serverErrorDecorator nw_simulator.ResponseDecorator
	serverErrorCode      int
	serverErrorEnabled   bool
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

func RateLimitSimulator(rate int, headerType nw_simulator.RetryAfterHeaderType, decorator nw_simulator.ResponseDecorator) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.rateLimitRate = rate
		o.rateLimitHeaderType = headerType
		o.rateLimitDecorator = decorator
		o.rateLimitEnabled = true
		return o
	}
}

func ServerErrorSimulator(rate int, code int, decorator nw_simulator.ResponseDecorator) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.serverErrorEnabled = true
		o.serverErrorRate = rate
		o.serverErrorCode = code
		o.serverErrorDecorator = decorator
		return o
	}
}

func New(opts ...ClientOpt) nw_client.Rest {
	l := esl.Default()

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

	var c0, c1, c2, c3 nw_client.Rest

	// Layer 0: capture
	c0 = nw_capture.New(hc)

	// Layer 1: rate limit simulator
	if co.rateLimitEnabled {
		l.Debug("Rate limit simulator enabled",
			esl.Int("Rate", co.rateLimitRate),
			esl.Int("HeaderType", int(co.rateLimitHeaderType)))
		c1 = nw_simulator.NewRateLimit(c0, co.rateLimitRate, co.rateLimitHeaderType, co.rateLimitDecorator)
	} else {
		c1 = c0
	}

	// Layer 2: server error simulator
	if co.serverErrorEnabled {
		l.Debug("Server error simulator enabled",
			esl.Int("Rate", co.serverErrorRate),
			esl.Int("Code", co.serverErrorCode))
		c2 = nw_simulator.NewServerError(c1, co.serverErrorRate, co.serverErrorCode, co.serverErrorDecorator)
	} else {
		c2 = c1
	}

	// Layer 3: assert
	c3 = NewAssert(co.Assert, c2)

	// Layer 4: retry
	return nw_retry.NewRetry(nw_retry.NewRatelimit(c3))
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
