package eq_bundle

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"github.com/watermint/toolbox/essentials/queue/eq_stat"
	"sync"
	"time"
)

var (
	ErrorOperationNotSupported = errors.New("operation not supported")
)

func NewPetty(logger esl.Logger, progress eq_progress.Progress, size int) Bundle {
	pf := eq_pipe.NewTransientSimple(logger)
	return &pettyImpl{
		logger:    logger,
		pipe:      pf.New("petty"),
		wip:       pf.New("wip"),
		pipeMutex: sync.Mutex{},
		pipeSize:  size,
		progress:  progress,
		stat:      eq_stat.New(),
	}
}

type pettyImpl struct {
	logger    esl.Logger
	pipe      eq_pipe.Pipe
	wip       eq_pipe.Pipe
	pipeMutex sync.Mutex
	pipeSize  int
	progress  eq_progress.Progress
	stat      eq_stat.Stat
}

func (z *pettyImpl) Enqueue(b Barrel) {
	bb := b.ToBytes()

	for z.pipeSize < z.pipe.Size() {
		time.Sleep(100 * time.Millisecond)
	}
	z.pipe.Enqueue(bb)
	z.stat.IncrEnqueue(b.MouldId, b.BatchId)
	if z.progress != nil {
		z.progress.OnEnqueue(b.MouldId, b.BatchId, z.stat)
	}
}

func (z *pettyImpl) Fetch() (b Barrel, found bool) {
	if d0 := z.pipe.Dequeue(); d0 != nil {
		d, err := FromBytes(d0)
		if err != nil {
			z.logger.Debug("Unable to unmarshal the message", esl.Error(err), esl.Binary("data", d0))
			return Barrel{}, false
		}
		z.wip.Enqueue(d0)
		return d, true
	}
	return Barrel{}, false
}

func (z *pettyImpl) Complete(b Barrel) {
	z.stat.IncrComplete(b.MouldId, b.BatchId)
	if z.progress != nil {
		z.progress.OnComplete(b.MouldId, b.BatchId, z.stat)
	}
	z.wip.Delete(b.ToBytes())
}

func (z *pettyImpl) Size() (total int) {
	return z.pipe.Size()
}

func (z *pettyImpl) SizeInProgress() int {
	return z.wip.Size()
}

func (z *pettyImpl) Close() {
	z.pipe.Close()
}

func (z *pettyImpl) Preserve() (session Session, err error) {
	return Session{}, ErrorOperationNotSupported
}
