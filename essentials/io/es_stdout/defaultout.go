package es_stdout

import (
	"github.com/mattn/go-colorable"
	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
	"github.com/watermint/toolbox/quality/infra/qt_secure"
	"io"
	"io/ioutil"
	"os"
)

func NewDefaultOut(test bool) io.WriteCloser {
	if test {
		if qt_secure.IsSecureEndToEndTest() {
			return &Discard{}
		} else {
			return &Stdout{}
		}
	} else {
		if es_terminfo.IsOutColorTerminal() {
			return &Colorable{co: colorable.NewColorableStdout()}
		} else {
			return &Stdout{}
		}
	}
}

type Colorable struct {
	co io.Writer
}

func (z Colorable) Write(p []byte) (n int, err error) {
	return z.co.Write(p)
}

func (z Colorable) Close() error {
	return nil
}

type Stdout struct {
}

func (z *Stdout) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (z *Stdout) Close() error {
	return nil
}

type Discard struct {
}

func (z *Discard) Write(p []byte) (n int, err error) {
	return ioutil.Discard.Write(p)
}

func (z *Discard) Close() error {
	return nil
}
