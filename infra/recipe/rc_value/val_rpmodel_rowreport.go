package rc_value

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"reflect"
)

func newValueRpModelRowReport(name string) rc_recipe.Value {
	n := strcase.ToSnake(name)
	v := &ValueRpModelRowReport{name: n}
	v.rep = rp_model_impl.NewRowReport(n)
	return v
}

type ValueRpModelRowReport struct {
	name string
	rep  *rp_model_impl.RowReport
}

func (z *ValueRpModelRowReport) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.rep), nil
}

func (z *ValueRpModelRowReport) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*rp_model.RowReport)(nil)).Elem()) {
		return newValueRpModelRowReport(name)
	}
	return nil
}

func (z *ValueRpModelRowReport) Bind() interface{} {
	return nil
}

func (z *ValueRpModelRowReport) Init() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelRowReport) ApplyPreset(v0 interface{}) {
	z.rep = v0.(*rp_model_impl.RowReport)
}

func (z *ValueRpModelRowReport) Apply() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelRowReport) SpinUp(ctl app_control.Control) error {
	// Report will not automatically open
	z.rep.SetCtl(ctl)
	return nil
}

func (z *ValueRpModelRowReport) SpinDown(ctl app_control.Control) error {
	z.rep.Close()
	return nil
}

func (z *ValueRpModelRowReport) Debug() interface{} {
	return nil
}

func (z *ValueRpModelRowReport) Report() (report rp_model.Report, valid bool) {
	return z.rep, true
}
