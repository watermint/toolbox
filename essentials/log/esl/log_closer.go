package esl

import (
	"github.com/watermint/toolbox/essentials/io/es_close"
	"io"
)

func NewLogCloser(level Level, flavor Flavor, w io.WriteCloser) LogCloser {
	wc := es_close.New(w)
	wl := &zapWrapper{
		zl: newZap(level, flavor, wc),
	}
	return &logCloserImpl{
		l: wl,
		n: wl.AddCallerSkip(1),
		w: wc,
	}
}

type logCloserImpl struct {
	w io.WriteCloser
	l Logger
	n Logger
}

func (z logCloserImpl) With(fields ...Field) Logger {
	return z.l.With(fields...)
}

func (z logCloserImpl) AddCallerSkip(n int) Logger {
	return z.l.AddCallerSkip(n)
}

func (z logCloserImpl) Debug(msg string, fields ...Field) {
	z.n.Debug(msg, fields...)
}

func (z logCloserImpl) Info(msg string, fields ...Field) {
	z.n.Info(msg, fields...)
}

func (z logCloserImpl) Warn(msg string, fields ...Field) {
	z.n.Warn(msg, fields...)
}

func (z logCloserImpl) Error(msg string, fields ...Field) {
	z.n.Error(msg, fields...)
}

func (z logCloserImpl) Sync() error {
	return z.l.Sync()
}

func (z logCloserImpl) Close() error {
	return z.w.Close()
}
