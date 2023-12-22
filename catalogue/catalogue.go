package catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
)

func NewCatalogue() rc_catalogue.Catalogue {
	var recipes = make([]rc_recipe.Recipe, 0)
	recipes = append(recipes, AutoDetectedRecipesClassic()...)
	for _, r := range AutoDetectedRecipesCitron() {
		recipes = append(recipes, r)
		rs := rc_spec.New(r)
		app_definitions.SecretRecipeCliPaths = append(app_definitions.SecretRecipeCliPaths, rs.CliPath())
	}

	return rc_catalogue_impl.NewCatalogue(
		recipes,
		AutoDetectedIngredients(),
		AutoDetectedMessageObjects(),
		AutoDetectedFeatures(),
	)
}
