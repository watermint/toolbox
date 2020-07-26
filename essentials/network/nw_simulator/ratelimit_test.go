package nw_simulator

import (
	"github.com/watermint/toolbox/essentials/network/nw_retry"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestNarrowClient_Call(t *testing.T) {
	{
		nc := NewRateLimit(&PanicClient{}, 100, RetryAfterHeaderRetryAfter, NoDecorator)
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
		nc := NewRateLimit(&PanicClient{}, 100, RetryAfterHeaderGitHub, NoDecorator)
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
		nc := NewRateLimit(&PanicClient{}, 100, RetryAfterHeaderIetfDraftTimestamp, NoDecorator)
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
		nc := NewRateLimit(&PanicClient{}, 100, RetryAfterHeaderIetfDraftSecond, NoDecorator)
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
