package rc_value

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"reflect"
)

func newValueRpModelTransactionReport(name string) rc_recipe.Value {
	n := strcase.ToSnake(name)
	v := &ValueRpModelTransactionReport{name: n}
	v.rep = rp_model_impl.NewTransactionReport(n)
	return v
}

type ValueRpModelTransactionReport struct {
	name string
	rep  *rp_model_impl.TransactionReport
}

func (z *ValueRpModelTransactionReport) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*rp_model.TransactionReport)(nil)).Elem()) {
		return newValueRpModelTransactionReport(name)
	}
	return nil
}

func (z *ValueRpModelTransactionReport) Bind() interface{} {
	return nil
}

func (z *ValueRpModelTransactionReport) Init() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelTransactionReport) ApplyPreset(v0 interface{}) {
	z.rep = v0.(*rp_model_impl.TransactionReport)
}

func (z *ValueRpModelTransactionReport) Apply() (v interface{}) {
	return z.rep
}

func (z *ValueRpModelTransactionReport) SpinUp(ctl app_control.Control) error {
	// Report will not automatically open
	z.rep.SetCtl(ctl)
	return nil
}

func (z *ValueRpModelTransactionReport) SpinDown(ctl app_control.Control) error {
	z.rep.Close()
	return nil
}

func (z *ValueRpModelTransactionReport) Debug() interface{} {
	return nil
}

func (z *ValueRpModelTransactionReport) Report() (report rp_model.Report, valid bool) {
	return z.rep, true
}
