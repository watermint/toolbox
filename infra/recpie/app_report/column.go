package app_report

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewColumn(row interface{}, ctl app_control.Control, opts ...ReportOpt) Column {
	ro := &ReportOpts{}
	for _, opt := range opts {
		opt(ro)
	}
	ri := &columnJsonImpl{
		ctl:  ctl,
		opts: ro,
	}
	_ = ri.Parse(row)

	return ri
}

type Column interface {
	Header() []string
	Values(r interface{}) []interface{}
	ValuesAsString(r interface{}) []string
}
