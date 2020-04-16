package catalogue

import (
	infra_api_api_api_auth_impl "github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	infra_recipe_rc_conn_impl "github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_local"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_callback"
	"github.com/watermint/toolbox/infra/control/app_feature"
	infra_control_app_workflow "github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	infra_kvs_kv_storageimpl "github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	infra_network_nw_diag "github.com/watermint/toolbox/infra/network/nw_diag"
	infra_recipe_rc_catalogue "github.com/watermint/toolbox/infra/recipe/rc_catalogue"
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
	ingredientbootstrap "github.com/watermint/toolbox/ingredient/bootstrap"
	ingredientfile "github.com/watermint/toolbox/ingredient/file"
	ingredientjob "github.com/watermint/toolbox/ingredient/job"
	ingredientteamnamespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	ingredientteamfolder "github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/recipe"
	recipeconfig "github.com/watermint/toolbox/recipe/config"
	recipeconnect "github.com/watermint/toolbox/recipe/connect"
	recipedev "github.com/watermint/toolbox/recipe/dev"
	recipedevciartifact "github.com/watermint/toolbox/recipe/dev/ci/artifact"
	recipedevciauth "github.com/watermint/toolbox/recipe/dev/ci/auth"
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
	recipeservicesgithub "github.com/watermint/toolbox/recipe/services/github"
	recipeservicesgithubissue "github.com/watermint/toolbox/recipe/services/github/issue"
	recipeservicesgithubrelease "github.com/watermint/toolbox/recipe/services/github/release"
	recipeservicesgithubreleaseasset "github.com/watermint/toolbox/recipe/services/github/release/asset"
	recipeservicesgithubtag "github.com/watermint/toolbox/recipe/services/github/tag"
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
	return infra_recipe_rc_catalogue.NewCatalogue(
		Recipes(),
		Ingredients(),
		Messages(),
		Features(),
	)
}

func Recipes() []infra_recipe_rc_recipe.Recipe {
	cat := []infra_recipe_rc_recipe.Recipe{
		infra_recipe_rc_recipe.Annotate(&recipe.License{}),
		infra_recipe_rc_recipe.Annotate(&recipe.Version{}),
		infra_recipe_rc_recipe.Annotate(&recipeconfig.Disable{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconfig.Enable{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconfig.Features{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessAudit{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessFile{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessInfo{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.BusinessMgmt{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipeconnect.UserFile{}, infra_recipe_rc_recipe.Console()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Async{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Catalogue{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Doc{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Dummy{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Echo{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedev.Preflight{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciartifact.Connect{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciartifact.Up{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciauth.Connect{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciauth.Export{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevciauth.Import{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Install{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Start{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Stop{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdesktop.Suspendupdate{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevdiag.Procmon{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevrelease.Candidate{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevrelease.Publish{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevspec.Diff{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevspec.Doc{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Monkey{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Recipe{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevtest.Resources{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevutil.Curl{}, infra_recipe_rc_recipe.Secret()),
		infra_recipe_rc_recipe.Annotate(&recipedevutil.Wait{}, infra_recipe_rc_recipe.Secret()),
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
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithub.Profile{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubissue.List{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubrelease.Draft{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubrelease.List{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubreleaseasset.List{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubreleaseasset.Up{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental()),
		infra_recipe_rc_recipe.Annotate(&recipeservicesgithubtag.Create{}, infra_recipe_rc_recipe.Console(), infra_recipe_rc_recipe.Experimental(), infra_recipe_rc_recipe.Irreversible()),
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
		infra_recipe_rc_recipe.Annotate(&recipeteamcontent.Member{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamcontent.Policy{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamdevice.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamdevice.Unlink{}, infra_recipe_rc_recipe.Irreversible()),
		infra_recipe_rc_recipe.Annotate(&recipeteamdiag.Explorer{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfilerequest.List{}),
		infra_recipe_rc_recipe.Annotate(&recipeteamfilerequest.Clone{}, infra_recipe_rc_recipe.Experimental(), infra_recipe_rc_recipe.Secret()),
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
		&ingredientbootstrap.Bootstrap{},
		&ingredientbootstrap.Autodelete{},
		&ingredientjob.Delete{},
	}
	return cat
}

func Features() []app_feature.OptIn {
	foi := []app_feature.OptIn{
		&api_auth_impl.FeatureRedirect{},
		&ingredientbootstrap.FeatureAutodelete{},
	}
	return foi
}

func Messages() []interface{} {
	msgs := []interface{}{
		api_auth_impl.MApiAuth,
		api_callback.MCallback,
		dbx_auth_attr.MAttr,
		fd_file_impl.MRowFeed,
		infra_api_api_api_auth_impl.MGenerated,
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
		ingredientteamnamespacefile.MList,
		ingredientteamnamespacefile.MSize,
		recipefile.MRestore,
		recipefiledispatch.MLocal,
		recipefileimportbatch.MUrl,
		recipegroupmember.MList,
		recipememberquota.MList,
		recipememberquota.MUpdate,
		recipememberquota.MUsage,
		recipememberupdate.MEmail,
		recipeservicesgithubreleaseasset.MUp,
		recipesharedfoldermember.MList,
		recipeteamactivity.MUser,
		recipeteamactivitybatch.MUser,
		recipeteamcontent.MScanMetadata,
		recipeteamdevice.MUnlink,
		recipeteamfilerequest.MList,
		recipeteamnamespacemember.MList,
		recipeteamsharedlink.MList,
		recipeteamsharedlinkupdate.MExpiry,
		rp_writer_impl.MSortedWriter,
		rp_writer_impl.MXlsxWriter,
		uc_compare_local.MCompare,
		uc_compare_paths.MCompare,
	}
	for _, m := range msgs {
		infra_ui_app_msg.Apply(m)
	}
	return msgs
}
