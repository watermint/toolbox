package eq_worker

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_registry"
	"sync"
)

type Worker interface {
	Startup(numWorker int)
	Wait()
}

func New(l esl.Logger, reg eq_registry.Registry, c chan eq_bundle.Barrel) Worker {
	return &workerImpl{
		reg:      reg,
		c:        c,
		l:        l,
		shutdown: make(chan struct{}),
	}
}

type workerImpl struct {
	reg          eq_registry.Registry
	c            chan eq_bundle.Barrel
	l            esl.Logger
	wg           sync.WaitGroup
	shutdown     chan struct{}
	shutdownOnce sync.Once
}

func (z *workerImpl) logger() esl.Logger {
	return z.l.With(esl.String("routine", es_goroutine.GetGoRoutineName()))
}

func (z *workerImpl) Startup(numWorker int) {
	l := z.logger().With(esl.Int("numWorker", numWorker))
	l.Debug("Startup worker")
	for i := 0; i < numWorker; i++ {
		l.Debug("Starting worker", esl.Int("workerId", i))
		z.wg.Add(1)
		go z.loop()
	}
}

func (z *workerImpl) Wait() {
	l := z.logger()
	l.Debug("Wait for shutdown")
	z.wg.Wait()
	l.Debug("Wait done")
}

func (z *workerImpl) loop() {
	l := z.logger()

	for barrel := range z.c {
		ll := l.With(esl.String("mouldId", barrel.MouldId), esl.String("batchId", barrel.BatchId))
		mould, found := z.reg.Get(barrel.MouldId)
		if !found {
			ll.Warn("Mould not found for mouldId, skip processing")
			continue
		}
		ll.Debug("Process: Start")
		mould.Process(barrel)
		ll.Debug("Process: Done")
	}
	l.Debug("End loop")
	z.wg.Done()
}
