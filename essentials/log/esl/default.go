package esl

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
	currentDefault Tee = newDefault()
	currentConsole     = newConsole(ConsoleDefaultLevel())
	capture            = newEmpty()
	loggerName     atomic.Int64
)

func Default() Logger {
	return currentDefault
}

func ConsoleOnly() Logger {
	return currentConsole
}

func Capture() Logger {
	return capture
}

func newEmpty() Logger {
	return New(ConsoleDefaultLevel(), FlavorConsole, ioutil.Discard)
}

func newDefault() Tee {
	t := NewTee()
	t.AddSubscriber(currentConsole)
	return t
}

func AddDefaultSubscriber(l Logger) {
	currentDefault.AddSubscriber(l)
}

func newConsole(level Level) Logger {
	return New(level, FlavorConsole, es_stdout.NewDefaultOut(false))
}

func New(level Level, flavor Flavor, w io.Writer) Logger {
	switch level {
	case LevelQuiet:
		return newEmpty()

	default:
		return &zapWrapper{zl: newZap(level, flavor, w)}
	}
}

func zapWithName(logger *zapuber.Logger) *zapuber.Logger {
	return logger.Named("z" + strconv.FormatInt(loggerName.Add(1), 10))
}

func zapWithFlavor(flavor Flavor, logger *zapuber.Logger) *zapuber.Logger {
	switch flavor {
	case FlavorFileStandard:
		return logger.WithOptions(zapuber.AddCaller())
	default:
		return logger
	}
}

func newZap(level Level, flavor Flavor, w io.Writer) *zapuber.Logger {
	core := zapcoreuber.NewCore(
		newFlavor(flavor),
		zapcoreuber.AddSync(w),
		zapLevel(level),
	)
	zl := zapuber.New(core, zapuber.AddCallerSkip(1))
	return zapWithName(zapWithFlavor(flavor, zl))
}
