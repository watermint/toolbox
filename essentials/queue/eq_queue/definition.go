package eq_queue

import (
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"sync"
)

type Definition interface {
	// Define a queue
	Define(queueId string, f interface{}, ctx ...interface{})

	// Traverse
	Each(f func(queueId string, f interface{}, ctx []interface{}))

	// Create a new container
	Current() Container

	// Restore container from the session
	Restore(session eq_bundle.Session) (Container, error)
}

func New(opt ...Opt) Definition {
	return &defImpl{
		opts:       defaultOpts().Apply(opt...),
		processors: make(map[string]interface{}),
		contexts:   make(map[string][]interface{}),
	}
}

type defImpl struct {
	opts                Opts
	container           Container
	containerCreateOnce sync.Once
	processors          map[string]interface{}
	contexts            map[string][]interface{}
}

func (z *defImpl) Current() Container {
	z.containerCreateOnce.Do(func() {
		z.container = newContainer(z, z.opts)
	})
	return z.container
}

func (z *defImpl) Restore(session eq_bundle.Session) (Container, error) {
	bundle, err := eq_bundle.RestoreSimple(z.opts.logger, z.opts.policy, z.opts.progress, z.opts.factory, session)
	if err != nil {
		return nil, err
	}
	z.container = newContainerWithBundle(z, bundle, z.opts)
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
