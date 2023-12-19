package rc_catalogue_impl

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
)

type catalogueImpl struct {
	recipes     []rc_recipe.Recipe
	ingredients []rc_recipe.Recipe
	messages    []interface{}
	features    []app_feature.OptIn
	root        rc_group.Group
}

func (z *catalogueImpl) Recipe(cliPath string) (recipe rc_recipe.Recipe, spec rc_recipe.Spec) {
	for _, r := range z.recipes {
		rs := rc_spec.New(r)
		if rs.CliPath() == cliPath {
			return r, rs
		}
	}
	return nil, nil
}

func (z *catalogueImpl) RecipeSpec(cliPath string) rc_recipe.Spec {
	_, spec := z.Recipe(cliPath)
	if spec == nil {
		panic(&rc_catalogue.RecipeNotFound{Path: cliPath})
	}
	return spec
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

func NewEmptyCatalogue() rc_catalogue.Catalogue {
	return &catalogueImpl{
		recipes:     []rc_recipe.Recipe{},
		ingredients: []rc_recipe.Recipe{},
		messages:    []interface{}{},
		features:    []app_feature.OptIn{},
		root:        rc_group_impl.NewGroup(),
	}
}

// RecipeAliasMap maps cli path to another cli path.
type RecipeAliasMap struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type CatalogueOpts struct {
	Aliases []RecipeAliasMap
}

func NewCatalogue(recipes, ingredients []rc_recipe.Recipe, messages []interface{}, features []app_feature.OptIn) rc_catalogue.Catalogue {
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
