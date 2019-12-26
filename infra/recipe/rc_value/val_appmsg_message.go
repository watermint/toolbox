package rc_value

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"reflect"
)

func newValueAppMsgMessage(name string, msg app_msg.Message) Value {
	return &ValueAppMsgMessage{name: name, msg: msg}
}

type ValueAppMsgMessage struct {
	name string
	msg  app_msg.Message
}

func (z *ValueAppMsgMessage) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*app_msg.Message)(nil)).Elem()) {
		return newValueAppMsgMessage(name, rc_recipe.RecipeMessage(r, strcase.ToSnake(name)))
	}
	return nil
}

func (z *ValueAppMsgMessage) Bind() interface{} {
	return nil
}

func (z *ValueAppMsgMessage) Init() (v interface{}) {
	return z.msg
}

func (z *ValueAppMsgMessage) Apply(v0 interface{}) (v interface{}) {
	return z.msg
}

func (z *ValueAppMsgMessage) Debug() interface{} {
	return z.name
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
