package history

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"time"
)

type Archive struct {
	Days               mo_int.RangeInt
	ProgressArchiving  app_msg.Message
	ErrorFailedArchive app_msg.Message
}

func (z *Archive) Exec(c app_control.Control) error {
	historian := app_job_impl.NewHistorian(c)
	histories := historian.Histories()
	threshold := time.Now().Add(time.Duration(-z.Days.Value()*24) * time.Hour)
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
			l.Debug("Skip: Time start is in range of retain", es_log.String("jobId", h.JobId()))
			continue
		}
		c.UI().Info(z.ProgressArchiving.With("JobId", h.JobId()))
		_, err := h.Archive()
		if err != nil {
			l.Debug("Unable to archive", es_log.Error(err), es_log.Any("history", h))
			c.UI().Error(z.ErrorFailedArchive.With("JobId", h.JobId()).With("Error", err.Error()))
		}
	}

	return nil
}

func (z *Archive) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Archive{}, func(r rc_recipe.Recipe) {
		m := r.(*Archive)
		m.Days.SetValue(7)
	})
}

func (z *Archive) Preset() {
	z.Days.SetRange(1, 3650, 7)
}
