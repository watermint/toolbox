package eq_queue

import (
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
)

type Queue interface {
	// Enqueue data into the queue.
	Enqueue(p interface{})

	// Create sub queue with batchId
	Batch(batchId string) Queue
}

func newQueue(mould eq_mould.Mould) Queue {
	return &qImpl{
		mould: mould,
	}
}

type qImpl struct {
	mould   eq_mould.Mould
	batchId string
}

func (z qImpl) Enqueue(p interface{}) {
	z.mould.Pour(p)
}

func (z qImpl) Batch(batchId string) Queue {
	z.mould = z.mould.Batch(batchId)
	z.batchId = batchId
	return z
}
