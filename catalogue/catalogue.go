package catalogue

import (
	infra_api_api_api_auth_impl "github.com/watermint/toolbox/infra/api/api_auth_impl"
	infra_recipe_rc_catalogue "github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	infra_recipe_rc_conn_impl "github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	infra_recipe_rc_group "github.com/watermint/toolbox/infra/recipe/rc_group"
	infra_recipe_rc_group_impl "github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	infra_recipe_rc_spec "github.com/watermint/toolbox/infra/recipe/rc_spec"
	infra_recipe_rcvalue "github.com/watermint/toolbox/infra/recipe/rc_value"
	infra_report_rpmodelimpl "github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	infra_ui_appui "github.com/watermint/toolbox/infra/ui/app_ui"
	infra_util_ut_doc "github.com/watermint/toolbox/infra/util/ut_doc"
	ingredientfile "github.com/watermint/toolbox/ingredient/file"
	ingredientteamnamespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	ingredientteamfolder "github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/recipe"
	recipedev "github.com/watermint/toolbox/recipe/dev"
	recipedevtest "github.com/watermint/toolbox/recipe/dev/test"
	recipefile "github.com/watermint/toolbox/recipe/file"
	recipefilecompare "github.com/watermint/toolbox/recipe/file/compare"
	recipefileexport "github.com/watermint/toolbox/recipe/file/export"
	recipefileimport "github.com/watermint/toolbox/recipe/file/import"
	recipefileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	recipefilesync "github.com/watermint/toolbox/recipe/file/sync"
	recipefilesyncpreflight "github.com/watermint/toolbox/recipe/file/sync/preflight"
	recipegroup "github.com/watermint/toolbox/recipe/group"
	recipegroupbatch "github.com/watermint/toolbox/recipe/group/batch"
	recipegroupmember "github.com/watermint/toolbox/recipe/group/member"
	recipemember "github.com/watermint/toolbox/recipe/member"
	recipememberquota "github.com/watermint/toolbox/recipe/member/quota"
	recipememberupdate "github.com/watermint/toolbox/recipe/member/update"
	recipesharedfolder "github.com/watermint/toolbox/recipe/sharedfolder"
	recipesharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	recipesharedlink "github.com/watermint/toolbox/recipe/sharedlink"
	recipeteam "github.com/watermint/toolbox/recipe/team"
	recipeteamactivity "github.com/watermint/toolbox/recipe/team/activity"
	recipeteamactivitydaily "github.com/watermint/toolbox/recipe/team/activity/daily"
	recipeteamdevice "github.com/watermint/toolbox/recipe/team/device"
	recipeteamdiag "github.com/watermint/toolbox/recipe/team/diag"
	recipeteamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	recipeteamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	recipeteamnamespace "github.com/watermint/toolbox/recipe/team/namespace"
	recipeteamnamespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	recipeteamnamespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	recipeteamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	recipeteamsharedlinkupdate "github.com/watermint/toolbox/recipe/team/sharedlink/update"
	recipeteamfolder "github.com/watermint/toolbox/recipe/teamfolder"
	recipeteamfolderbatch "github.com/watermint/toolbox/recipe/teamfolder/batch"
	recipeteamfolderfile "github.com/watermint/toolbox/recipe/teamfolder/file"
)

func NewCatalogue() infra_recipe_rc_catalogue.Catalogue {
	return infra_recipe_rc_catalogue.NewCatalogue(Recipes(), Ingredients(), Messages())
}

func Recipes() []rc_recipe.Recipe {
	cat := []rc_recipe.Recipe{
		rc_recipe.Annotate(&recipe.License{}),
		rc_recipe.Annotate(&recipedev.Async{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedev.Doc{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedev.Dummy{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedev.Preflight{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedevtest.Auth{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedevtest.Monkey{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedevtest.Recipe{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipedevtest.Resources{}, rc_recipe.Secret()),
		rc_recipe.Annotate(&recipefile.Copy{}),
		rc_recipe.Annotate(&recipefile.Delete{}),
		rc_recipe.Annotate(&recipefile.Download{}, rc_recipe.Experimental()),
		rc_recipe.Annotate(&recipefile.List{}),
		rc_recipe.Annotate(&recipefile.Merge{}),
		rc_recipe.Annotate(&recipefile.Move{}),
		rc_recipe.Annotate(&recipefile.Replication{}),
		rc_recipe.Annotate(&recipefile.Restore{}, rc_recipe.Experimental()),
		rc_recipe.Annotate(&recipefile.Upload{}),
		rc_recipe.Annotate(&recipefile.Watch{}),
		rc_recipe.Annotate(&recipefilecompare.Account{}),
		rc_recipe.Annotate(&recipefilecompare.Local{}),
		rc_recipe.Annotate(&recipefileexport.Doc{}, rc_recipe.Experimental()),
		rc_recipe.Annotate(&recipefileimport.Url{}),
		rc_recipe.Annotate(&recipefileimportbatch.Url{}),
		rc_recipe.Annotate(&recipefilesync.Up{}),
		rc_recipe.Annotate(&recipefilesyncpreflight.Up{}),
		rc_recipe.Annotate(&recipegroup.Delete{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipegroup.List{}),
		rc_recipe.Annotate(&recipegroupbatch.Delete{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipegroupmember.List{}),
		rc_recipe.Annotate(&recipemember.Delete{}),
		rc_recipe.Annotate(&recipemember.Detach{}),
		rc_recipe.Annotate(&recipemember.Invite{}),
		rc_recipe.Annotate(&recipemember.List{}),
		rc_recipe.Annotate(&recipemember.Replication{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipememberquota.List{}),
		rc_recipe.Annotate(&recipememberquota.Update{}),
		rc_recipe.Annotate(&recipememberquota.Usage{}),
		rc_recipe.Annotate(&recipememberupdate.Email{}),
		rc_recipe.Annotate(&recipememberupdate.Externalid{}),
		rc_recipe.Annotate(&recipememberupdate.Profile{}),
		rc_recipe.Annotate(&recipesharedfolder.List{}),
		rc_recipe.Annotate(&recipesharedfoldermember.List{}),
		rc_recipe.Annotate(&recipesharedlink.Create{}),
		rc_recipe.Annotate(&recipesharedlink.Delete{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipesharedlink.List{}),
		rc_recipe.Annotate(&recipeteam.Feature{}),
		rc_recipe.Annotate(&recipeteam.Info{}),
		rc_recipe.Annotate(&recipeteamactivity.Event{}),
		rc_recipe.Annotate(&recipeteamactivity.User{}),
		rc_recipe.Annotate(&recipeteamactivitydaily.Event{}),
		rc_recipe.Annotate(&recipeteamdevice.List{}),
		rc_recipe.Annotate(&recipeteamdevice.Unlink{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipeteamdiag.Explorer{}),
		rc_recipe.Annotate(&recipeteamfilerequest.List{}),
		rc_recipe.Annotate(&recipeteamfolder.Archive{}),
		rc_recipe.Annotate(&recipeteamfolder.List{}),
		rc_recipe.Annotate(&recipeteamfolder.Permdelete{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipeteamfolder.Replication{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipeteamfolderbatch.Archive{}),
		rc_recipe.Annotate(&recipeteamfolderbatch.Permdelete{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipeteamfolderbatch.Replication{}, rc_recipe.Irreversible()),
		rc_recipe.Annotate(&recipeteamfolderfile.List{}),
		rc_recipe.Annotate(&recipeteamfolderfile.Size{}),
		rc_recipe.Annotate(&recipeteamlinkedapp.List{}),
		rc_recipe.Annotate(&recipeteamnamespace.List{}),
		rc_recipe.Annotate(&recipeteamnamespacefile.List{}),
		rc_recipe.Annotate(&recipeteamnamespacefile.Size{}),
		rc_recipe.Annotate(&recipeteamnamespacemember.List{}),
		rc_recipe.Annotate(&recipeteamsharedlink.List{}),
		rc_recipe.Annotate(&recipeteamsharedlinkupdate.Expiry{}, rc_recipe.Irreversible()),
		//		&recipe.Web{},
	}
	return cat
}

func Ingredients() []rc_recipe.Recipe {
	cat := []rc_recipe.Recipe{
		&ingredientfile.Upload{},
		&ingredientteamfolder.Replication{},
		&ingredientteamnamespacefile.List{},
		&ingredientteamnamespacefile.Size{},
	}
	return cat
}

func Messages() []interface{} {
	msgs := []interface{}{
		infra_api_api_api_auth_impl.MCcAuth,
		infra_recipe_rc_conn_impl.MConnect,
		infra_recipe_rc_group.MHeader,
		infra_recipe_rc_group_impl.MGroup,
		infra_recipe_rc_spec.MSelfContained,
		infra_recipe_rcvalue.MRepository,
		infra_recipe_rcvalue.MValFdFileRowFeed,
		infra_report_rpmodelimpl.MTransactionReport,
		infra_report_rpmodelimpl.MXlsxWriter,
		infra_ui_appui.MConsole,
		infra_util_ut_doc.MDoc,
	}
	for _, m := range msgs {
		app_msg.Apply(m)
	}
	return msgs
}
