package es_stdout

import (
	"github.com/mattn/go-colorable"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"io"
	"io/ioutil"
	"os"
	"sync"
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
	if es_env.IsEnabled("TOOLBOX_ASYNC_CONSOLE") {
		return newAsync(co)
	} else {
		return newSync(co)
	}
}

type AsyncMsg struct {
	Wr      io.Writer
	Payload []byte
}

var (
	asyncQueue         = make(chan *AsyncMsg, 100)
	discardQueue       = make(chan *AsyncMsg)
	asyncQueueLauncher sync.Once
)

func newSync(co io.Writer) io.WriteCloser {
	return &syncOut{
		co: co,
	}
}

type syncOut struct {
	co io.Writer
}

func (z syncOut) Write(p []byte) (n int, err error) {
	return z.co.Write(p)
}

func (z syncOut) Close() error {
	return nil
}

func asyncLoop() {
	for m := range asyncQueue {
		m.Wr.Write(m.Payload)
	}
}
func discardLoop() {
	for range discardQueue {
	}
}

func newAsync(co io.Writer) io.WriteCloser {
	asyncQueueLauncher.Do(func() {
		go asyncLoop()
		go discardLoop()
	})
	return &asyncOut{
		co: co,
	}
}

type asyncOut struct {
	co io.Writer
}

func (z *asyncOut) Write(p []byte) (n int, err error) {
	m := &AsyncMsg{
		Wr:      z.co,
		Payload: p,
	}
	select {
	case asyncQueue <- m:
	case discardQueue <- m:
	}
	return len(p), nil
}

func (z *asyncOut) Close() error {
	return nil
}
