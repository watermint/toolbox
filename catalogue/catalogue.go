package catalogue

import (
	"github.com/watermint/toolbox/infra/recpie/rc_group"
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/recipe"
	"github.com/watermint/toolbox/recipe/dev"
	devtest "github.com/watermint/toolbox/recipe/dev/test"
	"github.com/watermint/toolbox/recipe/file"
	filecompare "github.com/watermint/toolbox/recipe/file/compare"
	fileimport "github.com/watermint/toolbox/recipe/file/import"
	fileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	filesync "github.com/watermint/toolbox/recipe/file/sync"
	filesyncpreflight "github.com/watermint/toolbox/recipe/file/sync/preflight"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	memberupdate "github.com/watermint/toolbox/recipe/member/update"
	"github.com/watermint/toolbox/recipe/sharedfolder"
	sharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	"github.com/watermint/toolbox/recipe/sharedlink"
	"github.com/watermint/toolbox/recipe/team"
	teamactivity "github.com/watermint/toolbox/recipe/team/activity"
	teamactivitydaily "github.com/watermint/toolbox/recipe/team/activity/daily"
	teamdevice "github.com/watermint/toolbox/recipe/team/device"
	teamdiag "github.com/watermint/toolbox/recipe/team/diag"
	teamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	teamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	teamnamespace "github.com/watermint/toolbox/recipe/team/namespace"
	teamnamespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	teamnamespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	teamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	teamsharedlinkupdate "github.com/watermint/toolbox/recipe/team/sharedlink/update"
	"github.com/watermint/toolbox/recipe/teamfolder"
)

func Recipes() []rc_recipe.Recipe {
	return []rc_recipe.Recipe{
		&dev.Async{},
		&dev.Doc{},
		&dev.Dummy{},
		&dev.Preflight{},
		&devtest.Auth{},
		&devtest.Recipe{},
		&devtest.Resources{},
		&file.Copy{},
		&file.Delete{},
		&file.Download{},
		&file.List{},
		&file.Merge{},
		&file.Move{},
		&file.Replication{},
		&file.Restore{},
		&file.Upload{},
		&filecompare.Account{},
		&filecompare.Local{},
		&fileimport.Url{},
		&fileimport.ViaApp{},
		&fileimportbatch.Url{},
		&filesync.Up{},
		&filesyncpreflight.Up{},
		&group.List{},
		&group.Delete{},
		&groupmember.List{},
		&member.Detach{},
		&member.Invite{},
		&member.List{},
		&member.Replication{},
		&member.Delete{},
		&memberquota.List{},
		&memberquota.Update{},
		&memberquota.Usage{},
		&memberupdate.Email{},
		&memberupdate.ExternalId{},
		&memberupdate.Profile{},
		&recipe.License{},
		&recipe.Web{},
		&sharedfolder.List{},
		&sharedfoldermember.List{},
		&sharedlink.Create{},
		&sharedlink.List{},
		&sharedlink.Delete{},
		&team.Feature{},
		&team.Info{},
		&teamdiag.Explorer{},
		&teamdiag.Path{},
		&teamactivity.Event{},
		&teamactivity.User{},
		&teamactivitydaily.Event{},
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
		&teamsharedlinkupdate.Expiry{},
	}
}

func Catalogue() *rc_group.Group {
	root := rc_group.NewGroup([]string{}, "")
	for _, r := range Recipes() {
		root.Add(r)
	}
	return root
}
