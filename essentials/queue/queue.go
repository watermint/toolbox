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
	// Enqueue data into the queue.
	Enqueue(p interface{})

	Batch(batchId string) Queue
}

type RootQueue interface {
	Queue
	Wait()
	Suspend() (session eq_bundle.Session, err error)
}

func New(l esl.Logger, numWorker int, factory eq_pipe.Factory, f interface{}, ctx ...interface{}) RootQueue {
	bundle := eq_bundle.NewSimple(l, factory)
	mould := eq_mould.New(bundle, f, ctx...)
	pump := eq_pump.New(l, bundle)
	pumpChan := pump.Start()
	worker := eq_worker.New(l, mould, pumpChan)
	worker.Startup(numWorker)

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

func (z queueImpl) Suspend() (session eq_bundle.Session, err error) {
	l := z.l

	l.Debug("Pump shutdown")
	z.pump.Shutdown()
	l.Debug("Pump shutdown: Done")

	l.Debug("Waiting for Worker")
	z.worker.Wait()
	l.Debug("Waiting for Worker: Done")

	l.Debug("Preserve")
	session, err = z.bundle.Preserve()
	l.Debug("Preserve: Done", esl.Any("session", session), esl.Error(err))
	return
}

func (z queueImpl) Wait() {
	l := z.l

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

func (z queueImpl) Enqueue(p interface{}) {
	z.mould.Pour(p)
}

func (z queueImpl) Batch(batchId string) Queue {
	z.mould = z.mould.Batch(batchId)
	return z
}
