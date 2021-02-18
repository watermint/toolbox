package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
	"strconv"
)

func newValueRangeInt() rc_recipe.Value {
	return &ValueMoIntRange{
		rangeInt: mo_int.NewRange(),
	}
}

type ValueMoIntRange struct {
	rangeInt mo_int.RangeInt
	valInt   int64
}

func (z *ValueMoIntRange) ValueText() string {
	return strconv.FormatInt(z.rangeInt.Value64(), 10)
}

func (z *ValueMoIntRange) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_int.RangeInt)(nil)).Elem()) {
		return newValueRangeInt()
	}
	return nil
}

func (z *ValueMoIntRange) Bind() interface{} {
	return &z.valInt
}

func (z *ValueMoIntRange) Init() (v interface{}) {
	return z.rangeInt
}

func (z *ValueMoIntRange) ApplyPreset(v0 interface{}) {
	z.rangeInt = v0.(mo_int.RangeInt)
	z.valInt = z.rangeInt.Value64()
}

func (z *ValueMoIntRange) Apply() (v interface{}) {
	z.rangeInt.SetValue(z.valInt)
	return z.rangeInt
}

func (z *ValueMoIntRange) Debug() interface{} {
	min, max := z.rangeInt.Range()
	return map[string]interface{}{
		"min": min,
		"max": max,
		"val": z.rangeInt.Value64(),
	}
}

func (z *ValueMoIntRange) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.valInt, nil
}

func (z *ValueMoIntRange) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.Number(); found {
		z.valInt = w.Int64()
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueMoIntRange) SpinUp(ctl app_control.Control) error {
	ui := ctl.UI()

	if !z.rangeInt.IsValid() {
		min, max := z.rangeInt.Range()
		ui.Error(MRepository.ErrorMoIntRangeInvalidRange.
			With("Min", min).With("Max", max).With("Value", z.rangeInt.Value()))
		return ErrorInvalidValue
	}
	return nil
}

func (z *ValueMoIntRange) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoIntRange) Spec() (typeName string, typeAttr interface{}) {
	min, max := z.rangeInt.Range()
	return es_reflect.Key(app.Pkg, z.rangeInt), map[string]interface{}{
		"min":   min,
		"max":   max,
		"value": z.rangeInt.Value(),
	}
}
