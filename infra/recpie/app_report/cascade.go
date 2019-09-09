package app_report

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func New(name string, row interface{}, ctl app_control.Control) (Report, error) {
	reports := make([]Report, 0)
	closeAll := func() {
		for _, r := range reports {
			r.Close()
		}
	}

	{
		csv, err := NewCsv(name, row, ctl)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, csv)
	}

	{
		js, err := NewJson(name, ctl)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, js)
	}

	{
		xl, err := NewXlsx(name, row, ctl)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, xl)
	}

	if ctl.IsQuiet() {
		// Output as JSON on quiet
		js, err := NewJsonForQuiet(name, ctl)
		if err != nil {
			closeAll()
			return nil, err
		}
		reports = append(reports, js)
	} else {
		// Output for UI
		ui, err := NewUI(name, row, ctl)
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
	Ctl     app_control.Control
	Reports []Report
}

func (z *Cascade) Success(input interface{}, result interface{}) {
	for _, r := range z.Reports {
		r.Success(input, result)
	}
}

func (z *Cascade) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	for _, r := range z.Reports {
		r.Failure(reason, input, result)
	}
}

func (z *Cascade) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	for _, r := range z.Reports {
		r.Skip(reason, input, result)
	}
}

func (z *Cascade) Row(row interface{}) {
	for _, r := range z.Reports {
		r.Row(row)
	}
}

func (z *Cascade) Flush() {
	for _, r := range z.Reports {
		r.Flush()
	}
}

func (z *Cascade) Close() {
	for _, r := range z.Reports {
		r.Close()
	}

	p, err := z.Ctl.Workspace().Descendant(reportPath)
	if err != nil {
		return
	}
	z.Ctl.UI().OpenArtifact(p)
}