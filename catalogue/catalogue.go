package catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
)

func NewCatalogue() rc_catalogue.Catalogue {
	var recipes []rc_recipe.Recipe
	flavor, found := os.LookupEnv(app_definitions.EnvNameToolboxRecipeFlavor)
	if found {
		switch flavor {
		case app_definitions.RecipeFlavorCitron:
			recipes = AutoDetectedRecipesCitron()
		default:
			recipes = AutoDetectedRecipesClassic()
		}
	} else {
		recipes = AutoDetectedRecipesClassic()
	}

	return rc_catalogue_impl.NewCatalogue(
		recipes,
		AutoDetectedIngredients(),
		AutoDetectedMessageObjects(),
		AutoDetectedFeatures(),
	)
}
