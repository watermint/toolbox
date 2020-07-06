package nw_simulator

import (
	"bytes"
	"fmt"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_concurrency"
	"github.com/watermint/toolbox/essentials/network/nw_ratelimit"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/infra/api/api_context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	maxRetryAfter = 4

	RetryAfterHeaderRetryAfter = iota
	RetryAfterHeaderGitHub
	RetryAfterHeaderIetfDraftSecond
	RetryAfterHeaderIetfDraftTimestamp
)

type RetryAfterHeaderType int
type ResponseDecorator func(res *http.Response)

func NoDecorator(res *http.Response) {
}

func New(client nw_client.Rest, rate int, headerType RetryAfterHeaderType, decorator ResponseDecorator) nw_client.Rest {
	return &narrowClient{
		rate:       rate,
		headerType: headerType,
		decorator:  decorator,
		client:     client,
	}
}

type narrowClient struct {
	// too many requests error rate in percent
	rate int

	// retry after header type
	headerType RetryAfterHeaderType

	// response decorator
	decorator ResponseDecorator

	// nested client
	client nw_client.Rest
}

func (z narrowClient) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res es_response.Response) {
	if rand.Intn(100) >= z.rate {
		return z.client.Call(ctx, req)
	} else {
		hr := &http.Response{}
		hr.StatusCode = http.StatusTooManyRequests
		hr.Header = make(map[string][]string)

		retryAfterSec := rand.Intn(maxRetryAfter) + 1

		switch z.headerType {
		case RetryAfterHeaderGitHub:
			hr.Header.Add(nw_retry.HeaderXRateLimitLimit, "100")
			hr.Header.Add(nw_retry.HeaderXRateLimitRemaining, "0")
			hr.Header.Add(nw_retry.HeaderXRateLimitReset, fmt.Sprintf("%d", time.Now().Add(time.Duration(retryAfterSec)*time.Second).Unix()))

		case RetryAfterHeaderIetfDraftTimestamp:
			hr.Header.Add(nw_retry.HeaderRateLimitLimit, "100")
			hr.Header.Add(nw_retry.HeaderRateLimitRemaining, "0")
			hr.Header.Add(nw_retry.HeaderRateLimitReset, time.Now().Add(time.Duration(retryAfterSec)*time.Second).Format(time.RFC1123))

		case RetryAfterHeaderIetfDraftSecond:
			hr.Header.Add(nw_retry.HeaderRateLimitLimit, "100")
			hr.Header.Add(nw_retry.HeaderRateLimitRemaining, "0")
			hr.Header.Add(nw_retry.HeaderRateLimitReset, fmt.Sprintf("%d", retryAfterSec))

		default:
			hr.Header.Add(nw_retry.HeaderRetryAfter, fmt.Sprintf("%d", retryAfterSec))
		}

		if z.decorator != nil {
			z.decorator(hr)
		}
		if hr.Body == nil {
			hr.Body = ioutil.NopCloser(&bytes.Buffer{})
		}

		nw_ratelimit.WaitIfRequired(ctx.ClientHash(), req.Endpoint())
		nw_concurrency.Start()
		res := es_response_impl.New(ctx, hr)
		nw_concurrency.End()
		return res
	}
}
