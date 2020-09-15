package nw_congestion

import (
	"container/list"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"runtime"
	"sync"
	"time"
)

const (
	hardLimitCongestionWindow = 8
)

var (
	monitorDuration          = 1 * time.Minute
	maxCongestionWindow      = runtime.NumCPU()
	initCongestionWindow     = 4
	minCongestionWindow      = 1
	currentImpl              = NewControl()
	thresholdSignificantWait = 59 * time.Second
)

func getReportInterval() time.Duration {
	if app.IsProduction() {
		return 1 * time.Minute
	} else {
		return 10 * time.Second
	}
}

func Start(hash, endpoint string) {
	currentImpl.Start(hash, endpoint)
}

func EndSuccess(hash, endpoint string) {
	currentImpl.EndSuccess(hash, endpoint)
}

func EndTransportError(hash, endpoint string) {
	currentImpl.EndTransportError(hash, endpoint)
}

func EndRateLimit(hash, endpoint string, reset time.Time) {
	currentImpl.EndRateLimit(hash, endpoint, reset)
}

func SetMaxCongestionWindow(w int, ignoreHardLimit bool) {
	if ignoreHardLimit {
		maxCongestionWindow = w
	} else {
		maxCongestionWindow = es_number.Min(w, hardLimitCongestionWindow).Int()
	}
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
	EndRateLimit(hash, endpoint string, reset time.Time)
}

func NewControl() CongestionControl {
	return &ccImpl{
		window:        make(map[string]int),
		lastRateLimit: make(map[string]time.Time),
		lastIncrement: make(map[string]time.Time),
		lastDecrement: make(map[string]time.Time),
		runners:       make(map[string]*ccRunner),
		monitorOnce:   &sync.Once{},
	}
}

type ccRunner struct {
	Key          string    `json:"key"`
	GoRoutineId  string    `json:"go_routine_id"`
	RunningSince time.Time `json:"running_since"`
}

// ref golang.org/x/sync/semaphore
type ccWaiter struct {
	Key         string          `json:"key"`
	Ready       chan<- struct{} `json:"-"`
	GoRoutineId string          `json:"go_routine_id"`
	WaitSince   time.Time       `json:"wait_since"`
}

type ccImpl struct {
	window        map[string]int
	lastRateLimit map[string]time.Time
	lastIncrement map[string]time.Time
	lastDecrement map[string]time.Time
	runners       map[string]*ccRunner
	waiters       list.List
	monitorOnce   *sync.Once
	mutex         sync.Mutex
}

func (z *ccImpl) key(hash, endpoint string) string {
	return hash + "-" + endpoint
}

func (z *ccImpl) Start(hash, endpoint string) {
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := esl.Default().With(esl.String("goroutine", goRoutineId),
		esl.String("hash", hash), esl.String("endpoint", endpoint))

	z.monitorOnce.Do(func() {
		go z.monitor()
	})

	key := z.key(hash, endpoint)
	z.mutex.Lock()

	if wnd, ok := z.window[key]; !ok {
		wnd = CurrentInitCongestionWindow()
		l.Debug("Congestion window not found, create new window", esl.Int("window", wnd))
		z.window[key] = wnd
		z.runners[goRoutineId] = &ccRunner{
			Key:          key,
			GoRoutineId:  goRoutineId,
			RunningSince: time.Now(),
		}
		z.mutex.Unlock()
		return
	} else {
		curConcurrency, curWaiters := z.noLockCalcConcurrency(key, false)
		concurrency := curConcurrency + curWaiters
		if concurrency < wnd {
			l.Debug("There is available window",
				esl.Int("window", wnd), esl.Int("concurrency", concurrency))
			z.runners[goRoutineId] = &ccRunner{
				Key:          key,
				GoRoutineId:  goRoutineId,
				RunningSince: time.Now(),
			}
			z.mutex.Unlock()
			return
		}

		l.Debug("There is no window available now. Wait for another process.", esl.Int("window", wnd), esl.Int("concurrency", concurrency))
		ready := make(chan struct{})
		waiter := &ccWaiter{
			Ready:       ready,
			Key:         key,
			WaitSince:   time.Now(),
			GoRoutineId: es_goroutine.GetGoRoutineName(),
		}
		z.waiters.PushBack(waiter)
		z.mutex.Unlock()

		select {
		case <-ready:
			z.mutex.Lock()
			//			newConcurrency := z.noLockCalcConcurrency(key, true) + 1
			goRoutineId = es_goroutine.GetGoRoutineName()
			z.runners[goRoutineId] = &ccRunner{
				Key:          key,
				GoRoutineId:  goRoutineId,
				RunningSince: time.Now(),
			}
			//z.concurrency[key] = newConcurrency
			l.Debug("Lock acquired.")
			z.mutex.Unlock()
			return
		}
	}
}

func (z *ccImpl) EndSuccess(hash, endpoint string) {
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := esl.Default().With(esl.String("goroutine", goRoutineId),
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
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := esl.Default().With(esl.String("goroutine", goRoutineId),
		esl.String("hash", hash), esl.String("endpoint", endpoint))

	l.Debug("Transport error; do not change window")
	key := z.key(hash, endpoint)
	z.mutex.Lock()
	z.noLockRelease(key)

	// do not increase window for while
	z.lastIncrement[key] = time.Now()
	z.mutex.Unlock()
}

func (z *ccImpl) isSignificantWait(reset time.Time) bool {
	return thresholdSignificantWait < reset.Sub(time.Now())
}

func (z *ccImpl) EndRateLimit(hash, endpoint string, reset time.Time) {
	isSignificant := z.isSignificantWait(reset)

	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("hash", hash),
		esl.String("endpoint", endpoint),
		esl.Bool("significant", isSignificant),
		esl.Time("reset", reset))

	key := z.key(hash, endpoint)
	z.mutex.Lock()
	z.noLockRelease(key)

	calcNewWindow := func(key string) int {
		if isSignificant {
			l.Debug("Set to minimum window when the signal was significant")
			return CurrentMinCongestionWindow()
		}

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

func (z *ccImpl) monitor() {
	ticker := time.NewTicker(getReportInterval())
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()))
	l.Debug("Monitor start")

	for {
		select {
		case <-ticker.C:
			z.mutex.Lock()
			waiters := make([]*ccWaiter, 0)
			for waiter := z.waiters.Front(); waiter != nil; waiter = waiter.Next() {
				waiters = append(waiters, waiter.Value.(*ccWaiter))
			}
			concurrencyMap := make(map[string]int)
			for _, runner := range z.runners {
				if c, ok := concurrencyMap[runner.Key]; ok {
					concurrencyMap[runner.Key] = c + 1
				} else {
					concurrencyMap[runner.Key] = 1
				}
			}
			l.Debug("WaiterStatus",
				esl.Any("runners", z.runners),
				esl.Int("numRunners", len(z.runners)),
				esl.Any("waiters", waiters),
				esl.Int("numWaiters", len(waiters)),
				esl.Any("window", z.window),
				esl.Any("concurrency", concurrencyMap),
			)
			z.mutex.Unlock()
		}
	}
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

func (z *ccImpl) noLockCalcConcurrency(key string, includeSelf bool) (concurrency, waiters int) {
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := esl.Default().With(esl.String("goroutine", goRoutineId),
		esl.String("key", key))

	concurrency = 0
	for _, r := range z.runners {
		if includeSelf && r.GoRoutineId == goRoutineId {
			continue
		}
		if r.Key == key {
			concurrency++
		}
	}
	waiters = 0
	for cur := z.waiters.Front(); cur != nil; cur = cur.Next() {
		w := cur.Value.(*ccWaiter)
		if includeSelf && w.GoRoutineId == goRoutineId {
			continue
		}
		if w.Key == key {
			waiters++
		}
	}
	l.Debug("Resolved concurrency",
		esl.String("key", key),
		esl.Int("concurrency", concurrency),
		esl.Int("waiters", waiters))
	return concurrency, waiters
}

func (z *ccImpl) noLockRelease(key string) {
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := esl.Default().With(esl.String("goroutine", goRoutineId),
		esl.String("key", key))

	delete(z.runners, goRoutineId)
	z.noLockNotifyWaiters(key)
	l.Debug("Lock released")
}

func (z *ccImpl) noLockNotifyWaiters(key string) {
	l := esl.Default().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()),
		esl.String("key", key))

	window := z.window[key]
	curConcurrency, curWaiters := z.noLockCalcConcurrency(key, false)
	l.Debug("NotifyToWaiters",
		esl.Int("concurrency", curConcurrency),
		esl.Int("waiters", curWaiters))
	numRelease := es_number.Max(0, window-curConcurrency).Int()
	numReleased := 0

	for current := z.waiters.Front(); current != nil; {
		if numRelease <= 0 {
			l.Debug("Waiters released", esl.Int("numReleased", numReleased))
			return
		}

		waiter := current.Value.(*ccWaiter)
		if waiter.Key == key {
			l.Debug("Waiter found. Notify", esl.Any("waiter", waiter))
			rel := current
			current = current.Next()
			z.waiters.Remove(rel)
			close(waiter.Ready)
			numRelease--
			numReleased++
		} else {
			current = current.Next()
		}
	}
}
