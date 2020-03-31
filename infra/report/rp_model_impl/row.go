package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"github.com/watermint/toolbox/infra/report/rp_writer_impl"
	"sync"
)

func NewRowReport(name string) *RowReport {
	return &RowReport{
		name: name,
	}
}

type RowReport struct {
	name  string
	ctl   app_control.Control
	w     rp_writer.Writer
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
		z.w = rp_writer_impl.NewCascade(z.name, z.ctl)
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
