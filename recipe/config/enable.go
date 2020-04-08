package config

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

var (
	ErrorInvalidKey = errors.New("invalid key")
)

type Enable struct {
	Key                        string
	ErrorInvalidKey            app_msg.Message
	ErrorUnableToEnableFeature app_msg.Message
}

func (z *Enable) Preset() {
}

func (z *Enable) Exec(c app_control.Control) error {
	ui := c.UI()
	found := false
	for _, k := range app_control.ConfigKeys {
		if k == z.Key {
			found = true
		}
	}
	if !found {
		ui.Error(z.ErrorInvalidKey.With("Key", z.Key))
		return ErrorInvalidKey
	}
	if err := c.Feature().Config().Put(z.Key, true); err != nil {
		ui.Error(z.ErrorUnableToEnableFeature.With("Key", z.Key))
		return err
	}
	return nil
}

func (z *Enable) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Enable{}, func(r rc_recipe.Recipe) {
		m := r.(*Enable)
		m.Key = app_control.ConfigKeyRecipeConfigEnableTest + "NoExistent"
	}); err != ErrorInvalidKey {
		return ErrorInvalidKey
	}

	if err := rc_exec.Exec(c, &Enable{}, func(r rc_recipe.Recipe) {
		m := r.(*Enable)
		m.Key = app_control.ConfigKeyRecipeConfigEnableTest
	}); err != nil {
		return err
	}
	return nil
}
