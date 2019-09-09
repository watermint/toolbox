package app_report

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewUI(name string, row interface{}, ctl app_control.Control) (Report, error) {
	parser := NewColumn(row, ctl)
	r := &UI{
		Ctl:    ctl,
		Table:  ctl.UI().InfoTable(name),
		Parser: parser,
	}
	return r, nil
}

type UI struct {
	Ctl    app_control.Control
	Table  app_ui.Table
	Parser Column
	Index  int
}

func (z *UI) Success(input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgSuccess.Key(), msgSuccess.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *UI) Failure(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgFailure.Key(), msgFailure.Params()...),
		Reason: z.Ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *UI) Skip(reason app_msg.Message, input interface{}, result interface{}) {
	z.Row(TransactionRow{
		Status: z.Ctl.UI().Text(msgSkip.Key(), msgFailure.Params()...),
		Reason: z.Ctl.UI().Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: result,
	})
}

func (z *UI) Row(row interface{}) {
	if z.Index == 0 {
		z.Table.HeaderRaw(z.Parser.Header()...)
	}
	z.Table.RowRaw(z.Parser.ValuesAsString(row)...)
	z.Index++
}

func (z *UI) Flush() {
	z.Table.Flush()
}

func (z *UI) Close() {
	z.Table.Flush()
}
