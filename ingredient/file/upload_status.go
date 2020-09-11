package file

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"math"
	"sync/atomic"
	"time"
)

type Status struct {
	summary Summary
}

func (z *Status) start() {
	z.summary.Start = time.Now()
}

func (z *Status) finish() {
	z.summary.End = time.Now()
}

func (z *Status) error() {
	atomic.AddInt64(&z.summary.NumFilesError, 1)
}

func (z *Status) skip() {
	atomic.AddInt64(&z.summary.NumFilesSkip, 1)
}

func (z *Status) upload(size int64, chunkSize int) {
	atomic.AddInt64(&z.summary.NumBytes, size)
	atomic.AddInt64(&z.summary.NumFilesTransferred, 1)

	apiCalls := es_number.Max(math.Ceil(float64(size)/float64(chunkSize)), 0).Int64()
	// Zero size file also consume API
	if size == 0 || apiCalls < 1 {
		apiCalls = 1
	}
	atomic.AddInt64(&z.summary.NumApiCall, apiCalls)
}

func (z *Status) download(size int64) {
	atomic.AddInt64(&z.summary.NumBytes, size)
	atomic.AddInt64(&z.summary.NumFilesTransferred, 1)
}

func (z *Status) createFolder() {
	atomic.AddInt64(&z.summary.NumFolderCreated, 1)
}

func (z *Status) delete() {
	atomic.AddInt64(&z.summary.NumDeleted, 1)
}
