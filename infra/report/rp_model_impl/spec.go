package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func newSpec(name string, model interface{}, opts []rp_model.ReportOpt) rp_model.Spec {
	var cols []string
	colDesc := make(map[string]app_msg.Message)

	cm := func(m interface{}, base string) []string {
		if m == nil {
			return []string{}
		}
		model := NewColumn(m, opts...)
		cols = model.Header()
		keyBase := ut_reflect.Key(app.Pkg, m)
		for _, col := range cols {
			colDesc[base+col] = app_msg.M(keyBase + "." + col + ".desc")
		}
		colsWithBase := make([]string, 0)
		for _, c := range cols {
			colsWithBase = append(colsWithBase, base+c)
		}
		return colsWithBase
	}

	switch md := model.(type) {
	case *rp_model.TransactionRow:
		cols = make([]string, 0)
		cols = append(cols, "status")
		cols = append(cols, "reason")
		cols = append(cols, cm(md.Input, "input.")...)
		cols = append(cols, cm(md.Result, "result.")...)

		colDesc["status"] = app_msg.M("infra.report.rp_model.transactionrow.status")
		colDesc["reason"] = app_msg.M("infra.report.rp_model.transactionrow.reason")

	default:
		cols = cm(model, "")
	}

	return &ColumnSpec{
		name:    name,
		model:   model,
		opts:    opts,
		cols:    cols,
		colDesc: colDesc,
	}
}

type ColumnSpec struct {
	name    string
	model   interface{}
	opts    []rp_model.ReportOpt
	cols    []string
	colDesc map[string]app_msg.Message
}

func (z *ColumnSpec) Name() string {
	return z.name
}

func (z *ColumnSpec) Model() interface{} {
	return z.model
}

func (z *ColumnSpec) Desc() app_msg.Message {
	key := ut_reflect.Key(app.Pkg, z.model) + ".desc"
	return app_msg.M(key)
}

func (z *ColumnSpec) Columns() []string {
	return z.cols
}

func (z *ColumnSpec) ColumnDesc(col string) app_msg.Message {
	return z.colDesc[col]
}

func (z *ColumnSpec) Options() []rp_model.ReportOpt {
	return z.opts
}
