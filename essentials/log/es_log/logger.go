package es_log

import (
	zapuber "go.uber.org/zap"
	"io"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelQuiet
)

type Level int

type Logger interface {
	With(fields ...Field) Logger
	AddCallerSkip(n int) Logger

	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)

	Sync() error
}

type LogCloser interface {
	Logger
	io.Closer
}

type Field interface {
	zap() zapuber.Field
}
