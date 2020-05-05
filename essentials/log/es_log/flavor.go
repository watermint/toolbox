package es_log

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
	"github.com/watermint/toolbox/infra/app"
	zapcoreuber "go.uber.org/zap/zapcore"
)

const (
	FlavorConsole      = iota
	FlavorFileStandard // for debug log
	FlavorFileCompact  // for capture log
)

type Flavor int

func newFlavor(f Flavor) zapcoreuber.Encoder {
	switch f {
	case FlavorConsole:
		return zapcoreuber.NewConsoleEncoder(zapcoreuber.EncoderConfig{
			LevelKey:       "level",
			MessageKey:     "msg",
			EncodeDuration: zapcoreuber.StringDurationEncoder,
			EncodeLevel:    zapTerminalEncodeLevel(),
		})
	case FlavorFileStandard:
		return zapcoreuber.NewJSONEncoder(zapcoreuber.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			EncodeLevel:    zapcoreuber.CapitalLevelEncoder,
			EncodeTime:     zapcoreuber.ISO8601TimeEncoder,
			EncodeDuration: zapcoreuber.StringDurationEncoder,
			EncodeCaller:   zapcoreuber.ShortCallerEncoder,
		})

	case FlavorFileCompact:
		return zapcoreuber.NewJSONEncoder(zapcoreuber.EncoderConfig{
			TimeKey:        "time",
			NameKey:        "name",
			EncodeTime:     zapcoreuber.ISO8601TimeEncoder,
			EncodeDuration: zapcoreuber.StringDurationEncoder,
			EncodeCaller:   zapcoreuber.ShortCallerEncoder,
		})

	default:
		panic(fmt.Sprintf("undefined flavor %d", f))
	}
}

func zapTerminalEncodeLevel() zapcoreuber.LevelEncoder {
	if es_terminfo.IsOutColorTerminal() {
		return zapcoreuber.CapitalLevelEncoder
	} else {
		return zapcoreuber.CapitalColorLevelEncoder
	}
}

func terminalDefaultLevel() Level {
	switch {
	case es_env.IsEnabled(app.EnvNameDebugVerbose):
		return LevelDebug

	default:
		return LevelInfo
	}
}
