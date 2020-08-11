package eq_pump

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"sync"
	"time"
)

type Pump interface {
	// Start the pump.
	Start() chan eq_bundle.Data

	// Wait & close the pump
	Close()

	// Stop the pump
	Shutdown()
}

func New(l esl.Logger, bundle eq_bundle.Bundle) Pump {
	ccl := &sync.Mutex{}
	return &pumpImpl{
		l:             l,
		bundle:        bundle,
		c:             make(chan eq_bundle.Data),
		closeCondLock: ccl,
		closeCond:     sync.NewCond(ccl),
	}
}

type pumpImpl struct {
	l             esl.Logger
	bundle        eq_bundle.Bundle
	startOnce     sync.Once
	c             chan eq_bundle.Data
	closeCondLock *sync.Mutex
	closeCond     *sync.Cond
	closeOnce     sync.Once
}

func (z *pumpImpl) logger() esl.Logger {
	return z.l.With(esl.String("routine", es_goroutine.GetGoRoutineName()))
}

func (z *pumpImpl) Start() chan eq_bundle.Data {
	l := z.logger()
	l.Debug("Try start")
	z.startOnce.Do(func() {
		l.Debug("Start")
		go z.loop()
	})
	return z.c
}

func (z *pumpImpl) Shutdown() {
	l := z.logger()

	l.Debug("Try Shutdown")
	z.closeOnce.Do(func() {
		l.Debug("Shutdown")
		close(z.c)
	})
}

func (z *pumpImpl) Close() {
	l := z.logger()

	l.Debug("Waiting for close condition")
	z.closeCondLock.Lock()
	z.closeCond.Wait()
	z.closeCondLock.Unlock()

	z.Shutdown()
}

func (z *pumpImpl) loop() {
	l := z.logger()
	for {
		d, found := z.bundle.Fetch()
		l.Debug("Fetch", esl.Bool("found", found))
		if found {
			if channelClosed := z.send(d); channelClosed {
				l.Debug("Channel closed, exit loop")
				return
			}
		} else {
			l.Debug("Data not found, broadcast close cond")
			z.closeCond.Broadcast()

			l.Debug("Data not found, waiting for data")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func (z *pumpImpl) send(d eq_bundle.Data) (channelClosed bool) {
	l := z.logger()
	defer func() {
		if err := recover(); err != nil {
			l.Debug("Caught an error", esl.Any("err", err))
			channelClosed = true
		}
	}()

	channelClosed = false
	l.Debug("Push to the channel")
	z.c <- d
	return channelClosed
}
