package nw_retry

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

const (
	// Retry-After: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
	HeaderRetryAfter = "Retry-After"

	// GitHub API
	HeaderXRateLimitLimit     = "X-RateLimit-Limit"
	HeaderXRateLimitRemaining = "X-RateLimit-Remaining"
	HeaderXRateLimitReset     = "X-RateLimit-Reset"

	// https://tools.ietf.org/id/draft-polli-ratelimit-headers-00.html
	HeaderRateLimitLimit     = "RateLimit-Limit"
	HeaderRateLimitRemaining = "RateLimit-Remaining"
	HeaderRateLimitReset     = "RateLimit-Reset"

	// Default time of retry-after, if the value is was not found or invalid.
	DefaultRetryAfterSec = 10

	// If the value exceeds below time (2000-01-01T00:00:00Z), assume the value is an unix time.
	RetryAfterSecOrFixDate = 946684800
)

func newErrorRateLimitFromHeadersRetryAfter(retryAfter string) *ErrorRateLimit {
	rat := strings.TrimSpace(retryAfter)
	l := app_root.Log().With(zap.String("retryAfter", retryAfter))

	if ra, err := strconv.ParseInt(rat, 10, 64); err == nil {
		if ra > RetryAfterSecOrFixDate {
			reset := time.Unix(ra, 0)
			l.Debug("Retry after fix date", zap.String("reset", reset.Format(time.RFC3339)))
			return &ErrorRateLimit{
				Reset: reset,
			}
		} else {
			l.Debug("Retry after second", zap.Int64("resetSec", ra))
			return &ErrorRateLimit{
				Reset: time.Now().Add(time.Duration(ra) * time.Second),
			}
		}
	}
	if reset, err := time.Parse(time.RFC1123, rat); err == nil {
		l.Debug("Retry after fix date", zap.String("reset", reset.Format(time.RFC3339)))
		return &ErrorRateLimit{
			Reset: reset,
		}
	}
	l.Debug("Unable to determine value for `Retry-after`. Fallback to default retry-after-sec", zap.Int("defaultRetryAfter", DefaultRetryAfterSec))
	return &ErrorRateLimit{
		Reset: time.Now().Add(DefaultRetryAfterSec * time.Second),
	}
}

func parseRateLimitQuota(limit string) int {
	els := strings.Split(limit, ",")
	if el, err := strconv.Atoi(strings.TrimSpace(els[0])); err != nil {
		return 0
	} else {
		return el
	}
}
func NewErrorRateLimitFromHeadersFallback(headers map[string]string) (erl *ErrorRateLimit) {
	if erl, found := NewErrorRateLimitFromHeaders(headers); found {
		return erl
	} else {
		return &ErrorRateLimit{
			Reset: time.Now().Add(DefaultRetryAfterSec * time.Second),
		}
	}
}

func NewErrorRateLimitFromHeaders(headers map[string]string) (erl *ErrorRateLimit, found bool) {
	headerLower := make(map[string]string)
	for k, v := range headers {
		headerLower[strings.ToLower(k)] = v
	}

	if retryAfter, ok := headerLower[strings.ToLower(HeaderRetryAfter)]; ok {
		return newErrorRateLimitFromHeadersRetryAfter(retryAfter), true
	}

	if reset, ok := headerLower[strings.ToLower(HeaderRateLimitReset)]; ok {
		e := newErrorRateLimitFromHeadersRetryAfter(reset)
		e.Limit = parseRateLimitQuota(headerLower[strings.ToLower(HeaderRateLimitLimit)])
		e.Remaining = parseRateLimitQuota(headerLower[strings.ToLower(HeaderRateLimitRemaining)])
		return e, true
	}

	if reset, ok := headerLower[strings.ToLower(HeaderXRateLimitReset)]; ok {
		e := newErrorRateLimitFromHeadersRetryAfter(reset)
		e.Limit = parseRateLimitQuota(headerLower[strings.ToLower(HeaderXRateLimitLimit)])
		e.Remaining = parseRateLimitQuota(headerLower[strings.ToLower(HeaderXRateLimitRemaining)])
		return e, true
	}

	return nil, false
}

func NewErrorRateLimit(limit, remaining int, reset time.Time) error {
	return &ErrorRateLimit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}
}
func NewErrorRateLimitResetOnly(reset time.Time) error {
	return &ErrorRateLimit{
		Reset: reset,
	}
}

type ErrorRateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

func (z ErrorRateLimit) Error() string {
	return "exceeds rate limit rule"
}
