package rc_value

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
	"strconv"
	"strings"
)

func newValueSelectString() rc_recipe.Value {
	return &ValueMoStringSelect{
		selStr: mo_string.NewSelect(),
	}
}

type ValueMoStringSelect struct {
	selStr mo_string.SelectString
	valStr string
}

func (z *ValueMoStringSelect) ValueText() string {
	return z.valStr
}

func (z *ValueMoStringSelect) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.selStr), map[string]interface{}{
		"options": z.selStr.Options(),
	}
}

func (z *ValueMoStringSelect) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_string.SelectString)(nil)).Elem()) {
		return newValueSelectString()
	}
	return nil
}

func (z *ValueMoStringSelect) Bind() interface{} {
	return &z.valStr
}

func (z *ValueMoStringSelect) Init() (v interface{}) {
	return z.selStr
}

func (z *ValueMoStringSelect) ApplyPreset(v0 interface{}) {
	z.selStr = v0.(mo_string.SelectString)
	z.valStr = z.selStr.Value()
}

func (z *ValueMoStringSelect) Apply() (v interface{}) {
	z.selStr.SetSelect(z.valStr)
	return z.selStr
}

func (z *ValueMoStringSelect) Debug() interface{} {
	return map[string]string{
		"selected": z.selStr.Value(),
		"is_valid": strconv.FormatBool(z.selStr.IsValid()),
		"opts":     strings.Join(z.selStr.Options(), ","),
	}
}

func (z *ValueMoStringSelect) SpinUp(ctl app_control.Control) error {
	ui := ctl.UI()

	if !z.selStr.IsValid() {
		ui.Error(MRepository.ErrorMoStringSelectInvalidOption.
			With("Selected", z.selStr.Value()).
			With("Options", strings.Join(z.selStr.Options(), ", ")))
		return ErrorInvalidValue
	}
	return nil
}

func (z *ValueMoStringSelect) SpinDown(ctl app_control.Control) error {
	return nil
}
