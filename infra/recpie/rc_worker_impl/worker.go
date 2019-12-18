package rc_worker_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_worker"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"sync"
)

func NewQueue(ctl app_control.Control, concurrency int) rc_worker.Queue {
	q := &Queue{
		ctl: ctl,
		wg:  sync.WaitGroup{},
		q:   make(chan rc_worker.Worker),
	}
	q.Launch(concurrency)
	return q
}

type Queue struct {
	ctl         app_control.Control
	wg          sync.WaitGroup
	q           chan rc_worker.Worker
	numWorker   int
	maxWorker   int
	workerMutex sync.Mutex
}

func (z *Queue) dequeue() {
	l := z.ctl.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))
	jobId := 0

	for w := range z.q {
		ll := l.With(zap.Int("Seq", jobId))
		jobId++

		ll.Debug("Run work")
		if err := w.Exec(); err != nil {
			ll.Debug("Failure", zap.Error(err))
		} else {
			ll.Debug("Success")
		}
		z.wg.Done()
	}
	l.Debug("Shutdown")
}

func (z *Queue) Launch(concurrency int) {
	l := z.ctl.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))
	if concurrency < 1 {
		l.Debug("RunConcurrently must positive number, use 1 as default", zap.Int("concurrency", concurrency))
		concurrency = 1
	}

	l.Debug("Launch first worker", zap.Int("concurrency", concurrency))
	z.maxWorker = concurrency
	z.numWorker = 1
	go z.dequeue()
}

func (z *Queue) Enqueue(w rc_worker.Worker) {
	l := z.ctl.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))

	z.workerMutex.Lock()
	defer z.workerMutex.Unlock()

	if z.numWorker < z.maxWorker {
		l.Debug("Launch additional worker", zap.Int("numWorker", z.numWorker), zap.Int("maxWorker", z.maxWorker))
		go z.dequeue()
		z.numWorker++
	}

	z.wg.Add(1)
	z.q <- w
}

func (z *Queue) Wait() {
	l := z.ctl.Log().With(zap.String("Routine", ut_runtime.GetGoRoutineName()))
	l.Debug("Waiting for worker shutdown")
	z.wg.Wait()
	close(z.q)
}
