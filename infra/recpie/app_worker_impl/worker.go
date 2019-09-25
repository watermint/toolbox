package app_worker_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_worker"
	"go.uber.org/zap"
	"sync"
)

func NewQueue(ctl app_control.Control) app_worker.Queue {
	return &Queue{
		ctl: ctl,
		wg:  sync.WaitGroup{},
		q:   make(chan app_worker.Worker),
	}
}

type Queue struct {
	ctl app_control.Control
	wg  sync.WaitGroup
	q   chan app_worker.Worker
}

func (z *Queue) dequeue(id int) {
	l := z.ctl.Log().With(zap.Int("Worker", id))
	jobId := 0

	for w := range z.q {
		ll := l.With(zap.Int("Job", jobId))
		jobId++

		ll.Debug("Run work")
		if err := w(z.ctl); err != nil {
			ll.Debug("Work finished with error", zap.Error(err))
		} else {
			ll.Debug("Work SUCCESS")
		}
		z.wg.Done()
	}
	l.Debug("Shutdown")
}

func (z *Queue) Launch(concurrency int) {
	l := z.ctl.Log()
	if concurrency < 1 {
		l.Error("Concurrency must grater than 1", zap.Int("concurrency", concurrency))
		z.ctl.Abort(app_control.Reason(app_control.FatalPanic))
	}

	for i := 0; i < concurrency; i++ {
		go z.dequeue(i)
	}
}

func (z *Queue) Enqueue(w app_worker.Worker) {
	z.wg.Add(1)
	z.q <- w
}

func (z *Queue) Wait() {
	z.wg.Wait()
}
