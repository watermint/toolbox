package app_log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

func NewConsoleLogger() *zap.Logger {
	return zap.New(NewConsoleLoggerCore())
}

func NewConsoleLoggerCore() zapcore.Core {
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

	level := zap.InfoLevel

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		level,
	)
}
