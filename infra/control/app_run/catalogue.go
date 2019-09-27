package app_run

import (
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/recipe"
	"github.com/watermint/toolbox/recipe/dev"
	"github.com/watermint/toolbox/recipe/file"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	"github.com/watermint/toolbox/recipe/sharedfolder"
	sharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	"github.com/watermint/toolbox/recipe/sharedlink"
	"github.com/watermint/toolbox/recipe/team"
	teamdevice "github.com/watermint/toolbox/recipe/team/device"
	teamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	teamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	teamnamespace "github.com/watermint/toolbox/recipe/team/namespace"
	teamnamespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	teamnamespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	teamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	teamsharedlinkcap "github.com/watermint/toolbox/recipe/team/sharedlink/cap"
	"github.com/watermint/toolbox/recipe/teamfolder"
)

func Recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.Quality{},
		&dev.Dummy{},
		&dev.EndToEnd{},
		&dev.Doc{},
		&dev.Async{},
		&file.List{},
		&group.List{},
		&groupmember.List{},
		&member.Invite{},
		&member.Detach{},
		&member.List{},
		&memberquota.List{},
		&team.Activity{},
		&team.Info{},
		&team.Feature{},
		&team.Diagnosis{},
		&teamdevice.List{},
		&teamfilerequest.List{},
		&teamlinkedapp.List{},
		&teamnamespace.List{},
		&teamnamespacefile.List{},
		&teamnamespacefile.Size{},
		&teamnamespacemember.List{},
		&teamsharedlink.List{},
		&teamsharedlinkcap.Expiry{},
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
