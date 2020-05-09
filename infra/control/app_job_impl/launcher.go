package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/stats/es_http"
	"github.com/watermint/toolbox/essentials/log/stats/es_memory"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
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

func (z launchImpl) recordStartLog() error {
	sl := app_job.StartLog{
		Name:        z.rcp.CliPath(),
		ValueObject: z.rcp.Debug(),
		TimeStart:   time.Now().Format(time.RFC3339),
		AppName:     app.Name,
		AppHash:     app.Hash,
		AppVersion:  app.Version,
	}
	return sl.Create(z.wb.Workspace())
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
	return rl.Create(z.wb.Workspace())
}

func (z launchImpl) Up() (ctl app_control.Control, err error) {
	lg := z.wb.Logger().Logger()
	sm := z.wb.Summary().Logger()
	fe := app_feature_impl.NewFeature(z.com, z.wb.Workspace(), z.rcp.IsTransient())
	ctl = app_control_impl.New(z.wb, z.ui, fe)

	if ctl.Feature().IsTransient() {
		return ctl, nil
	}

	if err := z.recordStartLog(); err != nil {
		return nil, err
	}

	// Launch monitor
	es_http.LaunchReporting(lg)
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

	artifacts := rp_artifact.Artifacts(ctl.Workspace())
	for _, artifact := range artifacts {
		ui.Link(artifact)
	}

	// Dump stats
	es_memory.DumpMemStats(sm)
	es_http.DumpStats(sm)

	_ = z.recordResultLog(err)

	sm.Debug("Down completed", esl.Error(err))

	// Close work bundle
	_ = z.wb.Close()
}
