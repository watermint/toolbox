package queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_pump"
	"github.com/watermint/toolbox/essentials/queue/eq_worker"
)

type Queue interface {
	Enqueue(p interface{})
	Batch(batchId string) Queue
	Wait()
	//Suspend()
}

func New(l esl.Logger, factory eq_pipe.Factory, f interface{}, ctx ...interface{}) Queue {
	bundle := eq_bundle.NewSimple(l, factory)
	mould := eq_mould.New(bundle, f, ctx...)
	pump := eq_pump.New(l, bundle)
	pumpChan := pump.Start()
	worker := eq_worker.New(l, mould, pumpChan)
	worker.Startup()

	return &queueImpl{
		l:      l,
		bundle: bundle,
		mould:  mould,
		pump:   pump,
		worker: worker,
	}
}

type queueImpl struct {
	l      esl.Logger
	bundle eq_bundle.Bundle
	mould  eq_mould.Mould
	pump   eq_pump.Pump
	worker eq_worker.Worker
}

func (z queueImpl) Wait() {
	l := z.l

	l.Debug("Waiting for Pump close")
	z.pump.Close()
	l.Debug("Waiting for Pump close: Done")

	l.Debug("Waiting for Worker")
	z.worker.Wait()
	l.Debug("Waiting for Worker: Done")
}

func (z queueImpl) Enqueue(p interface{}) {
	z.mould.Pour(p)
}

func (z queueImpl) Batch(batchId string) Queue {
	z.mould = z.mould.Batch(batchId)
	return z
}
