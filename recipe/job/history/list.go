package history

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"time"
)

type List struct {
	Log rp_model.RowReport
}

type JobRecord struct {
	JobId      string `json:"job_id"`
	AppVersion string `json:"app_version"`
	RecipeName string `json:"recipe_name"`
	TimeStart  string `json:"time_start"`
	TimeFinish string `json:"time_finish"`
}

func (z *List) Exec(c app_control.Control) error {
	historian := app_job_impl.NewHistorian(c)
	histories := historian.Histories()
	if err := z.Log.Open(); err != nil {
		return err
	}

	for _, h := range histories {
		ts := ""
		tf := ""
		if t, found := h.TimeStart(); found {
			ts = t.Format(time.RFC3339)
		}
		if t, found := h.TimeFinish(); found {
			tf = t.Format(time.RFC3339)
		}
		z.Log.Row(&JobRecord{
			JobId:      h.JobId(),
			AppVersion: h.AppVersion(),
			RecipeName: h.RecipeName(),
			TimeStart:  ts,
			TimeFinish: tf,
		})
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {})
}

func (z *List) Preset() {
	z.Log.SetModel(&JobRecord{})
}
