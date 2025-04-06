package es_stdout

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
)

const (
	DefaultTimeout = 5 * time.Second
)

var (
	ErrorGeneralIOFailure = errors.New("general i/o error")
)

type Feature interface {
	IsQuiet() bool
	IsTest() bool
}

type deadlineWriter interface {
	io.WriteCloser
	SetDeadline(t time.Time) error
}

func newDefaultOut(test, quiet bool) io.WriteCloser {
	if quiet {
		return &syncOut{co: io.Discard}
	}
	if test {
		if qt_secure.IsSecureEndToEndTest() {
			return &syncOut{co: io.Discard}
		} else {
			return &syncOut{co: os.Stdout}
		}
	} else {
		if es_terminfo.IsOutColorTerminal() {
			return newWriteCloser(colorable.NewColorableStdout())
		} else {
			return newWriteCloser(os.Stdout)
		}
	}
}

func newDefaultErr(test, quiet bool) io.WriteCloser {
	if quiet {
		return &syncOut{co: io.Discard}
	}
	if test {
		if qt_secure.IsSecureEndToEndTest() {
			return &syncOut{co: io.Discard}
		} else {
			return &syncOut{co: os.Stderr}
		}
	} else {
		if es_terminfo.IsOutColorTerminal() {
			return newWriteCloser(colorable.NewColorableStderr())
		} else {
			return newWriteCloser(os.Stderr)
		}
	}
}

func NewDiscard() io.WriteCloser {
	return &syncOut{co: io.Discard}
}

func NewDefaultOut(feature Feature) io.WriteCloser {
	return newDefaultOut(feature.IsTest(), feature.IsQuiet())
}

func NewTestOut() io.WriteCloser {
	return newDefaultOut(true, false)
}

func NewDirectOut() io.WriteCloser {
	return newDefaultOut(false, false)
}

func NewDirectErr() io.WriteCloser {
	return newDefaultErr(false, false)
}

func newWriteCloser(co io.Writer) io.WriteCloser {
	return newSync(co)
}

func newSync(co io.Writer) io.WriteCloser {
	return &syncOut{
		co: co,
	}
}

type syncOut struct {
	co io.Writer
}

func (z syncOut) Write(p []byte) (n int, err error) {
	// Recovery option for I/O error
	// https://github.com/watermint/toolbox/issues/411
	defer func() {
		if r := recover(); err != nil {
			if errVal, ok := r.(error); ok {
				err = errVal
			} else {
				err = ErrorGeneralIOFailure
			}
		}
	}()

	if w, ok := z.co.(deadlineWriter); ok {
		_ = w.SetDeadline(time.Now().Add(DefaultTimeout))
	}
	return z.co.Write(p)
}

func (z syncOut) Close() error {
	return nil
}
