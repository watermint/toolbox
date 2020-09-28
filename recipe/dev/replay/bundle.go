package replay

import (
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Bundle struct {
	rc_recipe.RemarkSecret
	ReplayPath mo_string.OptionalString
}

func (z *Bundle) Preset() {
}

func (z *Bundle) Exec(c app_control.Control) error {
	panic("implement me")
}

func (z *Bundle) Test(c app_control.Control) error {
	panic("implement me")
}
