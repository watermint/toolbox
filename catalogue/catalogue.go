package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
)

func NewCatalogue() rc_catalogue.Catalogue {
	return rc_catalogue_impl.NewCatalogue(
		AutoDetectedRecipes(),
		AutoDetectedIngredients(),
		AutoDetectedMessageObjects(),
		AutoDetectedFeatures(),
	)
}
