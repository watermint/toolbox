package nw_retry

import (
	"testing"
	"time"
)

func TestNewErrorRateLimit(t *testing.T) {
	reset := time.Now().Add(1 * 1000 * time.Millisecond)
	er := NewErrorRateLimit(2, 1, reset)
	if e, ok := er.(ErrorRateLimit); ok {
		if e.Limit != 2 || e.Remaining != 1 || e.Reset != reset {
			t.Error("invalid")
		}
	} else {
		t.Error("invalid")
	}
}

func TestNewErrorTransport(t *testing.T) {
	reset := time.Now().Add(1 * 1000 * time.Millisecond)
	er := NewErrorRateLimitResetOnly(reset)
	if e, ok := er.(ErrorRateLimit); ok {
		if e.Limit != 0 || e.Remaining != 0 || e.Reset != reset {
			t.Error("invalid")
		}
	} else {
		t.Error("invalid")
	}
}

func TestNewErrorRateLimitFromHeaders(t *testing.T) {
	// Retry-After : seconds
	{
		headers := map[string]string{
			"Retry-After": "3",
		}
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if ra.Reset.After(time.Now().Add(2*time.Second)) && ra.Reset.Before(time.Now().Add(5*time.Second)) {
			t.Error("invalid")
		}
		if ra.Remaining != 0 || ra.Limit != 0 {
			t.Error("invalid")
		}
	}

	// Retry-After : fix date
	{
		resetTime := "Mon, 20 Apr 2020 20:59:08 GMT"
		headers := map[string]string{
			"Retry-After": "Mon, 20 Apr 2020 20:59:08 GMT",
		}
		expected, err := time.Parse(time.RFC1123, resetTime)
		if err != nil {
			t.Error(err)
		}
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if !ra.Reset.Equal(expected) {
			t.Error("invalid")
		}
		if ra.Remaining != 0 || ra.Limit != 0 {
			t.Error("invalid")
		}
	}

	// RateLimit-Reset, in seconds
	{
		headers := map[string]string{
			"RateLimit-Reset":     "3",
			"RateLimit-Limit":     "10",
			"RateLimit-Remaining": "4",
		}
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if ra.Reset.After(time.Now().Add(2*time.Second)) && ra.Reset.Before(time.Now().Add(5*time.Second)) {
			t.Error("invalid")
		}
		if ra.Remaining != 4 || ra.Limit != 10 {
			t.Error("invalid")
		}
	}

	// RateLimit-Reset, in seconds
	{
		headers := map[string]string{
			"RateLimit-Reset":     "3",
			"RateLimit-Limit":     "10, 10;window=1;burst=1000, 1000;window=3600",
			"RateLimit-Remaining": "4",
		}
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if ra.Reset.After(time.Now().Add(2*time.Second)) && ra.Reset.Before(time.Now().Add(5*time.Second)) {
			t.Error("invalid")
		}
		if ra.Remaining != 4 || ra.Limit != 10 {
			t.Error("invalid")
		}
	}

	// RateLimit-Reset, in seconds
	{
		headers := map[string]string{
			"X-RateLimit-Reset":     "3",
			"X-RateLimit-Limit":     "10, 10;window=1;burst=1000, 1000;window=3600",
			"X-RateLimit-Remaining": "4",
		}
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if ra.Reset.After(time.Now().Add(2*time.Second)) && ra.Reset.Before(time.Now().Add(5*time.Second)) {
			t.Error("invalid")
		}
		if ra.Remaining != 4 || ra.Limit != 10 {
			t.Error("invalid")
		}
	}

	// RateLimit-Reset, in unix time
	{
		headers := map[string]string{
			"X-RateLimit-Reset":     "1587340800", // 2020-04-20T00:00:00Z
			"X-RateLimit-Limit":     "10, 10;window=1;burst=1000, 1000;window=3600",
			"X-RateLimit-Remaining": "4",
		}
		expected := time.Unix(1587340800, 0)
		ra, found := NewErrorRateLimitFromHeaders(headers)
		if !found {
			t.Error("invalid")
		}
		if !ra.Reset.Equal(expected) {
			t.Error("invalid")
		}
		if ra.Remaining != 4 || ra.Limit != 10 {
			t.Error("invalid")
		}
	}
}
