package app_report

import (
	"github.com/watermint/toolbox/app86/app_control"
	"github.com/watermint/toolbox/app86/app_ui"
)

func NewUI(name string, row interface{}, ctl app_control.Control) (Report, error) {
	parser := NewColumn(row, ctl)
	r := &UI{
		Ctl:    ctl,
		Table:  ctl.UI().InfoTable(false),
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
