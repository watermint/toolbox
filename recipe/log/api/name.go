package api

import (
	"errors"
	"github.com/watermint/toolbox/domain/core/dc_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"os"
	"path/filepath"
)

type Name struct {
	Name           string
	FullUrl        bool
	IntervalSecond int
	Population     rp_model.RowReport
	Latencies      rp_model.RowReport
	TimeSeries     rp_model.RowReport
}

func (z *Name) Preset() {
	z.Population.SetModel(&dc_log.Population{})
	z.Latencies.SetModel(&dc_log.Latency{})
	z.TimeSeries.SetModel(&dc_log.TimeSeries{})
	z.IntervalSecond = 3600
}

func (z *Name) Exec(c app_control.Control) error {
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

	if err := ca.AddByCliPath(z.Name); err != nil {
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

func (z *Name) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Name{}, rc_recipe.NoCustomValues)
	if errors.Is(err, dc_log.ErrorJobNotFound) {
		return nil
	}
	return err
}
