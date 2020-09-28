package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_essential"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoTimeTime(name string) rc_recipe.Value {
	v := &ValueMoTimeTime{name: name}
	v.time = mo_time.Zero()
	v.isOptional = false
	return v
}

func newValueMoTimeTimeOptional(name string) rc_recipe.Value {
	v := &ValueMoTimeTime{name: name}
	v.time = mo_time.Zero()
	v.isOptional = true
	return v
}

type ValueMoTimeTime struct {
	name       string
	dateTime   string
	time       mo_time.Time
	isOptional bool
}

func (z *ValueMoTimeTime) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.time), map[string]interface{}{
		"optional": z.isOptional,
	}
}

func (z *ValueMoTimeTime) IsOptional() bool {
	return z.isOptional
}

func (z *ValueMoTimeTime) ValueText() string {
	return z.dateTime
}

func (z *ValueMoTimeTime) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_time.TimeOptional)(nil)).Elem()) {
		return newValueMoTimeTimeOptional(name)
	}
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

func (z *ValueMoTimeTime) ApplyPreset(v0 interface{}) {
	z.time = v0.(mo_time.Time)
	if !z.time.IsZero() {
		z.dateTime = z.time.Value()
	}
}

func (z *ValueMoTimeTime) Apply() (v interface{}) {
	return z.time
}

func (z *ValueMoTimeTime) Debug() interface{} {
	return map[string]string{
		"dateTime": z.dateTime,
		"time":     z.time.Value(),
	}
}

func (z *ValueMoTimeTime) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.dateTime, nil
}

func (z *ValueMoTimeTime) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.String(); found {
		z.dateTime = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueMoTimeTime) SpinUp(ctl app_control.Control) (err error) {
	l := ctl.Log()

	// argument was't given, but applied on preset or custom value
	if z.dateTime == "" && !z.time.IsZero() {
		return nil
	}

	// optional case
	if z.dateTime == "" && z.isOptional {
		// mark as unset
		if to, ok := z.time.(mo_essential.OptionalMutable); ok {
			to.Unset()
		}
		return nil
	}

	ti := z.time.(*mo_time.TimeImpl)
	if err = ti.UpdateTime(z.dateTime); err != nil {
		l.Debug("Unable to parse", esl.Error(err), esl.String("dateTime", z.dateTime))
		return err
	}
	return nil
}

func (z *ValueMoTimeTime) SpinDown(ctl app_control.Control) error {
	return nil
}
