package app_worker_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_worker"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"sync"
)

func NewQueue(ctl app_control.Control, concurrency int) app_worker.Queue {
	q := &Queue{
		ctl: ctl,
		wg:  sync.WaitGroup{},
		q:   make(chan app_worker.Worker),
	}
	q.Launch(concurrency)
	return q
}

type Queue struct {
	ctl app_control.Control
	wg  sync.WaitGroup
	q   chan app_worker.Worker
}

func (z *Queue) dequeue() {
	l := z.ctl.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))
	jobId := 0

	for w := range z.q {
		ll := l.With(zap.Int("Job", jobId))
		jobId++

		ll.Debug("Run work")
		if err := w.Exec(); err != nil {
			ll.Debug("FAILURE: Work finished with error", zap.Error(err))
		} else {
			ll.Debug("SUCCESS: Done")
		}
		z.wg.Done()
	}
	l.Debug("Shutdown")
}

func (z *Queue) Launch(concurrency int) {
	l := z.ctl.Log()
	if concurrency < 1 {
		l.Debug("RunConcurrently must positive number, use 1 as default", zap.Int("concurrency", concurrency))
		concurrency = 1
	}

	l.Debug("Launch workers", zap.Int("concurrency", concurrency))
	for i := 0; i < concurrency; i++ {
		go z.dequeue()
	}
}

func (z *Queue) Enqueue(w app_worker.Worker) {
	z.wg.Add(1)
	z.q <- w
}

func (z *Queue) Wait() {
	z.wg.Wait()
	close(z.q)
}
