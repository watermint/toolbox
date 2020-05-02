package es_fallback

import (
	"github.com/watermint/toolbox/essentials/io/ut_io"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/infra/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	fallback = zap.New(newFallback())
)

// Fallback logger
func Fallback() *zap.Logger {
	return fallback
}

func newFallback() zapcore.Core {
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
	w := ut_io.NewDefaultOut(false)
	zo := zapcore.AddSync(w)

	var level zapcore.Level
	switch {
	case es_env.IsEnabled(app.EnvNameDebugVerbose):
		level = zap.DebugLevel

	default:
		level = zap.InfoLevel
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		level,
	)
}
