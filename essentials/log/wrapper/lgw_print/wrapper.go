package lgw_print

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type PrintLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

func New(l esl.Logger) PrintLogger {
	return &wrapperLogger{
		l: l,
	}
}

type wrapperLogger struct {
	l esl.Logger
}

func (z wrapperLogger) Print(v ...interface{}) {
	z.l.Debug(fmt.Sprint(v...))
}

func (z wrapperLogger) Printf(format string, v ...interface{}) {
	z.l.Debug(fmt.Sprintf(format, v...))
}
