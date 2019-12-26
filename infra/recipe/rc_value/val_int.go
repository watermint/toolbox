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

func (z *ValueInt) Apply() (v interface{}) {
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
