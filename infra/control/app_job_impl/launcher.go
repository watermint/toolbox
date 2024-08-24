package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_feature"
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
		CommonOpts:   z.com,
		TimeStart:    time.Now().Format(time.RFC3339),
		JobId:        ctl.Workspace().JobId(),
		AppName:      app_definitions.Name,
		AppHash:      app_definitions.BuildInfo.Hash,
		AppVersion:   app_definitions.BuildId,
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

func (z launchImpl) prepAuthDatabase(fe app_feature.Feature) (repo api_auth.Repository, err error) {
	// returns in memory database
	if fe.IsSecure() {
		return api_auth_repo.NewInMemory()
	}
	if fe.IsTransient() && fe.IsDefaultPathAuthRepository() {
		return api_auth_repo.NewInMemory()
	}

	return api_auth_repo.NewPersistent(fe.PathAuthRepository())
}

func (z launchImpl) Up() (ctl app_control.Control, err error) {
	lg := z.wb.Logger().Logger()
	sm := z.wb.Summary().Logger()
	st := z.wb.Stats().Logger()
	fe := app_feature_impl.NewFeature(z.com, z.wb.Workspace(), z.rcp.IsTransient())

	esl.SetStats(st)

	seq, er := app_queue.NewSequence(lg, fe, z.ui, z.wb)
	ar, err := z.prepAuthDatabase(fe)
	if err != nil {
		return nil, err
	}
	ctl = app_control_impl.New(z.wb, z.ui, fe, seq, ar, er)

	if err := er.Up(ctl); err != nil {
		return nil, err
	}

	if ctl.Feature().IsTransient() || ctl.Feature().IsSkipLogging() {
		_, err = z.rcp.Capture(ctl)
		return ctl, err
	}

	if err := z.recordStartLog(ctl); err != nil {
		return nil, err
	}

	// Launch monitor
	es_memory.LaunchReporting(st)

	sm.Debug("Up completed",
		esl.String("name", app_definitions.Name),
		esl.String("ver", app_definitions.BuildId),
		esl.String("hash", app_definitions.BuildInfo.Hash),
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
