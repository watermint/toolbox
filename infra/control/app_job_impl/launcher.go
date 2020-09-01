package app_job_impl

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe_preserve"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_error"
	"github.com/watermint/toolbox/infra/control/app_feature_impl"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"time"
)

func NewLauncher(ui app_ui.UI, wb app_workspace.Bundle, com app_opt.CommonOpts, rcp rc_recipe.Spec) app_job.Launcher {
	return &launchImpl{
		ui:  ui,
		wb:  wb,
		com: com,
		rcp: rcp,
	}
}

type launchImpl struct {
	ui  app_ui.UI
	wb  app_workspace.Bundle
	com app_opt.CommonOpts
	rcp rc_recipe.Spec
}

func (z launchImpl) recordStartLog(ctl app_control.Control) error {
	sl := app_job.StartLog{
		Name:        z.rcp.CliPath(),
		ValueObject: z.rcp.Debug(),
		JobId:       ctl.Workspace().JobId(),
		TimeStart:   time.Now().Format(time.RFC3339),
		AppName:     app.Name,
		AppHash:     app.Hash,
		AppVersion:  app.Version,
	}
	return sl.Write(z.wb.Workspace())
}

func (z launchImpl) recordResultLog(err error) error {
	errText := ""
	if err != nil {
		errText = err.Error()
	}
	rl := app_job.ResultLog{
		Success:    err == nil,
		TimeFinish: time.Now().Format(time.RFC3339),
		Error:      errText,
	}
	return rl.Write(z.wb.Workspace())
}

func (z launchImpl) Up() (ctl app_control.Control, err error) {
	lg := z.wb.Logger().Logger()
	sm := z.wb.Summary().Logger()
	fe := app_feature_impl.NewFeature(z.com, z.wb.Workspace(), z.rcp.IsTransient())

	preservePath := z.wb.Workspace().KVS()
	preserve := eq_pipe_preserve.NewFactory(lg, preservePath)
	factory := eq_pipe.NewSimple(lg, preserve)
	progress := eq_progress.NewBar(
		//		mpb.WithOutput(es_stdout.NewDefaultOut(z.feature)),
		mpb.WithWidth(72),
	)
	if fe.IsQuiet() {
		progress = nil
	}

	er := app_error.NewErrorReport(lg, z.wb, z.ui)

	batchPolicy := eq_bundle.FetchSequential
	if fe.Experiment(app.ExperimentBatchRandom) {
		batchPolicy = eq_bundle.FetchRandom
	}
	if fe.Experiment(app.ExperimentBatchSequential) {
		batchPolicy = eq_bundle.FetchSequential
	}

	seq := eq_sequence.New(
		eq_queue.Logger(lg),
		eq_queue.FetchPolicy(batchPolicy),
		eq_queue.Progress(progress),
		eq_queue.NumWorker(fe.Concurrency()),
		eq_queue.Factory(factory),
		eq_queue.ErrorHandler(er.ErrorHandler),
	)

	ctl = app_control_impl.New(z.wb, z.ui, fe, seq, er)

	if err := er.Up(ctl); err != nil {
		return nil, err
	}

	if ctl.Feature().IsTransient() {
		return ctl, nil
	}

	if err := z.recordStartLog(ctl); err != nil {
		return nil, err
	}

	// Launch monitor
	es_memory.LaunchReporting(lg)

	sm.Debug("Up completed",
		esl.String("name", app.Name),
		esl.String("ver", app.Version),
		esl.String("hash", app.Hash),
		esl.String("recipe", z.rcp.CliPath()),
	)

	return ctl, nil
}

func (z launchImpl) Down(err error, ctl app_control.Control) {
	if ctl.Feature().IsTransient() {
		return
	}

	sm := ctl.WorkBundle().Summary().Logger()
	ui := ctl.UI()

	if cc, ok := ctl.(app_control.ControlCloser); ok {
		cc.Close()
	}

	artifacts := rp_artifact.Artifacts(ctl.Workspace())
	for _, artifact := range artifacts {
		ui.Link(artifact)
	}

	// Dump stats
	es_memory.DumpMemStats(sm)

	_ = z.recordResultLog(err)

	sm.Debug("Down completed", esl.Error(err))

	// Close work bundle
	_ = z.wb.Close()
}
