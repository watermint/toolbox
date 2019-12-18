package rc_spec

import (
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/infra/recpie/rc_vo_impl"
)

func New(rcp rc_recipe.Recipe) rc_recipe.Spec {
	switch scr := rcp.(type) {
	case rc_recipe.SideCarRecipe:
		return newSideCar(scr)

	default:
		return nil
	}
}

func NewCommonValue() (sv rc_recipe.SpecValue, co *app_opt.CommonOpts, vc *rc_vo_impl.ValueContainer) {
	co = app_opt.NewDefaultCommonOpts()
	vc = rc_vo_impl.NewValueContainer(co)
	return newSideCarValue(vc), co, vc
}
