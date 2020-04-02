package catalogue

import (
	infra_api_api_api_auth_impl "github.com/watermint/toolbox/infra/api/dbx_auth"
	infra_control_app_workflow "github.com/watermint/toolbox/infra/control/app_workflow"
	infra_kvs_kv_storageimpl "github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	infra_network_nw_diag "github.com/watermint/toolbox/infra/network/nw_diag"
	infra_recipe_rc_catalogue "github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	infra_recipe_rc_conn_impl "github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	infra_recipe_rc_group "github.com/watermint/toolbox/infra/recipe/rc_group"
	infra_recipe_rc_group_impl "github.com/watermint/toolbox/infra/recipe/rc_group_impl"
	infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
	infra_recipe_rc_spec "github.com/watermint/toolbox/infra/recipe/rc_spec"
	infra_recipe_rcvalue "github.com/watermint/toolbox/infra/recipe/rc_value"
	infra_report_rpmodelimpl "github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/infra/report/rp_writer_impl"
	infra_ui_app_msg "github.com/watermint/toolbox/infra/ui/app_msg"
	infra_ui_appui "github.com/watermint/toolbox/infra/ui/app_ui"
	infra_util_ut_doc "github.com/watermint/toolbox/infra/util/ut_doc"
	ingredientfile "github.com/watermint/toolbox/ingredient/file"
	ingredientteamnamespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	ingredientteamfolder "github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/recipe"
	recipeconnect "github.com/watermint/toolbox/recipe/connect"
	recipedev "github.com/watermint/toolbox/recipe/dev"
	recipedevci "github.com/watermint/toolbox/recipe/dev/ci"
	recipedevciartifact "github.com/watermint/toolbox/recipe/dev/ci/artifact"
	recipedevdesktop "github.com/watermint/toolbox/recipe/dev/desktop"
	recipedevdiag "github.com/watermint/toolbox/recipe/dev/diag"
	recipedevrelease "github.com/watermint/toolbox/recipe/dev/release"
	recipedevspec "github.com/watermint/toolbox/recipe/dev/spec"
	recipedevtest "github.com/watermint/toolbox/recipe/dev/test"
	recipedevutil "github.com/watermint/toolbox/recipe/dev/util"
	recipefile "github.com/watermint/toolbox/recipe/file"
	recipefilecompare "github.com/watermint/toolbox/recipe/file/compare"
	recipefiledispatch "github.com/watermint/toolbox/recipe/file/dispatch"
	recipefileexport "github.com/watermint/toolbox/recipe/file/export"
	recipefileimport "github.com/watermint/toolbox/recipe/file/import"
	recipefileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	recipefilesearch "github.com/watermint/toolbox/recipe/file/search"
	recipefilesync "github.com/watermint/toolbox/recipe/file/sync"
	recipefilesyncpreflight "github.com/watermint/toolbox/recipe/file/sync/preflight"
	recipefilerequest "github.com/watermint/toolbox/recipe/filerequest"
	recipefilerequestdelete "github.com/watermint/toolbox/recipe/filerequest/delete"
	recipegroup "github.com/watermint/toolbox/recipe/group"
	recipegroupbatch "github.com/watermint/toolbox/recipe/group/batch"
	recipegroupmember "github.com/watermint/toolbox/recipe/group/member"
	recipejob "github.com/watermint/toolbox/recipe/job"
	recipejobhistory "github.com/watermint/toolbox/recipe/job/history"
	recipemember "github.com/watermint/toolbox/recipe/member"
	recipememberquota "github.com/watermint/toolbox/recipe/member/quota"
	recipememberupdate "github.com/watermint/toolbox/recipe/member/update"
	recipesharedfolder "github.com/watermint/toolbox/recipe/sharedfolder"
	recipesharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	recipesharedlink "github.com/watermint/toolbox/recipe/sharedlink"
	recipesharedlinkfile "github.com/watermint/toolbox/recipe/sharedlink/file"
	recipeteam "github.com/watermint/toolbox/recipe/team"
	recipeteamactivity "github.com/watermint/toolbox/recipe/team/activity"
	recipeteamactivitybatch "github.com/watermint/toolbox/recipe/team/activity/batch"
	recipeteamactivitydaily "github.com/watermint/toolbox/recipe/team/activity/daily"
	recipeteamcontent "github.com/watermint/toolbox/recipe/team/content"
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

func Recipes() []infra_recipe_rc_recipe.Recipe {
	cat := []infra_recipe_rc_recipe.Recipe{
		infra_recipe_rc_recipe.Annotate(&recipe.Version{}),
		infra_recipe_rc_recipe.Annotate(&recipe.License{}),
		infra_recipe_rc_recipe.Annotate(&recipe.Web{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessAudit{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessFile{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessInfo{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessMgmt{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.UserFile{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Async{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Doc{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Dummy{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Echo{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevrelease.Candidate{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevrelease.Publish{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Preflight{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevci.Auth{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciartifact.Up{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Install{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Start{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Stop{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Suspendupdate{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdiag.Procmon{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevspec.Diff{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevspec.Doc{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Monkey{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Recipe{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Resources{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevutil.Wait{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevutil.Curl{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipefile.Copy{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Delete{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Download{}, infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipefile.List{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Merge{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Move{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Replication{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Restore{}, infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipefile.Upload{}),
		infra_recipe_rc_recipe.Annotate(&recipefile.Watch{}),
		infra_recipe_rc_recipe.Annotate(&recipefilecompare.Account{}),
		infra_recipe_rc_recipe.Annotate(&recipefilecompare.Local{}),
		infra_recipe_rc_recipe.Annotate(&recipefiledispatch.Local{}),
		infra_recipe_rc_recipe.Annotate(&recipefileexport.Doc{}, infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipefileimport.Url{}),
		infra_recipe_rc_recipe.Annotate(&recipefileimportbatch.Url{}),
		infra_recipe_rc_recipe.Annotate(&recipefilerequest.Create{}),
		infra_recipe_rc_recipe.Annotate(&recipefilerequest.List{}),
		infra_recipe_rc_recipe.Annotate(&recipefilerequestdelete.Closed{}),
		infra_recipe_rc_recipe.Annotate(&recipefilerequestdelete.Url{}),
		infra_recipe_rc_recipe.Annotate(&recipefilesearch.Content{}),
		infra_recipe_rc_recipe.Annotate(&recipefilesearch.Name{}),
		infra_recipe_rc_recipe.Annotate(&recipefilesync.Up{}),
		infra_recipe_rc_recipe.Annotate(&recipefilesyncpreflight.Up{}),
		infra_recipe_rc_recipe.Annotate(&recipegroup.Add{}),
		infra_recipe_rc_recipe.Annotate(&recipegroup.Delete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipegroup.List{}),
		infra_recipe_rc_recipe.Annotate(&recipegroup.Rename{}),
		infra_recipe_rc_recipe.Annotate(&recipegroupbatch.Delete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipegroupmember.Add{}),
		infra_recipe_rc_recipe.Annotate(&recipegroupmember.Delete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipegroupmember.List{}),
		infra_recipe_rc_recipe.Annotate(&recipejob.Loop{}),
		infra_recipe_rc_recipe.Annotate(&recipejob.Run{}),
		infra_recipe_rc_recipe.Annotate(&recipejobhistory.Archive{}),
		infra_recipe_rc_recipe.Annotate(&recipejobhistory.Delete{}),
		infra_recipe_rc_recipe.Annotate(&recipejobhistory.List{}),
		infra_recipe_rc_recipe.Annotate(&recipejobhistory.Ship{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.Delete{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.Detach{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.Invite{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.List{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.Reinvite{}),
		infra_recipe_rc_recipe.Annotate(&recipemember.Replication{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipememberquota.List{}),
		infra_recipe_rc_recipe.Annotate(&recipememberquota.Update{}),
		infra_recipe_rc_recipe.Annotate(&recipememberquota.Usage{}),
		infra_recipe_rc_recipe.Annotate(&recipememberupdate.Email{}),
		infra_recipe_rc_recipe.Annotate(&recipememberupdate.Externalid{}),
		infra_recipe_rc_recipe.Annotate(&recipememberupdate.Profile{}),
		infra_recipe_rc_recipe.Annotate(&recipesharedfolder.List{}),
		infra_recipe_rc_recipe.Annotate(&recipesharedfoldermember.List{}),
		infra_recipe_rc_recipe.Annotate(&recipesharedlink.Create{}),
		infra_recipe_rc_recipe.Annotate(&recipesharedlink.Delete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipesharedlink.List{}),
		infra_recipe_rc_recipe.Annotate(&recipesharedlinkfile.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteam.Feature{}),
		infra_recipe_rc_recipe.Annotate(&recipeteam.Info{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamactivity.Event{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamactivity.User{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamactivitybatch.User{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamactivitydaily.Event{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamdevice.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamdevice.Unlink{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamdiag.Explorer{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfilerequest.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamcontent.Member{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamcontent.Policy{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolder.Archive{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolder.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolder.Permdelete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolder.Replication{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolderbatch.Archive{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolderbatch.Permdelete{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolderbatch.Replication{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolderfile.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfolderfile.Size{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamlinkedapp.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamnamespace.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamnamespacefile.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamnamespacefile.Size{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamnamespacemember.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamsharedlink.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamsharedlinkupdate.Expiry{}, infra_recipe_rc_recipe.Irreversible()),
	}
	return cat
}

func Ingredients() []infra_recipe_rc_recipe.Recipe {
	cat := []infra_recipe_rc_recipe.Recipe{
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
		infra_control_app_workflow.MRunBook,
		infra_kvs_kv_storageimpl.MStorage,
		infra_network_nw_diag.MNetwork,
		infra_recipe_rc_conn_impl.MConnect,
		infra_recipe_rc_group.MHeader,
		infra_recipe_rc_group_impl.MGroup,
		infra_recipe_rc_spec.MSelfContained,
		infra_recipe_rcvalue.MRepository,
		infra_recipe_rcvalue.MValFdFileRowFeed,
		infra_report_rpmodelimpl.MTransactionReport,
		infra_ui_appui.MConsole,
		infra_ui_appui.MProgress,
		infra_util_ut_doc.MDoc,
		recipefiledispatch.MLocal,
		recipeteamactivitybatch.MUser,
		recipeteamcontent.MScanMetadata,
		rp_writer_impl.MSortedWriter,
		rp_writer_impl.MXlsxWriter,
	}
	for _, m := range msgs {
		infra_ui_app_msg.Apply(m)
	}
	return msgs
}
