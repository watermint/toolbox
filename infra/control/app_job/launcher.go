package app_job

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

type Launcher interface {
	Up() (ctl app_control.Control, err error)

	// Shutdown job. Please specify err for result of the execution.
	Down(err error, ctl app_control.Control)
}
