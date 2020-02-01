package job

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"time"
)

type History struct {
	TableJobId      app_msg.Message
	TableAppVersion app_msg.Message
	TableRecipeName app_msg.Message
	TableTimeStart  app_msg.Message
	TableTimeFinish app_msg.Message
}

func (z *History) Exec(c app_control.Control) error {
	historian := app_job_impl.NewHistorian(c)
	histories := historian.Histories()
	t := c.UI().InfoTable("History")

	t.Header(z.TableJobId, z.TableAppVersion, z.TableRecipeName, z.TableTimeStart, z.TableTimeFinish)
	for _, h := range histories {
		ts := ""
		tf := ""
		if t, found := h.TimeStart(); found {
			ts = t.Format(time.RFC3339)
		}
		if t, found := h.TimeFinish(); found {
			tf = t.Format(time.RFC3339)
		}

		t.RowRaw(h.JobId(), h.AppVersion(), h.RecipeName(), ts, tf)
	}
	t.Flush()

	return nil
}

func (z *History) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &History{}, func(r rc_recipe.Recipe) {})
}

func (z *History) Preset() {
}
