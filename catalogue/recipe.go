package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
	recipe "github.com/watermint/toolbox/recipe"
	recipeconfig "github.com/watermint/toolbox/recipe/config"
	recipeconfigauth "github.com/watermint/toolbox/recipe/config/auth"
	recipedevbenchmark "github.com/watermint/toolbox/recipe/dev/benchmark"
	recipedevbuild "github.com/watermint/toolbox/recipe/dev/build"
	recipedevciartifact "github.com/watermint/toolbox/recipe/dev/ci/artifact"
	recipedevciauth "github.com/watermint/toolbox/recipe/dev/ci/auth"
	recipedevdiag "github.com/watermint/toolbox/recipe/dev/diag"
	recipedevkvs "github.com/watermint/toolbox/recipe/dev/kvs"
	recipedevmodule "github.com/watermint/toolbox/recipe/dev/module"
	recipedevrelease "github.com/watermint/toolbox/recipe/dev/release"
	recipedevreplay "github.com/watermint/toolbox/recipe/dev/replay"
	recipedevspec "github.com/watermint/toolbox/recipe/dev/spec"
	recipedevstage "github.com/watermint/toolbox/recipe/dev/stage"
	recipedevstagegui "github.com/watermint/toolbox/recipe/dev/stage/gui"
	recipedevtest "github.com/watermint/toolbox/recipe/dev/test"
	recipedevtestauth "github.com/watermint/toolbox/recipe/dev/test/auth"
	recipedevtestsetup "github.com/watermint/toolbox/recipe/dev/test/setup"
	recipedevutil "github.com/watermint/toolbox/recipe/dev/util"
	recipedevutilimage "github.com/watermint/toolbox/recipe/dev/util/image"
	recipefile "github.com/watermint/toolbox/recipe/file"
	recipefilecompare "github.com/watermint/toolbox/recipe/file/compare"
	recipefileexport "github.com/watermint/toolbox/recipe/file/export"
	recipefileimport "github.com/watermint/toolbox/recipe/file/import"
	recipefileimportbatch "github.com/watermint/toolbox/recipe/file/import/batch"
	recipefilelock "github.com/watermint/toolbox/recipe/file/lock"
	recipefilelockall "github.com/watermint/toolbox/recipe/file/lock/all"
	recipefilelockbatch "github.com/watermint/toolbox/recipe/file/lock/batch"
	recipefilepaper "github.com/watermint/toolbox/recipe/file/paper"
	recipefilerestore "github.com/watermint/toolbox/recipe/file/restore"
	recipefilerevision "github.com/watermint/toolbox/recipe/file/revision"
	recipefilesearch "github.com/watermint/toolbox/recipe/file/search"
	recipefileshare "github.com/watermint/toolbox/recipe/file/share"
	recipefilesync "github.com/watermint/toolbox/recipe/file/sync"
	recipefiletag "github.com/watermint/toolbox/recipe/file/tag"
	recipefiletemplateapply "github.com/watermint/toolbox/recipe/file/template/apply"
	recipefiletemplatecapture "github.com/watermint/toolbox/recipe/file/template/capture"
	recipefilerequest "github.com/watermint/toolbox/recipe/filerequest"
	recipefilerequestdelete "github.com/watermint/toolbox/recipe/filerequest/delete"
	recipegroup "github.com/watermint/toolbox/recipe/group"
	recipegroupbatch "github.com/watermint/toolbox/recipe/group/batch"
	recipegroupclear "github.com/watermint/toolbox/recipe/group/clear"
	recipegroupfolder "github.com/watermint/toolbox/recipe/group/folder"
	recipegroupmember "github.com/watermint/toolbox/recipe/group/member"
	recipegroupmemberbatch "github.com/watermint/toolbox/recipe/group/member/batch"
	recipejobhistory "github.com/watermint/toolbox/recipe/job/history"
	recipejoblog "github.com/watermint/toolbox/recipe/job/log"
	recipemember "github.com/watermint/toolbox/recipe/member"
	recipememberbatch "github.com/watermint/toolbox/recipe/member/batch"
	recipememberclear "github.com/watermint/toolbox/recipe/member/clear"
	recipememberfile "github.com/watermint/toolbox/recipe/member/file"
	recipememberfilelock "github.com/watermint/toolbox/recipe/member/file/lock"
	recipememberfilelockall "github.com/watermint/toolbox/recipe/member/file/lock/all"
	recipememberfolder "github.com/watermint/toolbox/recipe/member/folder"
	recipememberquota "github.com/watermint/toolbox/recipe/member/quota"
	recipememberupdate "github.com/watermint/toolbox/recipe/member/update"
	recipeservicesasanateam "github.com/watermint/toolbox/recipe/services/asana/team"
	recipeservicesasanateamproject "github.com/watermint/toolbox/recipe/services/asana/team/project"
	recipeservicesasanateamtask "github.com/watermint/toolbox/recipe/services/asana/team/task"
	recipeservicesasanaworkspace "github.com/watermint/toolbox/recipe/services/asana/workspace"
	recipeservicesasanaworkspaceproject "github.com/watermint/toolbox/recipe/services/asana/workspace/project"
	recipeservicesdropboxuser "github.com/watermint/toolbox/recipe/services/dropbox/user"
	recipeservicesgithub "github.com/watermint/toolbox/recipe/services/github"
	recipeservicesgithubcontent "github.com/watermint/toolbox/recipe/services/github/content"
	recipeservicesgithubissue "github.com/watermint/toolbox/recipe/services/github/issue"
	recipeservicesgithubrelease "github.com/watermint/toolbox/recipe/services/github/release"
	recipeservicesgithubreleaseasset "github.com/watermint/toolbox/recipe/services/github/release/asset"
	recipeservicesgithubtag "github.com/watermint/toolbox/recipe/services/github/tag"
	recipeservicesgooglecalendarevent "github.com/watermint/toolbox/recipe/services/google/calendar/event"
	recipeservicesgooglemailfilter "github.com/watermint/toolbox/recipe/services/google/mail/filter"
	recipeservicesgooglemailfilterbatch "github.com/watermint/toolbox/recipe/services/google/mail/filter/batch"
	recipeservicesgooglemaillabel "github.com/watermint/toolbox/recipe/services/google/mail/label"
	recipeservicesgooglemailmessage "github.com/watermint/toolbox/recipe/services/google/mail/message"
	recipeservicesgooglemailmessagelabel "github.com/watermint/toolbox/recipe/services/google/mail/message/label"
	recipeservicesgooglemailmessageprocessed "github.com/watermint/toolbox/recipe/services/google/mail/message/processed"
	recipeservicesgooglemailsendas "github.com/watermint/toolbox/recipe/services/google/mail/sendas"
	recipeservicesgooglemailthread "github.com/watermint/toolbox/recipe/services/google/mail/thread"
	recipeservicesgooglesheetssheet "github.com/watermint/toolbox/recipe/services/google/sheets/sheet"
	recipeservicesgooglesheetsspreadsheet "github.com/watermint/toolbox/recipe/services/google/sheets/spreadsheet"
	recipeserviceshellosignaccount "github.com/watermint/toolbox/recipe/services/hellosign/account"
	recipeservicesslackconversation "github.com/watermint/toolbox/recipe/services/slack/conversation"
	recipesharedfolder "github.com/watermint/toolbox/recipe/sharedfolder"
	recipesharedfoldermember "github.com/watermint/toolbox/recipe/sharedfolder/member"
	recipesharedfoldermount "github.com/watermint/toolbox/recipe/sharedfolder/mount"
	recipesharedlink "github.com/watermint/toolbox/recipe/sharedlink"
	recipesharedlinkfile "github.com/watermint/toolbox/recipe/sharedlink/file"
	recipeteam "github.com/watermint/toolbox/recipe/team"
	recipeteamactivity "github.com/watermint/toolbox/recipe/team/activity"
	recipeteamactivitybatch "github.com/watermint/toolbox/recipe/team/activity/batch"
	recipeteamactivitydaily "github.com/watermint/toolbox/recipe/team/activity/daily"
	recipeteamadmin "github.com/watermint/toolbox/recipe/team/admin"
	recipeteamadmingrouprole "github.com/watermint/toolbox/recipe/team/admin/group/role"
	recipeteamadminrole "github.com/watermint/toolbox/recipe/team/admin/role"
	recipeteamcontentlegacypaper "github.com/watermint/toolbox/recipe/team/content/legacypaper"
	recipeteamcontentmember "github.com/watermint/toolbox/recipe/team/content/member"
	recipeteamcontentmount "github.com/watermint/toolbox/recipe/team/content/mount"
	recipeteamcontentpolicy "github.com/watermint/toolbox/recipe/team/content/policy"
	recipeteamdevice "github.com/watermint/toolbox/recipe/team/device"
	recipeteamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	recipeteaminsightfile "github.com/watermint/toolbox/recipe/team/insight/file"
	recipeteamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	recipeteamnamespace "github.com/watermint/toolbox/recipe/team/namespace"
	recipeteamnamespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	recipeteamnamespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	recipeteamreport "github.com/watermint/toolbox/recipe/team/report"
	recipeteamrunasfile "github.com/watermint/toolbox/recipe/team/runas/file"
	recipeteamrunasfilebatch "github.com/watermint/toolbox/recipe/team/runas/file/batch"
	recipeteamrunasfilesyncbatch "github.com/watermint/toolbox/recipe/team/runas/file/sync/batch"
	recipeteamrunassharedfolder "github.com/watermint/toolbox/recipe/team/runas/sharedfolder"
	recipeteamrunassharedfolderbatch "github.com/watermint/toolbox/recipe/team/runas/sharedfolder/batch"
	recipeteamrunassharedfoldermemberbatch "github.com/watermint/toolbox/recipe/team/runas/sharedfolder/member/batch"
	recipeteamrunassharedfoldermount "github.com/watermint/toolbox/recipe/team/runas/sharedfolder/mount"
	recipeteamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	recipeteamsharedlinkcap "github.com/watermint/toolbox/recipe/team/sharedlink/cap"
	recipeteamsharedlinkdelete "github.com/watermint/toolbox/recipe/team/sharedlink/delete"
	recipeteamsharedlinkupdate "github.com/watermint/toolbox/recipe/team/sharedlink/update"
	recipeteamfolder "github.com/watermint/toolbox/recipe/teamfolder"
	recipeteamfolderbatch "github.com/watermint/toolbox/recipe/teamfolder/batch"
	recipeteamfolderfile "github.com/watermint/toolbox/recipe/teamfolder/file"
	recipeteamfolderfilelock "github.com/watermint/toolbox/recipe/teamfolder/file/lock"
	recipeteamfolderfilelockall "github.com/watermint/toolbox/recipe/teamfolder/file/lock/all"
	recipeteamfoldermember "github.com/watermint/toolbox/recipe/teamfolder/member"
	recipeteamfolderpartial "github.com/watermint/toolbox/recipe/teamfolder/partial"
	recipeteamfolderpolicy "github.com/watermint/toolbox/recipe/teamfolder/policy"
	recipeteamspaceasadminfile "github.com/watermint/toolbox/recipe/teamspace/asadmin/file"
	recipeteamspaceasadminfolder "github.com/watermint/toolbox/recipe/teamspace/asadmin/folder"
	recipeteamspaceasadminmember "github.com/watermint/toolbox/recipe/teamspace/asadmin/member"
	recipeteamspacefile "github.com/watermint/toolbox/recipe/teamspace/file"
	recipeutilarchive "github.com/watermint/toolbox/recipe/util/archive"
	recipeutildatabase "github.com/watermint/toolbox/recipe/util/database"
	recipeutildate "github.com/watermint/toolbox/recipe/util/date"
	recipeutildatetime "github.com/watermint/toolbox/recipe/util/datetime"
	recipeutildecode "github.com/watermint/toolbox/recipe/util/decode"
	recipeutilencode "github.com/watermint/toolbox/recipe/util/encode"
	recipeutilfile "github.com/watermint/toolbox/recipe/util/file"
	recipeutilgit "github.com/watermint/toolbox/recipe/util/git"
	recipeutilimage "github.com/watermint/toolbox/recipe/util/image"
	recipeutilmonitor "github.com/watermint/toolbox/recipe/util/monitor"
	recipeutilnet "github.com/watermint/toolbox/recipe/util/net"
	recipeutilqrcode "github.com/watermint/toolbox/recipe/util/qrcode"
	recipeutilrelease "github.com/watermint/toolbox/recipe/util/release"
	recipeutiltextcase "github.com/watermint/toolbox/recipe/util/text/case"
	recipeutiltextencoding "github.com/watermint/toolbox/recipe/util/text/encoding"
	recipeutiltidymove "github.com/watermint/toolbox/recipe/util/tidy/move"
	recipeutiltidypack "github.com/watermint/toolbox/recipe/util/tidy/pack"
	recipeutiltime "github.com/watermint/toolbox/recipe/util/time"
	recipeutilunixtime "github.com/watermint/toolbox/recipe/util/unixtime"
	recipeutilxlsx "github.com/watermint/toolbox/recipe/util/xlsx"
	recipeutilxlsxsheet "github.com/watermint/toolbox/recipe/util/xlsx/sheet"
)

func AutoDetectedRecipes() []infra_recipe_rc_recipe.Recipe {
	return []infra_recipe_rc_recipe.Recipe{
		&recipe.License{},
		&recipe.Version{},
		&recipeconfig.Disable{},
		&recipeconfig.Enable{},
		&recipeconfig.Features{},
		&recipeconfigauth.Delete{},
		&recipeconfigauth.List{},
		&recipedevbenchmark.Local{},
		&recipedevbenchmark.Upload{},
		&recipedevbenchmark.Uploadlink{},
		&recipedevbuild.Catalogue{},
		&recipedevbuild.Compile{},
		&recipedevbuild.Doc{},
		&recipedevbuild.Info{},
		&recipedevbuild.License{},
		&recipedevbuild.Package{},
		&recipedevbuild.Preflight{},
		&recipedevbuild.Readme{},
		&recipedevbuild.Target{},
		&recipedevciartifact.Up{},
		&recipedevciauth.Export{},
		&recipedevdiag.Endpoint{},
		&recipedevdiag.Throughput{},
		&recipedevkvs.Benchmark{},
		&recipedevkvs.Dump{},
		&recipedevmodule.List{},
		&recipedevrelease.Candidate{},
		&recipedevrelease.Doc{},
		&recipedevrelease.Publish{},
		&recipedevreplay.Approve{},
		&recipedevreplay.Bundle{},
		&recipedevreplay.Recipe{},
		&recipedevreplay.Remote{},
		&recipedevspec.Diff{},
		&recipedevspec.Doc{},
		&recipedevstage.Dbxfs{},
		&recipedevstage.Encoding{},
		&recipedevstage.Gmail{},
		&recipedevstage.Griddata{},
		&recipedevstage.HttpRange{},
		&recipedevstage.Scoped{},
		&recipedevstage.Teamfolder{},
		&recipedevstage.UploadAppend{},
		&recipedevstagegui.Launch{},
		&recipedevtest.Echo{},
		&recipedevtest.Panic{},
		&recipedevtest.Recipe{},
		&recipedevtest.Resources{},
		&recipedevtestauth.All{},
		&recipedevtestsetup.Massfiles{},
		&recipedevtestsetup.Teamsharedlink{},
		&recipedevutil.Anonymise{},
		&recipedevutil.Curl{},
		&recipedevutil.Wait{},
		&recipedevutilimage.Jpeg{},
		&recipefile.Copy{},
		&recipefile.Delete{},
		&recipefile.Info{},
		&recipefile.List{},
		&recipefile.Merge{},
		&recipefile.Move{},
		&recipefile.Replication{},
		&recipefile.Size{},
		&recipefile.Watch{},
		&recipefilecompare.Account{},
		&recipefilecompare.Local{},
		&recipefileexport.Doc{},
		&recipefileexport.Url{},
		&recipefileimport.Url{},
		&recipefileimportbatch.Url{},
		&recipefilelock.Acquire{},
		&recipefilelock.List{},
		&recipefilelock.Release{},
		&recipefilelockall.Release{},
		&recipefilelockbatch.Acquire{},
		&recipefilelockbatch.Release{},
		&recipefilepaper.Append{},
		&recipefilepaper.Create{},
		&recipefilepaper.Overwrite{},
		&recipefilepaper.Prepend{},
		&recipefilerestore.All{},
		&recipefilerevision.Download{},
		&recipefilerevision.List{},
		&recipefilerevision.Restore{},
		&recipefilesearch.Content{},
		&recipefilesearch.Name{},
		&recipefileshare.Info{},
		&recipefilesync.Down{},
		&recipefilesync.Online{},
		&recipefilesync.Up{},
		&recipefiletag.Add{},
		&recipefiletag.Delete{},
		&recipefiletag.List{},
		&recipefiletemplateapply.Local{},
		&recipefiletemplateapply.Remote{},
		&recipefiletemplatecapture.Local{},
		&recipefiletemplatecapture.Remote{},
		&recipefilerequest.Create{},
		&recipefilerequest.List{},
		&recipefilerequestdelete.Closed{},
		&recipefilerequestdelete.Url{},
		&recipegroup.Add{},
		&recipegroup.Delete{},
		&recipegroup.List{},
		&recipegroup.Rename{},
		&recipegroupbatch.Add{},
		&recipegroupbatch.Delete{},
		&recipegroupclear.Externalid{},
		&recipegroupfolder.List{},
		&recipegroupmember.Add{},
		&recipegroupmember.Delete{},
		&recipegroupmember.List{},
		&recipegroupmemberbatch.Add{},
		&recipegroupmemberbatch.Delete{},
		&recipegroupmemberbatch.Update{},
		&recipejobhistory.Archive{},
		&recipejobhistory.Delete{},
		&recipejobhistory.List{},
		&recipejobhistory.Ship{},
		&recipejoblog.Jobid{},
		&recipejoblog.Kind{},
		&recipejoblog.Last{},
		&recipemember.Delete{},
		&recipemember.Detach{},
		&recipemember.Feature{},
		&recipemember.Invite{},
		&recipemember.List{},
		&recipemember.Reinvite{},
		&recipemember.Replication{},
		&recipemember.Suspend{},
		&recipemember.Unsuspend{},
		&recipememberbatch.Suspend{},
		&recipememberbatch.Unsuspend{},
		&recipememberclear.Externalid{},
		&recipememberfile.Permdelete{},
		&recipememberfilelock.List{},
		&recipememberfilelock.Release{},
		&recipememberfilelockall.Release{},
		&recipememberfolder.List{},
		&recipememberfolder.Replication{},
		&recipememberquota.List{},
		&recipememberquota.Update{},
		&recipememberquota.Usage{},
		&recipememberupdate.Email{},
		&recipememberupdate.Externalid{},
		&recipememberupdate.Invisible{},
		&recipememberupdate.Profile{},
		&recipememberupdate.Visible{},
		&recipeservicesasanateam.List{},
		&recipeservicesasanateamproject.List{},
		&recipeservicesasanateamtask.List{},
		&recipeservicesasanaworkspace.List{},
		&recipeservicesasanaworkspaceproject.List{},
		&recipeservicesdropboxuser.Feature{},
		&recipeservicesdropboxuser.Info{},
		&recipeservicesgithub.Profile{},
		&recipeservicesgithubcontent.Get{},
		&recipeservicesgithubcontent.Put{},
		&recipeservicesgithubissue.List{},
		&recipeservicesgithubrelease.Draft{},
		&recipeservicesgithubrelease.List{},
		&recipeservicesgithubreleaseasset.Download{},
		&recipeservicesgithubreleaseasset.List{},
		&recipeservicesgithubreleaseasset.Upload{},
		&recipeservicesgithubtag.Create{},
		&recipeservicesgooglecalendarevent.List{},
		&recipeservicesgooglemailfilter.Add{},
		&recipeservicesgooglemailfilter.Delete{},
		&recipeservicesgooglemailfilter.List{},
		&recipeservicesgooglemailfilterbatch.Add{},
		&recipeservicesgooglemaillabel.Add{},
		&recipeservicesgooglemaillabel.Delete{},
		&recipeservicesgooglemaillabel.List{},
		&recipeservicesgooglemaillabel.Rename{},
		&recipeservicesgooglemailmessage.List{},
		&recipeservicesgooglemailmessage.Send{},
		&recipeservicesgooglemailmessagelabel.Add{},
		&recipeservicesgooglemailmessagelabel.Delete{},
		&recipeservicesgooglemailmessageprocessed.List{},
		&recipeservicesgooglemailsendas.Add{},
		&recipeservicesgooglemailsendas.Delete{},
		&recipeservicesgooglemailsendas.List{},
		&recipeservicesgooglemailthread.List{},
		&recipeservicesgooglesheetssheet.Append{},
		&recipeservicesgooglesheetssheet.Clear{},
		&recipeservicesgooglesheetssheet.Create{},
		&recipeservicesgooglesheetssheet.Delete{},
		&recipeservicesgooglesheetssheet.Export{},
		&recipeservicesgooglesheetssheet.Import{},
		&recipeservicesgooglesheetssheet.List{},
		&recipeservicesgooglesheetsspreadsheet.Create{},
		&recipeserviceshellosignaccount.Info{},
		&recipeservicesslackconversation.List{},
		&recipesharedfolder.Leave{},
		&recipesharedfolder.List{},
		&recipesharedfolder.Share{},
		&recipesharedfolder.Unshare{},
		&recipesharedfoldermember.Add{},
		&recipesharedfoldermember.Delete{},
		&recipesharedfoldermember.List{},
		&recipesharedfoldermount.Add{},
		&recipesharedfoldermount.Delete{},
		&recipesharedfoldermount.List{},
		&recipesharedfoldermount.Mountable{},
		&recipesharedlink.Create{},
		&recipesharedlink.Delete{},
		&recipesharedlink.Info{},
		&recipesharedlink.List{},
		&recipesharedlinkfile.List{},
		&recipeteam.Feature{},
		&recipeteam.Info{},
		&recipeteamactivity.Event{},
		&recipeteamactivity.User{},
		&recipeteamactivitybatch.User{},
		&recipeteamactivitydaily.Event{},
		&recipeteamadmin.List{},
		&recipeteamadmingrouprole.Add{},
		&recipeteamadmingrouprole.Delete{},
		&recipeteamadminrole.Add{},
		&recipeteamadminrole.Clear{},
		&recipeteamadminrole.Delete{},
		&recipeteamadminrole.List{},
		&recipeteamcontentlegacypaper.Count{},
		&recipeteamcontentlegacypaper.Export{},
		&recipeteamcontentlegacypaper.List{},
		&recipeteamcontentmember.List{},
		&recipeteamcontentmember.Size{},
		&recipeteamcontentmount.List{},
		&recipeteamcontentpolicy.List{},
		&recipeteamdevice.List{},
		&recipeteamdevice.Unlink{},
		&recipeteamfilerequest.Clone{},
		&recipeteamfilerequest.List{},
		&recipeteaminsightfile.Scan{},
		&recipeteamlinkedapp.List{},
		&recipeteamnamespace.List{},
		&recipeteamnamespace.Summary{},
		&recipeteamnamespacefile.List{},
		&recipeteamnamespacefile.Size{},
		&recipeteamnamespacemember.List{},
		&recipeteamreport.Activity{},
		&recipeteamreport.Devices{},
		&recipeteamreport.Membership{},
		&recipeteamreport.Storage{},
		&recipeteamrunasfile.List{},
		&recipeteamrunasfilebatch.Copy{},
		&recipeteamrunasfilesyncbatch.Up{},
		&recipeteamrunassharedfolder.Isolate{},
		&recipeteamrunassharedfolder.List{},
		&recipeteamrunassharedfolderbatch.Leave{},
		&recipeteamrunassharedfolderbatch.Share{},
		&recipeteamrunassharedfolderbatch.Unshare{},
		&recipeteamrunassharedfoldermemberbatch.Add{},
		&recipeteamrunassharedfoldermemberbatch.Delete{},
		&recipeteamrunassharedfoldermount.Add{},
		&recipeteamrunassharedfoldermount.Delete{},
		&recipeteamrunassharedfoldermount.List{},
		&recipeteamrunassharedfoldermount.Mountable{},
		&recipeteamsharedlink.List{},
		&recipeteamsharedlinkcap.Expiry{},
		&recipeteamsharedlinkcap.Visibility{},
		&recipeteamsharedlinkdelete.Links{},
		&recipeteamsharedlinkdelete.Member{},
		&recipeteamsharedlinkupdate.Expiry{},
		&recipeteamsharedlinkupdate.Password{},
		&recipeteamsharedlinkupdate.Visibility{},
		&recipeteamfolder.Add{},
		&recipeteamfolder.Archive{},
		&recipeteamfolder.List{},
		&recipeteamfolder.Permdelete{},
		&recipeteamfolder.Replication{},
		&recipeteamfolderbatch.Archive{},
		&recipeteamfolderbatch.Permdelete{},
		&recipeteamfolderbatch.Replication{},
		&recipeteamfolderfile.List{},
		&recipeteamfolderfile.Size{},
		&recipeteamfolderfilelock.List{},
		&recipeteamfolderfilelock.Release{},
		&recipeteamfolderfilelockall.Release{},
		&recipeteamfoldermember.Add{},
		&recipeteamfoldermember.Delete{},
		&recipeteamfoldermember.List{},
		&recipeteamfolderpartial.Replication{},
		&recipeteamfolderpolicy.List{},
		&recipeteamspaceasadminfile.List{},
		&recipeteamspaceasadminfolder.Add{},
		&recipeteamspaceasadminfolder.Delete{},
		&recipeteamspaceasadminfolder.Permdelete{},
		&recipeteamspaceasadminmember.List{},
		&recipeteamspacefile.List{},
		&recipeutilarchive.Unzip{},
		&recipeutilarchive.Zip{},
		&recipeutildatabase.Exec{},
		&recipeutildatabase.Query{},
		&recipeutildate.Today{},
		&recipeutildatetime.Now{},
		&recipeutildecode.Base32{},
		&recipeutildecode.Base64{},
		&recipeutilencode.Base32{},
		&recipeutilencode.Base64{},
		&recipeutilfile.Hash{},
		&recipeutilgit.Clone{},
		&recipeutilimage.Exif{},
		&recipeutilimage.Placeholder{},
		&recipeutilmonitor.Client{},
		&recipeutilnet.Download{},
		&recipeutilqrcode.Create{},
		&recipeutilqrcode.Wifi{},
		&recipeutilrelease.Install{},
		&recipeutiltextcase.Down{},
		&recipeutiltextcase.Up{},
		&recipeutiltextencoding.From{},
		&recipeutiltextencoding.To{},
		&recipeutiltidymove.Dispatch{},
		&recipeutiltidymove.Simple{},
		&recipeutiltidypack.Remote{},
		&recipeutiltime.Now{},
		&recipeutilunixtime.Format{},
		&recipeutilunixtime.Now{},
		&recipeutilxlsx.Create{},
		&recipeutilxlsxsheet.Export{},
		&recipeutilxlsxsheet.Import{},
		&recipeutilxlsxsheet.List{},
	}
}
