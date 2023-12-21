package app_bootstrap

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_bandwidth"
	"github.com/watermint/toolbox/essentials/network/nw_congestion"
	"github.com/watermint/toolbox/essentials/network/nw_diag"
	"github.com/watermint/toolbox/essentials/network/nw_proxy"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/essentials/runtime/es_open"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_lifecycle"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/ingredient/ig_bootstrap"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

type msgRun struct {
	ErrorInvalidArgument                  app_msg.Message
	ErrorTooManyArguments                 app_msg.Message
	ErrorRetainNoneSupportsJsonReportOnly app_msg.Message
	ErrorInterrupted                      app_msg.Message
	ErrorInterruptedInstruction           app_msg.Message
	ErrorUnableToFormatPath               app_msg.Message
	ErrorPanic                            app_msg.Message
	ErrorPanicInstruction                 app_msg.Message
	ErrorRecipeFailed                     app_msg.Message
	ErrorUnsupportedOutput                app_msg.Message
	ErrorInitialization                   app_msg.Message
	ErrorUnableToLoadExtra                app_msg.Message
	ProgressInterruptedShutdown           app_msg.Message
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
	var mc app_msg_container.Container
	var err error

	switch opt.Lang.Value() {
	case app_opt.LangEnglish:
		mc, err = app_msg_container_impl.NewSingle(es_lang.English)
		if err != nil {
			panic(err)
		}
	case app_opt.LangJapanese:
		mc, err = app_msg_container_impl.NewSingle(es_lang.Japanese)
		if err != nil {
			panic(err)
		}
	default:
		mc = app_msg_container_impl.NewContainer()
	}
	out := opt.Output.Value()
	lg := esl.Default()

	// Select UI
	switch {
	case opt.Quiet, out == app_opt.OutputNone:
		return app_ui.NewDiscard(mc, lg)

	case out == app_opt.OutputJson:
		return app_ui.NewDiscard(mc, lg)

	case out == app_opt.OutputText:
		w := es_stdout.NewDirectOut()
		return app_ui.NewConsole(mc, lg, w, es_dialogue.New(w))

	case out == app_opt.OutputMarkdown:
		w := es_stdout.NewDirectOut()
		return app_ui.NewMarkdown(mc, lg, w, es_dialogue.New(w))

	default:
		w := es_stdout.NewDirectOut()
		u := app_ui.NewConsole(mc, lg, w, es_dialogue.New(w))
		u.Error(MRun.ErrorUnsupportedOutput.With("Output", opt.Output))

		// fallback to regular console output
		return u
	}
}

func (z *bsImpl) verifyMessages(ui app_ui.UI, l esl.Logger) {
	// Test MRun message, due to unable to test because of package dependency
	if !app_definitions2.IsProduction() {
		for _, msg := range app_msg.Messages(MRun) {
			ui.Text(msg)
		}
		missing := qt_msgusage.Record().Missing()
		if len(missing) > 0 {
			w := es_stdout.NewDirectOut()
			for _, k := range missing {
				l.Error("Key missing", esl.String("key", k))
				_, _ = fmt.Fprintf(w, `"%s":"",\n`, k)
			}
		}
	}
}

func (z *bsImpl) Run(rcp rc_recipe.Spec, comSpec *rc_spec.CommonValues) {
	com := comSpec.Opts()

	ui := z.SelectUI(com)

	// Check binary build time
	app_lifecycle.LifecycleControl().Verify(ui)

	if exErr := com.ExtraLoad(); exErr != nil {
		ui.Failure(MRun.ErrorUnableToLoadExtra.With("Error", exErr).With("Path", com.Extra.Value()))
		app_exit.Abort(app_exit.FatalStartup)
	}

	if 0 < len(rcp.Reports()) &&
		com.RetainJobData.Value() != app_opt.RetainJobDataDefault &&
		com.Output.Value() != app_opt.OutputJson {
		ui.Failure(MRun.ErrorRetainNoneSupportsJsonReportOnly)
		app_exit.Abort(app_exit.FatalStartup)
	}

	clv := app_feature.ConsoleLogLevel(false, com.Debug)
	wb, err := app_workspace.NewBundle(
		com.Workspace.Value(),
		app_budget.Budget(com.BudgetStorage.Value()),
		clv,
		rcp.IsTransient(),
		rcp.IsTransient() || com.SkipLogging,
	)
	if err != nil {
		ui.Failure(MRun.ErrorInitialization.With("Error", err))
		app_exit.Abort(app_exit.FatalStartup)
	}
	z.verifyMessages(ui, wb.Logger().Logger())
	esl.AddDefaultSubscriber(wb.Logger().Core())

	jl := app_job_impl.NewLauncher(ui, wb, com, rcp)
	ctl, err := jl.Up()
	if err != nil {
		ui.Failure(MRun.ErrorInitialization.With("Error", err))
		app_exit.Abort(app_exit.FatalStartup)
	}

	// Notification
	switch {
	case com.Quiet,
		com.Output.Value() == app_opt.OutputJson,
		com.Output.Value() == app_opt.OutputNone,
		es_env.IsEnabled(app_definitions2.EnvNameDebugVerbose),
		ctl.Feature().Experiment(app_definitions2.ExperimentSuppressProgress):

		wb.Logger().Logger().Debug("Set indicators as silent mode")
		ea_indicator.SuppressIndicatorForce()
	default:
		wb.Logger().Logger().Debug("Start indicators")
		ea_indicator.StartIndicator()
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
	rc_group.AppHeader(ui, app_definitions2.BuildId)

	// Global settings
	nw_proxy.Setup("https://api.dropboxapi.com", com.Proxy.Value(), ctl.Log())
	nw_bandwidth.SetBandwidth(com.BandwidthKb)
	nw_congestion.SetMaxCongestionWindow(com.Concurrency,
		ctl.Feature().Experiment(app_definitions2.ExperimentCongestionWindowNoLimit))
	if ctl.Feature().Experiment(app_definitions2.ExperimentCongestionWindowAggressive) {
		ctl.Log().Debug("Enable aggressive initial window")
		nw_congestion.SetInitCongestionWindow(com.Concurrency)
	}

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

	// Bootstrap recipe
	if err := rc_exec.Exec(ctl, &ig_bootstrap.Bootstrap{}, rc_recipe.NoCustomValues); err != nil {
		ctl.Log().Error("Bootstrap failed with an error", esl.Error(err))
		ui.Failure(MRun.ErrorRecipeFailed.With("Error", err))
		jl.Down(err, ctl)
		app_exit.Abort(app_exit.FailureGeneral)
	}

	// Apply profiler
	var prof interface{ Stop() }
	if ctl.Feature().IsDebug() || ctl.Feature().Experiment(app_definitions2.ExperimentProfileMemory) {
		prof = profile.Start(
			profile.ProfilePath(ctl.Workspace().Log()),
			profile.MemProfile,
		)
	} else if ctl.Feature().Experiment(app_definitions2.ExperimentProfileCpu) {
		prof = profile.Start(
			profile.ProfilePath(ctl.Workspace().Log()),
			profile.CPUProfile,
		)
	}

	// Run
	var lastErr error
	ctl.WorkBundle().Summary().Logger().Debug("Run recipe", esl.Any("vo", rcp.Debug()), esl.Any("common", com))
	lastErr = rc_exec.ExecSpec(ctl, rcp, rc_recipe.NoCustomValues)

	// stop the profiler
	if prof != nil {
		prof.Stop()
	}

	// shutdown job
	jl.Down(lastErr, ctl)

	// abort on error
	if lastErr != nil {
		ctl.Log().Debug("Recipe failed with an error", esl.Error(lastErr))
		ui.Failure(MRun.ErrorRecipeFailed.With("Error", lastErr))
		app_exit.Abort(app_exit.FailureGeneral)
	}

	if ctl.Feature().IsAutoOpen() {
		artifacts := rp_artifact.Artifacts(wb.Workspace())
		if len(artifacts) > 0 {
			op := es_open.New()
			opErr := op.Open(wb.Workspace().Report())
			ctl.Log().Debug("open", esl.Error(opErr))
		}
	}

	app_exit.ExitSuccess()
}

func (z bsImpl) bootUI() app_ui.UI {
	lg := esl.Default()
	mc := app_msg_container_impl.NewContainer()
	wr := es_stdout.NewDirectOut()
	dg := es_dialogue.New(wr)
	return app_ui.NewConsole(mc, lg, wr, dg)
}

func (z *bsImpl) ParseCommon(args []string, ignoreErrors bool) (rem []string, com *rc_spec.CommonValues) {
	comSpec := rc_spec.NewCommonValue()
	f := flag.NewFlagSet(app_definitions2.Name, flag.ContinueOnError)
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
			grp.PrintUsage(ui, os.Args[0], app_definitions2.BuildId)
		} else {
			rg.PrintUsage(ui, os.Args[0], app_definitions2.BuildId)
		}
		app_exit.Abort(app_exit.FailureInvalidCommand)

	case rcp == nil:
		grp.PrintUsage(ui, os.Args[0], app_definitions2.BuildId)
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
		l.Info(ui.Text(MRun.ProgressInterruptedShutdown))
		l.Debug("Signal", esl.Any("signal", s))
		pc := make([]uintptr, 16)
		n := runtime.Callers(0, pc)
		pc = pc[:n]
		frames := runtime.CallersFrames(pc)
		for f := 0; ; f++ {
			frame, more := frames.Next()
			ctl.Log().Debug("Frame",
				esl.Int("Frame", f),
				esl.String("File", frame.File),
				esl.Int("Line", frame.Line),
				esl.String("Function", frame.Function),
			)
			if !more {
				break
			}
		}
		l.Error(ui.Text(MRun.ErrorInterrupted))
		l.Error(ui.Text(MRun.ErrorInterruptedInstruction.With("JobPath", ctl.Workspace().Job())))
		app_exit.Abort(app_exit.FatalInterrupted)
	}
}

func trapPanic(ctl app_control.Control) {
	l := ctl.Log()
	ui := ctl.UI()
	err := recover()
	if err != nil {
		l.Debug("Recovery from panic")
		l.Error(ui.Text(MRun.ErrorPanic.With("Error", err)))
		l.Error(ui.Text(MRun.ErrorPanicInstruction.With("JobPath", ctl.Workspace().Job())))

		for depth := 0; ; depth++ {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			l.Debug("Trace",
				esl.Int("Depth", depth),
				esl.String("File", file),
				esl.Int("Line", line),
			)
		}
		app_exit.Abort(app_exit.FatalPanic)
	}
}
