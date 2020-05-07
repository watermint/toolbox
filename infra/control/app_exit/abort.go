package app_exit

import (
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"os"
)

const (
	Success = iota
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
)

type AbortCode int

func Abort(code AbortCode) {
	app_shutdown.FlushShutdownHook()
	os.Exit(int(code))
}

func ExitSuccess() {
	app_shutdown.FlushSuccessShutdownHook()
	app_shutdown.FlushShutdownHook()
	os.Exit(Success)
}
