package nw_concurrency

import (
	"context"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"golang.org/x/sync/semaphore"
	"runtime"
	"sync"
)

var (
	masterConcurrency = newConcurrency()
)

func SetConcurrency(c int) {
	masterConcurrency.SetConcurrency(c)
}
func Start() {
	masterConcurrency.Start()
}
func End() {
	masterConcurrency.End()
}

type Concurrency interface {
	SetConcurrency(c int)
	Start()
	End()
}

func newConcurrency() Concurrency {
	return &concurrencyImpl{
		w:     semaphore.NewWeighted(int64(runtime.NumCPU())),
		mutex: sync.Mutex{},
	}
}

type concurrencyImpl struct {
	w     *semaphore.Weighted
	mutex sync.Mutex
}

func (z *concurrencyImpl) SetConcurrency(c int) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	l := es_log.Default()
	if c < 1 {
		l.Debug("Ignore setting concurrency for less than 1", es_log.Int("concurrency", c))
	} else {
		l.Debug("Set concurrency", es_log.Int("concurrency", c))
		z.w = semaphore.NewWeighted(int64(c))
	}
}

func (z *concurrencyImpl) Start() {
	err := z.w.Acquire(context.Background(), 1)
	if err != nil {
		l := es_log.Default()
		l.Debug("Unable to acquire semaphore", es_log.Error(err))
	}
}

func (z *concurrencyImpl) End() {
	z.w.Release(1)
}
