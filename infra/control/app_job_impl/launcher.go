package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_feature_impl"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_queue"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"time"
)

type MsgLauncher struct {
	ElapsedTimeOnEnd app_msg.Message
}

var (
	MLauncher = app_msg.Apply(&MsgLauncher{}).(*MsgLauncher)
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
	l := ctl.Log()
	l.Debug("Capture recipe values")
	rv, err := z.rcp.Capture(ctl)
	if err != nil {
		l.Debug("Unable to capture recipe values", esl.Error(err))
		return err
	}

	sl := app_job.StartLog{
		Name:         z.rcp.CliPath(),
		ValueObject:  z.rcp.Debug(),
		JobId:        ctl.Workspace().JobId(),
		TimeStart:    time.Now().Format(time.RFC3339),
		AppName:      app.Name,
		AppHash:      app.BuildInfo.Hash,
		AppVersion:   app.BuildId,
		RecipeValues: rv,
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

	seq, er := app_queue.NewSequence(lg, fe, z.ui, z.wb)

	ctl = app_control_impl.New(z.wb, z.ui, fe, seq, er)

	if err := er.Up(ctl); err != nil {
		return nil, err
	}

	if ctl.Feature().IsTransient() || ctl.Feature().IsSkipLogging() {
		return ctl, nil
	}

	if err := z.recordStartLog(ctl); err != nil {
		return nil, err
	}

	// Launch monitor
	es_memory.LaunchReporting(lg)

	sm.Debug("Up completed",
		esl.String("name", app.Name),
		esl.String("ver", app.BuildId),
		esl.String("hash", app.BuildInfo.Hash),
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
	rmJobData := z.com.ShouldDeleteJobData(err)

	if cc, ok := ctl.(app_control.ControlCloser); ok {
		cc.Close()
	}

	if !rmJobData {
		artifacts := rp_artifact.Artifacts(ctl.Workspace())
		for _, artifact := range artifacts {
			ui.Link(artifact)
		}
	}

	// Dump stats
	es_memory.DumpMemStats(sm)

	if !ctl.Feature().IsSkipLogging() {
		_ = z.recordResultLog(err)
	}

	timeEnd := time.Now()
	elapsedTime := timeEnd.Sub(z.wb.Workspace().JobStartTime()).Truncate(time.Millisecond)

	ui.Progress(MLauncher.ElapsedTimeOnEnd.With("Duration", elapsedTime.String()))

	sm.Debug("Down completed", esl.Error(err), esl.Bool("rmJobData", rmJobData))

	// Close work bundle
	_ = z.wb.Close()

	if rmJobData {
		z.deleteJobData()
	}
}

func (z launchImpl) deleteJobData() {
	path := z.wb.Workspace().Job()
	l := esl.ConsoleOnly()
	l.Debug("Remove job data", esl.String("jobPath", path))
	_ = os.RemoveAll(path)
}
