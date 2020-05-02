package app_root

import (
	"github.com/watermint/toolbox/infra/control/app_log"
	"go.uber.org/zap"
	"log"
)

var (
	rootLogger                = app_log.NewConsoleLogger(false, false)
	logWrapper                = app_log.NewLogWrapper(rootLogger)
	captureLogger *zap.Logger = nil
)

func InitLogger() {
	rootLogger = app_log.NewConsoleLogger(false, false)
	logWrapper = app_log.NewLogWrapper(rootLogger)
}

func SetLogger(logger *zap.Logger) {
	rootLogger = logger
	logWrapper = app_log.NewLogWrapper(logger)
	log.SetOutput(logWrapper)
}

func SetCapture(logger *zap.Logger) {
	captureLogger = logger
}

func Flush() {
	logWrapper.Flush()
	rootLogger.Sync()
}

func Log() *zap.Logger {
	return rootLogger
}

func Capture() *zap.Logger {
	return captureLogger
}
