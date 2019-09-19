package app_run

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_run_impl"
	"github.com/watermint/toolbox/infra/network/app_diag"
	"github.com/watermint/toolbox/infra/network/app_network"
	"github.com/watermint/toolbox/infra/quality/qt_control_impl"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/infra/recpie/app_vo_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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
	Secure    bool
}

func (z *CommonOpts) SetFlags(f *flag.FlagSet, mc app_msg_container.Container) {
	f.StringVar(&z.Workspace, "workspace", "", mc.Compile(app_msg.M("run.common.flag.workspace")))
	f.BoolVar(&z.Debug, "debug", false, mc.Compile(app_msg.M("run.common.flag.debug")))
	f.StringVar(&z.Workspace, "proxy", "", mc.Compile(app_msg.M("run.common.flag.proxy")))
	f.BoolVar(&z.Quiet, "quiet", false, mc.Compile(app_msg.M("run.common.flag.quiet")))
	f.BoolVar(&z.Secure, "secure", false, mc.Compile(app_msg.M("run.common.flag.secure")))
}

func Run(args []string, bx, web *rice.Box) (found bool) {
	// Initialize resources
	mc := app_run_impl.NewContainer(bx)
	ui := app_ui.NewConsole(mc, qt_control_impl.NewMessageMock(), false)
	cat := Catalogue()

	// Select recipe or group
	cmd, grp, rcp, rem, err := cat.Select(args)

	switch {
	case err != nil:
		//if grp != nil {
		//	grp.PrintUsage(ui)
		//} else {
		//	cat.PrintUsage(ui)
		//}
		//os.Exit(app_control.FailureInvalidCommand)
		return false

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
	recipeName := strings.Join(cmdPath, " ")

	vo := rcp.Requirement()
	f := flag.NewFlagSet(recipeName, flag.ContinueOnError)
	com := &CommonOpts{}
	com.SetFlags(f, mc)

	vc := app_vo_impl.NewValueContainer(vo)
	vc.MakeFlagSet(f)

	err = f.Parse(rem)
	if err != nil {
		os.Exit(app_control.FailureInvalidCommandFlags)
	}
	vc.Apply(vo)

	// Apply common flags
	// - Quiet
	if com.Quiet {
		ui = app_ui.NewQuiet(mc)
	}

	// Up
	so := make([]app_control.UpOpt, 0)
	if com.Workspace != "" {
		so = append(so, app_control.WorkspacePath(com.Workspace))
	}
	if com.Debug {
		so = append(so, app_control.Debug())
	}
	if com.Secure {
		so = append(so, app_control.Secure())
	}
	so = append(so, app_control.RecipeName(recipeName))

	ctl := app_control_impl.NewSingle(ui, bx, web, mc, com.Quiet, Recipes())
	err = ctl.Up(so...)
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Down()

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
			ui := ctl.UI(nil)
			ui.Error("run.error.panic", app_msg.P{"Reason": err})
			ui.Error("run.error.panic.instruction", app_msg.P{"JobPath": ctl.Workspace().Job()})
			ctl.Abort(app_control.Reason(app_control.FatalPanic))
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
			ui := ctl.UI(nil)
			ui.Error("run.error.interrupted")
			ui.Error("run.error.interrupted.instruction", app_msg.P{"JobPath": ctl.Workspace().Job()})
			ctl.Abort(app_control.Reason(app_control.FatalInterrupted))

			// in case the controller didn't fire exit..
			os.Exit(app_control.FatalInterrupted)
		}
	}()

	// - Proxy config
	app_network.SetHttpProxy(com.Proxy, ctl)

	// App Header
	app_recipe_group.AppHeader(ui)

	// Diagnosis
	err = app_diag.Runtime(ctl)
	if err != nil {
		ctl.Abort(app_control.Reason(app_control.FatalRuntime))
	}
	if ctl.IsProduction() {
		err = app_diag.Network(ctl)
		if err != nil {
			ctl.Abort(app_control.Reason(app_control.FatalNetwork))
		}
	}

	// Run
	ctl.Log().Debug("Run recipe", zap.Any("vo", vo))
	k := app_kitchen.NewKitchen(ctl, vo)
	err = rcp.Exec(k)
	if err != nil {
		os.Exit(app_control.FailureGeneral)
	}
	return true
}
