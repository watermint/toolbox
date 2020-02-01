package job

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"time"
)

type Archive struct {
	Days               int
	ProgressArchiving  app_msg.Message
	ErrorFailedArchive app_msg.Message
}

func (z *Archive) Exec(c app_control.Control) error {
	historian := app_job_impl.NewHistorian(c)
	histories := historian.Histories()
	threshold := time.Now().Add(time.Duration(-z.Days*24) * time.Hour)
	l := c.Log()

	for _, h := range histories {
		ts, found := h.TimeStart()
		if !found {
			l.Debug("Skip: Time start not found for the job")
			continue
		}
		if ts.After(threshold) {
			l.Debug("Skip: Time start is in range of retain", zap.String("jobId", h.JobId()))
			continue
		}
		c.UI().Info(z.ProgressArchiving.With("JobId", h.JobId()))
		_, err := h.Archive()
		if err != nil {
			l.Debug("Unable to archive", zap.Error(err), zap.Any("history", h))
			c.UI().Error(z.ErrorFailedArchive.With("JobId", h.JobId()).With("Error", err.Error()))
		}
	}

	return nil
}

func (z *Archive) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Archive{}, func(r rc_recipe.Recipe) {
		m := r.(*Archive)
		m.Days = 7
	})
}

func (z *Archive) Preset() {
	z.Days = 7
}
