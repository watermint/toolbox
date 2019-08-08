package app_log

import (
	"compress/gzip"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type FileLogContext struct {
	Logger *zap.Logger
	File   *os.File
	Writer *gzip.Writer
}

func (z *FileLogContext) Close() {
	if z.Writer != nil {
		z.Writer.Flush()
		z.Writer.Close()
	}
	if z.File != nil {
		z.File.Close()
	}
}

func NewFileLogger(path string, debug bool) (flc *FileLogContext, err error) {
	logPath := filepath.Join(path, "toolbox.log.gz")
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var zo zapcore.WriteSyncer
	f, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	g := gzip.NewWriter(f)
	zo = zapcore.AddSync(g)

	fileLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zo,
		zap.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(
			fileLoggerCore,
			NewConsoleLoggerCore(debug),
		),
	).WithOptions(zap.AddCaller())

	flc = &FileLogContext{
		Logger: logger,
		File:   f,
		Writer: g,
	}
	return flc, nil
}
