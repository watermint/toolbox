package rp_column_impl

import (
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

func NewModel(model interface{}, opts ...rp_model.ReportOpt) rp_column.Column {
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}
	if ro.ColumnModel != nil {
		return ro.ColumnModel
	} else {
		return NewStream(model, opts...)
	}
}
