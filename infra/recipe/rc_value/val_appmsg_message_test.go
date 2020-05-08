package rc_value

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueAppMsgMessageRecipe struct {
	Hello app_msg.Message
}

func (z *ValueAppMsgMessageRecipe) Preset() {
}

func (z *ValueAppMsgMessageRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueAppMsgMessageRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueAppMsgMessage(t *testing.T) {
	key := "infra.recipe.rc_value.value_app_msg_message_recipe.hello"

	rcp := &ValueAppMsgMessageRecipe{}
	repo := NewRepository(rcp)
	rcp1 := repo.Apply()
	mod1 := rcp1.(*ValueAppMsgMessageRecipe)
	if mod1.Hello == nil {
		t.Error(mod1)
	}
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueAppMsgMessageRecipe)
		if mod2.Hello == nil || mod2.Hello.Key() != key {
			t.Error(mod2)
		}
		c.UI().Info(mod2.Hello)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
