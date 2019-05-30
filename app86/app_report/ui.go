package app_report

import (
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_ui"
)

func NewUI(name string, ctl app_control.Control) (Report, error) {
	r := &UI{
		Ctl:   ctl,
		Table: ctl.UI().InfoTable(false),
	}
	return r, nil
}

type UI struct {
	Ctl    app_control.Control
	Table  app_ui.Table
	Parser Row
}

func (z *UI) Row(row interface{}) {
	if z.Parser == nil {
		z.Parser = NewRow(row, z.Ctl)
		z.Table.HeaderRaw(z.Parser.Header()...)
	}
	z.Table.RowRaw(z.Parser.ValuesAsString(row)...)
}

func (z *UI) Transaction(state State, input interface{}, result interface{}) {
	z.Row(Transaction{State: state(), Input: input, Result: result})
}

func (z *UI) Flush() {
	z.Table.Flush()
}

func (z *UI) Close() {
	z.Table.Flush()
}
