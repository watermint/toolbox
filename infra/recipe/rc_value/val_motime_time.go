package rc_value

import (
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoTimeTime(name string) Value {
	v := &ValueMoTimeTime{name: name}
	v.time = mo_time.Zero()
	return v
}

type ValueMoTimeTime struct {
	name     string
	dateTime string
	time     mo_time.Time
}

func (z *ValueMoTimeTime) ValueText() string {
	return z.dateTime
}

func (z *ValueMoTimeTime) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*mo_time.Time)(nil)).Elem()) {
		return newValueMoTimeTime(name)
	}
	return nil
}

func (z *ValueMoTimeTime) Bind() interface{} {
	return &z.dateTime
}

func (z *ValueMoTimeTime) Init() (v interface{}) {
	return z.time
}

func (z *ValueMoTimeTime) Apply(v0 interface{}) (v interface{}) {
	return z.time
}

func (z *ValueMoTimeTime) Debug() interface{} {
	return map[string]string{
		"dateTime": z.dateTime,
		"time":     z.time.String(),
	}
}

func (z *ValueMoTimeTime) SpinUp(ctl app_control.Control) (err error) {
	ti := z.time.(*mo_time.TimeImpl)
	if err = ti.UpdateTime(z.dateTime); err != nil {
		return err
	}
	return nil
}

func (z *ValueMoTimeTime) SpinDown(ctl app_control.Control) error {
	return nil
}