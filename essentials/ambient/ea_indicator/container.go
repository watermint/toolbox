package ea_indicator

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
	"github.com/watermint/toolbox/essentials/ambient/ea_notification"
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

type Container interface {
	// Add new indicator
	NewIndicator(total int64, opts ...mpb.BarOption) Indicator

	// Add new status
	NewStatus(name string, opts ...mpb.BarOption) StatusBar

	// Mark all indicators as done, and wait for shutdown
	Done()
}

var (
	global         = newSwitcherContainer(newContainer(esl.Default()))
	globalSuppress = false
)

func SuppressIndicatorForce() {
	globalSuppress = true
	SuppressIndicator()
}

func SuppressIndicator() {
	global.Done()
	global = newNopContainer()
}

func StartIndicator() {
	if !globalSuppress {
		global.Done()
		global = newSwitcherContainer(newContainer(esl.Default()))
	} else {
		global = newNopContainer()
	}
}

func Global() Container {
	return global
}

func newContainer(log esl.Logger) Container {
	return &containerImpl{
		log:        log,
		indicators: make(map[int]Indicator),
		statusBars: make(map[int]StatusBar),
		progress:   mpb.New(),
	}
}

func newNopContainer() Container {
	return nopContainer{}
}

type containerImpl struct {
	log             esl.Logger
	indicators      map[int]Indicator
	statusBars      map[int]StatusBar
	indicatorsMutex sync.Mutex
	progress        *mpb.Progress
}

func (z *containerImpl) NewStatus(name string, opts ...mpb.BarOption) StatusBar {
	z.indicatorsMutex.Lock()
	defer z.indicatorsMutex.Unlock()

	l := z.log.With(esl.String("name", name))
	l.Debug("Suppress notification")
	ea_notification.Global().Suppress()

	st := newStatus()

	decoStatus := func(statistics decor.Statistics) string {
		return st.CurrentStatus()
	}
	decoTitle := func(statistics decor.Statistics) string {
		return st.CurrentTitle()
	}

	barOpts := make([]mpb.BarOption, 0)
	barOpts = append(barOpts, opts...)
	barOpts = append(barOpts, mpb.BarWidth(1))
	barOpts = append(barOpts, mpb.PrependDecorators(
		decor.Any(decoTitle),
		decor.Name(" "),
		decor.Any(decoStatus)))

	bar := z.progress.AddBar(2, barOpts...)
	sb := newStatusBar(name, st, bar)
	z.statusBars[bar.ID()] = sb
	l.Debug("Add new status", esl.Int("barId", bar.ID()))

	return sb
}

func (z *containerImpl) NewIndicator(total int64, opts ...mpb.BarOption) Indicator {
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
	for _, status := range z.statusBars {
		status.Done()
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

func (n nopContainer) NewStatus(name string, opts ...mpb.BarOption) StatusBar {
	return NewNopStatus()
}

func (n nopContainer) NewIndicator(total int64, opts ...mpb.BarOption) Indicator {
	return NewNopIndicator()
}

func (n nopContainer) Done() {
}

func newSwitcherContainer(parent Container) Container {
	return &switcherContainer{
		container: parent,
	}
}

type switcherContainer struct {
	container Container
}

func (z switcherContainer) NewStatus(name string, opts ...mpb.BarOption) StatusBar {
	if globalSuppress {
		return NewNopStatus()
	}
	return z.container.NewStatus(name, opts...)
}

func (z switcherContainer) NewIndicator(total int64, opts ...mpb.BarOption) Indicator {
	if globalSuppress {
		return NewNopIndicator()
	}
	return z.container.NewIndicator(total, opts...)
}

func (z switcherContainer) Done() {
	z.container.Done()
}
