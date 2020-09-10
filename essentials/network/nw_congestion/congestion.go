package nw_congestion

import (
	"container/list"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"runtime"
	"sync"
	"time"
)

const (
	monitorDuration = 1 * time.Minute
)

var (
	maxCongestionWindow  = runtime.NumCPU()
	initCongestionWindow = 4
	minCongestionWindow  = 1
	currentImpl          = NewControl()
)

func Start(hash, endpoint string) {
	currentImpl.Start(hash, endpoint)
}

func EndSuccess(hash, endpoint string) {
	currentImpl.EndSuccess(hash, endpoint)
}

func EndTransportError(hash, endpoint string) {
	currentImpl.EndTransportError(hash, endpoint)
}

func EndRateLimit(hash, endpoint string) {
	currentImpl.EndRateLimit(hash, endpoint)
}

func SetMaxCongestionWindow(w int) {
	maxCongestionWindow = w
}

// Maximum congestion window size
func CurrentMaxCongestionWindow() int {
	return maxCongestionWindow
}

// Minimum congestion window size
func CurrentMinCongestionWindow() int {
	return minCongestionWindow
}

// Initial congestion window size
func CurrentInitCongestionWindow() int {
	return initCongestionWindow
}

type CongestionControl interface {
	// start and wait for QoS
	Start(hash, endpoint string)

	// mark transaction as success
	EndSuccess(hash, endpoint string)

	// mark transaction as transport error
	EndTransportError(hash, endpoint string)

	// mark transaction as failure and got a rate limit
	EndRateLimit(hash, endpoint string)
}

func NewControl() CongestionControl {
	return &ccImpl{
		window:        make(map[string]int),
		concurrency:   make(map[string]int),
		lastRateLimit: make(map[string]time.Time),
		lastIncrement: make(map[string]time.Time),
		lastDecrement: make(map[string]time.Time),
	}
}

// ref golang.org/x/sync/semaphore
type ccWaiter struct {
	key   string
	ready chan<- struct{}
}

type ccImpl struct {
	window        map[string]int
	concurrency   map[string]int
	lastRateLimit map[string]time.Time
	lastIncrement map[string]time.Time
	lastDecrement map[string]time.Time
	waiters       list.List
	mutex         sync.Mutex
}

func (z *ccImpl) key(hash, endpoint string) string {
	return hash + "-" + endpoint
}

func (z *ccImpl) Start(hash, endpoint string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("hash", hash), esl.String("endpoint", endpoint))

	key := z.key(hash, endpoint)
	z.mutex.Lock()

	if wnd, ok := z.window[key]; !ok {
		wnd = CurrentInitCongestionWindow()
		l.Debug("Congestion window not found, create new window", esl.Int("window", wnd))
		z.window[key] = wnd
		z.concurrency[key] = 1
		z.mutex.Unlock()
		return
	} else {
		concurrency := z.concurrency[key]
		if concurrency < wnd {
			l.Debug("There is available window",
				esl.Int("window", wnd), esl.Int("concurrency", concurrency))
			z.concurrency[key] = concurrency + 1
			z.mutex.Unlock()
			return
		}

		l.Debug("There is no window available now. Wait for another process.", esl.Int("window", wnd), esl.Int("concurrency", concurrency))
		ready := make(chan struct{})
		waiter := &ccWaiter{ready: ready, key: key}
		z.waiters.PushBack(waiter)
		z.mutex.Unlock()

		select {
		case <-ready:
			l.Debug("Lock acquired.")
			z.mutex.Lock()
			if concurrency, ok := z.concurrency[key]; ok {
				z.concurrency[key] = es_number.Max(concurrency-1, 0).Int()
			}
			z.mutex.Unlock()
			return
		}
	}
}

func (z *ccImpl) EndSuccess(hash, endpoint string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("hash", hash), esl.String("endpoint", endpoint))
	l.Debug("Process finished: success")
	key := z.key(hash, endpoint)
	z.mutex.Lock()
	z.noLockRelease(key)

	if z.noLockCanIncreaseWindow(key) {
		if wnd, ok := z.window[key]; ok {
			newWindow := es_number.Min(wnd+1, CurrentMaxCongestionWindow()).Int()
			z.lastIncrement[key] = time.Now()
			z.window[key] = newWindow
			l.Debug("Increase window", esl.Int("newWindow", newWindow))
		}
	}

	z.mutex.Unlock()
}

func (z *ccImpl) EndTransportError(hash, endpoint string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("hash", hash), esl.String("endpoint", endpoint))

	l.Debug("Transport error; do not change window")
	key := z.key(hash, endpoint)
	z.mutex.Lock()
	z.noLockRelease(key)
	z.mutex.Unlock()
}

func (z *ccImpl) EndRateLimit(hash, endpoint string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("hash", hash), esl.String("endpoint", endpoint))

	key := z.key(hash, endpoint)
	z.mutex.Lock()
	z.noLockRelease(key)

	calcNewWindow := func(key string) int {
		if wnd, ok := z.window[key]; ok {
			nw := es_number.Max(wnd-1,
				CurrentMinCongestionWindow()).Int()
			l.Debug("Slow start. New window", esl.Int("newWindow", nw), esl.Int("prevWindow", wnd))
			return nw
		} else {
			nw := es_number.Max(CurrentMaxCongestionWindow()/2,
				CurrentMinCongestionWindow()).Int()
			l.Debug("Current window not found (strange). But assume max window as current window", esl.Int("newWindow", nw))
			return nw
		}
	}

	monitorTime := time.Now().Add(-monitorDuration)
	var newWindow int
	if lastDecrement, ok := z.lastDecrement[key]; ok {
		if lastDecrement.Before(monitorTime) {
			newWindow = calcNewWindow(key)
		} else {
			l.Debug("Last decrease is before monitor time", esl.Time("monitorTime", monitorTime), esl.Time("lastDecrease", lastDecrement))
		}
	} else {
		newWindow = calcNewWindow(key)
	}

	z.lastRateLimit[key] = time.Now()
	if newWindow > 0 {
		z.window[key] = newWindow
		z.lastDecrement[key] = time.Now()
	}
	z.mutex.Unlock()
}

func (z *ccImpl) noLockCanIncreaseWindow(key string) bool {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("key", key))

	monitorTime := time.Now().Add(-monitorDuration)
	l.Debug("Verify", esl.Time("monitor", monitorTime))
	if lastIncrement, ok := z.lastIncrement[key]; ok {
		if lastIncrement.Before(monitorTime) {
			if lastRateLimit, ok := z.lastRateLimit[key]; ok {
				if lastRateLimit.Before(monitorTime) {
					l.Debug("Both last increment & last rate limit ok",
						esl.Time("lastIncrement", lastIncrement),
						esl.Time("lastRateLimit", lastRateLimit))
					return true
				} else {
					l.Debug("Last increment ok, but wait for last rate limit policy",
						esl.Time("lastIncrement", lastIncrement),
						esl.Time("lastRateLimit", lastRateLimit))
					return false
				}
			} else {
				l.Debug("Last increment ok",
					esl.Time("lastIncrement", lastIncrement),
					esl.Time("lastRateLimit", lastRateLimit))
				return true
			}
		} else {
			l.Debug("Last increment is before monitor time",
				esl.Time("lastIncrement", lastIncrement))
			return false
		}
	} else {
		l.Debug("Last increment is not found",
			esl.Time("lastIncrement", lastIncrement))
		return true
	}
}

func (z *ccImpl) noLockRelease(key string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("key", key))

	if concurrency, ok := z.concurrency[key]; ok {
		l.Debug("Release lock")
		z.concurrency[key] = es_number.Max(concurrency-1, 0).Int()
	} else {
		l.Debug("Lock not found")
	}
	z.noLockNotifyWaiters(key)
}

func (z *ccImpl) noLockNotifyWaiters(key string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("key", key))

	for current := z.waiters.Front(); current != nil; current = current.Next() {
		waiter := current.Value.(*ccWaiter)
		if waiter.key == key {
			l.Debug("Waiter found. Notify")
			z.waiters.Remove(current)
			close(waiter.ready)
			return
		}
	}
}
