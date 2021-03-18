package esl_rotate

import (
	"errors"
	"github.com/watermint/toolbox/essentials/concurrency/es_mutex"
	"github.com/watermint/toolbox/essentials/io/es_timeout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"time"
)

const (
	logFileExtension = ".log"
	logWriteTimeout  = 15 * time.Second
)

var (
	ErrorLogFileNotAvailable = errors.New("log file is not available")
)

type Writer interface {
	io.WriteCloser
	Open(opts ...RotateOpt) error
	UpdateOpt(opt RotateOpt)
}

func NewWriter(basePath, baseName string) Writer {
	w := writerImpl{
		m: es_mutex.New(),
	}
	w.ro = NewRotateOpts().Apply(
		BasePath(basePath),
		BaseName(baseName),
	)
	Startup()
	return &w
}

// This implementation is fully stateful
type writerImpl struct {
	current       *os.File
	currentWriter es_timeout.TimeoutWriter
	ro            RotateOpts
	written       int64
	m             es_mutex.MutexWithTimeout
}

func (z *writerImpl) UpdateOpt(opt RotateOpt) {
	z.ro = z.ro.Apply(opt)
}

func (z *writerImpl) Write(p []byte) (n int, err error) {
	z.m.Do(func() {
		if z.current == nil {
			// #507
			// err = ErrorLogFileNotAvailable
			return
		}
		n, err = z.currentWriter.Write(p)
		z.written += int64(n)
		if err != nil {
			return
		}
		if z.ro.ChunkSize() < z.written {
			err = z.rotate()
		}
		return
	})
	return
}

// this function must called from caller who owns mutex lock
func (z *writerImpl) rotate() (err error) {
	// close current log
	err = z.closeCurrent()
	if err != nil {
		return
	}

	// create new current log
	err = z.createCurrent()
	if err != nil {
		return
	}

	// Enqueue rotation
	enqueueRotate(MsgRotate{
		Opts: z.ro,
	})

	return
}

// this function must called from caller who owns mutex lock
func (z *writerImpl) createCurrent() (err error) {
	// Can't use es_log.Default(), because that cause race.
	// Use TOOLBOX_DEBUG_VERBOSE for view this log
	l := esl.ConsoleOnly()
	path := z.ro.CurrentPath()

	l.Debug("create", esl.String("path", path))
	z.current, err = os.Create(path)
	z.currentWriter = es_timeout.New(z.current, logWriteTimeout)
	return
}

// this function must called from caller who owns mutex lock
func (z *writerImpl) closeCurrent() (err error) {
	// Can't use es_log.Default(), because that cause race
	// Use TOOLBOX_DEBUG_VERBOSE for view this log
	l := esl.ConsoleOnly().With(esl.Int64("written", z.written))

	// flush written bytes
	z.written = 0

	// return on the file already closed
	if z.current == nil {
		l.Debug("already closed")
		return nil
	}

	// close
	name := z.current.Name()
	err = z.current.Close()

	l.Debug("Close", esl.String("name", name), esl.Error(err))
	if err != nil {
		// #507 changed from WARN to DEBUG
		l.Debug("Unable to close current log file", esl.String("path", name), esl.Error(err))
	}
	z.current = nil
	z.currentWriter = nil

	ror := rotateOut(MsgOut{
		Path: name,
		Opts: z.ro,
	})
	if !ror {
		// #507 changed from WARN to DEBUG
		l.Debug("Unable to enqueue rotate out", esl.String("path", name))
	}

	return
}

func (z *writerImpl) Close() (err error) {
	z.m.Do(func() {
		err = z.closeCurrent()
	})
	return
}

func (z *writerImpl) Open(opts ...RotateOpt) (err error) {
	z.m.Do(func() {
		// Do nothing when it already opened
		if z.current != nil {
			return
		}
		z.ro = z.ro.Apply(opts...)
		err = z.createCurrent()
	})
	return
}
