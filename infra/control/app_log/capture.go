package app_log

import (
	"github.com/watermint/toolbox/essentials/log/es_rotate"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CaptureContext struct {
	Logger *zap.Logger
	File   es_rotate.Writer
}

func (z *CaptureContext) Close() {
	if z.File != nil {
		z.File.Close()
	}
}

func NewCaptureLogger(path string) (cc *CaptureContext, err error) {
	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		NameKey:        "name",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var zo zapcore.WriteSyncer
	app_shutdown.AddShutdownHook(es_rotate.Shutdown)
	rw := es_rotate.NewWriter(path, "capture")
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
		),
	).WithOptions(zap.AddCaller())

	cc = &CaptureContext{
		Logger: logger,
		File:   rw,
	}
	return cc, nil
}
