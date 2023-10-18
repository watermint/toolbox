package lgw_badger

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"testing"
)

func TestNewLogWrapper(t *testing.T) {
	l := NewLogWrapper(esl.Default())
	l.Infof("Hello, World [%s]", "info")
	l.Warningf("Hello, World [%s]", "warn")
	l.Errorf("Hello, World [%s]", "error")
	l.Debugf("Hello, World [%s]", "debug")
}
