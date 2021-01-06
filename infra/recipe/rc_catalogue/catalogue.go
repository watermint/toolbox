package rc_catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Catalogue interface {
	// All recipes
	Recipes() []rc_recipe.Recipe

	// Recipes by cliPath. Returns nil if no recipe registered for cliPath.
	Recipe(cliPath string) (recipe rc_recipe.Recipe, spec rc_recipe.Spec)

	// Recipe spec by cliPath. Panic when the spec not found.
	RecipeSpec(cliPath string) rc_recipe.Spec

	// All ingredients
	Ingredients() []rc_recipe.Recipe

	// All messages
	Messages() []interface{}

	// Root recipe group
	RootGroup() rc_group.Group

	// Features
	Features() []app_feature.OptIn
}

type RecipeNotFound struct {
	Path string
}

func (z RecipeNotFound) Error() string {
	return "recipe not found for the path : " + z.Path
}
