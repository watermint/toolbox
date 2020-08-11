package eq_bundle

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"sync"
)

func NewSimple(logger esl.Logger, factory eq_pipe.Factory) Bundle {
	return &simpleImpl{
		logger:     logger,
		factory:    factory,
		pipes:      make(map[string]eq_pipe.Pipe),
		pipesMutex: &sync.Mutex{},
		wip:        factory.New(""),
	}
}

type simpleImpl struct {
	logger         esl.Logger
	factory        eq_pipe.Factory
	pipes          map[string]eq_pipe.Pipe
	pipesMutex     *sync.Mutex
	wip            eq_pipe.Pipe
	batchId        string
	currentBatchId string
}

func (z *simpleImpl) Complete(d Data) {
	l := z.logger
	l.Debug("Mark as completed")

	z.wip.Delete(d.ToBytes())
}

func (z *simpleImpl) Close() {
	l := z.logger
	l.Debug("Shutdown bundle")

	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	for _, pipe := range z.pipes {
		pipe.Close()
	}
	z.pipes = make(map[string]eq_pipe.Pipe)
}

func (z *simpleImpl) Enqueue(d Data) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l := z.logger.With(esl.String("batchId", d.BatchId))
	l.Debug("Enqueue bundle")

	pipe, ok := z.pipes[d.BatchId]
	if !ok {
		l.Debug("A pipe not found for the BatchId, create a new pipe")
		pipe = z.factory.New(d.BatchId)
		z.pipes[d.BatchId] = pipe
	}
	pipe.Enqueue(d.ToBytes())
	l.Debug("Enqueue bundle: Done")
}

func (z *simpleImpl) Fetch() (d Data, found bool) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l := z.logger.With(esl.String("currentBatchId", z.currentBatchId))
	l.Debug("Fetch from currentBatchId")

	for {
		if pipe, ok := z.pipes[z.currentBatchId]; ok {
			l.Debug("The pipe found with currentBatchId")
			if d0 := pipe.Dequeue(); d0 != nil {
				l.Debug("Data found, dequeue success")
				d, err := FromBytes(d0)
				if err != nil {
					l.Debug("Unable to unmarshal the message", esl.Error(err), esl.Binary("data", d0))
					return d, false
				}
				z.wip.Enqueue(d0)
				return d, true
			}

			l.Debug("Data not found, closing the pipe")
			pipe.Close()
			delete(z.pipes, z.currentBatchId)

			// fall thru
		}

		// finish when no more pipes found
		if len(z.pipes) < 1 {
			l.Debug("No more pipes")
			z.currentBatchId = ""
			return d, false
		}

		// find next
		l.Debug("Find next currentBatchId")
		for b := range z.pipes {
			l.Debug("Next pipe", esl.String("batchId", b))
			z.currentBatchId = b
			break
		}
	}
}

func (z *simpleImpl) Size() (sizes map[string]int, total int) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	sizes = make(map[string]int)
	for b, pipe := range z.pipes {
		s := pipe.Size()
		sizes[b] = s
		total += s
	}

	return
}
