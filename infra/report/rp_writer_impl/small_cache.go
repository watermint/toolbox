package rp_writer_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"sync"
)

const (
	smallCacheThreshold = 1000
)

func NewSmallCache(name string, writer rp_writer.Writer) rp_writer.Writer {
	return NewSmallCacheWithThreshold(name, writer, smallCacheThreshold)
}

func NewSmallCacheWithThreshold(name string, writer rp_writer.Writer, threshold int64) rp_writer.Writer {
	return &smallCache{
		name:           name,
		cache:          make([]interface{}, 0),
		cacheThreshold: threshold,
		numRows:        0,
		writer:         writer,
	}
}

// Cache first X rows.
// Pass through to the child writer once row exceeds threshold (no cache).
type smallCache struct {
	name           string
	cache          []interface{}
	cacheMutex     sync.Mutex
	cacheThreshold int64
	numRows        int64
	writer         rp_writer.Writer
}

func (z *smallCache) noLockFlush() {
	for _, row := range z.cache {
		z.writer.Row(row)
	}
	z.cache = make([]interface{}, 0)
}

func (z *smallCache) Name() string {
	return z.name
}

func (z *smallCache) Row(r interface{}) {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	if z.numRows < z.cacheThreshold {
		z.cache = append(z.cache, r)
	} else if z.numRows == z.cacheThreshold {
		z.noLockFlush()
		z.writer.Row(r)
	} else {
		z.writer.Row(r)
	}
	z.numRows++
}

func (z *smallCache) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	return z.writer.Open(ctl, model, opts...)
}

func (z *smallCache) Close() {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	z.noLockFlush()
	z.writer.Close()
}
