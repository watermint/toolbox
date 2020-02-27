package app_root

import (
	"github.com/watermint/toolbox/infra/control/app_log"
	"go.uber.org/zap"
	"log"
	"sync"
)

type Hook func()

var (
	rootLogger                           = app_log.NewConsoleLogger(false, false)
	logWrapper                           = app_log.NewLogWrapper(rootLogger)
	captureLogger            *zap.Logger = nil
	ready                                = false
	successShutdownHook                  = make([]Hook, 0)
	successShutdownHookMutex             = sync.Mutex{}
)

func Ready() bool {
	return ready
}

func InitLogger() {
	rootLogger = app_log.NewConsoleLogger(false, false)
	logWrapper = app_log.NewLogWrapper(rootLogger)
}

func SetLogger(logger *zap.Logger) {
	rootLogger = logger
	logWrapper = app_log.NewLogWrapper(logger)
	log.SetOutput(logWrapper)
	ready = true
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

func AddSuccessShutdownHook(hook Hook) {
	successShutdownHookMutex.Lock()
	defer successShutdownHookMutex.Unlock()
	successShutdownHook = append(successShutdownHook, hook)
}

func FlushSuccessShutdownHook() {
	successShutdownHookMutex.Lock()
	defer successShutdownHookMutex.Unlock()

	for _, h := range successShutdownHook {
		h()
	}
	successShutdownHook = make([]Hook, 0)
}
