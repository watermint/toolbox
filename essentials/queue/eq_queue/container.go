package eq_queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"github.com/watermint/toolbox/essentials/queue/eq_pump"
	"github.com/watermint/toolbox/essentials/queue/eq_registry"
	"github.com/watermint/toolbox/essentials/queue/eq_worker"
)

type Container interface {
	// Get queue for queueId
	Get(queueId string) (q Queue, found bool)

	// Get queue for queueId. Panic when a queue is not found for the id.
	MustGet(queueId string) Queue

	// Wait for complete all queues
	Wait()

	// Suspend queue. Workers will be shutdown, but do not wait for worker shutdown.
	Suspend() (session eq_bundle.Session, err error)
}

func newContainer(definition Definition, opts Opts) Container {
	bundle := eq_bundle.NewSimple(opts.logger, opts.progress, opts.factory)
	return newContainerWithBundle(definition, bundle, opts)
}

func newContainerWithBundle(definition Definition, bundle eq_bundle.Bundle, opts Opts) Container {
	reg := eq_registry.New(bundle)
	pump := eq_pump.New(opts.logger, bundle)
	pumpChan := pump.Start()
	worker := eq_worker.New(opts.logger, reg, pumpChan)
	worker.Startup(opts.numWorker)

	container := &conImpl{
		reg:    reg,
		bundle: bundle,
		queues: make(map[string]Queue),
		pump:   pump,
		worker: worker,
		opts:   opts,
	}
	definition.Each(func(queueId string, f interface{}, ctx []interface{}) {
		container.define(queueId, f, ctx...)
	})
	return container
}

type conImpl struct {
	reg    eq_registry.Registry
	bundle eq_bundle.Bundle
	queues map[string]Queue
	pump   eq_pump.Pump
	worker eq_worker.Worker
	opts   Opts
}

func (z *conImpl) define(queueId string, f interface{}, ctx ...interface{}) Queue {
	mould := z.reg.Define(queueId, f, ctx...)
	queue := newQueue(mould)
	z.queues[queueId] = queue
	return queue
}

func (z *conImpl) Get(queueId string) (q Queue, found bool) {
	q, found = z.queues[queueId]
	return
}

func (z *conImpl) MustGet(queueId string) Queue {
	if q, found := z.Get(queueId); found {
		return q
	}
	z.opts.logger.Warn("Queue not found for the queueId", esl.String("queueId", queueId))
	panic("queue not found")
}

func (z *conImpl) Wait() {
	l := z.opts.logger

	l.Debug("Waiting for Pump close")
	z.pump.Close()
	l.Debug("Waiting for Pump close: Done")

	l.Debug("Waiting for Worker")
	z.worker.Wait()
	l.Debug("Waiting for Worker: Done")

	l.Debug("Waiting for Bundle close")
	z.bundle.Close()
	l.Debug("Waiting for Bundle close: Done")
}

func (z *conImpl) Suspend() (session eq_bundle.Session, err error) {
	l := z.opts.logger

	l.Debug("Pump shutdown")
	z.pump.Shutdown()
	l.Debug("Pump shutdown: Done")

	l.Debug("Preserve")
	session, err = z.bundle.Preserve()
	l.Debug("Preserve: Done", esl.Any("session", session), esl.Error(err))
	return
}

type Opts struct {
	logger    esl.Logger
	numWorker int
	factory   eq_pipe.Factory
	progress  eq_progress.Progress
}

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

func defaultOpts() Opts {
	return Opts{
		logger:    esl.Default(),
		numWorker: 1,
		progress:  nil,
		factory:   eq_pipe.NewTransientSimple(esl.Default()),
	}
}

type Opt func(o Opts) Opts

func Logger(l esl.Logger) Opt {
	return func(o Opts) Opts {
		o.logger = l
		return o
	}
}

func NumWorker(n int) Opt {
	return func(o Opts) Opts {
		o.numWorker = n
		return o
	}
}

func Progress(p eq_progress.Progress) Opt {
	return func(o Opts) Opts {
		o.progress = p
		return o
	}
}

func Factory(f eq_pipe.Factory) Opt {
	return func(o Opts) Opts {
		o.factory = f
		return o
	}
}
