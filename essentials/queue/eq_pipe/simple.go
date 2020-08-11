package eq_pipe

import (
	"container/list"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
	"sync"
)

var (
	ErrorPreserveIsNotSupported = errors.New("this pipe does not support preserve")
)

func NewSimple(l esl.Logger) Factory {
	return &simpleFactory{
		l: l,
	}
}

type simpleFactory struct {
	l esl.Logger
}

func (z simpleFactory) New(batchId string) Pipe {
	return &simpleImpl{
		lg: z.l.With(esl.String("batchId", batchId)),
		l:  list.New(),
	}
}

type simpleImpl struct {
	lg esl.Logger
	l  *list.List
	m  sync.Mutex
}

func (z *simpleImpl) Preserve() (id SessionId, err error) {
	return "", ErrorPreserveIsNotSupported
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
	l := z.lg

	z.m.Lock()
	defer z.m.Unlock()
	l.Debug("Enqueue")
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

	l.Debug("Dequeue")
	e := z.l.Front()
	d = e.Value.([]byte)
	z.l.Remove(e)
	return
}

func (z *simpleImpl) Size() int {
	return z.l.Len()
}
