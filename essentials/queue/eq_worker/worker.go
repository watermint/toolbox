package eq_worker

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"sync"
)

type Worker interface {
	Startup()
	Wait()
}

func New(l esl.Logger, m eq_mould.Mould, c chan eq_bundle.Data) Worker {
	return &workerImpl{
		c:        c,
		l:        l,
		m:        m,
		shutdown: make(chan struct{}),
	}
}

type workerImpl struct {
	c            chan eq_bundle.Data
	l            esl.Logger
	m            eq_mould.Mould
	wg           sync.WaitGroup
	shutdown     chan struct{}
	shutdownOnce sync.Once
}

func (z *workerImpl) logger() esl.Logger {
	return z.l.With(esl.String("routine", es_goroutine.GetGoRoutineName()))
}

func (z *workerImpl) Startup() {
	l := z.logger()
	l.Debug("Startup worker")
	z.wg.Add(1)
	go z.loop()
}

func (z *workerImpl) Wait() {
	l := z.logger()
	l.Debug("Wait for shutdown")
	z.wg.Wait()
	l.Debug("Wait done")
}

func (z *workerImpl) loop() {
	l := z.logger()

	for d := range z.c {
		ll := l.With(esl.String("batchId", d.BatchId))
		ll.Debug("Process: Start")
		z.m.Process(d)
		ll.Debug("Process: Done")
	}
	l.Debug("End loop")
	z.wg.Done()
}
