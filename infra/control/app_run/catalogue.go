package app_run

import (
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/recipe"
	"github.com/watermint/toolbox/recipe/dev"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	"github.com/watermint/toolbox/recipe/sharedfolder"
	sharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	"github.com/watermint/toolbox/recipe/sharedlink"
	"github.com/watermint/toolbox/recipe/team"
	teamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	"github.com/watermint/toolbox/recipe/teamfolder"
)

func Recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.Quality{},
		&dev.Dummy{},
		&group.List{},
		&groupmember.List{},
		&member.Invite{},
		&member.Detach{},
		&member.List{},
		&team.Info{},
		&team.Feature{},
		&teamsharedlink.List{},
		&teamfolder.List{},
		&sharedlink.List{},
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
