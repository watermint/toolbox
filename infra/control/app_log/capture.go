package app_log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type CaptureContext struct {
	Logger *zap.Logger
	File   *os.File
}

func (z *CaptureContext) Close() {
	if z.File != nil {
		z.File.Close()
	}
}

func NewCaptureLogger(path string) (cc *CaptureContext, err error) {
	logPath := filepath.Join(path, "capture.log")
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		NameKey:        "name",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var zo zapcore.WriteSyncer
	f, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	zo = zapcore.AddSync(f)

	fileLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zo,
		zap.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(
			fileLoggerCore,
		),
	).WithOptions(zap.AddCaller())

	cc = &CaptureContext{
		Logger: logger,
		File:   f,
	}
	return cc, nil
}
