package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	citronasanateam "github.com/watermint/toolbox/citron/asana/team"
	citronasanateamproject "github.com/watermint/toolbox/citron/asana/team/project"
	citronasanateamtask "github.com/watermint/toolbox/citron/asana/team/task"
	citronasanaworkspace "github.com/watermint/toolbox/citron/asana/workspace"
	citronasanaworkspaceproject "github.com/watermint/toolbox/citron/asana/workspace/project"
	citrondeepltranslate "github.com/watermint/toolbox/citron/deepl/translate"
	citrondropboxfile "github.com/watermint/toolbox/citron/dropbox/file"
	citrondropboxfileaccount "github.com/watermint/toolbox/citron/dropbox/file/account"
	citrondropboxfilecompare "github.com/watermint/toolbox/citron/dropbox/file/compare"
	citrondropboxfileexport "github.com/watermint/toolbox/citron/dropbox/file/export"
	citrondropboxfileimport "github.com/watermint/toolbox/citron/dropbox/file/import"
	citrondropboxfileimportbatch "github.com/watermint/toolbox/citron/dropbox/file/import/batch"
	citrondropboxfilelock "github.com/watermint/toolbox/citron/dropbox/file/lock"
	citrondropboxfilelockall "github.com/watermint/toolbox/citron/dropbox/file/lock/all"
	citrondropboxfilelockbatch "github.com/watermint/toolbox/citron/dropbox/file/lock/batch"
	citrondropboxfilerequest "github.com/watermint/toolbox/citron/dropbox/file/request"
	citrondropboxfilerequestdelete "github.com/watermint/toolbox/citron/dropbox/file/request/delete"
	citrondropboxfilerestore "github.com/watermint/toolbox/citron/dropbox/file/restore"
	citrondropboxfilerevision "github.com/watermint/toolbox/citron/dropbox/file/revision"
	citrondropboxfilesearch "github.com/watermint/toolbox/citron/dropbox/file/search"
	citrondropboxfileshare "github.com/watermint/toolbox/citron/dropbox/file/share"
	citrondropboxfilesharedfolder "github.com/watermint/toolbox/citron/dropbox/file/sharedfolder"
	citrondropboxfilesharedfoldermember "github.com/watermint/toolbox/citron/dropbox/file/sharedfolder/member"
	citrondropboxfilesharedfoldermount "github.com/watermint/toolbox/citron/dropbox/file/sharedfolder/mount"
	citrondropboxfilesharedlink "github.com/watermint/toolbox/citron/dropbox/file/sharedlink"
	citrondropboxfilesharedlinkfile "github.com/watermint/toolbox/citron/dropbox/file/sharedlink/file"
	citrondropboxfilesync "github.com/watermint/toolbox/citron/dropbox/file/sync"
	citrondropboxfiletag "github.com/watermint/toolbox/citron/dropbox/file/tag"
	citrondropboxfiletemplate "github.com/watermint/toolbox/citron/dropbox/file/template"
	citrondropboxpaper "github.com/watermint/toolbox/citron/dropbox/paper"
	citrondropboxsignaccount "github.com/watermint/toolbox/citron/dropbox/sign/account"
	citrondropboxsignrequest "github.com/watermint/toolbox/citron/dropbox/sign/request"
	citrondropboxsignrequestsignature "github.com/watermint/toolbox/citron/dropbox/sign/request/signature"
	citrondropboxteam "github.com/watermint/toolbox/citron/dropbox/team"
	citrondropboxteamactivity "github.com/watermint/toolbox/citron/dropbox/team/activity"
	citrondropboxteamactivitybatch "github.com/watermint/toolbox/citron/dropbox/team/activity/batch"
	citrondropboxteamactivitydaily "github.com/watermint/toolbox/citron/dropbox/team/activity/daily"
	citrondropboxteamadmin "github.com/watermint/toolbox/citron/dropbox/team/admin"
	citrondropboxteamadmingrouprole "github.com/watermint/toolbox/citron/dropbox/team/admin/group/role"
	citrondropboxteamadminrole "github.com/watermint/toolbox/citron/dropbox/team/admin/role"
	citrondropboxteambackupdevice "github.com/watermint/toolbox/citron/dropbox/team/backup/device"
	citrondropboxteamcontentlegacypaper "github.com/watermint/toolbox/citron/dropbox/team/content/legacypaper"
	citrondropboxteamcontentmember "github.com/watermint/toolbox/citron/dropbox/team/content/member"
	citrondropboxteamcontentmount "github.com/watermint/toolbox/citron/dropbox/team/content/mount"
	citrondropboxteamcontentpolicy "github.com/watermint/toolbox/citron/dropbox/team/content/policy"
	citrondropboxteamdevice "github.com/watermint/toolbox/citron/dropbox/team/device"
	citrondropboxteamfilerequest "github.com/watermint/toolbox/citron/dropbox/team/filerequest"
	citrondropboxteamgroup "github.com/watermint/toolbox/citron/dropbox/team/group"
	citrondropboxteamgroupbatch "github.com/watermint/toolbox/citron/dropbox/team/group/batch"
	citrondropboxteamgroupclear "github.com/watermint/toolbox/citron/dropbox/team/group/clear"
	citrondropboxteamgroupfolder "github.com/watermint/toolbox/citron/dropbox/team/group/folder"
	citrondropboxteamgroupmember "github.com/watermint/toolbox/citron/dropbox/team/group/member"
	citrondropboxteamgroupmemberbatch "github.com/watermint/toolbox/citron/dropbox/team/group/member/batch"
	citrondropboxteamgroupupdate "github.com/watermint/toolbox/citron/dropbox/team/group/update"
	citrondropboxteaminsight "github.com/watermint/toolbox/citron/dropbox/team/insight"
	citrondropboxteaminsightreport "github.com/watermint/toolbox/citron/dropbox/team/insight/report"
	citrondropboxteamlegalhold "github.com/watermint/toolbox/citron/dropbox/team/legalhold"
	citrondropboxteamlegalholdmember "github.com/watermint/toolbox/citron/dropbox/team/legalhold/member"
	citrondropboxteamlegalholdmemberbatch "github.com/watermint/toolbox/citron/dropbox/team/legalhold/member/batch"
	citrondropboxteamlegalholdrevision "github.com/watermint/toolbox/citron/dropbox/team/legalhold/revision"
	citrondropboxteamlegalholdupdate "github.com/watermint/toolbox/citron/dropbox/team/legalhold/update"
	citrondropboxteamlinkedapp "github.com/watermint/toolbox/citron/dropbox/team/linkedapp"
	citrondropboxteammember "github.com/watermint/toolbox/citron/dropbox/team/member"
	citrondropboxteammemberbatch "github.com/watermint/toolbox/citron/dropbox/team/member/batch"
	citrondropboxteammemberclear "github.com/watermint/toolbox/citron/dropbox/team/member/clear"
	citrondropboxteammemberfile "github.com/watermint/toolbox/citron/dropbox/team/member/file"
	citrondropboxteammemberfilelock "github.com/watermint/toolbox/citron/dropbox/team/member/file/lock"
	citrondropboxteammemberfilelockall "github.com/watermint/toolbox/citron/dropbox/team/member/file/lock/all"
	citrondropboxteammemberfolder "github.com/watermint/toolbox/citron/dropbox/team/member/folder"
	citrondropboxteammemberquota "github.com/watermint/toolbox/citron/dropbox/team/member/quota"
	citrondropboxteammemberquotabatch "github.com/watermint/toolbox/citron/dropbox/team/member/quota/batch"
	citrondropboxteammemberupdatebatch "github.com/watermint/toolbox/citron/dropbox/team/member/update/batch"
	citrondropboxteamnamespace "github.com/watermint/toolbox/citron/dropbox/team/namespace"
	citrondropboxteamnamespacefile "github.com/watermint/toolbox/citron/dropbox/team/namespace/file"
	citrondropboxteamnamespacemember "github.com/watermint/toolbox/citron/dropbox/team/namespace/member"
	citrondropboxteamreport "github.com/watermint/toolbox/citron/dropbox/team/report"
	citrondropboxteamrunasfile "github.com/watermint/toolbox/citron/dropbox/team/runas/file"
	citrondropboxteamrunasfilebatch "github.com/watermint/toolbox/citron/dropbox/team/runas/file/batch"
	citrondropboxteamrunasfilesyncbatch "github.com/watermint/toolbox/citron/dropbox/team/runas/file/sync/batch"
	citrondropboxteamrunassharedfolder "github.com/watermint/toolbox/citron/dropbox/team/runas/sharedfolder"
	citrondropboxteamrunassharedfolderbatch "github.com/watermint/toolbox/citron/dropbox/team/runas/sharedfolder/batch"
	citrondropboxteamrunassharedfoldermemberbatch "github.com/watermint/toolbox/citron/dropbox/team/runas/sharedfolder/member/batch"
	citrondropboxteamrunassharedfoldermount "github.com/watermint/toolbox/citron/dropbox/team/runas/sharedfolder/mount"
	citrondropboxteamsharedlink "github.com/watermint/toolbox/citron/dropbox/team/sharedlink"
	citrondropboxteamsharedlinkcap "github.com/watermint/toolbox/citron/dropbox/team/sharedlink/cap"
	citrondropboxteamsharedlinkdelete "github.com/watermint/toolbox/citron/dropbox/team/sharedlink/delete"
	citrondropboxteamsharedlinkupdate "github.com/watermint/toolbox/citron/dropbox/team/sharedlink/update"
	citrondropboxteamteamfolder "github.com/watermint/toolbox/citron/dropbox/team/teamfolder"
	citrondropboxteamteamfolderbatch "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/batch"
	citrondropboxteamteamfolderfile "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/file"
	citrondropboxteamteamfolderfilelock "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/file/lock"
	citrondropboxteamteamfolderfilelockall "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/file/lock/all"
	citrondropboxteamteamfoldermember "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/member"
	citrondropboxteamteamfolderpartial "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/partial"
	citrondropboxteamteamfolderpolicy "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/policy"
	citrondropboxteamteamfoldersyncsetting "github.com/watermint/toolbox/citron/dropbox/team/teamfolder/sync/setting"
	citronfigmaaccount "github.com/watermint/toolbox/citron/figma/account"
	citronfigmafile "github.com/watermint/toolbox/citron/figma/file"
	citronfigmafileexport "github.com/watermint/toolbox/citron/figma/file/export"
	citronfigmafileexportall "github.com/watermint/toolbox/citron/figma/file/export/all"
	citronfigmaproject "github.com/watermint/toolbox/citron/figma/project"
	citrongithub "github.com/watermint/toolbox/citron/github"
	citrongithubcontent "github.com/watermint/toolbox/citron/github/content"
	citrongithubissue "github.com/watermint/toolbox/citron/github/issue"
	citrongithubrelease "github.com/watermint/toolbox/citron/github/release"
	citrongithubreleaseasset "github.com/watermint/toolbox/citron/github/release/asset"
	citrongithubtag "github.com/watermint/toolbox/citron/github/tag"
	citrongooglecalendarevent "github.com/watermint/toolbox/citron/google/calendar/event"
	citrongooglemailfilter "github.com/watermint/toolbox/citron/google/mail/filter"
	citrongooglemailfilterbatch "github.com/watermint/toolbox/citron/google/mail/filter/batch"
	citrongooglemaillabel "github.com/watermint/toolbox/citron/google/mail/label"
	citrongooglemailmessage "github.com/watermint/toolbox/citron/google/mail/message"
	citrongooglemailmessagelabel "github.com/watermint/toolbox/citron/google/mail/message/label"
	citrongooglemailmessageprocessed "github.com/watermint/toolbox/citron/google/mail/message/processed"
	citrongooglemailsendas "github.com/watermint/toolbox/citron/google/mail/sendas"
	citrongooglemailthread "github.com/watermint/toolbox/citron/google/mail/thread"
	citrongooglesheetssheet "github.com/watermint/toolbox/citron/google/sheets/sheet"
	citrongooglesheetsspreadsheet "github.com/watermint/toolbox/citron/google/sheets/spreadsheet"
	citrongoogletranslate "github.com/watermint/toolbox/citron/google/translate"
	citronlocalfiletemplate "github.com/watermint/toolbox/citron/local/file/template"
	citronslackconversation "github.com/watermint/toolbox/citron/slack/conversation"
	infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

func AutoDetectedRecipesCitron() []infra_recipe_rc_recipe.Recipe {
	return []infra_recipe_rc_recipe.Recipe{
		&citronasanateam.List{},
		&citronasanateamproject.List{},
		&citronasanateamtask.List{},
		&citronasanaworkspace.List{},
		&citronasanaworkspaceproject.List{},
		&citrondeepltranslate.Text{},
		&citrondropboxfile.Copy{},
		&citrondropboxfile.Delete{},
		&citrondropboxfile.Info{},
		&citrondropboxfile.List{},
		&citrondropboxfile.Merge{},
		&citrondropboxfile.Move{},
		&citrondropboxfile.Replication{},
		&citrondropboxfile.Size{},
		&citrondropboxfile.Watch{},
		&citrondropboxfileaccount.Feature{},
		&citrondropboxfileaccount.Filesystem{},
		&citrondropboxfileaccount.Info{},
		&citrondropboxfilecompare.Account{},
		&citrondropboxfilecompare.Local{},
		&citrondropboxfileexport.Doc{},
		&citrondropboxfileexport.Url{},
		&citrondropboxfileimport.Url{},
		&citrondropboxfileimportbatch.Url{},
		&citrondropboxfilelock.Acquire{},
		&citrondropboxfilelock.List{},
		&citrondropboxfilelock.Release{},
		&citrondropboxfilelockall.Release{},
		&citrondropboxfilelockbatch.Acquire{},
		&citrondropboxfilelockbatch.Release{},
		&citrondropboxfilerequest.Create{},
		&citrondropboxfilerequest.List{},
		&citrondropboxfilerequestdelete.Closed{},
		&citrondropboxfilerequestdelete.Url{},
		&citrondropboxfilerestore.All{},
		&citrondropboxfilerevision.Download{},
		&citrondropboxfilerevision.List{},
		&citrondropboxfilerevision.Restore{},
		&citrondropboxfilesearch.Content{},
		&citrondropboxfilesearch.Name{},
		&citrondropboxfileshare.Info{},
		&citrondropboxfilesharedfolder.Leave{},
		&citrondropboxfilesharedfolder.List{},
		&citrondropboxfilesharedfolder.Share{},
		&citrondropboxfilesharedfolder.Info{},
		&citrondropboxfilesharedfolder.Unshare{},
		&citrondropboxfilesharedfoldermember.Add{},
		&citrondropboxfilesharedfoldermember.Delete{},
		&citrondropboxfilesharedfoldermember.List{},
		&citrondropboxfilesharedfoldermount.Add{},
		&citrondropboxfilesharedfoldermount.Delete{},
		&citrondropboxfilesharedfoldermount.List{},
		&citrondropboxfilesharedfoldermount.Mountable{},
		&citrondropboxfilesharedlink.Create{},
		&citrondropboxfilesharedlink.Delete{},
		&citrondropboxfilesharedlink.Info{},
		&citrondropboxfilesharedlink.List{},
		&citrondropboxfilesharedlinkfile.List{},
		&citrondropboxfilesync.Down{},
		&citrondropboxfilesync.Online{},
		&citrondropboxfilesync.Up{},
		&citrondropboxfiletag.Add{},
		&citrondropboxfiletag.Delete{},
		&citrondropboxfiletag.List{},
		&citrondropboxfiletemplate.Apply{},
		&citrondropboxfiletemplate.Capture{},
		&citrondropboxpaper.Append{},
		&citrondropboxpaper.Create{},
		&citrondropboxpaper.Overwrite{},
		&citrondropboxpaper.Prepend{},
		&citrondropboxsignaccount.Info{},
		&citrondropboxsignrequest.List{},
		&citrondropboxsignrequestsignature.List{},
		&citrondropboxteam.Feature{},
		&citrondropboxteam.Filesystem{},
		&citrondropboxteam.Info{},
		&citrondropboxteamactivity.Event{},
		&citrondropboxteamactivity.User{},
		&citrondropboxteamactivitybatch.User{},
		&citrondropboxteamactivitydaily.Event{},
		&citrondropboxteamadmin.List{},
		&citrondropboxteamadmingrouprole.Add{},
		&citrondropboxteamadmingrouprole.Delete{},
		&citrondropboxteamadminrole.Add{},
		&citrondropboxteamadminrole.Clear{},
		&citrondropboxteamadminrole.Delete{},
		&citrondropboxteamadminrole.List{},
		&citrondropboxteambackupdevice.Status{},
		&citrondropboxteamcontentlegacypaper.Count{},
		&citrondropboxteamcontentlegacypaper.Export{},
		&citrondropboxteamcontentlegacypaper.List{},
		&citrondropboxteamcontentmember.List{},
		&citrondropboxteamcontentmember.Size{},
		&citrondropboxteamcontentmount.List{},
		&citrondropboxteamcontentpolicy.List{},
		&citrondropboxteamdevice.List{},
		&citrondropboxteamdevice.Unlink{},
		&citrondropboxteamfilerequest.Clone{},
		&citrondropboxteamfilerequest.List{},
		&citrondropboxteamgroup.Add{},
		&citrondropboxteamgroup.Delete{},
		&citrondropboxteamgroup.List{},
		&citrondropboxteamgroup.Rename{},
		&citrondropboxteamgroupbatch.Add{},
		&citrondropboxteamgroupbatch.Delete{},
		&citrondropboxteamgroupclear.Externalid{},
		&citrondropboxteamgroupfolder.List{},
		&citrondropboxteamgroupmember.Add{},
		&citrondropboxteamgroupmember.Delete{},
		&citrondropboxteamgroupmember.List{},
		&citrondropboxteamgroupmemberbatch.Add{},
		&citrondropboxteamgroupmemberbatch.Delete{},
		&citrondropboxteamgroupmemberbatch.Update{},
		&citrondropboxteamgroupupdate.Type{},
		&citrondropboxteaminsight.Scan{},
		&citrondropboxteaminsight.Scanretry{},
		&citrondropboxteaminsight.Summarize{},
		&citrondropboxteaminsightreport.Teamfoldermember{},
		&citrondropboxteamlegalhold.Add{},
		&citrondropboxteamlegalhold.List{},
		&citrondropboxteamlegalhold.Release{},
		&citrondropboxteamlegalholdmember.List{},
		&citrondropboxteamlegalholdmemberbatch.Update{},
		&citrondropboxteamlegalholdrevision.List{},
		&citrondropboxteamlegalholdupdate.Desc{},
		&citrondropboxteamlegalholdupdate.Name{},
		&citrondropboxteamlinkedapp.List{},
		&citrondropboxteammember.Feature{},
		&citrondropboxteammember.List{},
		&citrondropboxteammember.Replication{},
		&citrondropboxteammember.Suspend{},
		&citrondropboxteammember.Unsuspend{},
		&citrondropboxteammemberbatch.Delete{},
		&citrondropboxteammemberbatch.Detach{},
		&citrondropboxteammemberbatch.Invite{},
		&citrondropboxteammemberbatch.Reinvite{},
		&citrondropboxteammemberbatch.Suspend{},
		&citrondropboxteammemberbatch.Unsuspend{},
		&citrondropboxteammemberclear.Externalid{},
		&citrondropboxteammemberfile.Permdelete{},
		&citrondropboxteammemberfilelock.List{},
		&citrondropboxteammemberfilelock.Release{},
		&citrondropboxteammemberfilelockall.Release{},
		&citrondropboxteammemberfolder.List{},
		&citrondropboxteammemberfolder.Replication{},
		&citrondropboxteammemberquota.List{},
		&citrondropboxteammemberquota.Usage{},
		&citrondropboxteammemberquotabatch.Update{},
		&citrondropboxteammemberupdatebatch.Email{},
		&citrondropboxteammemberupdatebatch.Externalid{},
		&citrondropboxteammemberupdatebatch.Invisible{},
		&citrondropboxteammemberupdatebatch.Profile{},
		&citrondropboxteammemberupdatebatch.Visible{},
		&citrondropboxteamnamespace.List{},
		&citrondropboxteamnamespace.Summary{},
		&citrondropboxteamnamespacefile.List{},
		&citrondropboxteamnamespacefile.Size{},
		&citrondropboxteamnamespacemember.List{},
		&citrondropboxteamreport.Activity{},
		&citrondropboxteamreport.Devices{},
		&citrondropboxteamreport.Membership{},
		&citrondropboxteamreport.Storage{},
		&citrondropboxteamrunasfile.List{},
		&citrondropboxteamrunasfilebatch.Copy{},
		&citrondropboxteamrunasfilesyncbatch.Up{},
		&citrondropboxteamrunassharedfolder.Isolate{},
		&citrondropboxteamrunassharedfolder.List{},
		&citrondropboxteamrunassharedfolderbatch.Leave{},
		&citrondropboxteamrunassharedfolderbatch.Share{},
		&citrondropboxteamrunassharedfolderbatch.Unshare{},
		&citrondropboxteamrunassharedfoldermemberbatch.Add{},
		&citrondropboxteamrunassharedfoldermemberbatch.Delete{},
		&citrondropboxteamrunassharedfoldermount.Add{},
		&citrondropboxteamrunassharedfoldermount.Delete{},
		&citrondropboxteamrunassharedfoldermount.List{},
		&citrondropboxteamrunassharedfoldermount.Mountable{},
		&citrondropboxteamsharedlink.List{},
		&citrondropboxteamsharedlinkcap.Expiry{},
		&citrondropboxteamsharedlinkcap.Visibility{},
		&citrondropboxteamsharedlinkdelete.Links{},
		&citrondropboxteamsharedlinkdelete.Member{},
		&citrondropboxteamsharedlinkupdate.Expiry{},
		&citrondropboxteamsharedlinkupdate.Password{},
		&citrondropboxteamsharedlinkupdate.Visibility{},
		&citrondropboxteamteamfolder.Add{},
		&citrondropboxteamteamfolder.Archive{},
		&citrondropboxteamteamfolder.List{},
		&citrondropboxteamteamfolder.Permdelete{},
		&citrondropboxteamteamfolder.Replication{},
		&citrondropboxteamteamfolderbatch.Archive{},
		&citrondropboxteamteamfolderbatch.Permdelete{},
		&citrondropboxteamteamfolderbatch.Replication{},
		&citrondropboxteamteamfolderfile.List{},
		&citrondropboxteamteamfolderfile.Size{},
		&citrondropboxteamteamfolderfilelock.List{},
		&citrondropboxteamteamfolderfilelock.Release{},
		&citrondropboxteamteamfolderfilelockall.Release{},
		&citrondropboxteamteamfoldermember.Add{},
		&citrondropboxteamteamfoldermember.Delete{},
		&citrondropboxteamteamfoldermember.List{},
		&citrondropboxteamteamfolderpartial.Replication{},
		&citrondropboxteamteamfolderpolicy.List{},
		&citrondropboxteamteamfoldersyncsetting.List{},
		&citrondropboxteamteamfoldersyncsetting.Update{},
		&citronfigmaaccount.Info{},
		&citronfigmafile.Info{},
		&citronfigmafile.List{},
		&citronfigmafileexport.Frame{},
		&citronfigmafileexport.Node{},
		&citronfigmafileexport.Page{},
		&citronfigmafileexportall.Page{},
		&citronfigmaproject.List{},
		&citrongithub.Profile{},
		&citrongithubcontent.Get{},
		&citrongithubcontent.Put{},
		&citrongithubissue.List{},
		&citrongithubrelease.Draft{},
		&citrongithubrelease.List{},
		&citrongithubreleaseasset.Download{},
		&citrongithubreleaseasset.List{},
		&citrongithubreleaseasset.Upload{},
		&citrongithubtag.Create{},
		&citrongooglecalendarevent.List{},
		&citrongooglemailfilter.Add{},
		&citrongooglemailfilter.Delete{},
		&citrongooglemailfilter.List{},
		&citrongooglemailfilterbatch.Add{},
		&citrongooglemaillabel.Add{},
		&citrongooglemaillabel.Delete{},
		&citrongooglemaillabel.List{},
		&citrongooglemaillabel.Rename{},
		&citrongooglemailmessage.List{},
		&citrongooglemailmessage.Send{},
		&citrongooglemailmessagelabel.Add{},
		&citrongooglemailmessagelabel.Delete{},
		&citrongooglemailmessageprocessed.List{},
		&citrongooglemailsendas.Add{},
		&citrongooglemailsendas.Delete{},
		&citrongooglemailsendas.List{},
		&citrongooglemailthread.List{},
		&citrongooglesheetssheet.Append{},
		&citrongooglesheetssheet.Clear{},
		&citrongooglesheetssheet.Create{},
		&citrongooglesheetssheet.Delete{},
		&citrongooglesheetssheet.Export{},
		&citrongooglesheetssheet.Import{},
		&citrongooglesheetssheet.List{},
		&citrongooglesheetsspreadsheet.Create{},
		&citrongoogletranslate.Text{},
		&citronlocalfiletemplate.Apply{},
		&citronlocalfiletemplate.Capture{},
		&citronslackconversation.History{},
		&citronslackconversation.List{},
	}
}
