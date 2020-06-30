package es_mutex

import "time"

const (
	ShortTimeOut = 500 * time.Millisecond
)

type MutexWithTimeout interface {
	Do(f func())
}

func New() MutexWithTimeout {
	return NewWithShortTimeout()
}

func NewWithShortTimeout() MutexWithTimeout {
	return NewWithTimeout(ShortTimeOut)
}

func NewWithTimeout(timeout time.Duration) MutexWithTimeout {
	return &mutexImpl{
		lock:           make(chan struct{}, 1),
		timeout:        timeout,
		timeoutHandler: func() {},
		maxRetries:     1,
	}
}

func NewWithTimeoutRetry(timeout time.Duration, retries int, timeoutHandler func()) MutexWithTimeout {
	return &mutexImpl{
		lock:           make(chan struct{}, 1),
		timeout:        timeout,
		timeoutHandler: timeoutHandler,
		maxRetries:     retries,
	}
}

type mutexImpl struct {
	lock           chan struct{}
	timeout        time.Duration
	timeoutHandler func()
	maxRetries     int
}

func (z *mutexImpl) Do(f func()) {
	for i := 0; i < z.maxRetries; i++ {
		select {
		case z.lock <- struct{}{}:
			// do
			f()

			// unlock
			<-z.lock
			return

		case <-time.After(z.timeout):
			z.timeoutHandler()
		}
	}
}
