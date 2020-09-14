package ea_indicator

import (
	"github.com/vbauerster/mpb/v5"
	"sync/atomic"
)

type Indicator interface {
	AddTotal(total int64)
	AddProgress(progress int64)
	UpdateProgress(progress int64)
	UpdateTotal(total int64)
	Done()
}

func newIndicator(total int64, bar *mpb.Bar) Indicator {
	return &indicatorImpl{
		progress: 0,
		total:    total,
		bar:      bar,
	}
}

type indicatorImpl struct {
	progress int64
	total    int64
	bar      *mpb.Bar
}

func (z *indicatorImpl) AddTotal(total int64) {
	atomic.AddInt64(&z.total, total)
	z.UpdateTotal(z.total)
}

func (z *indicatorImpl) AddProgress(progress int64) {
	atomic.AddInt64(&z.progress, progress)
	z.UpdateProgress(z.progress)
}

func (z *indicatorImpl) UpdateProgress(progress int64) {
	z.progress = progress
	z.bar.SetCurrent(progress)
}

func (z *indicatorImpl) UpdateTotal(total int64) {
	z.total = total
	z.bar.SetTotal(total, false)
}

func (z *indicatorImpl) Done() {
	z.bar.SetTotal(z.total, true)
}

func NewNopIndicator() Indicator {
	return &nopIndicatorImpl{}
}

type nopIndicatorImpl struct {
}

func (z nopIndicatorImpl) AddTotal(total int64) {
}

func (z nopIndicatorImpl) AddProgress(progress int64) {
}

func (z nopIndicatorImpl) UpdateProgress(progress int64) {
}

func (z nopIndicatorImpl) UpdateTotal(total int64) {
}

func (z nopIndicatorImpl) Done() {
}
