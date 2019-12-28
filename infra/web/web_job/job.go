package web_job

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"os"
)

type WebJobRun struct {
	Name      string
	JobId     string
	Recipe    rc_recipe.Spec
	UC        app_control.Control
	UiLogFile *os.File
}

func Runner(ctl app_control.Control, jc <-chan *WebJobRun) {
	ui := ctl.UI()
	for job := range jc {
		l := ctl.Log().With(zap.String("name", job.Name), zap.String("jobId", job.JobId))
		l.Debug("Start a new job")
		if rcp, err := job.Recipe.SpinUp(ctl, rc_recipe.NoCustomValues); err != nil {
			l.Error("Unable to start the job", zap.Error(err))
			ui.Failure("web.job.result.failure", app_msg.P{"ErrorK": err.Error()})
		} else {
			err := rcp.Exec(ctl)
			if err != nil {
				l.Error("Unable to finish the job", zap.Error(err))
				ui.Failure("web.job.result.failure", app_msg.P{"ErrorK": err.Error()})
			} else {
				ui.Success("web.job.result.success")
			}
		}
		l.Debug("Closing log file")
		job.UiLogFile.Close()

		l.Debug("Job spin down")
		job.UC.Down()

		l.Debug("The job finished")
	}
}
