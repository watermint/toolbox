package api

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
)

type ApiTeam struct {
	Context *ApiContext
}

func (a *ApiTeam) Compat() team.Client {
	return team.New(a.Context.compatConfig())
}

func (a *ApiTeam) DevicesListMemberDevices(arg *team.ListMemberDevicesArg) (res *team.ListMemberDevicesResult, err error) {
	return a.Compat().DevicesListMemberDevices(arg)
}
func (a *ApiTeam) DevicesListMembersDevices(arg *team.ListMembersDevicesArg) (res *team.ListMembersDevicesResult, err error) {
	return a.Compat().DevicesListMembersDevices(arg)
}
func (a *ApiTeam) DevicesListTeamDevices(arg *team.ListTeamDevicesArg) (res *team.ListTeamDevicesResult, err error) {
	return a.Compat().DevicesListTeamDevices(arg)
}
func (a *ApiTeam) DevicesRevokeDeviceSession(arg *team.RevokeDeviceSessionArg) (err error) {
	return a.Compat().DevicesRevokeDeviceSession(arg)
}
func (a *ApiTeam) DevicesRevokeDeviceSessionBatch(arg *team.RevokeDeviceSessionBatchArg) (res *team.RevokeDeviceSessionBatchResult, err error) {
	return a.Compat().DevicesRevokeDeviceSessionBatch(arg)
}
func (a *ApiTeam) FeaturesGetValues(arg *team.FeaturesGetValuesBatchArg) (res *team.FeaturesGetValuesBatchResult, err error) {
	return a.Compat().FeaturesGetValues(arg)
}
func (a *ApiTeam) GetInfo() (res *team.TeamGetInfoResult, err error) {
	return a.Compat().GetInfo()
}
func (a *ApiTeam) GroupsCreate(arg *team.GroupCreateArg) (res *team.GroupFullInfo, err error) {
	return a.Compat().GroupsCreate(arg)
}
func (a *ApiTeam) GroupsDelete(arg *team.GroupSelector) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().GroupsDelete(arg)
}
func (a *ApiTeam) GroupsGetInfo(arg *team.GroupsSelector) (res []*team.GroupsGetInfoItem, err error) {
	return a.Compat().GroupsGetInfo(arg)
}
func (a *ApiTeam) GroupsJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	return a.Compat().GroupsJobStatusGet(arg)
}
func (a *ApiTeam) GroupsList(arg *team.GroupsListArg) (res *team.GroupsListResult, err error) {
	return a.Compat().GroupsList(arg)
}
func (a *ApiTeam) GroupsListContinue(arg *team.GroupsListContinueArg) (res *team.GroupsListResult, err error) {
	return a.Compat().GroupsListContinue(arg)
}
func (a *ApiTeam) GroupsMembersAdd(arg *team.GroupMembersAddArg) (res *team.GroupMembersChangeResult, err error) {
	return a.Compat().GroupsMembersAdd(arg)
}
func (a *ApiTeam) GroupsMembersList(arg *team.GroupsMembersListArg) (res *team.GroupsMembersListResult, err error) {
	return a.Compat().GroupsMembersList(arg)
}
func (a *ApiTeam) GroupsMembersListContinue(arg *team.GroupsMembersListContinueArg) (res *team.GroupsMembersListResult, err error) {
	return a.Compat().GroupsMembersListContinue(arg)
}
func (a *ApiTeam) GroupsMembersRemove(arg *team.GroupMembersRemoveArg) (res *team.GroupMembersChangeResult, err error) {
	return a.Compat().GroupsMembersRemove(arg)
}
func (a *ApiTeam) GroupsMembersSetAccessType(arg *team.GroupMembersSetAccessTypeArg) (res []*team.GroupsGetInfoItem, err error) {
	return a.Compat().GroupsMembersSetAccessType(arg)
}
func (a *ApiTeam) GroupsUpdate(arg *team.GroupUpdateArgs) (res *team.GroupFullInfo, err error) {
	return a.Compat().GroupsUpdate(arg)
}
func (a *ApiTeam) LinkedAppsListMemberLinkedApps(arg *team.ListMemberAppsArg) (res *team.ListMemberAppsResult, err error) {
	return a.Compat().LinkedAppsListMemberLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsListMembersLinkedApps(arg *team.ListMembersAppsArg) (res *team.ListMembersAppsResult, err error) {
	return a.Compat().LinkedAppsListMembersLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsListTeamLinkedApps(arg *team.ListTeamAppsArg) (res *team.ListTeamAppsResult, err error) {
	return a.Compat().LinkedAppsListTeamLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsRevokeLinkedApp(arg *team.RevokeLinkedApiAppArg) (err error) {
	return a.Compat().LinkedAppsRevokeLinkedApp(arg)
}
func (a *ApiTeam) LinkedAppsRevokeLinkedAppBatch(arg *team.RevokeLinkedApiAppBatchArg) (res *team.RevokeLinkedAppBatchResult, err error) {
	return a.Compat().LinkedAppsRevokeLinkedAppBatch(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersAdd(arg *team.ExcludedUsersUpdateArg) (res *team.ExcludedUsersUpdateResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersAdd(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersList(arg *team.ExcludedUsersListArg) (res *team.ExcludedUsersListResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersList(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersListContinue(arg *team.ExcludedUsersListContinueArg) (res *team.ExcludedUsersListResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersListContinue(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersRemove(arg *team.ExcludedUsersUpdateArg) (res *team.ExcludedUsersUpdateResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersRemove(arg)
}
func (a *ApiTeam) MemberSpaceLimitsGetCustomQuota(arg *team.CustomQuotaUsersArg) (res []*team.CustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsGetCustomQuota(arg)
}
func (a *ApiTeam) MemberSpaceLimitsRemoveCustomQuota(arg *team.CustomQuotaUsersArg) (res []*team.RemoveCustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsRemoveCustomQuota(arg)
}
func (a *ApiTeam) MemberSpaceLimitsSetCustomQuota(arg *team.SetCustomQuotaArg) (res []*team.CustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsSetCustomQuota(arg)
}
func (a *ApiTeam) MembersAdd(arg *team.MembersAddArg) (res *team.MembersAddLaunch, err error) {
	if r, err := a.Context.NewApiRpcRequest("team/members/add", nil, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseMembersAddLaunch(r)
	}
}
func (a *ApiTeam) MembersAddJobStatusGet(arg *async.PollArg) (res *team.MembersAddJobStatus, err error) {
	if r, err := a.Context.NewApiRpcRequest("/team/members/add/job_status/get", nil, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseMembersAddJobStatus(r)
	}
}
func (a *ApiTeam) MembersGetInfo(arg *team.MembersGetInfoArgs) (res []*team.MembersGetInfoItem, err error) {
	return a.Compat().MembersGetInfo(arg)
}
func (a *ApiTeam) MembersList(arg *team.MembersListArg) (res *team.MembersListResult, err error) {
	return a.Compat().MembersList(arg)
}
func (a *ApiTeam) MembersListContinue(arg *team.MembersListContinueArg) (res *team.MembersListResult, err error) {
	return a.Compat().MembersListContinue(arg)
}
func (a *ApiTeam) MembersRecover(arg *team.MembersRecoverArg) (err error) {
	return a.Compat().MembersRecover(arg)
}
func (a *ApiTeam) MembersRemove(arg *team.MembersRemoveArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().MembersRemove(arg)
}
func (a *ApiTeam) MembersRemoveJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	return a.Compat().MembersRemoveJobStatusGet(arg)
}
func (a *ApiTeam) MembersSendWelcomeEmail(arg *team.UserSelectorArg) (err error) {
	return a.Compat().MembersSendWelcomeEmail(arg)
}
func (a *ApiTeam) MembersSetAdminPermissions(arg *team.MembersSetPermissionsArg) (res *team.MembersSetPermissionsResult, err error) {
	return a.Compat().MembersSetAdminPermissions(arg)
}
func (a *ApiTeam) MembersSetProfile(arg *team.MembersSetProfileArg) (res *team.TeamMemberInfo, err error) {
	return a.Compat().MembersSetProfile(arg)
}
func (a *ApiTeam) MembersSuspend(arg *team.MembersDeactivateArg) (err error) {
	return a.Compat().MembersSuspend(arg)
}
func (a *ApiTeam) MembersUnsuspend(arg *team.MembersUnsuspendArg) (err error) {
	return a.Compat().MembersUnsuspend(arg)
}
func (a *ApiTeam) NamespacesList(arg *team.TeamNamespacesListArg) (res *team.TeamNamespacesListResult, err error) {
	return a.Compat().NamespacesList(arg)
}
func (a *ApiTeam) NamespacesListContinue(arg *team.TeamNamespacesListContinueArg) (res *team.TeamNamespacesListResult, err error) {
	return a.Compat().NamespacesListContinue(arg)
}
func (a *ApiTeam) PropertiesTemplateAdd(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().PropertiesTemplateAdd(arg)
}
func (a *ApiTeam) PropertiesTemplateGet(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().PropertiesTemplateGet(arg)
}
func (a *ApiTeam) PropertiesTemplateList() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().PropertiesTemplateList()
}
func (a *ApiTeam) PropertiesTemplateUpdate(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().PropertiesTemplateUpdate(arg)
}
func (a *ApiTeam) ReportsGetActivity(arg *team.DateRange) (res *team.GetActivityReport, err error) {
	return a.Compat().ReportsGetActivity(arg)
}
func (a *ApiTeam) ReportsGetDevices(arg *team.DateRange) (res *team.GetDevicesReport, err error) {
	return a.Compat().ReportsGetDevices(arg)
}
func (a *ApiTeam) ReportsGetMembership(arg *team.DateRange) (res *team.GetMembershipReport, err error) {
	return a.Compat().ReportsGetMembership(arg)
}
func (a *ApiTeam) ReportsGetStorage(arg *team.DateRange) (res *team.GetStorageReport, err error) {
	return a.Compat().ReportsGetStorage(arg)
}
func (a *ApiTeam) TeamFolderActivate(arg *team.TeamFolderIdArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderActivate(arg)
}
func (a *ApiTeam) TeamFolderArchive(arg *team.TeamFolderArchiveArg) (res *team.TeamFolderArchiveLaunch, err error) {
	return a.Compat().TeamFolderArchive(arg)
}
func (a *ApiTeam) TeamFolderArchiveCheck(arg *async.PollArg) (res *team.TeamFolderArchiveJobStatus, err error) {
	return a.Compat().TeamFolderArchiveCheck(arg)
}
func (a *ApiTeam) TeamFolderCreate(arg *team.TeamFolderCreateArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderCreate(arg)
}
func (a *ApiTeam) TeamFolderGetInfo(arg *team.TeamFolderIdListArg) (res []*team.TeamFolderGetInfoItem, err error) {
	return a.Compat().TeamFolderGetInfo(arg)
}
func (a *ApiTeam) TeamFolderList(arg *team.TeamFolderListArg) (res *team.TeamFolderListResult, err error) {
	return a.Compat().TeamFolderList(arg)
}
func (a *ApiTeam) TeamFolderListContinue(arg *team.TeamFolderListContinueArg) (res *team.TeamFolderListResult, err error) {
	return a.Compat().TeamFolderListContinue(arg)
}
func (a *ApiTeam) TeamFolderPermanentlyDelete(arg *team.TeamFolderIdArg) (err error) {
	return a.Compat().TeamFolderPermanentlyDelete(arg)
}
func (a *ApiTeam) TeamFolderRename(arg *team.TeamFolderRenameArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderRename(arg)
}
func (a *ApiTeam) TokenGetAuthenticatedAdmin() (res *team.TokenGetAuthenticatedAdminResult, err error) {
	return a.Compat().TokenGetAuthenticatedAdmin()
}
