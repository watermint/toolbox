package ut_io

import (
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
		return &Stdout{}
	}
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
