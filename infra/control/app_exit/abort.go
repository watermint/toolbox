package app_exit

import (
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"os"
)

const (
	Success AbortCode = iota
	FatalGeneral
	FatalStartup
	FatalPanic
	FatalInterrupted
	FatalRuntime
	FatalNetwork
	FatalResourceUnavailable

	// Failures
	FailureGeneral
	FailureInvalidCommand
	FailureInvalidCommandFlags
	FailureAuthenticationFailedOrCancelled
	FailureBinaryExpired
)

type AbortCode int

var (
	testMode = false
)

func Abort(code AbortCode) {
	app_shutdown.FlushShutdownHook()
	if testMode {
		panic(code)
	} else {
		os.Exit(int(code))
	}
}

func ExitSuccess() {
	app_shutdown.FlushSuccessShutdownHook()
	app_shutdown.FlushShutdownHook()
	if testMode {
		panic(Success)
	} else {
		os.Exit(int(Success))
	}
}

// Panic instead of os.Exit if it set.
func SetTestMode(enabled bool) {
	testMode = enabled
}
