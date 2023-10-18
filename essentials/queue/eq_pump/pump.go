package eq_pump

import (
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"sync"
	"time"
)

var (
	// Duration of sleep when the bundle is empty.
	PollInterval = 50 * time.Millisecond
)

type Pump interface {
	// Start the pump.
	Start() chan eq_bundle.Barrel

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
		c:             make(chan eq_bundle.Barrel),
		closeCondLock: ccl,
		closeCond:     sync.NewCond(ccl),
	}
}

type pumpImpl struct {
	l             esl.Logger
	bundle        eq_bundle.Bundle
	startOnce     sync.Once
	c             chan eq_bundle.Barrel
	closeCondLock *sync.Mutex
	closeCond     *sync.Cond
	closeOnce     sync.Once
	pumpRunning   sync.WaitGroup
	pumpShutdown  bool
}

func (z *pumpImpl) logger() esl.Logger {
	return z.l.With(esl.String("routine", es_goroutine.GetGoRoutineName()))
}

func (z *pumpImpl) Start() chan eq_bundle.Barrel {
	l := z.logger()
	l.Debug("Try start")
	z.startOnce.Do(func() {
		l.Debug("Start")
		z.pumpRunning.Add(1)
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
	z.pumpShutdown = true
	z.pumpRunning.Wait()
}

func (z *pumpImpl) Close() {
	l := z.logger()
	if z.pumpShutdown {
		panic("Pump already shutdown")
	}

	l.Debug("Waiting for close condition")
	z.closeCondLock.Lock()
	z.closeCond.Wait()
	z.closeCondLock.Unlock()

	z.Shutdown()
}

func (z *pumpImpl) loop() {
	l := z.logger()
	defer func() {
		l.Debug("Exit from the loop")
		z.pumpRunning.Done()
	}()

	for {
		if z.pumpShutdown {
			l.Debug("Shutdown the loop")
			return
		}
		barrel, found := z.bundle.Fetch()
		//l.Debug("Fetch", esl.Bool("found", found))
		if found {
			if channelClosed := z.send(barrel); channelClosed {
				l.Debug("Channel closed, exit loop")
				return
			}
		} else {
			if s := z.bundle.SizeInProgress(); s < 1 {
				l.Debug("Data not found, broadcast close cond")
				z.closeCond.Broadcast()
			} else {
				l.Debug("The worker still in progress", esl.Int("InProgressSize", s))
			}

			l.Debug("Data not found, waiting for data")
			time.Sleep(PollInterval)
		}
	}
}

func (z *pumpImpl) send(barrel eq_bundle.Barrel) (channelClosed bool) {
	l := z.logger()
	defer func() {
		if err := recover(); err != nil {
			l.Debug("Caught an error", esl.Any("err", err))
			channelClosed = true
		}
	}()

	channelClosed = false
	//l.Debug("Push to the channel")
	z.c <- barrel
	return channelClosed
}
