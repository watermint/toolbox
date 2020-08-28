package es_stdout

import (
	"errors"
	"github.com/mattn/go-colorable"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"io"
	"io/ioutil"
	"os"
)

var (
	ErrorGeneralIOFailure = errors.New("general i/o error")
)

type Feature interface {
	IsQuiet() bool
	IsTest() bool
}

func newDefaultOut(test, quiet bool) io.WriteCloser {
	if quiet {
		return &syncOut{co: ioutil.Discard}
	}
	if test {
		if qt_secure.IsSecureEndToEndTest() {
			return &syncOut{co: ioutil.Discard}
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

func NewDiscard() io.WriteCloser {
	return &syncOut{co: ioutil.Discard}
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
			l := esl.Default()
			l.Debug("Recovery from the syncOut error", esl.Any("r", r))
			if errVal, ok := r.(error); ok {
				err = errVal
			} else {
				err = ErrorGeneralIOFailure
			}
		}
	}()
	return z.co.Write(p)
}

func (z syncOut) Close() error {
	return nil
}
