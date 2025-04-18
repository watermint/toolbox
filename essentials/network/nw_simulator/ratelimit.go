package nw_simulator

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/essentials/network/nw_throttle"
)

const (
	maxRetryAfter = 4

	RetryAfterHeaderRetryAfter = iota
	RetryAfterHeaderGitHub
	RetryAfterHeaderIetfDraftSecond
	RetryAfterHeaderIetfDraftTimestamp
)

type RetryAfterHeaderType int
type ResponseDecorator func(endpoint string, res *http.Response)

func NoDecorator(endpoint string, res *http.Response) {
}

func NewRateLimit(client nw_client.Rest, rate int, headerType RetryAfterHeaderType, decorator ResponseDecorator) nw_client.Rest {
	return &rateLimitClient{
		rate:       rate,
		headerType: headerType,
		decorator:  decorator,
		client:     client,
	}
}

type rateLimitClient struct {
	// too many requests error rate in percent
	rate int

	// retry after header type
	headerType RetryAfterHeaderType

	// response decorator
	decorator ResponseDecorator

	// nested client
	client nw_client.Rest
}

func (z rateLimitClient) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
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
