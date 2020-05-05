package rc_catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
)

type Catalogue interface {
	Recipes() []rc_recipe.Recipe
	Ingredients() []rc_recipe.Recipe
	Messages() []interface{}
	RootGroup() rc_group.Group
	Features() []app_feature.OptIn
}

type catalogueImpl struct {
	recipes     []rc_recipe.Recipe
	ingredients []rc_recipe.Recipe
	messages    []interface{}
	features    []app_feature.OptIn
	root        rc_group.Group
}

func (z *catalogueImpl) Features() []app_feature.OptIn {
	return z.features
}

func (z *catalogueImpl) Recipes() []rc_recipe.Recipe {
	return z.recipes
}

func (z *catalogueImpl) Ingredients() []rc_recipe.Recipe {
	return z.ingredients
}

func (z *catalogueImpl) Messages() []interface{} {
	return z.messages
}

func (z *catalogueImpl) RootGroup() rc_group.Group {
	return z.root
}

func NewEmptyCatalogue() Catalogue {
	return &catalogueImpl{
		recipes:     []rc_recipe.Recipe{},
		ingredients: []rc_recipe.Recipe{},
		messages:    []interface{}{},
		features:    []app_feature.OptIn{},
		root:        rc_group_impl.NewGroup(),
	}
}

func NewCatalogue(recipes, ingredients []rc_recipe.Recipe, messages []interface{}, features []app_feature.OptIn) Catalogue {
	root := rc_group_impl.NewGroup()
	for _, r := range recipes {
		s := rc_spec.New(r)
		root.Add(s)
	}

	return &catalogueImpl{
		recipes:     recipes,
		ingredients: ingredients,
		messages:    messages,
		root:        root,
		features:    features,
	}
}
