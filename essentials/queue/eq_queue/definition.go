package eq_queue

import (
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"sync"
)

type ErrorListener eq_mould.ErrorListener

type Definition interface {
	// Define a queue
	Define(queueId string, f interface{}, ctx ...interface{})

	// Traverse
	Each(f func(queueId string, f interface{}, ctx []interface{}))

	// Create a new container
	Current() Container

	// Restore container from the session
	Restore(session eq_bundle.Session) (Container, error)

	// Add error handler listener
	AddErrorListener(h ErrorListener)
}

func ExecWithQueue(f func(q Definition)) {
	q := New()
	f(q)
	q.Current().Wait()
}

func New(opt ...Opt) Definition {
	return &defImpl{
		handlers:            make([]ErrorListener, 0),
		opts:                defaultOpts().Apply(opt...),
		containerCreateOnce: sync.Once{},
		processors:          make(map[string]interface{}),
		contexts:            make(map[string][]interface{}),
	}
}

type defImpl struct {
	handlers            []ErrorListener
	opts                Opts
	container           Container
	containerCreateOnce sync.Once
	processors          map[string]interface{}
	contexts            map[string][]interface{}
}

func (z *defImpl) AddErrorListener(h ErrorListener) {
	z.handlers = append(z.handlers, h)
}

func (z *defImpl) notifyError(err error, mouldId string, batchId string, p interface{}) {
	for _, h := range z.handlers {
		h(err, mouldId, batchId, p)
	}
	for _, h := range z.opts.errorHandler {
		h(err, mouldId, batchId, p)
	}
}

func (z *defImpl) Current() Container {
	z.containerCreateOnce.Do(func() {
		z.container = newContainer(z, z.opts, z.notifyError)
	})
	return z.container
}

func (z *defImpl) Restore(session eq_bundle.Session) (Container, error) {
	bundle, err := eq_bundle.RestoreSimple(z.opts.logger, z.opts.policy, z.opts.progress, z.opts.factory, session)
	if err != nil {
		return nil, err
	}
	z.container = newContainerWithBundle(z, bundle, z.opts, z.notifyError)
	return z.container, nil
}

func (z *defImpl) Each(f func(queueId string, f interface{}, ctx []interface{})) {
	for queueId, g := range z.processors {
		c := z.contexts[queueId]
		f(queueId, g, c)
	}
}

func (z *defImpl) Define(queueId string, f interface{}, ctx ...interface{}) {
	z.processors[queueId] = f
	z.contexts[queueId] = ctx
}
