package app_run

import (
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_recipe_group"
	"github.com/watermint/toolbox/experimental/recipe"
	"github.com/watermint/toolbox/experimental/recipe/dev"
	"github.com/watermint/toolbox/experimental/recipe/group"
	groupmember "github.com/watermint/toolbox/experimental/recipe/group/member"
	"github.com/watermint/toolbox/experimental/recipe/member"
	memberupdate "github.com/watermint/toolbox/experimental/recipe/member/update"
	"github.com/watermint/toolbox/experimental/recipe/sharedfolder"
	sharedfoldermember "github.com/watermint/toolbox/experimental/recipe/sharedfolder/member"
	"github.com/watermint/toolbox/experimental/recipe/teamfolder"
)

func Recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.LongRunning{},
		&dev.Dummy{},
		&group.List{},
		&groupmember.List{},
		&member.Invite{},
		&member.List{},
		&memberupdate.Email{},
		&teamfolder.List{},
		&sharedfolder.List{},
		&sharedfoldermember.List{},
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
