package es_log

import (
	"github.com/watermint/toolbox/essentials/io/es_close"
	"go.uber.org/zap"
	"io"
)

func NewLogCloser(level Level, flavor Flavor, w io.WriteCloser) LogCloser {
	wc := es_close.New(w)
	wl := &zapWrapper{
		zl: zapWithName(newZap(level, flavor, wc).WithOptions(zap.AddCallerSkip(1))),
	}
	return &logCloserImpl{
		zl: wl,
		w:  wc,
	}
}

type logCloserImpl struct {
	w  io.WriteCloser
	zl Logger
}

func (z logCloserImpl) With(fields ...Field) Logger {
	return z.zl.With(fields...)
}

func (z logCloserImpl) AddCallerSkip(n int) Logger {
	return z.zl.AddCallerSkip(n)
}

func (z logCloserImpl) Debug(msg string, fields ...Field) {
	z.zl.Debug(msg, fields...)
}

func (z logCloserImpl) Info(msg string, fields ...Field) {
	z.zl.Info(msg, fields...)
}

func (z logCloserImpl) Warn(msg string, fields ...Field) {
	z.zl.Warn(msg, fields...)
}

func (z logCloserImpl) Error(msg string, fields ...Field) {
	z.zl.Error(msg, fields...)
}

func (z logCloserImpl) Sync() error {
	return z.zl.Sync()
}

func (z logCloserImpl) Close() error {
	return z.w.Close()
}
