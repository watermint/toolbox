package eq_sequence

import (
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/infra/control/app_ambient"
)

type Stage interface {
	Define(queueId string, f interface{}, ctx ...interface{})

	Get(queueId string) eq_queue.Queue
}

type StageController interface {
	Stage

	Exec()
}

type Sequence interface {
	Do(exec func(s Stage))

	DoThen(exec func(s Stage)) Sequence
}

func newStg(d eq_queue.Definition) Stage {
	return &stgImpl{
		d: d,
	}
}

type stgImpl struct {
	d eq_queue.Definition
}

func (z *stgImpl) Define(queueId string, f interface{}, ctx ...interface{}) {
	z.d.Define(queueId, f, ctx...)
}

func (z *stgImpl) Get(queueId string) eq_queue.Queue {
	return z.d.Current().MustGet(queueId)
}

type seqImpl struct {
	opt []eq_queue.Opt
}

func (z *seqImpl) Do(exec func(s Stage)) {
	_ = z.DoThen(exec)
}

func (z *seqImpl) DoThen(exec func(s Stage)) Sequence {
	d := eq_queue.New(z.opt...)
	s := newStg(d)
	app_ambient.Current.SuppressProgress()
	exec(s)
	d.Current().Wait()
	app_ambient.Current.ResumeProgress()

	return z
}

func New(opt ...eq_queue.Opt) Sequence {
	return &seqImpl{
		opt: opt,
	}
}
