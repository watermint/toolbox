package app_bootstrap

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_concurrency"
	"github.com/watermint/toolbox/infra/network/nw_diag"
	"github.com/watermint/toolbox/infra/network/nw_proxy"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/ingredient/bootstrap"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

type msgRun struct {
	ErrorInvalidArgument        app_msg.Message
	ErrorTooManyArguments       app_msg.Message
	ErrorInterrupted            app_msg.Message
	ErrorInterruptedInstruction app_msg.Message
	ErrorUnableToFormatPath     app_msg.Message
	ErrorPanic                  app_msg.Message
	ErrorPanicInstruction       app_msg.Message
	ErrorRecipeFailed           app_msg.Message
	ErrorUnsupportedOutput      app_msg.Message
	ErrorInitialization         app_msg.Message
}

var (
	MRun = app_msg.Apply(&msgRun{}).(*msgRun)
)

// Mutable object for running recipe
type Bootstrap interface {
	// Parse arguments for recipe & common values
	Parse(args ...string) (rcp rc_recipe.Spec, com *rc_spec.CommonValues)

	// Parse arguments for common values only
	ParseCommon(args []string, ignoreErrors bool) (rem []string, com *rc_spec.CommonValues)

	// Run recipe
	Run(rcp rc_recipe.Spec, com *rc_spec.CommonValues)
}

func NewBootstrap() Bootstrap {
	return &bsImpl{}
}

type bsImpl struct {
}

func (z *bsImpl) SelectUI(opt app_opt.CommonOpts) (ui app_ui.UI) {
	mc := app_msg_container_impl.NewContainer()
	out := opt.Output.Value()
	lg := es_log.Default()

	// Select UI
	switch {
	case opt.Quiet, out == app_opt.OutputNone:
		return app_ui.NewDiscard(mc, lg)

	case out == app_opt.OutputJson:
		return app_ui.NewDiscard(mc, lg)

	case out == app_opt.OutputText:
		w := es_stdout.NewDefaultOut(false)
		return app_ui.NewConsole(mc, lg, w, es_dialogue.New(w))

	case out == app_opt.OutputMarkdown:
		w := es_stdout.NewDefaultOut(false)
		return app_ui.NewMarkdown(mc, lg, w, es_dialogue.New(w))

	default:
		w := es_stdout.NewDefaultOut(false)
		u := app_ui.NewConsole(mc, lg, w, es_dialogue.New(w))
		u.Error(MRun.ErrorUnsupportedOutput.With("Output", opt.Output))

		// fallback to regular console output
		return u
	}
}

func (z *bsImpl) verifyMessages(ui app_ui.UI, l es_log.Logger) {
	// Test MRun message, due to unable to test because of package dependency
	if !app.IsProduction() {
		for _, msg := range app_msg.Messages(MRun) {
			ui.Text(msg)
		}
		missing := qt_missingmsg.Record().Missing()
		if len(missing) > 0 {
			w := es_stdout.NewDefaultOut(false)
			for _, k := range missing {
				l.Error("Key missing", es_log.String("key", k))
				_, _ = fmt.Fprintf(w, `"%s":"",\n`, k)
			}
		}
	}
}

func (z *bsImpl) Run(rcp rc_recipe.Spec, comSpec *rc_spec.CommonValues) {
	com := comSpec.Opts()
	ui := z.SelectUI(com)

	clv := app_feature.ConsoleLogLevel(false, com.Debug)
	wb, err := app_workspace.NewBundle(com.Workspace.Value(), app_budget.Budget(com.BudgetStorage.Value()), clv)
	if err != nil {
		ui.Failure(MRun.ErrorInitialization.With("Error", err))
		app_exit.Abort(app_exit.FatalStartup)
	}
	z.verifyMessages(ui, wb.Logger().Logger())

	jl := app_job_impl.NewLauncher(ui, wb, com, rcp)
	ctl, err := jl.Up()
	if err != nil {
		ui.Failure(MRun.ErrorInitialization.With("Error", err))
		app_exit.Abort(app_exit.FatalStartup)
	}

	// Recover
	if ctl.Feature().IsProduction() {
		defer trapPanic(ctl)
	}

	// Trap signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)
	go trapSignal(sig, ctl)

	// App Header
	rc_group.AppHeader(ui, app.Version)

	// Global settings
	nw_proxy.SetHttpProxy(com.Proxy.Value(), ctl)
	nw_bandwidth.SetBandwidth(com.BandwidthKb)
	nw_concurrency.SetConcurrency(com.Concurrency)

	// Diagnosis
	if err = nw_diag.Runtime(ctl); err != nil {
		jl.Down(err, ctl)
		app_exit.Abort(app_exit.FatalRuntime)
	}
	if ctl.Feature().IsProduction() && len(rcp.ConnScopes()) > 0 {
		if err = nw_diag.Network(ctl); err != nil {
			jl.Down(err, ctl)
			app_exit.Abort(app_exit.FatalNetwork)
		}
	}

	// Apply profiler
	if com.Debug {
		defer profile.Start(
			profile.ProfilePath(ctl.Workspace().Log()),
			profile.MemProfile,
		).Stop()
	}

	// Bootstrap recipe
	if err := rc_exec.Exec(ctl, &bootstrap.Bootstrap{}, rc_recipe.NoCustomValues); err != nil {
		ctl.Log().Error("Bootstrap failed with an error", es_log.Error(err))
		ui.Failure(MRun.ErrorRecipeFailed.With("Error", err))
		jl.Down(err, ctl)
		app_exit.Abort(app_exit.FailureGeneral)
	}

	// Run
	var lastErr error
	ctl.WorkBundle().Summary().Logger().Debug("Run recipe", es_log.Any("vo", rcp.Debug()), es_log.Any("common", com))
	lastErr = rc_exec.ExecSpec(ctl, rcp, rc_recipe.NoCustomValues)

	// shutdown job
	jl.Down(lastErr, ctl)

	// abort on error
	if lastErr != nil {
		ctl.Log().Error("Recipe failed with an error", es_log.Error(lastErr))
		ui.Failure(MRun.ErrorRecipeFailed.With("Error", lastErr))
		app_exit.Abort(app_exit.FailureGeneral)
	}

	app_exit.ExitSuccess()
}

func (z bsImpl) bootUI() app_ui.UI {
	lg := es_log.Default()
	mc := app_msg_container_impl.NewContainer()
	wr := es_stdout.NewDefaultOut(false)
	dg := es_dialogue.New(wr)
	return app_ui.NewConsole(mc, lg, wr, dg)
}

func (z *bsImpl) ParseCommon(args []string, ignoreErrors bool) (rem []string, com *rc_spec.CommonValues) {
	comSpec := rc_spec.NewCommonValue()
	f := flag.NewFlagSet(app.Name, flag.ContinueOnError)
	ui := z.bootUI()
	comSpec.SetFlags(f, ui)
	err := f.Parse(rem)
	rem = f.Args()
	if err != nil {
		if ignoreErrors {
			return []string{}, nil
		} else {
			ui.Error(MRun.ErrorInvalidArgument.With("Args", strings.Join(args, " ")).With("Error", err))
			app_exit.Abort(app_exit.FailureInvalidCommandFlags)
		}
	}
	return rem, comSpec
}

func (z *bsImpl) Parse(args ...string) (rcp rc_recipe.Spec, com *rc_spec.CommonValues) {
	ui := z.bootUI()
	cat := app_catalogue.Current()
	rg := cat.RootGroup()

	// Select recipe or group
	grp, rcp, rem, err := rg.Select(args)

	switch {
	case err != nil:
		ui.Error(MRun.ErrorInvalidArgument.With("Args", strings.Join(args, " ")))
		if grp != nil {
			grp.PrintUsage(ui, os.Args[0], app.Version)
		} else {
			rg.PrintUsage(ui, os.Args[0], app.Version)
		}
		app_exit.Abort(app_exit.FailureInvalidCommand)

	case rcp == nil:
		grp.PrintUsage(ui, os.Args[0], app.Version)
		app_exit.ExitSuccess()
	}

	comSpec := rc_spec.NewCommonValue()

	f := flag.NewFlagSet(rcp.CliPath(), flag.ContinueOnError)
	fBuf := &bytes.Buffer{}
	f.SetOutput(fBuf)

	comSpec.SetFlags(f, ui)
	rcp.SetFlags(f, ui)

	err = f.Parse(rem)
	rem2 := f.Args()

	switch {
	case err != nil:
		ui.Error(MRun.ErrorInvalidArgument.
			With("Args", strings.Join(args, " ")).
			With("Error", err))
		rcp.PrintUsage(ui)
		app_exit.Abort(app_exit.FailureInvalidCommandFlags)

	case len(rem2) > 0 && rem2[0] == "help":
		rcp.PrintUsage(ui)
		app_exit.ExitSuccess()

	case len(rem2) > 0:
		ui.Error(MRun.ErrorTooManyArguments.
			With("Args", strings.Join(rem2, " ")).
			With("AllArgs", strings.Join(args, " ")))
		rcp.PrintUsage(ui)
		app_exit.Abort(app_exit.FailureInvalidCommandFlags)

	}

	comSpec.Apply()
	return rcp, comSpec
}

func trapSignal(sig chan os.Signal, ctl app_control.Control) {
	l := ctl.Log()
	ui := ctl.UI()
	for s := range sig {
		// fatal shutdown
		l.Debug("Signal", es_log.Any("signal", s))
		pc := make([]uintptr, 16)
		n := runtime.Callers(0, pc)
		pc = pc[:n]
		frames := runtime.CallersFrames(pc)
		for f := 0; ; f++ {
			frame, more := frames.Next()
			ctl.Log().Debug("Frame",
				es_log.Int("Frame", f),
				es_log.String("File", frame.File),
				es_log.Int("Line", frame.Line),
				es_log.String("Function", frame.Function),
			)
			if !more {
				break
			}
		}
		ui.Error(MRun.ErrorInterrupted)
		ui.Error(MRun.ErrorInterruptedInstruction.With("JobPath", ctl.Workspace().Job()))
		app_exit.Abort(app_exit.FatalInterrupted)
	}
}

func trapPanic(ctl app_control.Control) {
	l := ctl.Log()
	ui := ctl.UI()
	err := recover()
	if err != nil {
		l.Debug("Recovery from panic")
		ui.Error(MRun.ErrorPanic.With("Error", err))
		ui.Error(MRun.ErrorPanicInstruction.With("JobPath", ctl.Workspace().Job()))

		for depth := 0; ; depth++ {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			l.Debug("Trace",
				es_log.Int("Depth", depth),
				es_log.String("File", file),
				es_log.Int("Line", line),
			)
		}
		app_exit.Abort(app_exit.FatalPanic)
	}
}
