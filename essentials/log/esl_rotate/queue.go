package esl_rotate

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/watermint/toolbox/essentials/concurrency/es_mutex"
	"github.com/watermint/toolbox/essentials/concurrency/es_timeout"
	"github.com/watermint/toolbox/essentials/file/es_gzip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
)

const (
	RotateEnqueueTimeout = 10 * time.Second
)

type MsgPurge struct {
	Path string
	Opts RotateOpts
}

type MsgOut struct {
	Path string
	Opts RotateOpts
}

type MsgRotate struct {
	Opts RotateOpts
}

var (
	queuePurge        chan MsgPurge
	queuePurgeStatus  sync.WaitGroup
	queueOut          chan MsgOut
	queueOutStatus    sync.WaitGroup
	queueRotate       chan MsgRotate
	queueRotateStatus sync.WaitGroup
	queueMutex        = es_mutex.New()
	queueAvailable    = false

	// logs which unable to remove
	poisonLogs = make(map[string]error)
)

func purgeLoop() {
	for m := range queuePurge {
		// consume poison message
		if _, ok := poisonLogs[m.Path]; ok {
			continue
		}

		// ensure file exists & not a directory
		ls, err := os.Lstat(m.Path)
		if err != nil || ls.IsDir() {
			continue
		}

		// execute hook
		if m.Opts.rotateHook != nil {
			m.Opts.rotateHook(m.Path)
		}

		l := esl.ConsoleOnly()
		// clean up
		l.Debug("Removing the old log that exceeds the quota", esl.String("path", m.Path))
		err = os.Remove(m.Path)
		if err != nil {
			l.Debug("Unable to remove log file", esl.String("path", m.Path), esl.Error(err))
			poisonLogs[m.Path] = err
		}
	}
	queuePurgeStatus.Done()
}

func execRotate(m MsgRotate) {
	l := esl.ConsoleOnly()

	targets, err := m.Opts.PurgeTargets()
	if err != nil {
		l.Debug("Unable to read log directory", esl.String("path", m.Opts.BasePath()), esl.Error(err))
		return
	}

	for _, path := range targets {
		l.Debug("Purge log", esl.String("entry", path))
		enqueuePurge(MsgPurge{
			Path: path,
			Opts: m.Opts,
		})
	}
}

func rotateLoop() {
	for m := range queueRotate {
		execRotate(m)
	}
	queueRotateStatus.Done()
}

func outLoop() {
	for m := range queueOut {
		execOut(m)
	}
	queueOutStatus.Done()
}

func execOut(m MsgOut) {
	if !m.Opts.IsCompress() {
		return
	}

	// ignore errors. the original file retains if compression failed.
	// that will be processed on purge process.
	_, _ = es_gzip.Compress(m.Path)
}

func enqueuePurge(m MsgPurge) {
	if queuePurge != nil {
		queuePurge <- m
	}
}

func rotateOut(m MsgOut) (ok bool) {
	ok = false
	es_timeout.DoWithTimeout(RotateEnqueueTimeout, func(ctx context.Context) {
		if queueOut != nil {
			queueOut <- m
			ok = true
		}
	})
	return
}

func enqueueRotate(m MsgRotate) {
	if queueRotate != nil {
		queueRotate <- m
	}
}

func Startup() {
	queueMutex.Do(func() {
		if queueAvailable {
			return
		}
		queueRotate = make(chan MsgRotate)
		queueOut = make(chan MsgOut)
		queuePurge = make(chan MsgPurge)
		queueRotateStatus.Add(1)
		queueOutStatus.Add(1)
		queuePurgeStatus.Add(1)
		go rotateLoop()
		go outLoop()
		go purgeLoop()
		queueAvailable = true
	})
	app_shutdown.AddShutdownHook(Shutdown)
}

func Shutdown() {
	queueMutex.Do(func() {
		l := esl.ConsoleOnly()
		if queueRotate != nil {
			l.Debug("Shutdown queue rotate")
			close(queueRotate)
			queueRotateStatus.Wait()
			queueRotate = nil
		}

		if queueOut != nil {
			l.Debug("Shutdown queue out")
			close(queueOut)
			queueOutStatus.Wait()
			queueOut = nil
		}

		if queuePurge != nil {
			l.Debug("Shutdown queue purge")
			close(queuePurge)
			queuePurgeStatus.Wait()
			queuePurge = nil
		}
		queueAvailable = false
	})
}
