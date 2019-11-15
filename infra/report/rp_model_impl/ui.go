package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sync"
)

func NewUI(name string, row interface{}, ctl app_control.Control, opts ...rp_model.ReportOpt) (rp_model.Report, error) {
	parser := NewColumn(row, opts...)
	ui := ctl.UI()
	r := &UI{
		ctl:    ctl,
		table:  ui.InfoTable(name),
		parser: parser,
	}
	return r, nil
}

type UI struct {
	ctl    app_control.Control
	table  app_ui.Table
	parser Column
	index  int
	mutex  sync.Mutex
}

func (z *UI) Success(input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSuccess.Key(), rp_model.MsgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *UI) Failure(err error, input interface{}) {
	z.Row(rowForFailure(z.ctl.UI(), err, input))
}

func (z *UI) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	ui := z.ctl.UI()
	z.Row(rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgSkip.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *UI) Row(row interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.index == 0 {
		z.table.HeaderRaw(z.parser.Header()...)
	}
	z.table.RowRaw(z.parser.ValuesAsString(row)...)
	z.index++
}

func (z *UI) Close() {
	z.table.Flush()
}
