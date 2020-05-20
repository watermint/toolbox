package lgw_badger

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func New(l esl.Logger) badger.Logger {
	return &badgerLogger{
		l: l.AddCallerSkip(1),
	}
}

type badgerLogger struct {
	l esl.Logger
}

func (z *badgerLogger) Errorf(f string, p ...interface{}) {
	z.l.Warn(fmt.Sprintf(f, p...), esl.String("level", "error"))
}

func (z *badgerLogger) Warningf(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), esl.String("level", "warn"))
}

func (z *badgerLogger) Infof(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), esl.String("level", "info"))
}

func (z *badgerLogger) Debugf(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), esl.String("level", "debug"))
}
