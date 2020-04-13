package app_control

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_template"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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
	TestResource(key string) (data gjson.Result, found bool)
	Workspace() app_workspace.Workspace
	Messages() app_msg_container.Container
	Feature() app_feature.Feature

	NewQueue() rc_worker.Queue
}

type ControlTestExtension interface {
	TestValue(key string) interface{}
	SetTestValue(key string, v interface{})
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
	RecipeOptions map[string]interface{}
	CommonOptions map[string]interface{}
	Concurrency   int
	LowMemory     bool
	AutoOpen      bool
	Quiet         bool
	UIFormat      string
}

func (z *UpOpts) Clone() *UpOpts {
	return &UpOpts{
		WorkspacePath: z.WorkspacePath,
		Workspace:     z.Workspace,
		Debug:         z.Debug,
		Test:          z.Test,
		Secure:        z.Secure,
		CommonOptions: z.CommonOptions,
		Concurrency:   z.Concurrency,
		LowMemory:     z.LowMemory,
		AutoOpen:      z.AutoOpen,
		Quiet:         z.Quiet,
		UIFormat:      z.UIFormat,
	}
}

func Quiet(enabled bool) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Quiet = enabled
		return opt
	}
}
func UIFormat(format string) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.UIFormat = format
		return opt
	}
}
func AutoOpen(enabled bool) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.AutoOpen = enabled
		return opt
	}
}
func LowMemory(enabled bool) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.LowMemory = enabled
		return opt
	}
}
func Concurrency(c int) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.Concurrency = c
		return opt
	}
}
func RecipeName(name string) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.RecipeName = name
		return opt
	}
}
func RecipeOptions(vo map[string]interface{}) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.RecipeOptions = vo
		return opt
	}
}
func CommonOptions(vo map[string]interface{}) UpOpt {
	return func(opt *UpOpts) *UpOpts {
		opt.CommonOptions = vo
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
