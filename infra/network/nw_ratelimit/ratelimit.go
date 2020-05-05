package nw_ratelimit

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"sync"
	"time"
)

const (
	MaxLastErrors                  = 100
	SameErrorThreshold             = 80
	LastErrorRecycleTimeMultiplier = 100
	ShortWait                      = 3 * 1000 * time.Millisecond
	LongWait                       = 1 * time.Minute
	LongWaitAlertThreshold         = 5 * time.Minute
)

const (
	RetryActionShortWait = iota
	RetryActionLongWait
	RetryActionAbort
)

var (
	rateLimitState = newLimitState()
	isTest         = false
)

func SetTestMode(enabled bool) {
	isTest = enabled
}
func AddError(hash, endpoint string, err error) (abort bool) {
	return rateLimitState.AddError(hash, endpoint, err)
}
func UpdateRetryAfter(hash, endpoint string, retryAfter time.Time) {
	rateLimitState.UpdateRetryAfter(hash, endpoint, retryAfter)
}
func WaitIfRequired(hash, endpoint string) {
	rateLimitState.WaitIfRequired(hash, endpoint)
}

type LimitState interface {
	// Add error that happened with request. Blocks operation for retry SLA.
	// When abort == true, recommend abort retry loop if certain threshold exceed by given err.
	AddError(hash, endpoint string, err error) (abort bool)

	// Update retry after, if Rewind-After header present in the response
	UpdateRetryAfter(hash, endpoint string, retryAfter time.Time)

	// Wait if required. Waits for Max(Rewind-After, RetryActionWait).
	WaitIfRequired(hash, endpoint string)
}

func newLimitState() LimitState {
	return &limitStateImpl{
		maxLastErrors:              MaxLastErrors,
		sameErrorThreshold:         SameErrorThreshold,
		durationShortWait:          ShortWait,
		durationLongWait:           LongWait,
		lastErrorRecycleMultiplier: LastErrorRecycleTimeMultiplier,
		lastErrorDict:              make(map[string][]error),
		lastErrorTime:              make(map[string]time.Time),
		retryAfter:                 make(map[string]time.Time),
		retryAction:                make(map[string]int),
	}
}

type limitStateImpl struct {
	maxLastErrors              int
	sameErrorThreshold         int
	durationShortWait          time.Duration
	durationLongWait           time.Duration
	lastErrorRecycleMultiplier int
	lastErrorDict              map[string][]error
	lastErrorTime              map[string]time.Time
	lastErrorMutex             sync.Mutex
	retryAction                map[string]int
	retryActionMutex           sync.Mutex
	retryAfter                 map[string]time.Time
	retryAfterMutex            sync.Mutex
}

func (z *limitStateImpl) keyHash(hash, endpoint string) string {
	return hash + endpoint
}

func (z *limitStateImpl) retryActionFor(key string) (action int, wait, recycle time.Duration) {
	durForAction := func(a int) (wait, recycle time.Duration) {
		switch a {
		case RetryActionShortWait:
			return z.durationShortWait, z.durationShortWait * time.Duration(z.lastErrorRecycleMultiplier)
		case RetryActionLongWait:
			return z.durationLongWait, z.durationLongWait * time.Duration(z.lastErrorRecycleMultiplier)
		default:
			return z.durationLongWait, z.durationLongWait * time.Duration(z.lastErrorRecycleMultiplier)
		}
	}

	z.retryActionMutex.Lock()
	defer z.retryActionMutex.Unlock()

	if ra, ok := z.retryAction[key]; ok {
		w, r := durForAction(ra)
		return ra, w, r
	} else {
		ra = RetryActionShortWait
		z.retryAction[key] = ra
		w, r := durForAction(ra)
		return ra, w, r
	}
}

func (z *limitStateImpl) logger(hash, endpoint string) es_log.Logger {
	return es_log.Default().With(
		es_log.String("hash", hash),
		es_log.String("endpoint", endpoint),
		es_log.String("Routine", ut_runtime.GetGoRoutineName()),
	)
}

func (z *limitStateImpl) AddError(hash, endpoint string, err error) (abort bool) {
	l := z.logger(hash, endpoint)
	if err == nil {
		l.Debug("Skip adding error (nil)")
		return false
	}

	key := z.keyHash(hash, endpoint)

	// promote to next retryAction level:
	retryActionPromote := false
	retryAction, wait, recycle := z.retryActionFor(key)

	z.lastErrorMutex.Lock()
	defer z.lastErrorMutex.Unlock()
	z.retryActionMutex.Lock()
	defer z.retryActionMutex.Unlock()

	purgeLastError := false
	if let, ok := z.lastErrorTime[key]; ok {
		if let.Add(recycle).Before(time.Now()) {
			purgeLastError = true
		}
	}
	if purgeLastError {
		z.lastErrorDict[key] = make([]error, 0)
		z.retryAction[key] = RetryActionShortWait
		retryAction = RetryActionShortWait
		wait = 0
	}
	le, ok := z.lastErrorDict[key]

	switch {
	case purgeLastError, !ok:
		le := make([]error, 0)
		le = append(le, err)
		z.lastErrorDict[key] = le

	default:
		if len(le) >= z.maxLastErrors {
			le = le[1:]
		}
		le = append(le, err)
		z.lastErrorDict[key] = le
		numSameErrors := 0
		for _, e := range le {
			if e.Error() == err.Error() {
				numSameErrors++
			}
		}
		if numSameErrors > z.sameErrorThreshold {
			retryActionPromote = true
		}
	}
	z.lastErrorTime[key] = time.Now()

	if retryActionPromote {
		switch retryAction {
		case RetryActionShortWait:
			retryAction = RetryActionLongWait
		case RetryActionLongWait, RetryActionAbort:
			retryAction = RetryActionAbort
		}
		z.retryAction[key] = retryAction
		z.lastErrorDict[key] = make([]error, 0)
	}

	abort = retryAction == RetryActionAbort
	l.Debug("Wait for SLA",
		es_log.Error(err),
		es_log.Int("retryAction", retryAction),
		es_log.Bool("retryActionPromote", retryActionPromote),
		es_log.Bool("abort", abort),
		es_log.Bool("purgeLastError", purgeLastError),
		es_log.String("wait", wait.String()),
	)
	z.UpdateRetryAfter(hash, endpoint, time.Now().Add(wait))
	time.Sleep(wait)

	return abort
}

func (z *limitStateImpl) UpdateRetryAfter(hash, endpoint string, retryAfter time.Time) {
	z.retryAfterMutex.Lock()
	defer z.retryAfterMutex.Unlock()

	key := z.keyHash(hash, endpoint)
	ra, ok := z.retryAfter[key]
	if !ok || ra.After(retryAfter) {
		z.retryAfter[key] = retryAfter
	}
}

func (z *limitStateImpl) WaitIfRequired(hash, endpoint string) {
	l := z.logger(hash, endpoint)
	key := z.keyHash(hash, endpoint)

	z.retryAfterMutex.Lock()
	retryAfter, ok := z.retryAfter[key]
	z.retryAfterMutex.Unlock()

	if ok {
		dur := retryAfter.Sub(time.Now())
		if dur > LongWaitAlertThreshold {
			l.Warn("Waiting for server rate limit", es_log.String("duration", dur.String()))
			if isTest {
				panic("server rate limit exceeds threshold")
			}
		}
		if dur > 0 {
			l.Debug("Waiting",
				es_log.String("retryAfter", retryAfter.String()),
				es_log.String("duration", dur.String()),
			)
			time.Sleep(dur)
		}
	}
}
