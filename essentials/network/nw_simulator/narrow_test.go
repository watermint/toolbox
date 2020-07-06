package nw_simulator

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"github.com/watermint/toolbox/infra/api/api_context"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type PanicClient struct {
}

func (p PanicClient) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res es_response.Response) {
	panic("always panic!")
}

type MockApiContext struct {
}

func (z MockApiContext) ClientHash() string {
	return ""
}

func (z MockApiContext) Log() esl.Logger {
	return esl.Default()
}

func (z MockApiContext) Capture() esl.Logger {
	return esl.Default()
}

type MockReqBuilder struct {
}

func (z MockReqBuilder) Build() (*http.Request, error) {
	return http.NewRequest("POST", z.Endpoint(), &bytes.Buffer{})
}

func (z MockReqBuilder) Endpoint() string {
	return "http://www.example.com"
}

func (z MockReqBuilder) Param() string {
	return ""
}

func TestNarrowClient_Call(t *testing.T) {
	{
		nc := New(&PanicClient{}, 100, RetryAfterHeaderRetryAfter, NoDecorator)
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusTooManyRequests {
			t.Error(res.Code())
		}
		v := res.Header(nw_retry.HeaderRetryAfter)
		if va, err := strconv.Atoi(v); err != nil || va < 1 {
			t.Error(err, va)
		}
	}

	{
		nc := New(&PanicClient{}, 100, RetryAfterHeaderGitHub, NoDecorator)
		now := time.Now().Unix()
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusTooManyRequests {
			t.Error(res.Code())
		}
		v := res.Header(nw_retry.HeaderXRateLimitReset)
		if va, err := strconv.Atoi(v); err != nil || int64(va) < now {
			t.Error(err, va)
		}
	}

	{
		nc := New(&PanicClient{}, 100, RetryAfterHeaderIetfDraftTimestamp, NoDecorator)
		now := time.Now()
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusTooManyRequests {
			t.Error(res.Code())
		}
		v := res.Header(nw_retry.HeaderRateLimitReset)
		if va, err := time.Parse(time.RFC1123, v); err != nil || va.Before(now) {
			t.Error(err, va)
		}
	}

	{
		nc := New(&PanicClient{}, 100, RetryAfterHeaderIetfDraftSecond, NoDecorator)
		res := nc.Call(&MockApiContext{}, &MockReqBuilder{})
		if res.IsSuccess() {
			t.Error(res.IsSuccess())
		}
		if res.Code() != http.StatusTooManyRequests {
			t.Error(res.Code())
		}
		v := res.Header(nw_retry.HeaderRateLimitReset)
		if va, err := strconv.Atoi(v); err != nil || va < 1 {
			t.Error(err, va)
		}
	}
}
