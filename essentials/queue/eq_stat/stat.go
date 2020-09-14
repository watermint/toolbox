package eq_stat

import "sync"

type Stat interface {
	// Increment for enqueue event
	IncrEnqueue(mouldId, batchId string)

	// Increment for fetch event
	IncrFetch(mouldId, batchId string)

	// Increment for complete event
	IncrComplete(mouldId, batchId string)

	// Statistics of tasks.
	// Returns 0, 0 if the task not found for the mouldId
	StatTask(mouldId string) (completed, total int)

	// Statistics of batch.
	// Returns 0, 0 if the task not found for the mouldId
	StatBatch(mouldId string) (completed, total int)
}

func New() Stat {
	return &statImpl{
		taskTotal:      make(map[string]int),
		taskCompleted:  make(map[string]int),
		batchTotal:     make(map[string]map[string]int),
		batchCompleted: make(map[string]map[string]int),
	}
}

type statImpl struct {
	// mouldId -> total tasks per mould
	taskTotal map[string]int
	// mouldId -> completed tasks per mould
	taskCompleted map[string]int

	// mouldId -> batchId -> total tasks per batch
	batchTotal map[string]map[string]int
	// mouldId -> batchId -> completed tasks per batch
	batchCompleted map[string]map[string]int

	lockTotal     sync.Mutex
	lockCompleted sync.Mutex
}

func (z *statImpl) IncrEnqueue(mouldId, batchId string) {
	z.lockTotal.Lock()
	defer z.lockTotal.Unlock()

	if v, ok := z.taskTotal[mouldId]; ok {
		z.taskTotal[mouldId] = v + 1
	} else {
		z.taskTotal[mouldId] = 1
	}

	if batch, ok := z.batchTotal[mouldId]; ok {
		if v, ok := batch[batchId]; ok {
			z.batchTotal[mouldId][batchId] = v + 1
		} else {
			z.batchTotal[mouldId][batchId] = 1
		}
	} else {
		z.batchTotal[mouldId] = make(map[string]int)
		z.batchTotal[mouldId][batchId] = 1
	}
}

func (z *statImpl) IncrFetch(mouldId, batchId string) {
	// nop
}

func (z *statImpl) IncrComplete(mouldId, batchId string) {
	z.lockCompleted.Lock()
	defer z.lockCompleted.Unlock()

	if v, ok := z.taskCompleted[mouldId]; ok {
		z.taskCompleted[mouldId] = v + 1
	} else {
		z.taskCompleted[mouldId] = 1
	}

	if batch, ok := z.batchCompleted[mouldId]; ok {
		if v, ok := batch[batchId]; ok {
			z.batchCompleted[mouldId][batchId] = v + 1
		} else {
			z.batchCompleted[mouldId][batchId] = 1
		}
	} else {
		z.batchCompleted[mouldId] = make(map[string]int)
		z.batchCompleted[mouldId][batchId] = 1
	}
}

func (z *statImpl) StatTask(mouldId string) (completed, total int) {
	z.lockTotal.Lock()
	z.lockCompleted.Lock()
	defer z.lockTotal.Unlock()
	defer z.lockCompleted.Unlock()

	completed = z.taskCompleted[mouldId]
	total = z.taskTotal[mouldId]
	return
}

func (z *statImpl) StatBatch(mouldId string) (completed, total int) {
	z.lockTotal.Lock()
	z.lockCompleted.Lock()
	defer z.lockTotal.Unlock()
	defer z.lockCompleted.Unlock()

	batchTotal, ok := z.batchTotal[mouldId]
	if !ok {
		return
	}
	total = len(batchTotal)
	batchCompleted, ok := z.batchCompleted[mouldId]
	if !ok {
		return 0, total
	}

	for batchId, batchCompleted := range batchCompleted {
		if v, ok := batchTotal[batchId]; ok {
			if v == batchCompleted {
				completed++
			}
		}
	}
	return
}
