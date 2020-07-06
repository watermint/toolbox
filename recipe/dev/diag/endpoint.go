package diag

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type EndpointRow struct {
	Endpoint     string `json:"endpoint"`
	Count        int    `json:"count"`
	CountSuccess int    `json:"count_success"`
	CountFailure int    `json:"count_failure"`
}

type Endpoint struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkConsole
	JobId  mo_string.OptionalString
	Path   mo_string.OptionalString
	Report rp_model.RowReport

	stats map[string]*EndpointRow
}

func (z *Endpoint) Preset() {
	z.Report.SetModel(&EndpointRow{})
}

func (z *Endpoint) handleRecord(history app_job.History, rec nw_capture.Record) {
	if stat, ok := z.stats[rec.Req.RequestUrl]; ok {
		stat.Count++
		if rec.IsSuccess() {
			stat.CountSuccess++
		} else {
			stat.CountFailure++
		}
	} else {
		if rec.IsSuccess() {
			z.stats[rec.Req.RequestUrl] = &EndpointRow{
				Endpoint:     rec.Req.RequestUrl,
				Count:        1,
				CountSuccess: 1,
				CountFailure: 0,
			}
		} else {
			z.stats[rec.Req.RequestUrl] = &EndpointRow{
				Endpoint:     rec.Req.RequestUrl,
				Count:        1,
				CountSuccess: 0,
				CountFailure: 1,
			}
		}
	}
}

func (z *Endpoint) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Report.Open(); err != nil {
		return err
	}

	z.stats = make(map[string]*EndpointRow)

	loader := CaptureLoader{
		Ctl:   c,
		JobId: z.JobId,
		Path:  z.Path,
	}
	err := loader.Load(z.handleRecord)
	if err != nil {
		l.Debug("Unable to load", esl.Error(err))
		return err
	}

	for _, stat := range z.stats {
		z.Report.Row(stat)
	}

	return nil
}

func (z *Endpoint) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Endpoint{}, rc_recipe.NoCustomValues)
}
