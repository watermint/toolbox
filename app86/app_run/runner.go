package app_run

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_control_impl"
	"github.com/watermint/toolbox/app86/app_diag"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_network"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/app_vo"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

type CommonOpts struct {
	Workspace string
	Debug     bool
	Proxy     string
	Quiet     bool
}

func (z *CommonOpts) SetFlags(f *flag.FlagSet, mc app_msg_container.Container) {
	f.StringVar(&z.Workspace, "workspace", "", mc.Compile(app_msg.M("run.common.flag.workspace")))
	f.BoolVar(&z.Debug, "debug", false, mc.Compile(app_msg.M("run.common.flag.debug")))
	f.StringVar(&z.Workspace, "proxy", "", mc.Compile(app_msg.M("run.common.flag.proxy")))
	f.BoolVar(&z.Quiet, "quiet", false, mc.Compile(app_msg.M("run.common.flag.quiet")))
}

func Run(args []string, bx *rice.Box) {
	// Initialize resources
	mc := NewContainer(bx)
	ui := app_ui.NewConsole(mc)
	cat := Catalogue()

	// Select recipe or group
	cmd, grp, rcp, rem, err := cat.Select(args)

	switch {
	case err != nil:
		if grp != nil {
			grp.PrintUsage(ui)
		} else {
			cat.PrintUsage(ui)
		}
		os.Exit(app_control.FailureInvalidCommand)

	case rcp == nil:
		grp.PrintUsage(ui)
		os.Exit(app_control.Success)
	}

	// Initialize recipe value object
	cmdPath := make([]string, 0)
	cmdPath = append(cmdPath, grp.Path...)
	if cmd != "" {
		cmdPath = append(cmdPath, cmd)
	}

	vo := rcp.Requirement()
	f := flag.NewFlagSet(strings.Join(cmdPath, " "), flag.ContinueOnError)
	com := &CommonOpts{}
	com.SetFlags(f, mc)

	vc := app_vo.NewValueContainer(vo)
	vc.MakeFlagSet(f)

	err = f.Parse(rem)
	if err != nil {
		os.Exit(app_control.FailureInvalidCommandFlags)
	}
	vc.Apply(vo)

	// Apply common flags
	// - Quiet
	if com.Quiet {
		ui = app_ui.NewQuiet()
	}

	// Startup
	so := make([]app_control.StartupOpt, 0)
	if com.Workspace != "" {
		so = append(so, app_control.Workspace(com.Workspace))
	}
	if com.Debug {
		so = append(so, app_control.Debug())
	}

	ctl := app_control_impl.NewControl(ui, bx)
	err = ctl.Startup(so...)
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Shutdown()

	// - Quiet
	if qui, ok := ui.(*app_ui.Quiet); ok {
		qui.SetLogger(ctl.Log())
	}

	// Recover
	defer func() {
		err := recover()
		if err != nil {
			ctl.Log().Debug("Recovery from panic")
			for depth := 0; ; depth++ {
				_, file, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				ctl.Log().Debug("Trace",
					zap.Int("Depth", depth),
					zap.String("File", file),
					zap.Int("Line", line),
				)
			}
			ctl.UI().Error("run.error.panic", app_msg.P("Reason", err))
			ctl.UI().Error("run.error.panic.instruction", app_msg.P("JobPath", ctl.Workspace().Job()))
			ctl.Fatal(app_control.Reason(app_control.FatalPanic))
		}
	}()

	// Trap signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)
	go func() {
		for s := range sig {
			// fatal shutdown
			ctl.Log().Debug("Signal", zap.Any("signal", s))
			pc := make([]uintptr, 16)
			n := runtime.Callers(0, pc)
			pc = pc[:n]
			frames := runtime.CallersFrames(pc)
			for f := 0; ; f++ {
				frame, more := frames.Next()
				ctl.Log().Debug("Frame",
					zap.Int("Frame", f),
					zap.String("File", frame.File),
					zap.Int("Line", frame.Line),
					zap.String("Function", frame.Function),
				)
				if !more {
					break
				}
			}
			ctl.UI().Error("run.error.interrupted")
			ctl.UI().Error("run.error.interrupted.instruction", app_msg.P("JobPath", ctl.Workspace().Job()))
			ctl.Fatal(app_control.Reason(app_control.FatalInterrupted))

			// in case the controller didn't fire exit..
			os.Exit(app_control.FatalInterrupted)
		}
	}()

	// - Proxy config
	app_network.SetHttpProxy(com.Proxy, ctl)

	// App Header
	AppHeader(ui)

	// Diagnosis
	err = app_diag.Runtime(ctl)
	if err != nil {
		ctl.Fatal(app_control.Reason(app_control.FatalRuntime))
	}
	err = app_diag.Network(ctl)
	if err != nil {
		ctl.Fatal(app_control.Reason(app_control.FatalNetwork))
	}

	// Run
	ctl.Log().Debug("Run recipe", zap.Any("vo", vo))
	k := app_recipe.NewKitchen(ctl, vo)
	err = rcp.Exec(k)
	if err != nil {
		os.Exit(app_control.FailureGeneral)
	}
}
