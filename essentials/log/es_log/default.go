package es_log

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"go.uber.org/atomic"
	zapuber "go.uber.org/zap"
	zapcoreuber "go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"strconv"
)

var (
	fallback   = newDefault()
	capture    = newEmpty()
	loggerName atomic.Int64
)

func Default() Logger {
	return fallback
}

func DefaultConsole() Logger {
	return fallback
}

func Capture() Logger {
	return capture
}

func newEmpty() Logger {
	return New(terminalDefaultLevel(), FlavorConsole, ioutil.Discard)
}

func newDefault() Logger {
	return newTerminal(terminalDefaultLevel())
}

func newTerminal(level Level) Logger {
	return New(level, FlavorConsole, es_stdout.NewDefaultOut(false))
}

func New(level Level, flavor Flavor, w io.Writer) Logger {
	return &zapWrapper{
		zl: newZap(level, flavor, w),
	}
}

func zapWithName(logger *zapuber.Logger) *zapuber.Logger {
	return logger.Named("z" + strconv.FormatInt(loggerName.Add(1), 10))
}

func zapWithFlavor(flavor Flavor, logger *zapuber.Logger) *zapuber.Logger {
	switch flavor {
	case FlavorFileStandard:
		return logger.WithOptions(zapuber.AddCaller(), zapuber.AddCallerSkip(1))
	default:
		return logger.WithOptions(zapuber.AddCallerSkip(1))
	}
}

func newZap(level Level, flavor Flavor, w io.Writer) *zapuber.Logger {
	core := zapcoreuber.NewCore(
		newFlavor(flavor),
		zapcoreuber.AddSync(w),
		zapLevel(level),
	)
	return zapWithName(zapWithFlavor(flavor, zapuber.New(core)))
}
