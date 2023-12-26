package catalogue

import (
	_ "embed"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

//go:embed catalogue_compatibility.json
var compatibilityDefinitionsData []byte

func NewCatalogue() rc_catalogue.Catalogue {
	var recipes = make([]rc_recipe.Recipe, 0)

	// Load compatibility definitions
	cds, err := rc_compatibility.ParseCompatibilityDefinition(compatibilityDefinitionsData)
	if err != nil {
		panic(err)
	}
	rc_compatibility.Definitions = cds

	// Load recipes
	recipes = append(recipes, AutoDetectedRecipesClassic()...)
	for _, r := range AutoDetectedRecipesCitron() {
		recipes = append(recipes, r)
		//
		//rs := rc_spec.New(r)
		//app_definitions.SecretRecipeCliPaths = append(app_definitions.SecretRecipeCliPaths, rs.CliPath())
	}

	return rc_catalogue_impl.NewCatalogue(
		recipes,
		AutoDetectedIngredients(),
		AutoDetectedMessageObjects(),
		AutoDetectedFeatures(),
	)
}
