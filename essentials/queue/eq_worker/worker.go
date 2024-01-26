package eq_worker

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_registry"
	"sync"
	"time"
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
		status:   make(map[int]*workerStatus),
	}
}

type workerStatus struct {
	MouldId     string `json:"mould_id"`
	BatchId     string `json:"batch_id"`
	Time        string `json:"time"`
	GoRoutineId string `json:"go_routine_id"`
}

type workerImpl struct {
	reg           eq_registry.Registry
	c             chan eq_bundle.Barrel
	l             esl.Logger
	wg            sync.WaitGroup
	shutdown      chan struct{}
	shutdownOnce  sync.Once
	status        map[int]*workerStatus
	statusMutex   sync.Mutex
	statusControl chan struct{}
	running       bool
}

func (z *workerImpl) logger() esl.Logger {
	return z.l.With(esl.String("routine", es_goroutine.GetGoRoutineName()))
}

func (z *workerImpl) Startup(numWorker int) {
	l := z.logger().With(esl.Int("numWorker", numWorker))
	if numWorker < 1 {
		l.Warn("numWorker less than 1, fallback to 1")
		numWorker = 1
	}
	if z.running {
		l.Debug("The worker is already running")
		return
	}
	l.Debug("Startup worker")
	for i := 0; i < numWorker; i++ {
		l.Debug("Starting worker", esl.Int("workerId", i))
		z.wg.Add(1)
		go z.loop(i)
	}
	z.statusControl = make(chan struct{})
	go z.monitor()
	z.running = true
}

func (z *workerImpl) Wait() {
	l := z.logger()
	if !z.running {
		l.Debug("The worker is not running")
		return
	}

	l.Debug("Wait for shutdown")
	z.wg.Wait()
	close(z.statusControl)
	l.Debug("Wait done")
}

func (z *workerImpl) monitor() {
	l := z.logger().With(esl.String("goroutine", es_goroutine.GetGoRoutineName()))
	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-ticker.C:
			z.statusMutex.Lock()
			l.Debug("Worker status", esl.Any("workers", z.status))
			z.statusMutex.Unlock()

		case <-z.statusControl:
			ticker.Stop()
			//l.Debug("Monitor shutdown")
		}
	}
}

func (z *workerImpl) loop(workerId int) {
	goRoutineId := es_goroutine.GetGoRoutineName()
	l := z.logger().With(esl.Int("workerId", workerId), esl.String("goroutine", goRoutineId))

	for barrel := range z.c {
		ll := l.With(esl.String("mouldId", barrel.MouldId), esl.String("batchId", barrel.BatchId))
		mould, found := z.reg.Get(barrel.MouldId)
		if !found {
			ll.Warn("Mould not found for mouldId, skip processing")
			continue
		}
		ll.Debug("Process: Start")
		z.statusMutex.Lock()
		z.status[workerId] = &workerStatus{
			MouldId:     barrel.MouldId,
			BatchId:     barrel.BatchId,
			Time:        time.Unix(barrel.Time, 0).Format(time.RFC3339),
			GoRoutineId: goRoutineId,
		}
		z.statusMutex.Unlock()

		mould.Process(barrel)

		z.statusMutex.Lock()
		z.status[workerId] = nil
		z.statusMutex.Unlock()
		ll.Debug("Process: Done")
	}
	l.Debug("End loop")
	z.wg.Done()
}
