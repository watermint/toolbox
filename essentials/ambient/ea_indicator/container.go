package ea_indicator

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/watermint/toolbox/essentials/ambient/ea_notification"
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

type Container interface {
	// Add new indicator
	Add(total int64, opts ...mpb.BarOption) Indicator

	// Mark all indicators as done, and wait for shutdown
	Done()
}

var (
	global = newNopContainer()
)

func SuppressIndicator() {
	global.Done()
	global = nopContainer{}
}

func StartIndicator() {
	global.Done()
	global = newContainer(esl.Default())
}

func Global() Container {
	return global
}

func newContainer(log esl.Logger) Container {
	return &containerImpl{
		log:        log,
		indicators: make(map[int]Indicator),
		progress:   mpb.New(),
	}
}

func newNopContainer() Container {
	return nopContainer{}
}

type containerImpl struct {
	log             esl.Logger
	indicators      map[int]Indicator
	indicatorsMutex sync.Mutex
	progress        *mpb.Progress
}

func (z *containerImpl) Add(total int64, opts ...mpb.BarOption) Indicator {
	z.indicatorsMutex.Lock()
	defer z.indicatorsMutex.Unlock()

	l := z.log.With(esl.Int64("total", total))
	l.Debug("Suppress notification")

	ea_notification.Global().Suppress()

	bar := z.progress.AddBar(total, opts...)
	indicator := newIndicator(total, bar)
	z.indicators[bar.ID()] = indicator
	l.Debug("Add new indicator", esl.Int("barId", bar.ID()))

	return indicator
}

func (z *containerImpl) Done() {
	z.indicatorsMutex.Lock()
	defer z.indicatorsMutex.Unlock()

	l := z.log
	l.Debug("Mark as done for all indicators")
	for _, indicator := range z.indicators {
		indicator.Done()
	}

	z.indicators = make(map[int]Indicator)

	l.Debug("Wait for progress shutdown")
	z.progress.Wait()
	z.progress = mpb.New()

	l.Debug("Progress shutdown completed, resume notification")
	ea_notification.Global().Resume()
}

type nopContainer struct {
}

func (n nopContainer) Add(total int64, opts ...mpb.BarOption) Indicator {
	return NewNopIndicator()
}

func (n nopContainer) Done() {
}
