package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Writer interface {
	Name() string
	Row(r interface{})
	Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error
	Close()
}

func NewRowReport(name string) *RowReport {
	return &RowReport{
		name: name,
	}
}

func NewTransactionReport(name string) *TransactionReport {
	return &TransactionReport{
		name: name,
	}
}

type RowReport struct {
	name  string
	ctl   app_control.Control
	w     Writer
	model interface{}
	opts  []rp_model.ReportOpt
}

func (z *RowReport) Spec() rp_model.Spec {
	return newSpec(z.name, z.model, z.opts)
}

func (z *RowReport) Fork(ctl app_control.Control) rp_model.RowReport {
	return &RowReport{
		name:  z.name,
		ctl:   ctl,
		w:     nil, // clear writers on fork
		model: z.model,
		opts:  z.opts,
	}
}

func (z *RowReport) Open(opts ...rp_model.ReportOpt) error {
	if z.w == nil {
		z.w = newCascade(z.name, z.ctl)
	}
	allOpts := make([]rp_model.ReportOpt, 0)
	allOpts = append(allOpts, z.opts...)
	allOpts = append(allOpts, opts...)
	return z.w.Open(z.ctl, z.model, allOpts...)
}

func (z *RowReport) Close() {
	z.w.Close()
}

func (z *RowReport) Row(row interface{}) {
	z.w.Row(row)
}

func (z *RowReport) Model(row interface{}, opts ...rp_model.ReportOpt) {
	z.model = row
	z.opts = opts
}

type TransactionReport struct {
	name  string
	ctl   app_control.Control
	w     Writer
	model interface{}
	opts  []rp_model.ReportOpt
}

func (z *TransactionReport) Spec() rp_model.Spec {
	return newSpec(z.name, z.model, z.opts)
}

func (z *TransactionReport) Fork(ctl app_control.Control) rp_model.TransactionReport {
	return &TransactionReport{
		name:  z.name,
		ctl:   ctl,
		w:     nil, // clear writers on fork
		model: z.model,
		opts:  z.opts,
	}
}

func (z *TransactionReport) Open(opts ...rp_model.ReportOpt) error {
	if z.w == nil {
		z.w = newCascade(z.name, z.ctl)
	}
	allOpts := make([]rp_model.ReportOpt, 0)
	allOpts = append(allOpts, z.opts...)
	allOpts = append(allOpts, opts...)
	return z.w.Open(z.ctl, z.model, allOpts...)
}

func (z *TransactionReport) Close() {
	z.w.Close()
}

func (z *TransactionReport) Success(input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.w.Row(&rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSuccess.Key(), rp_model.MsgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *TransactionReport) Failure(err error, input interface{}) {
	ui := z.ctl.UI()
	reason := api_util.MsgFromError(err)
	if ui.TextOrEmpty(reason.Key()) == "" {
		summary := api_util.ErrorSummary(err)
		if summary == "" {
			summary = err.Error()
		}
		reason = app_msg.M("dbx.err.general_error", app_msg.P{"Error": summary})
	}
	z.w.Row(&rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgFailure.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: nil,
	})
}

func (z *TransactionReport) Skip(reason app_msg.Message, input interface{}) {
	ui := z.ctl.UI()
	z.w.Row(&rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSkip.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: nil,
	})
}

func (z *TransactionReport) Model(input interface{}, result interface{}, opts ...rp_model.ReportOpt) {
	z.model = &rp_model.TransactionRow{Input: input, Result: result}
	z.opts = opts
}
