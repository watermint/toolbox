package rp_column_impl

import (
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

func NewStream(model interface{}, opts ...rp_model.ReportOpt) rp_column.Column {
	ro := &rp_model.ReportOpts{}
	for _, opt := range opts {
		opt(ro)
	}
	ri := &columnStream{
		opts: ro,
	}
	_ = ri.Parse(model)

	return ri
}

type columnStream struct {
	header []string
	opts   *rp_model.ReportOpts
}

func (z *columnStream) Parse(r interface{}) (err error) {
	z.header, err = Headers(r, z.opts.IsHiddenColumn)
	return err
}

func (z *columnStream) Header() []string {
	return z.header
}

func (z *columnStream) Values(r interface{}) (cols []interface{}) {
	if s, err := Parse(r); err != nil {
		return
	} else {
		cols = make([]interface{}, 0)
		for _, p := range z.header {
			cols = append(cols, s.Get(p).Value())
		}
		return cols
	}
}

func (z *columnStream) ValueStrings(r interface{}) (cols []string) {
	if s, err := Parse(r); err != nil {
		return
	} else {
		cols = make([]string, 0)
		for _, p := range z.header {
			cols = append(cols, s.Get(p).String())
		}
		return cols
	}
}
