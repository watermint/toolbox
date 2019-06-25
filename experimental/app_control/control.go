package app_control

import (
	"github.com/watermint/toolbox/experimental/app_template"
	"github.com/watermint/toolbox/experimental/app_ui"
	"github.com/watermint/toolbox/experimental/app_workspace"
	"go.uber.org/zap"
	"net/http"
)

type Control interface {
	Up(opts ...UpOpt) error
	Down()
	Abort(opts ...AbortOpt)

	UI() app_ui.UI
	Log() *zap.Logger
	Capture() *zap.Logger
	Resource(key string) (bin []byte, err error)
	Workspace() app_workspace.Workspace

	IsProduction() bool
	IsTest() bool
	IsQuiet() bool
	IsSecure() bool
}

type ControlHttpFileSystem interface {
	HttpFileSystem() http.FileSystem
	Template() app_template.Template
}

type UpOpt func(opt *UpOpts) *UpOpts
type UpOpts struct {
	WorkspacePath string
	Workspace     app_workspace.Workspace
	Debug         bool
	Test          bool
	Secure        bool
	RecipeName    string
}

func RecipeName(name string) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.RecipeName = name
		return opt
	}
}
func Secure() UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Secure = true
		return opt
	}
}
func Debug() UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Debug = true
		return opt
	}
}
func Test() UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Test = true
		return opt
	}
}
func WorkspacePath(path string) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.WorkspacePath = path
		return opt
	}
}
func Workspace(ws app_workspace.Workspace) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Workspace = ws
		return opt
	}
}

type AbortOpt func(opt *AbortOpts) *AbortOpts
type AbortOpts struct {
	Reason *int
}

func Reason(reason int) AbortOpt {
	return func(opt *AbortOpts) *AbortOpts {
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
	FatalResourceUnavailable

	// Failures
	FailureGeneral
	FailureInvalidCommand
	FailureInvalidCommandFlags
	FailureAuthenticationFailedOrCancelled
)
