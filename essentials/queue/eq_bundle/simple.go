package eq_bundle

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"sync"
)

func NewSimple(logger esl.Logger,
	progress eq_progress.Progress,
	factory eq_pipe.Factory) Bundle {
	return newSimple(logger, progress, factory, factory.New(""), make(map[string]eq_pipe.Pipe))
}

func newSimple(logger esl.Logger,
	progress eq_progress.Progress,
	factory eq_pipe.Factory,
	wip eq_pipe.Pipe,
	pipes map[string]eq_pipe.Pipe) Bundle {
	return &simpleImpl{
		logger:         logger,
		progress:       progress,
		factory:        factory,
		pipes:          pipes,
		pipesMutex:     &sync.Mutex{},
		statsTotal:     make(map[string]int),
		statsCompleted: make(map[string]int),
		wip:            wip,
		batchId:        "",
		currentBatchId: "",
	}
}

func RestoreSimple(logger esl.Logger,
	progress eq_progress.Progress,
	factory eq_pipe.Factory,
	session Session) (b Bundle, err error) {

	l := logger.With(esl.Any("session", session))

	l.Debug("Restore InProgress Pipe", esl.String("sessionId", string(session.InProgress)))
	wip, err := factory.Restore(session.InProgress)
	if err != nil {
		l.Debug("Unable to restore InProgress Pipe", esl.Error(err))
		return nil, err
	}

	pipes := make(map[string]eq_pipe.Pipe)
	for batchId, sessionId := range session.Pipes {
		ll := l.With(esl.String("batchId", batchId), esl.String("sessionId", string(sessionId)))
		ll.Debug("Restore Pipe")

		pipe, err := factory.Restore(sessionId)
		if err != nil {
			ll.Debug("Unable to restore Pipe", esl.Error(err))
			return nil, err
		}

		pipes[batchId] = pipe
	}

	// Enqueue In progress data into the pipe
	b = newSimple(logger, progress, factory, wip, pipes)

	l.Debug("Dequeue from In Progress pipe")
	for p := wip.Dequeue(); p != nil; p = wip.Dequeue() {
		l.Debug("Dequeue data", esl.Binary("p", p))
		d, err := FromBytes(p)
		if err != nil {
			l.Debug("Unable to unmarshal bytes", esl.Error(err))
			return nil, err
		}

		l.Debug("Enqueue data", esl.String("batchId", d.BatchId))
		b.Enqueue(d)
	}

	return b, nil
}

type simpleImpl struct {
	logger         esl.Logger
	progress       eq_progress.Progress
	factory        eq_pipe.Factory
	pipes          map[string]eq_pipe.Pipe
	pipesMutex     *sync.Mutex
	statsTotal     map[string]int
	statsCompleted map[string]int
	wip            eq_pipe.Pipe
	batchId        string
	currentBatchId string
}

func (z *simpleImpl) SizeInProgress() int {
	return z.wip.Size()
}

func (z *simpleImpl) Preserve() (session Session, err error) {
	session.Pipes = make(map[string]eq_pipe.SessionId)
	l := z.logger
	l.Debug("Preserve")

	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	for batchId, pipe := range z.pipes {
		ll := l.With(esl.String("batchId", batchId))
		ll.Debug("Preserve")
		sessionId, err := pipe.Preserve()
		if err != nil {
			l.Debug("Unable to preserve pipe", esl.Error(err))
			return session, err
		}
		session.Pipes[batchId] = sessionId
	}

	l.Debug("Preserve In progress pipe")
	session.InProgress, err = z.wip.Preserve()
	if err != nil {
		l.Debug("Unable to preserve pipe", esl.Error(err))
		return session, err
	}

	l.Debug("Bundle preserved", esl.Any("session", session))
	return session, nil
}

func (z *simpleImpl) noLockCallHandler(batchBarrel, mouldId, batchId string, handler func(mouldId, batchId string, completed, total int)) {
	l := z.logger
	l.Debug("Mark as completed")

	if handler != nil {
		if total, ok := z.statsTotal[batchBarrel]; ok {
			if completed, ok := z.statsCompleted[batchBarrel]; ok {
				l.Debug("Call handler", esl.Int("completed", completed), esl.Int("total", total))
				handler(mouldId, batchId, completed, total)
			}
		}
	}
}

func (z *simpleImpl) Complete(b Barrel) {
	bb := b.BarrelBatch()

	l := z.logger.With(esl.String("barrelBatch", bb))
	l.Debug("Mark as completed")

	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	if completed, ok := z.statsCompleted[bb]; ok {
		z.statsCompleted[bb] = completed + 1
	} else {
		z.statsCompleted[bb] = 1
	}

	// Call OnCompletedHandler
	if z.progress != nil {
		z.noLockCallHandler(bb, b.MouldId, b.BatchId, z.progress.OnComplete)
	}

	z.wip.Delete(b.ToBytes())
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

func (z *simpleImpl) Enqueue(b Barrel) {
	bb := b.BarrelBatch()

	l := z.logger.With(esl.String("barrelBatch", bb))
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l.Debug("Enqueue bundle")

	if total, ok := z.statsTotal[bb]; ok {
		z.statsTotal[bb] = total + 1
	} else {
		z.statsTotal[bb] = 1
	}

	// Call OnCompletedHandler
	if z.progress != nil {
		z.noLockCallHandler(bb, b.MouldId, b.BatchId, z.progress.OnEnqueue)
	}

	pipe, ok := z.pipes[bb]
	if !ok {
		l.Debug("A pipe not found for the BatchId, create a new pipe")
		pipe = z.factory.New(bb)
		z.pipes[bb] = pipe
	}
	pipe.Enqueue(b.ToBytes())
	l.Debug("Enqueue bundle: Done")
}

func (z *simpleImpl) Fetch() (b Barrel, found bool) {
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
			return b, false
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
