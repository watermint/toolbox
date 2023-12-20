package app_feature

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
)

func ConsoleLogLevel(test, debug bool) esl.Level {
	switch {
	case es_env.IsEnabled(app_definitions2.EnvNameDebugVerbose):
		return esl.LevelDebug

	case es_env.IsEnabled(app_definitions2.EnvNameTestQuiet), app_definitions2.IsProduction() && test:
		return esl.LevelQuiet

	case test:
		return esl.LevelInfo

	case debug:
		return esl.LevelDebug

	default:
		return esl.LevelInfo
	}
}
