package app_shutdown

import "sync"

type Hook func()

var (
	successShutdownHook      = make([]Hook, 0)
	successShutdownHookMutex = sync.Mutex{}
	shutdownHook             = make([]Hook, 0)
	shutdownHookMutex        = sync.Mutex{}
)

func AddShutdownHook(hook Hook) {
	shutdownHookMutex.Lock()
	shutdownHook = append(shutdownHook, hook)
	shutdownHookMutex.Unlock()
}

func FlushShutdownHook() {
	shutdownHookMutex.Lock()
	for _, h := range shutdownHook {
		h()
	}
	shutdownHookMutex.Unlock()
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
