package esl

import (
	"fmt"
	zapuber "go.uber.org/zap"
	zapcoreuber "go.uber.org/zap/zapcore"
)

type zapWrapper struct {
	zl *zapuber.Logger
}

func (z zapWrapper) AddCallerSkip(n int) Logger {
	return &zapWrapper{
		zl: zapWithName(z.zl.WithOptions(zapuber.AddCallerSkip(n))),
	}
}

func (z zapWrapper) Sync() error {
	return z.zl.Sync()
}

func (z zapWrapper) zapFields(fields []Field) []zapuber.Field {
	zfs := make([]zapuber.Field, len(fields))
	for i, f := range fields {
		zfs[i] = f.zap()
	}
	return zfs
}

func (z zapWrapper) With(fields ...Field) Logger {
	return &zapWrapper{
		zl: zapWithName(z.zl.With(z.zapFields(fields)...)),
	}
}

func (z zapWrapper) Debug(msg string, fields ...Field) {
	z.zl.Debug(msg, z.zapFields(fields)...)
}

func (z zapWrapper) Info(msg string, fields ...Field) {
	z.zl.Info(msg, z.zapFields(fields)...)
}

func (z zapWrapper) Warn(msg string, fields ...Field) {
	z.zl.Warn(msg, z.zapFields(fields)...)
}

func (z zapWrapper) Error(msg string, fields ...Field) {
	z.zl.Error(msg, z.zapFields(fields)...)
}

func zapLevel(level Level) zapcoreuber.Level {
	switch level {
	case LevelDebug:
		return zapuber.DebugLevel
	case LevelInfo:
		return zapuber.InfoLevel
	case LevelWarn:
		return zapuber.WarnLevel
	case LevelError:
		return zapuber.ErrorLevel
	default:
		panic(fmt.Sprintf("Undefined level: %d", level))
	}
}
