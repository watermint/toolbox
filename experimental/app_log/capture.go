package app_log

import (
	"compress/gzip"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type CaptureContext struct {
	Logger *zap.Logger
	File   *os.File
	Writer *gzip.Writer
}

func (z *CaptureContext) Close() {
	if z.Writer != nil {
		z.Writer.Flush()
		z.Writer.Close()
	}
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
		),
	).WithOptions(zap.AddCaller())

	cc = &CaptureContext{
		Logger: logger,
		File:   f,
		Writer: g,
	}
	return cc, nil
}
