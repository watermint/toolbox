package file

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"math"
	"sync/atomic"
)

type UploadStatus struct {
	summary UploadSummary
}

func (z *UploadStatus) error() {
	atomic.AddInt64(&z.summary.NumFilesError, 1)
}

func (z *UploadStatus) skip() {
	atomic.AddInt64(&z.summary.NumFilesSkip, 1)
}

func (z *UploadStatus) upload(size int64, chunkSize int) {
	atomic.AddInt64(&z.summary.NumBytes, size)
	atomic.AddInt64(&z.summary.NumFilesUpload, 1)

	apiCalls := es_number.Max(math.Ceil(float64(size)/float64(chunkSize)), 0).Int64()
	// Zero size file also consume API
	if size == 0 || apiCalls < 1 {
		apiCalls = 1
	}
	atomic.AddInt64(&z.summary.NumApiCall, apiCalls)
}
