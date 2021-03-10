package eq_queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
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

func newContainer(definition Definition, opts Opts, el eq_mould.ErrorListener) Container {
	bundle := eq_bundle.NewSimple(opts.logger, opts.policy, opts.progress, opts.factory)
	return newContainerWithBundle(definition, bundle, opts, el)
}

func newContainerWithBundle(definition Definition, bundle eq_bundle.Bundle, opts Opts, el eq_mould.ErrorListener) Container {
	reg := eq_registry.New(bundle, []eq_mould.ErrorListener{el})
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
		container.define(queueId, f, opts.mouldOpts, ctx...)
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

func (z *conImpl) define(queueId string, f interface{}, opts eq_mould.Opts, ctx ...interface{}) Queue {
	mould := z.reg.Define(queueId, f, opts, ctx...)
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

	l.Debug("Flush progress bars")
	if z.opts.progress != nil {
		z.opts.progress.Flush()
	}
	l.Debug("Flush progress bars: Done")
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
