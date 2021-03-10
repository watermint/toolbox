package eq_bundle

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"github.com/watermint/toolbox/essentials/queue/eq_stat"
	"math/rand"
	"sync"
)

const (
	FetchSequential FetchPolicy = "sequential"
	FetchRandom     FetchPolicy = "random"
	FetchBalance    FetchPolicy = "balance"
)

var (
	FetchPolicies = []FetchPolicy{FetchSequential, FetchRandom, FetchBalance}
)

type FetchPolicy string

func NewSimple(logger esl.Logger,
	policy FetchPolicy,
	progress eq_progress.Progress,
	factory eq_pipe.Factory) Bundle {
	return newSimple(logger, policy, progress, factory, factory.New(""), make(map[string]eq_pipe.Pipe))
}

func newSimple(logger esl.Logger,
	policy FetchPolicy,
	progress eq_progress.Progress,
	factory eq_pipe.Factory,
	wip eq_pipe.Pipe,
	pipes map[string]eq_pipe.Pipe) Bundle {
	return &simpleImpl{
		logger:                       logger,
		policy:                       policy,
		progress:                     progress,
		factory:                      factory,
		pipes:                        pipes,
		pipesMutex:                   &sync.Mutex{},
		stat:                         eq_stat.New(),
		wip:                          wip,
		batchId:                      "",
		sequentialCurrentBatchBarrel: "",
	}
}

func RestoreSimple(logger esl.Logger,
	policy FetchPolicy,
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
	b = newSimple(logger, policy, progress, factory, wip, pipes)

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
	logger                       esl.Logger
	policy                       FetchPolicy
	progress                     eq_progress.Progress
	factory                      eq_pipe.Factory
	pipes                        map[string]eq_pipe.Pipe
	pipesMutex                   *sync.Mutex
	stat                         eq_stat.Stat
	wip                          eq_pipe.Pipe
	batchId                      string
	sequentialCurrentBatchBarrel string
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

func (z *simpleImpl) Complete(b Barrel) {
	bb := b.BarrelBatch()

	l := z.logger.With(esl.String("barrelBatch", bb))
	l.Debug("Mark as completed")

	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	z.stat.IncrComplete(b.MouldId, b.BatchId)

	// Call OnCompletedHandler
	if z.progress != nil {
		z.progress.OnComplete(b.MouldId, b.BatchId, z.stat)
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

	z.stat.IncrEnqueue(b.MouldId, b.BatchId)

	// Call OnCompletedHandler
	if z.progress != nil {
		z.progress.OnEnqueue(b.MouldId, b.BatchId, z.stat)
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

func (z *simpleImpl) fetchSequential() (b Barrel, found bool) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l := z.logger.With(esl.String("currentBatchId", z.sequentialCurrentBatchBarrel))
	l.Debug("Sequential fetch from currentBatchId")

	for {
		if pipe, ok := z.pipes[z.sequentialCurrentBatchBarrel]; ok {
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
			delete(z.pipes, z.sequentialCurrentBatchBarrel)

			// fall thru
		}

		// finish when no more pipes found
		if len(z.pipes) < 1 {
			l.Debug("No more pipes")
			z.sequentialCurrentBatchBarrel = ""
			return b, false
		}

		// find next
		l.Debug("Find next currentBatchId")
		for b := range z.pipes {
			l.Debug("Next pipe", esl.String("batchId", b))
			z.sequentialCurrentBatchBarrel = b
			break
		}
	}
}

func (z *simpleImpl) fetchRandom() (b Barrel, found bool) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l := z.logger

	for {
		if len(z.pipes) < 1 {
			l.Debug("No more pipes")
			return b, false
		}

		r := rand.Intn(len(z.pipes))
		keys := make([]string, 0)
		for k := range z.pipes {
			keys = append(keys, k)
		}
		key := keys[r]
		ll := l.With(esl.String("batchBarrel", key))
		ll.Debug("Pick pipe key")
		pipe := z.pipes[key]

		ll.Debug("The pipe found with currentBatchId")
		if d0 := pipe.Dequeue(); d0 != nil {
			ll.Debug("Data found, dequeue success")
			d, err := FromBytes(d0)
			if err != nil {
				ll.Debug("Unable to unmarshal the message", esl.Error(err), esl.Binary("data", d0))
				return d, false
			}
			z.wip.Enqueue(d0)
			return d, true
		}

		ll.Debug("Data not found, closing the pipe")
		pipe.Close()
		delete(z.pipes, key)
	}
}

func (z *simpleImpl) fetchBalance() (b Barrel, found bool) {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	l := z.logger

	for {
		if len(z.pipes) < 1 {
			l.Debug("No more pipes")
			return b, false
		}

		var maxPipeLen int = -1
		var maxPipe string

		for k, p := range z.pipes {
			if maxPipeLen < p.Size() {
				maxPipeLen = p.Size()
				maxPipe = k
			}
		}

		l.Debug("Pipe selected", esl.String("maxPipe", maxPipe), esl.Int("pipeLen", maxPipeLen))

		if p, ok := z.pipes[maxPipe]; ok {
			if d0 := p.Dequeue(); d0 != nil {
				l.Debug("Data found, dequeue success")
				d, err := FromBytes(d0)
				if err != nil {
					l.Debug("Unable to unmarshal the message", esl.Error(err), esl.Binary("data", d0))
					return d, false
				}
				z.wip.Enqueue(d0)
				return d, true
			}

			l.Debug("Data not found, closing the pipe", esl.String("pipe", maxPipe))
			p.Close()
			delete(z.pipes, maxPipe)
		}
	}
}

func (z *simpleImpl) Fetch() (b Barrel, found bool) {
	switch z.policy {
	case FetchSequential:
		return z.fetchSequential()
	case FetchRandom:
		return z.fetchRandom()
	case FetchBalance:
		return z.fetchBalance()
	default:
		z.logger.Debug("Unknown fetch policy, fallback to sequential policy",
			esl.Any("policy", z.policy))
		return z.fetchSequential()
	}
}

func (z *simpleImpl) Size() int {
	z.pipesMutex.Lock()
	defer z.pipesMutex.Unlock()

	total := z.wip.Size()
	for _, pipe := range z.pipes {
		s := pipe.Size()
		total += s
	}

	return total
}
