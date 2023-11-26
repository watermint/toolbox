package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"reflect"
)

func newValueAppMsgMessage(name string, msg app_msg.Message) rc_recipe.Value {
	return &ValueAppMsgMessage{name: name, msg: msg}
}

type ValueAppMsgMessage struct {
	name string
	msg  app_msg.Message
}

func (z *ValueAppMsgMessage) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.msg), nil
}

func (z *ValueAppMsgMessage) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*app_msg.Message)(nil)).Elem()) {
		return newValueAppMsgMessage(name, app_msg.ObjMessage(v0, ecase.ToLowerSnakeCase(name)))
	}
	return nil
}

func (z *ValueAppMsgMessage) Bind() interface{} {
	return nil
}

func (z *ValueAppMsgMessage) Init() (v interface{}) {
	return z.msg
}

func (z *ValueAppMsgMessage) ApplyPreset(v0 interface{}) {
	z.msg = v0.(app_msg.Message)
}

func (z *ValueAppMsgMessage) Apply() (v interface{}) {
	return z.msg
}

func (z *ValueAppMsgMessage) Debug() interface{} {
	return z.name
}

func (z *ValueAppMsgMessage) Capture(ctl app_control.Control) (v interface{}, err error) {
	return
}

func (z *ValueAppMsgMessage) Restore(v es_json.Json, ctl app_control.Control) error {
	return nil
}

func (z *ValueAppMsgMessage) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueAppMsgMessage) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueAppMsgMessage) Message() (msg app_msg.Message, valid bool) {
	return z.msg, true
}
