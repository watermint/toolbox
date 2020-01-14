package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"sync"
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
	mutex sync.Mutex
}

func (z *RowReport) Spec() rp_model.Spec {
	return newSpec(z.name, z.model, z.opts)
}

func (z *RowReport) SetCtl(ctl app_control.Control) {
	z.ctl = ctl
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

func (z *RowReport) OpenNew(opts ...rp_model.ReportOpt) (r rp_model.RowReport, err error) {
	r = z.Fork(z.ctl)
	if err := r.Open(opts...); err != nil {
		return nil, err
	}
	return r, nil
}

func (z *RowReport) Open(opts ...rp_model.ReportOpt) error {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.w == nil {
		z.w = newCascade(z.name, z.ctl)
	}
	allOpts := make([]rp_model.ReportOpt, 0)
	allOpts = append(allOpts, z.opts...)
	allOpts = append(allOpts, opts...)
	return z.w.Open(z.ctl, z.model, allOpts...)
}

func (z *RowReport) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.w != nil {
		z.w.Close()
		z.w = nil
	}
}

func (z *RowReport) Row(row interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.w.Row(row)
}

func (z *RowReport) SetModel(row interface{}, opts ...rp_model.ReportOpt) {
	z.model = row
	z.opts = opts
}

type TransactionReport struct {
	name  string
	ctl   app_control.Control
	w     Writer
	model interface{}
	opts  []rp_model.ReportOpt
	mutex sync.Mutex
}

func (z *TransactionReport) Spec() rp_model.Spec {
	return newSpec(z.name, z.model, z.opts)
}

func (z *TransactionReport) SetCtl(ctl app_control.Control) {
	z.ctl = ctl
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

func (z *TransactionReport) OpenNew(opts ...rp_model.ReportOpt) (r rp_model.TransactionReport, err error) {
	r = z.Fork(z.ctl)
	if err := r.Open(opts...); err != nil {
		return nil, err
	}
	return r, nil
}

func (z *TransactionReport) Open(opts ...rp_model.ReportOpt) error {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.w == nil {
		z.w = newCascade(z.name, z.ctl)
	}
	allOpts := make([]rp_model.ReportOpt, 0)
	allOpts = append(allOpts, z.opts...)
	allOpts = append(allOpts, opts...)
	allOpts = append(allOpts, rp_model.HiddenColumns("status_tag"))
	return z.w.Open(z.ctl, z.model, allOpts...)
}

func (z *TransactionReport) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.w != nil {
		z.w.Close()
		z.w = nil
	}
}

func (z *TransactionReport) Success(input interface{}, result interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	ui := z.ctl.UI()
	z.w.Row(&rp_model.TransactionRow{
		Status:    ui.Text(MTransactionReport.Success),
		StatusTag: rp_model.StatusTagSuccess,
		Input:     input,
		Result:    result,
	})
}

func (z *TransactionReport) Failure(err error, input interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	ui := z.ctl.UI()
	reason := api_util.MsgFromError(err)
	if ui.TextOrEmpty(reason) == "" {
		summary := api_util.ErrorSummary(err)
		if summary == "" {
			summary = err.Error()
		}
		reason = app_msg.M("dbx.err.general_error", app_msg.P{"Error": summary})
	}
	z.w.Row(&rp_model.TransactionRow{
		Status:    ui.Text(MTransactionReport.Failure),
		StatusTag: rp_model.StatusTagFailure,
		Reason:    ui.Text(reason),
		Input:     input,
		Result:    nil,
	})
}

func (z *TransactionReport) Skip(reason app_msg.Message, input interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	ui := z.ctl.UI()
	z.w.Row(&rp_model.TransactionRow{
		Status:    ui.Text(MTransactionReport.Skip),
		StatusTag: rp_model.StatusTagSkip,
		Reason:    ui.Text(reason),
		Input:     input,
		Result:    nil,
	})
}

func (z *TransactionReport) SetModel(input interface{}, result interface{}, opts ...rp_model.ReportOpt) {
	z.model = &rp_model.TransactionRow{Input: input, Result: result}
	z.opts = opts
}
