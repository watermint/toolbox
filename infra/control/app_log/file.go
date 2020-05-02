package app_log

import (
	"github.com/watermint/toolbox/essentials/log/es_rotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogContext struct {
	Logger *zap.Logger
	File   es_rotate.Writer
}

func (z *FileLogContext) Close() {
	if z.File != nil {
		z.File.Close()
		z.File = nil
	}
}

func NewFileLogger(path string, debug bool, test bool) (flc *FileLogContext, err error) {
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
	es_rotate.Startup()
	rw := es_rotate.NewWriter(path, "toolbox")
	if err := rw.Open(); err != nil {
		return nil, err
	}
	zo = zapcore.AddSync(rw)

	fileLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zo,
		zap.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(
			fileLoggerCore,
			NewConsoleLoggerCore(debug, test),
		),
	).WithOptions(zap.AddCaller())

	flc = &FileLogContext{
		Logger: logger,
		File:   rw,
	}
	return flc, nil
}
