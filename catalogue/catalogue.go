package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
)

func NewCatalogue() rc_catalogue.Catalogue {
	return rc_catalogue.NewCatalogue(
		AutoDetectedRecipes(),
		AutoDetectedIngredients(),
		AutoDetectedMessageObjects(),
		AutoDetectedFeatures(),
	)
}
