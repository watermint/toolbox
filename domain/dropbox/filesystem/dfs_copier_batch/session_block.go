package dfs_copier_batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/log/esl"
	"go.uber.org/atomic"
	"sync"
)

type BlockSession interface {
	// AddBlock Add block
	AddBlock(sessionId string, offset int64)

	// FinishSuccess Tell finish operation
	FinishSuccess(sessionId string, offset int64)

	// FinishFailure Tell failure
	FinishFailure(sessionId string, err error)
}

func NewBlockSession(ctx dbx_client.Client, bs BatchSessions) BlockSession {
	cbs := &copierBlockSession{
		ctx:                ctx,
		bs:                 bs,
		backlog:            atomic.Uint32{},
		backlogOffsetMutex: sync.Mutex{},
		backlogOffsets:     make(map[string]map[int64]bool),
	}
	return cbs
}

type copierBlockSession struct {
	ctx                dbx_client.Client
	bs                 BatchSessions
	backlog            atomic.Uint32
	backlogOffsetMutex sync.Mutex
	backlogOffsets     map[string]map[int64]bool
}

func (z *copierBlockSession) AddBlock(sessionId string, offset int64) {
	l := z.ctx.Log().With(esl.String("sessionId", sessionId), esl.Int64("offset", offset))
	l.Debug("Add the block")

	z.backlogOffsetMutex.Lock()
	if bo, ok := z.backlogOffsets[sessionId]; ok {
		bo[offset] = true
	} else {
		bo := make(map[int64]bool)
		bo[offset] = true
		z.backlogOffsets[sessionId] = bo
	}

	z.backlogOffsetMutex.Unlock()
}

func (z *copierBlockSession) FinishSuccess(sessionId string, offset int64) {
	l := z.ctx.Log().With(esl.String("sessionId", sessionId), esl.Int64("offset", offset))

	allBlockFinished := false
	z.backlogOffsetMutex.Lock()
	if bo, ok := z.backlogOffsets[sessionId]; ok {
		delete(bo, offset)
		if len(bo) < 1 {
			l.Debug("look like it's the last block of this session")
			allBlockFinished = true
			delete(z.backlogOffsets, sessionId)
		} else {
			z.backlogOffsets[sessionId] = bo
		}
	} else {
		l.Warn("The session backlog data not found, try mark as it's the last block")
		allBlockFinished = true
		z.noLockRecalculateBacklogs()
	}
	z.backlogOffsetMutex.Unlock()

	l.Debug("Finish success", esl.Bool("allBlockFinished", allBlockFinished))

	if allBlockFinished {
		l.Debug("Report finish the session")
		z.bs.FinishBlockUploads(sessionId)
	}
}

func (z *copierBlockSession) noLockRecalculateBacklogs() {
	l := z.ctx.Log()
	numBacklog := 0
	for _, bo := range z.backlogOffsets {
		numBacklog += len(bo)
	}
	l.Debug("New number of backlogs", esl.Int("backlogs", numBacklog))
	z.backlog.Store(uint32(numBacklog))
}

func (z *copierBlockSession) FinishFailure(sessionId string, err error) {
	l := z.ctx.Log().With(esl.String("sessionId", sessionId))
	l.Debug("Finish failure", esl.Error(err))

	// recalculate number of backlogs
	z.backlogOffsetMutex.Lock()
	delete(z.backlogOffsets, sessionId)
	z.noLockRecalculateBacklogs()
	z.backlogOffsetMutex.Unlock()

	z.bs.AbortSession(sessionId, err)
}
