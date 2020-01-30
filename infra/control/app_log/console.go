package app_log

import (
	"github.com/watermint/toolbox/infra/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewConsoleLogger(debug bool) *zap.Logger {
	return zap.New(NewConsoleLoggerCore(debug))
}

func NewConsoleLoggerCore(debug bool) zapcore.Core {
	en := zapcore.EncoderConfig{
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	if app.IsWindows() {
		en.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		en.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	zo := zapcore.AddSync(os.Stdout)

	var level zapcore.Level
	if debug {
		level = zap.DebugLevel
	} else {
		level = zap.InfoLevel
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		level,
	)
}
