package rp_model_deprecated

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func New(name string, row interface{}, ctl app_control.Control, opts ...rp_model.ReportOpt) (rp_model.SideCarReport, error) {
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}
	reportName := name + ro.ReportSuffix

	reports := make([]rp_model.SideCarReport, 0)
	closeAll := func() {
		for _, r := range reports {
			r.Close()
		}
	}

	if !ctl.IsLowMemory() {
		csv, err := NewCsv(reportName, row, ctl, opts...)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, csv)
	}

	{
		js, err := NewJson(reportName, ctl, opts...)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, js)
	}

	if !ctl.IsLowMemory() {
		xl, err := NewXlsx(reportName, row, ctl, opts...)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, xl)
	}

	if ctl.IsQuiet() {
		// Output as JSON on quiet
		js, err := NewJsonForQuiet(reportName, ctl)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, js)
	} else if !ctl.IsLowMemory() {
		// Output for UI
		ui, err := NewUI(reportName, row, ctl, opts...)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, ui)
	}

	r := &Cascade{
		Ctl:     ctl,
		Reports: reports,
	}
	return r, nil
}

type Cascade struct {
	Ctl      app_control.Control
	Reports  []rp_model.SideCarReport
	isClosed bool
}

func (z *Cascade) Success(input interface{}, result interface{}) {
	for _, r := range z.Reports {
		r.Success(input, result)
	}
}

func (z *Cascade) Failure(err error, input interface{}) {
	for _, r := range z.Reports {
		r.Failure(err, input)
	}
}

func (z *Cascade) Skip(reason app_msg.Message, input interface{}) {
	for _, r := range z.Reports {
		r.Skip(reason, input)
	}
}

func (z *Cascade) Row(row interface{}) {
	if z.isClosed {
		z.Ctl.Log().Error("The report already closed")
	}

	for _, r := range z.Reports {
		r.Row(row)
	}
}

func (z *Cascade) Close() {
	ui := z.Ctl.UI()
	for _, r := range z.Reports {
		r.Close()
	}

	p := z.Ctl.Workspace().Report()
	ui.OpenArtifact(p)
	z.isClosed = true
}
