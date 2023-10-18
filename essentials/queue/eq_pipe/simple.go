package eq_pipe

import (
	"container/list"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe_preserve"
	"reflect"
	"sync"
)

var (
	ErrorIncompatibleSessionData = errors.New("incompatible session data")
)

const (
	SimpleFormatVersionFirst = iota + 1
)

func NewTransientSimple(l esl.Logger) Factory {
	return &simpleFactory{
		l: l,
		f: eq_pipe_preserve.NopFactory(),
	}
}

func NewSimple(l esl.Logger, f eq_pipe_preserve.Factory) Factory {
	return &simpleFactory{
		l: l,
		f: f,
	}
}

type simpleFactory struct {
	l esl.Logger
	f eq_pipe_preserve.Factory
}

func (z simpleFactory) Restore(id SessionId) (pipe Pipe, err error) {
	l := z.l.With(esl.String("sessionId", string(id)))
	l.Debug("Restore")

	restore := z.f.NewRestorer(string(id))
	info := &simplePipeMeta{}
	pipeData := list.New()

	infoLoader := func(d []byte) error {
		l.Debug("Load info")
		if err := json.Unmarshal(d, info); err != nil {
			l.Debug("Unable to unmarshal info", esl.Error(err))
			return err
		}

		if info.Version != SimpleFormatVersionFirst {
			l.Debug("Incompatible version found", esl.Any("info", info))
			return ErrorIncompatibleSessionData
		}

		return nil
	}

	loader := func(d []byte) error {
		l.Debug("Load entry", esl.ByteString("entry", d))
		pipeData.PushBack(d)
		return nil
	}

	if err := restore.Restore(infoLoader, loader); err != nil {
		l.Debug("Unable to restore", esl.Error(err))
		return nil, err
	}

	pipe = &simpleImpl{
		batchId: info.BatchId,
		pr:      z.f.NewPreserver(),
		lg:      z.l.With(esl.String("batchId", info.BatchId)),
		l:       pipeData,
	}

	l.Debug("Pipe restored", esl.Any("info", info), esl.Int("entries", pipeData.Len()))

	return pipe, nil
}

func (z simpleFactory) New(batchId string) Pipe {
	return &simpleImpl{
		batchId: batchId,
		pr:      z.f.NewPreserver(),
		lg:      z.l.With(esl.String("batchId", batchId)),
		l:       list.New(),
	}
}

type simplePipeMeta struct {
	Version int    `json:"version"`
	BatchId string `json:"batch_id"`
}

type simpleImpl struct {
	batchId string
	pr      eq_pipe_preserve.Preserver
	lg      esl.Logger
	l       *list.List
	m       sync.Mutex
}

func (z *simpleImpl) Preserve() (id SessionId, err error) {
	l := z.lg
	z.m.Lock()
	defer z.m.Unlock()

	l.Debug("Start preserve")
	if err := z.pr.Start(); err != nil {
		l.Debug("Unable to start preserve session", esl.Error(err))
		return "", err
	}

	for p := z.l.Front(); p != nil; p = p.Next() {
		if err := z.pr.Add(p.Value.([]byte)); err != nil {
			l.Debug("Unable to add data", esl.Error(err))
			return "", err
		}
	}

	info := &simplePipeMeta{
		Version: SimpleFormatVersionFirst,
		BatchId: z.batchId,
	}
	infoData, err := json.Marshal(info)
	if err != nil {
		l.Debug("Unable to create info data", esl.Error(err))
		return "", err
	}

	l.Debug("Commit")
	sessionId, err := z.pr.Commit(infoData)
	l.Debug("Committed", esl.String("sessionId", sessionId), esl.Error(err))
	return SessionId(sessionId), err
}

func (z *simpleImpl) Delete(d []byte) {
	l := z.lg
	l.Debug("Delete")
	z.m.Lock()
	defer z.m.Unlock()

	l.Debug("Delete: Scan from the front")
	p := z.l.Front()
	for {
		if reflect.DeepEqual(d, p.Value) {
			l.Debug("Delete: found, remove")
			z.l.Remove(p)
			return
		}

		p = p.Next()
		if p == nil {
			l.Debug("Delete: Not found")
			return
		}
	}
}

func (z *simpleImpl) Close() {
	l := z.lg
	l.Debug("Close")
}

func (z *simpleImpl) Enqueue(d []byte) {
	z.m.Lock()
	defer z.m.Unlock()
	z.l.PushBack(d)
}

func (z *simpleImpl) Dequeue() (d []byte) {
	l := z.lg

	z.m.Lock()
	defer z.m.Unlock()

	if z.l.Len() < 1 {
		l.Debug("Not found")
		return nil
	}

	e := z.l.Front()
	d = e.Value.([]byte)
	z.l.Remove(e)
	return
}

func (z *simpleImpl) Size() int {
	return z.l.Len()
}
