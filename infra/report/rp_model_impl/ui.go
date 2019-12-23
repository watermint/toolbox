package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sync"
)

func newUIWriter(name string, ctl app_control.Control) Writer {
	return &uiWriter{
		name: name,
		ctl:  ctl,
	}
}

type uiWriter struct {
	name     string
	ctl      app_control.Control
	table    app_ui.Table
	colModel Column
	index    int
	mutex    sync.Mutex
}

func (z *uiWriter) Name() string {
	return z.name
}

func (z *uiWriter) Row(r interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.index == 0 {
		z.table.HeaderRaw(z.colModel.Header()...)
	}
	z.table.RowRaw(z.colModel.ValueStrings(r)...)
	z.index++
}

func (z *uiWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	z.ctl = ctl
	z.colModel = NewColumn(model, opts...)
	z.table = ctl.UI().InfoTable(z.name)
	return nil
}

func (z *uiWriter) Close() {
	z.table.Flush()
}
