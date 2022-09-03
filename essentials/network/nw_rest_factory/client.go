package nw_rest_factory

import (
	"context"
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_assert"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_http"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/essentials/network/nw_simulator"
	"net/http"
	"time"
)

type ClientOpts struct {
	Assert     nw_assert.AssertResponse
	Mock       bool
	ReplayMock []nw_replay.Response

	oAuthApp         api_auth2.OAuthAppData
	oAuthKeyResolver api_auth2.OAuthKeyResolver
	oAuthEntity      api_auth2.OAuthEntity
	authFactory      func(client nw_client.Rest) (rest nw_client.Rest)

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

func OAuthEntity(app api_auth2.OAuthAppData, resolver api_auth2.OAuthKeyResolver, entity api_auth2.OAuthEntity) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.oAuthApp = app
		o.oAuthKeyResolver = resolver
		o.oAuthEntity = entity
		return o
	}
}

func Auth(factory func(client nw_client.Rest) (rest nw_client.Rest)) ClientOpt {
	return func(o ClientOpts) ClientOpts {
		o.authFactory = factory
		return o
	}
}

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

func Assert(ar nw_assert.AssertResponse) ClientOpt {
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

	co := ClientOpts{}
	co.oAuthEntity = api_auth2.NewNoAuthOAuthEntity()
	co = co.Apply(opts...)

	var hc nw_client.Http

	switch {
	case co.Mock:
		hc = nw_http.Mock{}
	case len(co.ReplayMock) > 0:
		hc = nw_replay.NewSequentialReplay(co.ReplayMock)
	default:
		if co.oAuthEntity.IsNoAuth() {
			hc = nw_http.NewClient(&http.Client{Timeout: 1 * time.Minute})
		} else {
			cfg := co.oAuthApp.Config(co.oAuthEntity.Scopes, co.oAuthKeyResolver)
			hc = nw_http.NewClient(
				cfg.Client(
					context.Background(),
					co.oAuthEntity.Token.OAuthToken(),
				),
			)
		}
	}

	var c0, c1, c2, c3, c4, c5 nw_client.Rest

	// Layer 0: capture
	c0 = nw_capture.New(hc)

	// Layer 1: auth
	if co.authFactory != nil {
		c1 = co.authFactory(c0)
	} else {
		c1 = c0
	}

	// Layer 2: rate limit simulator
	if co.rateLimitEnabled {
		l.Debug("Rate limit simulator enabled",
			esl.Int("Rate", co.rateLimitRate),
			esl.Int("HeaderType", int(co.rateLimitHeaderType)))
		c2 = nw_simulator.NewRateLimit(c1, co.rateLimitRate, co.rateLimitHeaderType, co.rateLimitDecorator)
	} else {
		c2 = c1
	}

	// Layer 3: server error simulator
	if co.serverErrorEnabled {
		l.Debug("Server error simulator enabled",
			esl.Int("Rate", co.serverErrorRate),
			esl.Int("Code", co.serverErrorCode))
		c3 = nw_simulator.NewServerError(c2, co.serverErrorRate, co.serverErrorCode, co.serverErrorDecorator)
	} else {
		c3 = c2
	}

	// Layer 4: assert
	c4 = nw_assert.NewAssert(co.Assert, c3)

	// Layer 5: rate limit
	c5 = nw_retry.NewRatelimit(c4)

	// Layer 6: retry
	return nw_retry.NewRetry(c5)
}
