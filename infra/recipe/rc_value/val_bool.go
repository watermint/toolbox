package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueBool() rc_recipe.Value {
	return &ValueBool{}
}

type ValueBool struct {
	v bool
}

func (z *ValueBool) Spec() (typeName string, typeAttr interface{}) {
	return "bool", nil
}

func (z *ValueBool) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Kind() == reflect.Bool {
		return newValueBool()
	}
	return nil
}

func (z *ValueBool) Bind() interface{} {
	return &z.v
}

func (z *ValueBool) Init() (v interface{}) {
	return z.v
}

func (z *ValueBool) ApplyPreset(v0 interface{}) {
	z.v = v0.(bool)
}

func (z *ValueBool) Apply() (v interface{}) {
	return z.v
}

func (z *ValueBool) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.v, nil
}

func (z *ValueBool) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.Bool(); found {
		z.v = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueBool) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueBool) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueBool) Debug() interface{} {
	return z.v
}
