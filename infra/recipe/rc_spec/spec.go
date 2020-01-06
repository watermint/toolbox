package rc_spec

import (
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

func New(rcp rc_recipe.Recipe) rc_recipe.Spec {
	switch scr := rcp.(type) {
	case rc_recipe.SelfContainedRecipe:
		return newSelfContained(scr)

	default:
		return nil
	}
}
