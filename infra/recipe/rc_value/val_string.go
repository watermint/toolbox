package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueString() Value {
	return &ValueString{}
}

type ValueString struct {
	v string
}

func (z *ValueString) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Kind() == reflect.String {
		return newValueString()
	}
	return nil
}

func (z *ValueString) Fork(ctl app_control.Control) Value {
	v := &ValueString{}
	v.v = z.v
	return v
}

func (z *ValueString) Bind() interface{} {
	return &z.v
}

func (z *ValueString) Init() (v interface{}) {
	return z.v
}

func (z *ValueString) ApplyPreset(v0 interface{}) {
	z.v = v0.(string)
}

func (z *ValueString) Apply() (v interface{}) {
	return z.v
}

func (z *ValueString) Debug() interface{} {
	return z.v
}

func (z *ValueString) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueString) SpinDown(ctl app_control.Control) error {
	return nil
}
