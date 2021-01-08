package eq_sequence

import (
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func New(opt ...eq_queue.Opt) Sequence {
	return &seqImpl{
		opt: opt,
	}
}

// Batch sequence stage:
type Stage interface {
	// Define function. This function must be called before Get.
	Define(queueId string, f interface{}, ctx ...interface{})

	// Get queue by id.
	Get(queueId string) eq_queue.Queue
}

// Batch sequence
type Sequence interface {
	// Do single stage
	Do(exec func(s Stage), opts ...DoOpt)

	// Do single stage, then returns next stage.
	DoThen(exec func(s Stage), opts ...DoOpt) Sequence
}

type DoOpts struct {
	ErrorHandler eq_mould.ErrorHandler
	SingleThread bool
}

func (z DoOpts) Apply(opts []DoOpt) DoOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type DoOpt func(o DoOpts) DoOpts

func SingleThread() DoOpt {
	return func(o DoOpts) DoOpts {
		o.SingleThread = true
		return o
	}
}

func ErrorHandler(h eq_mould.ErrorHandler) DoOpt {
	return func(o DoOpts) DoOpts {
		o.ErrorHandler = h
		return o
	}
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

func (z *seqImpl) Do(exec func(s Stage), opts ...DoOpt) {
	_ = z.DoThen(exec, opts...)
}

func (z *seqImpl) DoThen(exec func(s Stage), opts ...DoOpt) Sequence {
	dOps := DoOpts{}.Apply(opts)
	qOps := make([]eq_queue.Opt, 0)
	qOps = append(qOps, z.opt...)
	if dOps.ErrorHandler != nil {
		qOps = append(qOps, eq_queue.ErrorHandler(dOps.ErrorHandler))
	}
	if dOps.SingleThread {
		qOps = append(qOps, eq_queue.NumWorker(1))
	}
	d := eq_queue.New(qOps...)
	s := newStg(d)
	exec(s)
	d.Current().Wait()

	return z
}
