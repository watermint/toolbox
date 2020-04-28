package app_log

import (
	"github.com/watermint/toolbox/essentials/io/ut_io"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/infra/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewConsoleLogger(debug bool, test bool) *zap.Logger {
	return zap.New(NewConsoleLoggerCore(debug, test))
}

func NewConsoleLoggerCore(debug bool, test bool) zapcore.Core {
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
	w := ut_io.NewDefaultOut(test)
	zo := zapcore.AddSync(w)

	var level zapcore.Level
	switch {
	case es_env.IsEnabled(app.EnvNameDebugVerbose):
		level = zap.DebugLevel

	case app.IsProduction() && test:
		level = zap.FatalLevel

	case debug, test:
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
