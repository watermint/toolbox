package eq_progress

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
	"github.com/watermint/toolbox/essentials/queue/eq_stat"
	"sync"
	"time"
)

const (
	// Should not exceed 15 fps
	minRefreshFps      = 15
	minRefreshInterval = time.Second / minRefreshFps
)

func NewBar(opts ...mpb.ContainerOption) Progress {
	return &barImpl{
		barLock:       sync.Mutex{},
		barBatch:      make(map[string]*mpb.Bar),
		barTask:       make(map[string]*mpb.Bar),
		barTotalTask:  make(map[string]int64),
		barTotalBatch: make(map[string]int64),
		barOpts:       opts,
		container:     mpb.New(opts...),
	}
}

type barImpl struct {
	barLock       sync.Mutex
	barTask       map[string]*mpb.Bar
	barBatch      map[string]*mpb.Bar
	barTotalTask  map[string]int64
	barTotalBatch map[string]int64
	barOpts       []mpb.ContainerOption
	container     *mpb.Progress
	lastRefresh   time.Time
}

func (z *barImpl) Flush() {
	z.barLock.Lock()
	defer z.barLock.Unlock()

	for k, b := range z.barTask {
		b.SetTotal(z.barTotalTask[k], true)
	}
	for k, b := range z.barBatch {
		b.SetTotal(z.barTotalBatch[k], true)
	}

	z.barBatch = make(map[string]*mpb.Bar)
	z.barTask = make(map[string]*mpb.Bar)
	z.barTotalTask = make(map[string]int64)
	z.barTotalBatch = make(map[string]int64)
	z.container.Wait()

	z.container = mpb.New(z.barOpts...)
}

func (z *barImpl) noLockNewBar(mouldId string, total int, typeName string) *mpb.Bar {
	mouldName := mouldId
	digestLen := 24
	if len(mouldName) > digestLen {
		mouldName = mouldName[len(mouldName)-digestLen:]
	}

	return z.container.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name(mouldName+" ", decor.WC{W: digestLen}),
			decor.Name(typeName+" ", decor.WC{W: 5}),
			decor.Elapsed(decor.ET_STYLE_MMSS),
		),
		mpb.AppendDecorators(
			decor.Counters(0, " %d / %d "),
			decor.OnComplete(
				decor.Spinner(mpb.DefaultSpinnerStyle), "DONE",
			),
		),
	)
}

func (z *barImpl) noLockGetBar(mouldId string, totalTask, totalBatch int) (barTask, barBatch *mpb.Bar, new bool) {
	if barTask, ok := z.barTask[mouldId]; ok {
		if barBatch, ok := z.barBatch[mouldId]; ok {
			return barTask, barBatch, false
		}
	}

	barBatch = z.noLockNewBar(mouldId, totalBatch, "Batch")
	barTask = z.noLockNewBar(mouldId, totalTask, "Task ")
	z.barBatch[mouldId] = barBatch
	z.barTask[mouldId] = barTask
	return barTask, barBatch, true
}

func (z *barImpl) onChange(mouldId, batchId string, stat eq_stat.Stat) {
	z.barLock.Lock()
	defer z.barLock.Unlock()

	batchCompleted, batchTotal := stat.StatBatch(mouldId)
	taskCompleted, taskTotal := stat.StatTask(mouldId)

	z.barTotalBatch[mouldId] = int64(batchTotal)
	z.barTotalTask[mouldId] = int64(taskTotal)

	barTask, barBatch, newBar := z.noLockGetBar(mouldId, taskTotal, batchTotal)

	if newBar ||
		batchCompleted == batchTotal ||
		taskCompleted == taskTotal ||
		z.lastRefresh.Add(minRefreshInterval).Before(time.Now()) {

		barTask.SetCurrent(int64(taskCompleted))
		barTask.SetTotal(int64(taskTotal), false)
		barBatch.SetCurrent(int64(batchCompleted))
		barBatch.SetTotal(int64(batchTotal), false)

		z.lastRefresh = time.Now()
	}
}

func (z *barImpl) OnComplete(mouldId, batchId string, stat eq_stat.Stat) {
	z.onChange(mouldId, batchId, stat)
}

func (z *barImpl) OnEnqueue(mouldId, batchId string, stat eq_stat.Stat) {
	z.onChange(mouldId, batchId, stat)
}
