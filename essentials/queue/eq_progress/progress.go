package eq_progress

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/queue/eq_stat"
	"sync"
)

type Progress interface {
	OnEnqueue(mouldId, batchId string, stat eq_stat.Stat)
	OnComplete(mouldId, batchId string, stat eq_stat.Stat)
	Flush()
}

func NewProgress(container ea_indicator.Container) Progress {
	return &barImpl{
		barLock:   sync.Mutex{},
		barBatch:  make(map[string]ea_indicator.Indicator),
		barTask:   make(map[string]ea_indicator.Indicator),
		container: container,
	}
}

type barImpl struct {
	barLock   sync.Mutex
	barTask   map[string]ea_indicator.Indicator
	barBatch  map[string]ea_indicator.Indicator
	container ea_indicator.Container
}

func (z *barImpl) Flush() {
	z.barLock.Lock()
	defer z.barLock.Unlock()

	for _, b := range z.barTask {
		b.Done()
	}
	for _, b := range z.barBatch {
		b.Done()
	}

	z.barBatch = make(map[string]ea_indicator.Indicator)
	z.barTask = make(map[string]ea_indicator.Indicator)
}

func (z *barImpl) noLockNewBar(mouldId string, total int, typeName string) ea_indicator.Indicator {
	mouldName := mouldId
	digestLen := 20
	if len(mouldName) > digestLen {
		mouldName = mouldName[len(mouldName)-digestLen:]
	}

	return z.container.Add(int64(total),
		mpb.PrependDecorators(
			decor.Name(mouldName+" ", decor.WC{W: digestLen}),
			decor.Name(typeName+" ", decor.WC{W: 5}),
			decor.Elapsed(decor.ET_STYLE_MMSS),
		),
		mpb.AppendDecorators(
			decor.Counters(0, "%5d/%5d", decor.WC{W: 12}),
			decor.OnComplete(
				decor.Spinner(mpb.DefaultSpinnerStyle, decor.WC{W: 5}), "DONE",
			),
		),
	)
}

func (z *barImpl) noLockGetBar(mouldId string, totalTask, totalBatch int) (barTask, barBatch ea_indicator.Indicator) {
	if barTask, ok := z.barTask[mouldId]; ok {
		if barBatch, ok := z.barBatch[mouldId]; ok {
			return barTask, barBatch
		}
	}

	barBatch = z.noLockNewBar(mouldId, totalBatch, "Batch")
	barTask = z.noLockNewBar(mouldId, totalTask, "Task ")
	z.barBatch[mouldId] = barBatch
	z.barTask[mouldId] = barTask
	return barTask, barBatch
}

func (z *barImpl) onChange(mouldId, batchId string, stat eq_stat.Stat) {
	z.barLock.Lock()
	defer z.barLock.Unlock()

	batchCompleted, batchTotal := stat.StatBatch(mouldId)
	taskCompleted, taskTotal := stat.StatTask(mouldId)

	barTask, barBatch := z.noLockGetBar(mouldId, taskTotal, batchTotal)

	barTask.UpdateProgress(int64(taskCompleted))
	barTask.UpdateTotal(int64(taskTotal))
	barBatch.UpdateProgress(int64(batchCompleted))
	barBatch.UpdateTotal(int64(batchTotal))
}

func (z *barImpl) OnComplete(mouldId, batchId string, stat eq_stat.Stat) {
	z.onChange(mouldId, batchId, stat)
}

func (z *barImpl) OnEnqueue(mouldId, batchId string, stat eq_stat.Stat) {
	z.onChange(mouldId, batchId, stat)
}
