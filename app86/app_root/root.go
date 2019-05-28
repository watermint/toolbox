package app_root

import (
	"github.com/watermint/toolbox/app86/app_log"
	"go.uber.org/zap"
	"log"
)

var (
	rootLogger = app_log.NewConsoleLogger()
	logWrapper = app_log.NewLogWrapper(rootLogger)
)

func SetLogger(logger *zap.Logger) {
	rootLogger = logger
	logWrapper = app_log.NewLogWrapper(logger)
	log.SetOutput(logWrapper)
}

func Flush() {
	logWrapper.Flush()
	rootLogger.Sync()
}

func Log() *zap.Logger {
	return rootLogger
}
