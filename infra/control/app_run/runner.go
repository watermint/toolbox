package app_run

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	"github.com/watermint/toolbox/catalogue"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_concurrency"
	"github.com/watermint/toolbox/infra/network/nw_diag"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"github.com/watermint/toolbox/infra/network/nw_proxy"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"github.com/watermint/toolbox/infra/util/ut_memory"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg_impl"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

type MsgRun struct {
	ErrorInvalidArgument app_msg.Message
}

var (
	MRun = app_msg.Apply(&MsgRun{}).(*MsgRun)
)

func runRecipe(mc app_msg_container.Container, ui app_ui.UI, rcpSpec rc_recipe.Spec, grp rc_group.Group, rem []string, bx, web *rice.Box) (found bool) {
	comSpec := rc_spec.NewCommonValue()

	f := flag.NewFlagSet(rcpSpec.CliPath(), flag.ContinueOnError)

	comSpec.SetFlags(f, ui)
	rcpSpec.SetFlags(f, ui)

	err := f.Parse(rem)
	rem2 := f.Args()
	if err != nil || (len(rem2) > 0 && rem2[0] == "help") {
		grp.PrintRecipeUsage(ui, rcpSpec, f)
		os.Exit(app_control.FailureInvalidCommandFlags)
	}
	comSpec.Apply()
	com := comSpec.Opts()

	// Apply common flags

	// - Quiet
	if com.Quiet {
		ui = app_ui.NewQuiet(mc)
	}

	// Up
	so := make([]app_control.UpOpt, 0)
	if com.Workspace != "" {
		wsPath, err := ut_filepath.FormatPathWithPredefinedVariables(com.Workspace)
		if err != nil {
			ui.ErrorK("run.error.unable_to_format_path", app_msg.P{
				"Error": err.Error(),
			})
			os.Exit(app_control.FailureInvalidCommandFlags)
		}
		so = append(so, app_control.WorkspacePath(wsPath))
	}
	if com.Debug {
		so = append(so, app_control.Debug())
	}
	if com.Secure {
		so = append(so, app_control.Secure())
	}
	so = append(so, app_control.LowMemory(com.LowMemory))
	so = append(so, app_control.Concurrency(com.Concurrency))
	so = append(so, app_control.RecipeName(rcpSpec.CliPath()))
	so = append(so, app_control.CommonOptions(comSpec.Debug()))
	so = append(so, app_control.RecipeOptions(rcpSpec.Debug()))

	ctl := app_control_impl.NewSingle(ui, bx, web, mc, com.Quiet, catalogue.NewCatalogue())
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
	if ctl.IsProduction() {
		defer func() {
			err := recover()
			if err != nil {
				l := ctl.Log()
				l.Debug("Recovery from panic")
				l.Error(ctl.UI().TextK("run.error.panic"),
					zap.Any("error", err),
				)
				l.Error(ctl.UI().TextK("run.error.panic.instruction"),
					zap.String("JobPath", ctl.Workspace().Job()),
				)

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
				ctl.Abort(app_control.Reason(app_control.FatalPanic))
			}
		}()
	}

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
			ui := ctl.UI()
			ui.ErrorK("run.error.interrupted")
			ui.ErrorK("run.error.interrupted.instruction", app_msg.P{"JobPath": ctl.Workspace().Job()})
			ctl.Abort(app_control.Reason(app_control.FatalInterrupted))

			// in case the controller didn't fire exit..
			os.Exit(app_control.FatalInterrupted)
		}
	}()

	// - Proxy config
	nw_proxy.SetHttpProxy(com.Proxy, ctl)

	// App Header
	rc_group.AppHeader(ui, app.Version)

	// Diagnosis
	err = nw_diag.Runtime(ctl)
	if err != nil {
		ctl.Abort(app_control.Reason(app_control.FatalRuntime))
	}
	if ctl.IsProduction() {
		err = nw_diag.Network(ctl)
		if err != nil {
			ctl.Abort(app_control.Reason(app_control.FatalNetwork))
		}
	}

	// Launch monitor
	nw_monitor.LaunchReporting(ctl.Log())
	ut_memory.LaunchReporting(ctl.Log())

	// Set bandwidth
	nw_bandwidth.SetBandwidth(com.BandwidthKb)
	nw_concurrency.SetConcurrency(com.Concurrency)

	// Apply profiler
	if com.Debug {
		defer profile.Start(
			profile.ProfilePath(ctl.Workspace().Log()),
			profile.MemProfile,
		).Stop()
	}

	// Run
	var lastErr error
	ctl.Log().Debug("Run recipe", zap.Any("vo", rcpSpec.Debug()), zap.Any("common", com))
	lastErr = rc_exec.ExecSpec(ctl, rcpSpec, rc_recipe.NoCustomValues)
	if lastErr != nil {
		ctl.Log().Error("Recipe failed with an error", zap.Error(lastErr))
		ui.Failure("run.error.recipe.failed", app_msg.P{"Error": lastErr.Error()})
		os.Exit(app_control.FailureGeneral)
	}

	// Dump stats
	ut_memory.DumpStats(ctl.Log())
	nw_monitor.DumpStats(ctl.Log())

	app_root.FlushSuccessShutdownHook()

	return true
}

func Run(args []string, bx, web *rice.Box) (found bool) {
	// Initialize resources
	mc := app_msg_container_impl.NewContainer(bx)
	ui := app_ui.NewConsole(mc, qt_missingmsg_impl.NewMessageMemory(), false)
	cat := catalogue.NewCatalogue()
	rg := cat.RootGroup()

	// Select recipe or group
	_, grp, rcp, rem, err := rg.Select(args)

	switch {
	case err != nil:
		ui.Error(MRun.ErrorInvalidArgument.With("Args", strings.Join(args, " ")))
		if grp != nil {
			grp.PrintGroupUsage(ui, os.Args[0], app.Version)
		} else {
			rg.PrintGroupUsage(ui, os.Args[0], app.Version)
		}
		os.Exit(app_control.FailureInvalidCommand)

	case rcp == nil:
		grp.PrintGroupUsage(ui, os.Args[0], app.Version)
		os.Exit(app_control.Success)
	}

	return runRecipe(mc, ui, rcp, grp, rem, bx, web)
}
