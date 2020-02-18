package app_run

import (
	"bytes"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	"github.com/watermint/toolbox/catalogue"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_concurrency"
	"github.com/watermint/toolbox/infra/network/nw_diag"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"github.com/watermint/toolbox/infra/network/nw_proxy"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
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
	ErrorInvalidArgument        app_msg.Message
	ErrorTooManyArguments       app_msg.Message
	ErrorInterrupted            app_msg.Message
	ErrorInterruptedInstruction app_msg.Message
	ErrorUnableToFormatPath     app_msg.Message
	ErrorPanic                  app_msg.Message
	ErrorPanicInstruction       app_msg.Message
	ErrorRecipeFailed           app_msg.Message
	ErrorUnsupportedOutput      app_msg.Message
}

var (
	MRun = app_msg.Apply(&MsgRun{}).(*MsgRun)
)

// Mutable object for running recipe
type Bootstrap interface {
	// Select UI for the option
	SelectUI(mc app_msg_container.Container, opt *app_opt.CommonOpts) (uiFormat string, ui app_ui.UI)

	// Parse arguments for recipe & common values
	Parse(args ...string) (rcp rc_recipe.Spec, com *rc_spec.CommonValues)

	// Parse arguments for common values only
	ParseCommon(args []string, ignoreErrors bool) (rem []string, com *rc_spec.CommonValues)

	// Run recipe
	Run(rcp rc_recipe.Spec, com *rc_spec.CommonValues)
}

func NewBootstrap(bx, web *rice.Box) Bootstrap {
	mc := app_msg_container_impl.NewContainer(bx)
	ui := app_ui.NewConsole(mc, qt_missingmsg_impl.NewMessageMemory(), false)
	cat := catalogue.NewCatalogue()
	rg := cat.RootGroup()

	return &bootstrapImpl{
		boxResource:  bx,
		boxWeb:       web,
		msgContainer: mc,
		currentUI:    ui,
		catalogue:    cat,
		rootGroup:    rg,
	}
}

type bootstrapImpl struct {
	boxResource *rice.Box
	boxWeb      *rice.Box

	// Mutable field. This field point to current UI only. UI may change on flags.
	currentUI app_ui.UI

	msgContainer app_msg_container.Container
	catalogue    rc_catalogue.Catalogue
	rootGroup    rc_group.Group
}

func (z *bootstrapImpl) SelectUI(mc app_msg_container.Container, opt *app_opt.CommonOpts) (uiFormat string, ui app_ui.UI) {
	output := strings.ToLower(opt.Output)

	// Select UI
	switch {
	case opt.Quiet, output == app_opt.OutputNone:
		return app_opt.OutputNone, app_ui.NewQuiet(mc)

	case output == app_opt.OutputJson:
		return app_opt.OutputJson, app_ui.NewQuiet(mc)

	case output == app_opt.OutputText:
		return app_opt.OutputText, app_ui.NewConsole(mc, qt_missingmsg_impl.NewMessageMemory(), false)

	case output == app_opt.OutputMarkdown:
		return app_opt.OutputMarkdown, app_ui.NewMarkdown(mc, os.Stdout, true)

	default:
		u := app_ui.NewConsole(mc, qt_missingmsg_impl.NewMessageMemory(), false)
		u.Error(MRun.ErrorUnsupportedOutput.With("Output", opt.Output))

		// fallback to regular console output
		return app_opt.OutputText, u
	}
}

func (z *bootstrapImpl) Run(rcp rc_recipe.Spec, comSpec *rc_spec.CommonValues) {
	com := comSpec.Opts()
	ufmt, ui := z.SelectUI(z.msgContainer, com)

	// Up
	so := make([]app_control.UpOpt, 0)
	if com.Workspace != "" {
		wsPath, err := ut_filepath.FormatPathWithPredefinedVariables(com.Workspace)
		if err != nil {
			ui.Error(MRun.ErrorUnableToFormatPath.With("Error", err))
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
	so = append(so, app_control.UIFormat(ufmt))
	so = append(so, app_control.LowMemory(com.LowMemory))
	so = append(so, app_control.Concurrency(com.Concurrency))
	so = append(so, app_control.AutoOpen(com.AutoOpen))
	so = append(so, app_control.RecipeName(rcp.CliPath()))
	so = append(so, app_control.CommonOptions(comSpec.Debug()))
	so = append(so, app_control.RecipeOptions(rcp.Debug()))

	ctl := app_control_impl.NewSingle(
		ui,
		z.boxResource,
		z.boxWeb,
		z.msgContainer,
		com.Quiet,
		z.catalogue,
	)

	err := ctl.Up(so...)
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Down()

	// - Quiet
	if qui, ok := ui.(*app_ui.Quiet); ok {
		qui.SetLogger(ctl.Log())
	}

	// Test MRun message, due to unable to test because of package dependency
	if !app.IsProduction() {
		for _, msg := range app_msg.Messages(MRun) {
			ui.Text(msg)
		}
		if qm, ok := ctl.Messages().(app_msg_container.Quality); ok {
			missing := qm.MissingKeys()
			if len(missing) > 0 {
				for _, k := range missing {
					ctl.Log().Error("Key missing", zap.String("key", k))
				}
			}
		}
	}

	// Recover
	if ctl.IsProduction() {
		defer func() {
			err := recover()
			if err != nil {
				l := ctl.Log()
				l.Debug("Recovery from panic")
				ctl.UI().Error(MRun.ErrorPanic.With("Error", err))
				ctl.UI().Error(MRun.ErrorPanicInstruction.With("JobPath", ctl.Workspace().Job()))

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
			ui.Error(MRun.ErrorInterrupted)
			ui.Error(MRun.ErrorInterruptedInstruction.With("JobPath", ctl.Workspace().Job()))
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
	if ctl.IsProduction() && len(rcp.ConnScopes()) > 0 {
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
	ctl.Log().Debug("Run recipe", zap.Any("vo", rcp.Debug()), zap.Any("common", com))
	lastErr = rc_exec.ExecSpec(ctl, rcp, rc_recipe.NoCustomValues)
	if lastErr != nil {
		ctl.Log().Error("Recipe failed with an error", zap.Error(lastErr))
		ui.Failure(MRun.ErrorRecipeFailed.With("Error", lastErr))
		os.Exit(app_control.FailureGeneral)
	}

	// Dump stats
	ut_memory.DumpStats(ctl.Log())
	nw_monitor.DumpStats(ctl.Log())

	app_root.FlushSuccessShutdownHook()
}

func (z *bootstrapImpl) ParseCommon(args []string, ignoreErrors bool) (rem []string, com *rc_spec.CommonValues) {
	comSpec := rc_spec.NewCommonValue()
	f := flag.NewFlagSet(app.Name, flag.ContinueOnError)
	comSpec.SetFlags(f, z.currentUI)
	err := f.Parse(rem)
	rem = f.Args()
	if err != nil {
		if ignoreErrors {
			return []string{}, nil
		} else {
			z.currentUI.Error(MRun.ErrorInvalidArgument.With("Args", strings.Join(args, " ")).With("Error", err))
			os.Exit(app_control.FailureInvalidCommandFlags)
		}
	}
	return rem, comSpec
}

func (z *bootstrapImpl) Parse(args ...string) (rcp rc_recipe.Spec, com *rc_spec.CommonValues) {
	// Select recipe or group
	grp, rcp, rem, err := z.rootGroup.Select(args)

	switch {
	case err != nil:
		z.currentUI.Error(MRun.ErrorInvalidArgument.With("Args", strings.Join(args, " ")))
		if grp != nil {
			grp.PrintUsage(z.currentUI, os.Args[0], app.Version)
		} else {
			z.rootGroup.PrintUsage(z.currentUI, os.Args[0], app.Version)
		}
		os.Exit(app_control.FailureInvalidCommand)

	case rcp == nil:
		grp.PrintUsage(z.currentUI, os.Args[0], app.Version)
		os.Exit(app_control.Success)
	}

	comSpec := rc_spec.NewCommonValue()

	f := flag.NewFlagSet(rcp.CliPath(), flag.ContinueOnError)
	fBuf := &bytes.Buffer{}
	f.SetOutput(fBuf)

	comSpec.SetFlags(f, z.currentUI)
	rcp.SetFlags(f, z.currentUI)

	err = f.Parse(rem)
	rem2 := f.Args()

	switch {
	case err != nil:
		z.currentUI.Error(MRun.ErrorInvalidArgument.
			With("Args", strings.Join(args, " ")).
			With("Error", err))
		rcp.PrintUsage(z.currentUI)
		os.Exit(app_control.FailureInvalidCommandFlags)

	case len(rem2) > 0 && rem2[0] == "help":
		rcp.PrintUsage(z.currentUI)
		os.Exit(app_control.Success)

	case len(rem2) > 0:
		z.currentUI.Error(MRun.ErrorTooManyArguments.
			With("Args", strings.Join(rem2, " ")).
			With("AllArgs", strings.Join(args, " ")))
		rcp.PrintUsage(z.currentUI)
		os.Exit(app_control.FailureInvalidCommandFlags)

	}

	comSpec.Apply()
	return rcp, comSpec
}
