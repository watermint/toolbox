package rc_value

import (
	"reflect"
	"strconv"

	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

func newValueInt() rc_recipe.Value {
	return &ValueInt{}
}

type ValueInt struct {
	v int64
}

func (z *ValueInt) Spec() (typeName string, typeAttr interface{}) {
	return "int", nil
}

func (z *ValueInt) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Kind() == reflect.Int || t.Kind() == reflect.Int64 {
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

func (z *ValueInt) ApplyPreset(v0 interface{}) {
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
}

func (z *ValueInt) Apply() interface{} {
	return z.v
}

func (z *ValueInt) Debug() interface{} {
	return z.v
}

func (z *ValueInt) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.v, nil
}

func (z *ValueInt) Restore(v es_json.Json, ctl app_control.Control) error {
	if numStr, found := v.Number(); found {
		var err error
		z.v, err = strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return rc_recipe.ErrorValueRestoreFailed
		}
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueInt) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueInt) SpinDown(ctl app_control.Control) error {
	return nil
}
