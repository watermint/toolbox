package rp_model_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/report/rp_column_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
	"go.uber.org/zap"
)

type MsgColumnSpec struct {
	TransactionRowStatus app_msg.Message
	TransactionRowReason app_msg.Message
}

var (
	MColumnSpec = app_msg.Apply(&MsgColumnSpec{}).(*MsgColumnSpec)
)

func newSpec(name string, model interface{}, opts []rp_model.ReportOpt) rp_model.Spec {
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}
	cols := make([]string, 0)
	colDesc := make(map[string]app_msg.Message)

	cm := func(m interface{}, base string) []string {
		if m == nil {
			return []string{}
		}
		model := rp_column_impl.NewStream(m, opts...)
		headers := model.Header()
		visibleHeaders := make([]string, 0)
		for _, h := range headers {
			if !ro.IsHiddenColumn(base + h) {
				visibleHeaders = append(visibleHeaders, h)
			}
		}

		keyBase := ut_reflect.Key(app.Pkg, m)
		for _, col := range visibleHeaders {
			colDesc[base+col] = app_msg.M(keyBase + "." + col + ".desc")
		}
		colsWithBase := make([]string, 0)
		for _, c := range visibleHeaders {
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

		colDesc["status"] = MColumnSpec.TransactionRowStatus
		colDesc["reason"] = MColumnSpec.TransactionRowReason

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

func (z *ColumnSpec) Doc(ui app_ui.UI) *rc_doc.Report {
	cols := make([]*rc_doc.ReportColumn, 0)
	for _, col := range z.Columns() {
		cols = append(cols, &rc_doc.ReportColumn{
			Name: col,
			Desc: ui.TextOrEmpty(z.ColumnDesc(col)),
		})
	}
	return &rc_doc.Report{
		Name:    z.Name(),
		Desc:    ui.TextOrEmpty(z.Desc()),
		Columns: cols,
	}
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
	if m, ok := z.colDesc[col]; !ok {
		app_root.Log().Error("Column description not found", zap.String("col", col))
		return app_msg.Raw(col)
	} else {
		return m
	}
}

func (z *ColumnSpec) Options() []rp_model.ReportOpt {
	return z.opts
}
