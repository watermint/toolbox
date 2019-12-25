package rc_exec

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
)

func Exec(ctl app_control.Control, r rc_recipe.Recipe, custom func(r rc_recipe.Recipe)) error {
	spec := rc_spec.New(r)
	if spec == nil {
		return errors.New("no spec found")
	}
	scr, _, err := spec.ApplyValues(ctl, custom)
	if err != nil {
		return err
	}
	custom(scr)
	return scr.Exec(rc_kitchen.NewKitchen(ctl, scr))
}
