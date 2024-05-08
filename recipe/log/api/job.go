package api

import (
	"errors"
	"github.com/watermint/toolbox/domain/core/dc_log"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"os"
	"path/filepath"
)

type Job struct {
	JobId          mo_string.OptionalString
	FullUrl        bool
	IntervalSecond int
	Population     rp_model.RowReport
	Latencies      rp_model.RowReport
	TimeSeries     rp_model.RowReport
}

func (z *Job) Preset() {
	z.Population.SetModel(&dc_log.Population{})
	z.Latencies.SetModel(&dc_log.Latency{})
	z.TimeSeries.SetModel(&dc_log.TimeSeries{})
	z.IntervalSecond = 3600
}

func (z *Job) Exec(c app_control.Control) error {
	if err := z.Population.Open(); err != nil {
		return err
	}
	if err := z.Latencies.Open(); err != nil {
		return err
	}
	if err := z.TimeSeries.Open(); err != nil {
		return err
	}
	if err := os.MkdirAll(c.Workspace().Database(), 0755); err != nil {
		return err
	}
	db, err := c.NewOrm(filepath.Join(c.Workspace().Database(), "capture.db"))
	if err != nil {
		return err
	}
	ca := dc_log.NewCaptureAggregator(db, c,
		dc_log.OptShorten(!z.FullUrl),
		dc_log.OptTimeInterval(z.IntervalSecond),
	)

	var jobId string
	if z.JobId.IsExists() {
		jobId = z.JobId.Value()
	} else {
		histories, err := app_job_impl.NewHistorian(c.Workspace()).Histories()
		if err != nil {
			return err
		}
		last := histories[len(histories)-1]
		jobId = last.JobId()
	}

	if err := ca.AddById(jobId); err != nil {
		return err
	}

	if err = ca.AggregatePopulation(func(r *dc_log.Population) {
		z.Population.Row(r)
	}); err != nil {
		return err
	}

	if err = ca.AggregateLatency(func(r *dc_log.Latency) {
		z.Latencies.Row(r)
	}); err != nil {
		return err
	}

	if err = ca.AggregateTimeSeries(func(r *dc_log.TimeSeries) {
		z.TimeSeries.Row(r)
	}); err != nil {
		return err
	}

	return nil
}

func (z *Job) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Job{}, rc_recipe.NoCustomValues)
	if errors.Is(err, dc_log.ErrorJobNotFound) {
		return nil
	}
	return err
}
