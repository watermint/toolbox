package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDaGridDataOutput(recipe interface{}, name string) rc_recipe.Value {
	return &ValueDaGridDataOutput{
		gdOutput: da_griddata.NewOutput(recipe, name),
	}
}

type ValueDaGridDataOutput struct {
	gdOutput da_griddata.GridDataOutput
}

func (z *ValueDaGridDataOutput) ValueText() string {
	return ""
}

func (z *ValueDaGridDataOutput) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*da_griddata.GridDataOutput)(nil)).Elem()) {
		return newValueDaGridDataOutput(recipe, name)
	}
	return nil
}

func (z *ValueDaGridDataOutput) Bind() interface{} {
	return z.gdOutput
}

func (z *ValueDaGridDataOutput) Init() (v interface{}) {
	return z.gdOutput
}

func (z *ValueDaGridDataOutput) ApplyPreset(v0 interface{}) {
	z.gdOutput = v0.(da_griddata.GridDataOutput)
}

func (z *ValueDaGridDataOutput) Apply() (v interface{}) {
	return z.gdOutput
}

func (z *ValueDaGridDataOutput) Debug() interface{} {
	return z.gdOutput.Debug()
}

func (z *ValueDaGridDataOutput) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.gdOutput.Capture(), nil
}

func (z *ValueDaGridDataOutput) Restore(v es_json.Json, ctl app_control.Control) error {
	return z.gdOutput.Restore(v)
}

func (z *ValueDaGridDataOutput) SpinUp(ctl app_control.Control) error {
	return z.gdOutput.Open(ctl)
}

func (z *ValueDaGridDataOutput) SpinDown(ctl app_control.Control) error {
	z.gdOutput.Close()
	return nil
}

func (z *ValueDaGridDataOutput) Spec() (typeName string, typeAttr interface{}) {
	return z.gdOutput.Spec().Name(), nil
}

func (z *ValueDaGridDataOutput) GridDataOutput() (gd da_griddata.GridDataOutput, valid bool) {
	return z.gdOutput, true
}
