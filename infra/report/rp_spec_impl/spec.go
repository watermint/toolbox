package rp_spec_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func New(recipe app_recipe.Recipe, ctl app_control.Control) *Specs {
	specs := make(map[string]rp_spec.ReportSpec)
	for _, rs := range recipe.Reports() {
		specs[rs.Name()] = rs
	}

	s := &Specs{
		recipe: recipe,
		ctl:    ctl,
		specs:  specs,
	}
	return s
}

type Specs struct {
	recipe app_recipe.Recipe
	ctl    app_control.Control
	specs  map[string]rp_spec.ReportSpec
}

func (z *Specs) Open(name string, opt ...rp_model.ReportOpt) (rep rp_model.Report, err error) {
	if rs, ok := z.specs[name]; ok {
		opts := make([]rp_model.ReportOpt, 0)
		if rs.Options() != nil {
			opts = append(opts, rs.Options()...)
		}
		if opt != nil {
			opts = append(opts, opt...)
		}
		return rp_model_impl.New(rs.Name(), rs.Row(), z.ctl, opts...)
	}

	return nil, errors.New("report specification not found")
}

func (z *Specs) Spec(name string) rp_spec.ReportSpec {
	if rs, ok := z.specs[name]; ok {
		return rs
	} else {
		return &EmptySpec{}
	}
}

func Spec(name string, row interface{}, opt ...rp_model.ReportOpt) rp_spec.ReportSpec {
	return &ReportSpec{
		name: name,
		row:  row,
		opts: opt,
	}
}

type EmptySpec struct {
}

func (e EmptySpec) Name() string {
	return ""
}

func (e EmptySpec) Row() interface{} {
	return struct{}{}
}

func (e EmptySpec) Desc() app_msg.Message {
	return app_msg.M("report.empty_spec.desc")
}

func (e EmptySpec) Columns() []string {
	return []string{}
}

func (e EmptySpec) ColumnDesc(col string) app_msg.Message {
	return app_msg.M("report.empty_spec.desc_column", app_msg.P{"Column": col})
}

func (e EmptySpec) Options() []rp_model.ReportOpt {
	return []rp_model.ReportOpt{}
}

func (e EmptySpec) Open(opts ...rp_model.ReportOpt) (rp_model.Report, error) {
	return nil, errors.New("no report spec")
}

type ReportSpec struct {
	name string
	row  interface{}
	opts []rp_model.ReportOpt
}

func (z *ReportSpec) Open(opts ...rp_model.ReportOpt) (rp_model.Report, error) {
	return nil, errors.New("not enough resource")
}

func (z *ReportSpec) Options() []rp_model.ReportOpt {
	return z.opts
}

func (z *ReportSpec) Name() string {
	return z.name
}

func (z *ReportSpec) Row() interface{} {
	return z.row
}

func (z *ReportSpec) Desc() app_msg.Message {
	key := ut_reflect.Key(app.Pkg, z.row)
	return app_msg.M(key + ".desc")
}

func (z *ReportSpec) Columns() []string {
	panic("implement me")
}

func (z *ReportSpec) ColumnDesc(col string) app_msg.Message {
	panic("implement me")
}

type ReportSpecWithControl struct {
	spec rp_spec.ReportSpec
	ctl  app_control.Control
}

func (z *ReportSpecWithControl) Name() string {
	return z.spec.Name()
}

func (z *ReportSpecWithControl) Row() interface{} {
	return z.spec.Row()
}

func (z *ReportSpecWithControl) Desc() app_msg.Message {
	return z.spec.Desc()
}

func (z *ReportSpecWithControl) Columns() []string {
	return z.spec.Columns()
}

func (z *ReportSpecWithControl) ColumnDesc(col string) app_msg.Message {
	return z.spec.ColumnDesc(col)
}

func (z *ReportSpecWithControl) Options() []rp_model.ReportOpt {
	return z.spec.Options()
}

func (z *ReportSpecWithControl) Open(opts ...rp_model.ReportOpt) (rp_model.Report, error) {
	ros := make([]rp_model.ReportOpt, 0)
	if z.spec.Options() != nil {
		ros = append(ros, z.spec.Options()...)
	}
	if opts != nil {
		ros = append(ros, opts...)
	}
	return rp_model_impl.New(z.Name(), z.Row(), z.ctl, opts...)
}
