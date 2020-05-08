package rc_catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Catalogue interface {
	Recipes() []rc_recipe.Recipe
	Ingredients() []rc_recipe.Recipe
	Messages() []interface{}
	RootGroup() rc_group.Group
	Features() []app_feature.OptIn
}
