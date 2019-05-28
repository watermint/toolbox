package app_run

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_control_impl"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/app_vo"
	"go.uber.org/zap"
	"os"
	"runtime"
)

type CommonOpts struct {
	Workspace string
	Debug     bool
	Proxy     string
}

func (z *CommonOpts) SetFlags(f *flag.FlagSet, mc app_msg_container.Container) {
	f.StringVar(&z.Workspace, "workspace", "", mc.Compile(app_msg.M("run.common.flag.workspace")))
	f.BoolVar(&z.Debug, "debug", false, mc.Compile(app_msg.M("run.common.flag.debug")))
	f.StringVar(&z.Workspace, "proxy", "", mc.Compile(app_msg.M("run.common.flag.proxy")))
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

	case grp != nil:
		grp.PrintUsage(ui)
		os.Exit(app_control.Success)
	}

	// Initialize recipe value object
	if cmd == "" {
		cmd = "toolbox"
	}

	vo := rcp.Requirement()
	f := flag.NewFlagSet(cmd, flag.ContinueOnError)
	com := &CommonOpts{}
	com.SetFlags(f, mc)

	vc := app_vo.NewValueContainer(vo)
	vc.MakeFlagSet(f)

	err = f.Parse(rem)
	if err != nil {
		os.Exit(app_control.FailureInvalidCommandFlags)
	}
	vc.Apply(vo)

	// Startup
	so := make([]app_control.StartupOpt, 0)
	if com.Workspace != "" {
		so = append(so, app_control.Workspace(com.Workspace))
	}

	ctl := app_control_impl.NewControl(ui, bx)
	err = ctl.Startup(so...)
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Shutdown()

	// Recover
	defer func() {
		err := recover()
		if err != nil {
			for depth := 0; ; depth++ {
				_, file, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				ctl.Log().Debug("Trace", zap.Int("Depth", depth), zap.String("file", file), zap.Int("line", line))
			}
			ctl.UI().Error("run.error.panic", app_msg.P("Reason", err))
			ctl.UI().Error("run.error.panic.instruction", app_msg.P("JobPath", ctl.Workspace().Job()))
			ctl.Fatal(app_control.Reason(app_control.FatalPanic))
		}
	}()

	// Run
	ctl.Log().Debug("Run recipe", zap.Any("vo", vo))
	k := app_recipe.NewKitchen(ctl, vo)
	err = rcp.Exec(k)
	if err != nil {
		os.Exit(app_control.FailureGeneral)
	}
}
