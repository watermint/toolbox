package app_run

import (
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_recipe_group"
	"github.com/watermint/toolbox/experimental/recipe"
	"github.com/watermint/toolbox/experimental/recipe/dev"
	"github.com/watermint/toolbox/experimental/recipe/group"
	"github.com/watermint/toolbox/experimental/recipe/member"
	"github.com/watermint/toolbox/experimental/recipe/member/update"
	"github.com/watermint/toolbox/experimental/recipe/teamfolder"
)

func Recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.LongRunning{},
		&dev.Dummy{},
		&group.List{},
		&member.Invite{},
		&member.List{},
		&update.Email{},
		&teamfolder.List{},
		&recipe.Web{},
	}
}

func Catalogue() *app_recipe_group.Group {
	root := app_recipe_group.NewGroup([]string{}, "")
	for _, r := range Recipes() {
		root.Add(r)
	}
	return root
}
