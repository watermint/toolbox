package app_root

import (
	"github.com/watermint/toolbox/app86/app_log"
	"go.uber.org/zap"
)

var (
	Logger = app_log.NewConsoleLogger()
)

func Log() *zap.Logger {
	return Logger
}
