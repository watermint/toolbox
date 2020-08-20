package eq_progress

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
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
		barLock:   sync.Mutex{},
		bars:      make(map[string]*mpb.Bar),
		container: mpb.New(opts...),
	}
}

type barImpl struct {
	barLock     sync.Mutex
	bars        map[string]*mpb.Bar
	container   *mpb.Progress
	lastRefresh time.Time
}

func (z *barImpl) noLockNewBar(mouldId, batchId string, total int) *mpb.Bar {
	mouldName := mouldId
	digestLen := 16
	if len(mouldName) > digestLen {
		mouldName = mouldName[len(mouldName)-digestLen:]
	}
	batchName := batchId
	if len(mouldName) > digestLen {
		batchName = batchName[len(batchName)-digestLen:]
	}

	return z.container.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name(mouldName+" ", decor.WC{W: digestLen}),
			decor.Name(batchName+" ", decor.WC{W: digestLen}),
			decor.OnComplete(decor.AverageETA(decor.ET_STYLE_HHMM, decor.WC{W: 5}), "Done"),
		),
		mpb.AppendDecorators(decor.Counters(0, "%d / %d")),
	)
}

func (z *barImpl) noLockGetBar(mouldId, batchId string, total int) (bar *mpb.Bar, new bool) {
	batchBarrel := mouldId + "/" + batchId
	if bar, ok := z.bars[batchBarrel]; ok {
		return bar, false
	}

	bar = z.noLockNewBar(mouldId, batchId, total)
	z.bars[batchBarrel] = bar
	return bar, true
}

func (z *barImpl) onChange(mouldId, batchId string, completed, total int) {
	z.barLock.Lock()
	defer z.barLock.Unlock()

	bar, newBar := z.noLockGetBar(mouldId, batchId, total)

	if newBar || completed == total || z.lastRefresh.Add(minRefreshInterval).Before(time.Now()) {
		bar.SetCurrent(int64(completed))
		bar.SetTotal(int64(total), false)

		z.lastRefresh = time.Now()
	}
}

func (z *barImpl) OnComplete(mouldId, batchId string, completed, total int) {
	z.onChange(mouldId, batchId, completed, total)
}

func (z *barImpl) OnEnqueue(mouldId, batchId string, completed, total int) {
	z.onChange(mouldId, batchId, completed, total)
}
