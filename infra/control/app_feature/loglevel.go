package app_feature

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/infra/app"
)

func ConsoleLogLevel(test, debug bool) es_log.Level {
	switch {
	case es_env.IsEnabled(app.EnvNameDebugVerbose):
		return es_log.LevelDebug

	case app.IsProduction() && test:
		return es_log.LevelQuiet

	case debug, test:
		return es_log.LevelDebug

	default:
		return es_log.LevelInfo
	}
}
