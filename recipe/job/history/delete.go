package history

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"time"
)

type Delete struct {
	Days              int
	ProgressDeleting  app_msg.Message
	ErrorFailedDelete app_msg.Message
}

func (z *Delete) Exec(c app_control.Control) error {
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
		if h.JobId() == c.Workspace().JobId() {
			continue
		}
		if ts.After(threshold) {
			l.Debug("Skip: Time start is in range of retain", zap.String("jobId", h.JobId()))
			continue
		}
		c.UI().Info(z.ProgressDeleting.With("JobId", h.JobId()))
		err := h.Delete()
		if err != nil {
			l.Debug("Unable to archive", zap.Error(err), zap.Any("history", h))
			c.UI().Error(z.ProgressDeleting.With("JobId", h.JobId()).With("Error", err.Error()))
		}
	}
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Days = 365
	})
}

func (z *Delete) Preset() {
	z.Days = 28
}