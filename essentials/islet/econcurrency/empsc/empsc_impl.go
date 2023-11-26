package empsc

import "sync"

type channelInternal interface {
	Channel

	Start()

	// Send a value to the consumer. This func will block if a channel is full.
	Send(v interface{})
	OnConsumeError(v interface{}, err error)
	OnConsumeSuccess(v interface{})
	OnConsumeStart(v interface{})
	OnProducerStart(v interface{})
	OnProducerSuccess(v interface{})
	OnOverflowIgnore(v interface{})
	OnClosed()
}

type channelOpts struct {
	BufSize int
}

func (z channelOpts) Apply(opts []ChannelOpt) channelOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type ChannelOpt func(o channelOpts) channelOpts

func ChannelBufSize(size int) ChannelOpt {
	return func(o channelOpts) channelOpts {
		o.BufSize = size
		return o
	}
}

func New(c Consumer, opts ...ChannelOpt) Channel {
	co := channelOpts{}.Apply(opts)
	ch := &channelImpl{
		mutex:    sync.Mutex{},
		closed:   false,
		queue:    make(chan interface{}, co.BufSize),
		consumer: c,
	}
	ch.Start()
	return ch
}

type channelImpl struct {
	mutex    sync.Mutex
	closed   bool
	queue    chan interface{}
	consumer Consumer
	//overflow Consumer
}

func (z *channelImpl) OnConsumeError(v interface{}, err error) {
}

func (z *channelImpl) OnConsumeSuccess(v interface{}) {
}

func (z *channelImpl) OnConsumeStart(v interface{}) {
}

func (z *channelImpl) OnProducerStart(v interface{}) {
}

func (z *channelImpl) OnProducerSuccess(v interface{}) {
}

func (z *channelImpl) OnOverflowIgnore(v interface{}) {
}

func (z *channelImpl) OnClosed() {
}

func (z *channelImpl) Start() {
	go z.consumeLoop()
}

func (z *channelImpl) consumeLoop() {
	for v := range z.queue {
		z.OnConsumeStart(v)
		if cErr := z.consumer(v); cErr != nil {
			z.OnConsumeError(v, cErr)
		} else {
			z.OnConsumeSuccess(v)
		}
	}
	z.OnClosed()
}

func (z *channelImpl) Close() error {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.closed {
		return nil
	}
	close(z.queue)
	z.closed = true
	z.OnClosed()

	return nil
}

func (z *channelImpl) Producer() Producer {
	return &producerImpl{
		c: z,
	}
}

func (z *channelImpl) Send(v interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if !z.closed {
		z.OnProducerStart(v)
		z.queue <- v
		z.OnProducerSuccess(v)
	} else {
		//if z.overflow != nil {
		//	go func() {
		//		z.OnOverflowSuccess(v)
		//		if oErr := z.overflow(v); oErr != nil {
		//			z.OnOverflowError(v, oErr)
		//		} else {
		//			z.OnOverflowSuccess(v)
		//		}
		//	}()
		//} else {
		z.OnOverflowIgnore(v)
		//}
	}
}

type producerImpl struct {
	c channelInternal
}

func (z producerImpl) Send(v interface{}) {
	z.c.Send(v)
}
