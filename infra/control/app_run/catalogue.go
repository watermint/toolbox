package app_run

import (
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_recipe_group"
	"github.com/watermint/toolbox/recipe"
	"github.com/watermint/toolbox/recipe/dev"
	devtest "github.com/watermint/toolbox/recipe/dev/test"
	"github.com/watermint/toolbox/recipe/file"
	filecompare "github.com/watermint/toolbox/recipe/file/compare"
	fileimport "github.com/watermint/toolbox/recipe/file/import"
	fileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	memberupdate "github.com/watermint/toolbox/recipe/member/update"
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
		&dev.Async{},
		&dev.Doc{},
		&dev.Dummy{},
		&devtest.Auth{},
		&devtest.Recipe{},
		&devtest.Resources{},
		&file.Copy{},
		&file.List{},
		&file.Merge{},
		&file.Move{},
		&file.Replication{},
		&filecompare.Account{},
		&fileimport.Url{},
		&fileimportbatch.Url{},
		&group.List{},
		&group.Remove{},
		&groupmember.List{},
		&member.Detach{},
		&member.Invite{},
		&member.List{},
		&memberquota.List{},
		&memberquota.Usage{},
		&memberupdate.Email{},
		&recipe.License{},
		&recipe.Web{},
		&sharedfolder.List{},
		&sharedfoldermember.List{},
		&sharedlink.Create{},
		&sharedlink.List{},
		&sharedlink.Remove{},
		&team.Activity{},
		&team.Diagnosis{},
		&team.Feature{},
		&team.Info{},
		&teamdevice.List{},
		&teamdevice.Unlink{},
		&teamfilerequest.List{},
		&teamfolder.Archive{},
		&teamfolder.List{},
		&teamfolder.PermDelete{},
		&teamfolder.Replication{},
		&teamlinkedapp.List{},
		&teamnamespace.List{},
		&teamnamespacefile.List{},
		&teamnamespacefile.Size{},
		&teamnamespacemember.List{},
		&teamsharedlink.List{},
		&teamsharedlinkcap.Expiry{},
	}
}

func Catalogue() *app_recipe_group.Group {
	root := app_recipe_group.NewGroup([]string{}, "")
	for _, r := range Recipes() {
		root.Add(r)
	}
	return root
}
