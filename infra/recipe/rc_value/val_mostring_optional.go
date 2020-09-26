package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
	"strconv"
)

func newValueOptionalString() rc_recipe.Value {
	return &ValueMoStringOptional{
		optStr: mo_string.NewOptional(""),
	}
}

type ValueMoStringOptional struct {
	optStr mo_string.OptionalString
	valStr string
}

func (z *ValueMoStringOptional) ValueText() string {
	return z.valStr
}

func (z *ValueMoStringOptional) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.optStr), nil
}

func (z *ValueMoStringOptional) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_string.OptionalString)(nil)).Elem()) {
		return newValueOptionalString()
	}
	return nil
}

func (z *ValueMoStringOptional) Bind() interface{} {
	return &z.valStr
}

func (z *ValueMoStringOptional) Init() (v interface{}) {
	return z.optStr
}

func (z *ValueMoStringOptional) ApplyPreset(v0 interface{}) {
	z.optStr = v0.(mo_string.OptionalString)
	z.valStr = z.optStr.Value()
}

func (z *ValueMoStringOptional) Apply() (v interface{}) {
	z.optStr = mo_string.NewOptional(z.valStr)
	return z.optStr
}

func (z *ValueMoStringOptional) Debug() interface{} {
	return map[string]string{
		"str":    z.optStr.Value(),
		"exists": strconv.FormatBool(z.optStr.IsExists()),
	}
}

func (z *ValueMoStringOptional) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.valStr, nil
}

func (z *ValueMoStringOptional) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.String(); found {
		z.valStr = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueMoStringOptional) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoStringOptional) SpinDown(ctl app_control.Control) error {
	return nil
}
