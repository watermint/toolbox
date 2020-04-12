package app_log

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/util/ut_io"
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
	if debug || (test && !app.IsProduction()) {
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
