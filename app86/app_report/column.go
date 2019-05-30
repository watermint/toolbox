package app_report

import (
	"github.com/watermint/toolbox/app86/app_control"
)

func NewColumn(row interface{}, ctl app_control.Control) Column {
	ri := &columnJsonImpl{
		ctl: ctl,
	}
	_ = ri.Parse(row)

	return ri
}

type Column interface {
	Header() []string
	Values(r interface{}) []interface{}
	ValuesAsString(r interface{}) []string
}
