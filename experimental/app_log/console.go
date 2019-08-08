package app_log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
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
	if runtime.GOOS == "windows" {
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
