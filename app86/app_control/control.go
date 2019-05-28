package app_control

import (
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/app_workspace"
	"go.uber.org/zap"
)

type Control interface {
	Startup(opts ...StartupOpt) error
	Shutdown()
	Fatal(opts ...FatalOpt)
	UI() app_ui.UI
	Log() *zap.Logger
	Resource(key string) (bin []byte, err error)
	Workspace() app_workspace.Workspace
	IsTest() bool
}

type StartupOpt func(opt *StartupOpts) *StartupOpts
type StartupOpts struct {
	WorkspacePath string
	Debug         bool
	Test          bool
}

func Debug() StartupOpt {
	return func(opt *StartupOpts) *StartupOpts {
		opt.Debug = true
		return opt
	}
}
func Test() StartupOpt {
	return func(opt *StartupOpts) *StartupOpts {
		opt.Test = true
		return opt
	}
}
func Workspace(path string) StartupOpt {
	return func(opt *StartupOpts) *StartupOpts {
		opt.WorkspacePath = path
		return opt
	}
}

type FatalOpt func(opt *FatalOpts) *FatalOpts
type FatalOpts struct {
	Reason *int
}

func Reason(reason int) FatalOpt {
	return func(opt *FatalOpts) *FatalOpts {
		opt.Reason = &reason
		return opt
	}
}

const (
	Success = iota
	FatalGeneral
	FatalStartup
	FatalPanic
	FatalInterrupted
	FatalRuntime
	FatalNetwork

	// Failures
	FailureGeneral = iota + 1000
	FailureInvalidCommand
	FailureInvalidCommandFlags
)
