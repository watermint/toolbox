package app_root

import (
	"github.com/watermint/toolbox/infra/control/app_log"
	"go.uber.org/zap"
	"log"
)

var (
	rootLogger                = app_log.NewConsoleLogger(false)
	logWrapper                = app_log.NewLogWrapper(rootLogger)
	captureLogger *zap.Logger = nil
	ready                     = false
)

func Ready() bool {
	return ready
}

func SetLogger(logger *zap.Logger) {
	rootLogger = logger
	logWrapper = app_log.NewLogWrapper(logger)
	log.SetOutput(logWrapper)
	ready = true
}

func SetCompatibleLogger(logger *zap.Logger) {
	rootLogger = logger
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
