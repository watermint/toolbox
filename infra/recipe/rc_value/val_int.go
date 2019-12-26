package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueInt() Value {
	return &ValueInt{}
}

type ValueInt struct {
	v int64
}

func (z *ValueInt) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Kind() == reflect.Int {
		return newValueInt()
	}
	return nil
}

func (z *ValueInt) Bind() interface{} {
	return &z.v
}

func (z *ValueInt) Init() (v interface{}) {
	return z.v
}

func (z *ValueInt) Apply(v0 interface{}) (v interface{}) {
	switch v1 := v0.(type) {
	case int:
		z.v = int64(v1)
	case int8:
		z.v = int64(v1)
	case int16:
		z.v = int64(v1)
	case int32:
		z.v = int64(v1)
	case int64:
		z.v = v1
	default:
		// nop
	}
	return z.v
}

func (z *ValueInt) Debug() interface{} {
	return z.v
}

func (z *ValueInt) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueInt) SpinDown(ctl app_control.Control) error {
	return nil
}