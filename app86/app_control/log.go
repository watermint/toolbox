package app_control

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

func newConsoleLogger() *zap.Logger {
	return zap.New(newConsoleLoggerCore())
}

func newConsoleLoggerCore() zapcore.Core {
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

	level := zap.DebugLevel

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		level,
	)
}
