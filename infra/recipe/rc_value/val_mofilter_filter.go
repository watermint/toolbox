package rc_value

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoFilter(name string) rc_recipe.Value {
	return &ValueMoFilterFilter{
		filter: mo_filter.New(name),
	}
}

type ValueMoFilterFilter struct {
	filter mo_filter.Filter
}

func (z *ValueMoFilterFilter) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_filter.Filter)(nil)).Elem()) {
		return newValueMoFilter(name)
	}
	return nil
}

func (z *ValueMoFilterFilter) Bind() interface{} {
	return z.filter
}

func (z *ValueMoFilterFilter) Init() (v interface{}) {
	return z.filter
}

func (z *ValueMoFilterFilter) ApplyPreset(v0 interface{}) {
	z.filter = v0.(mo_filter.Filter)
}

func (z *ValueMoFilterFilter) Apply() (v interface{}) {
	return z.filter
}

func (z *ValueMoFilterFilter) Debug() interface{} {
	return z.filter // TODO: debug info
}

func (z *ValueMoFilterFilter) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoFilterFilter) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoFilterFilter) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.filter), nil /// TODO: type attr
}
