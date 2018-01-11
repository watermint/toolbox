package api

import (
	"encoding/json"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_requests"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/paper"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
)

func parseLaunchEmptyResult(res *ApiRpcResponse) (r *async.LaunchEmptyResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseLaunchResultBase(res *ApiRpcResponse) (r *async.LaunchResultBase, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parsePollEmptyResult(res *ApiRpcResponse) (r *async.PollEmptyResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseAddTemplateResult(res *ApiRpcResponse) (r *file_properties.AddTemplateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetTemplateResult(res *ApiRpcResponse) (r *file_properties.GetTemplateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListTemplateResult(res *ApiRpcResponse) (r *file_properties.ListTemplateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parsePropertiesSearchResult(res *ApiRpcResponse) (r *file_properties.PropertiesSearchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseUpdateTemplateResult(res *ApiRpcResponse) (r *file_properties.UpdateTemplateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFileRequest(res *ApiRpcResponse) (r *file_requests.FileRequest, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFileRequestsResult(res *ApiRpcResponse) (r *file_requests.ListFileRequestsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseCreateFolderResult(res *ApiRpcResponse) (r *files.CreateFolderResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseDeleteBatchJobStatus(res *ApiRpcResponse) (r *files.DeleteBatchJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseDeleteBatchLaunch(res *ApiRpcResponse) (r *files.DeleteBatchLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseDeleteResult(res *ApiRpcResponse) (r *files.DeleteResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseDownloadZipResult(res *ApiRpcResponse) (r *files.DownloadZipResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFileMetadata(res *ApiRpcResponse) (r *files.FileMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFolderMetadata(res *ApiRpcResponse) (r *files.FolderMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetCopyReferenceResult(res *ApiRpcResponse) (r *files.GetCopyReferenceResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetTemporaryLinkResult(res *ApiRpcResponse) (r *files.GetTemporaryLinkResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetThumbnailBatchResult(res *ApiRpcResponse) (r *files.GetThumbnailBatchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFolderGetLatestCursorResult(res *ApiRpcResponse) (r *files.ListFolderGetLatestCursorResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFolderLongpollResult(res *ApiRpcResponse) (r *files.ListFolderLongpollResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFolderResult(res *ApiRpcResponse) (r *files.ListFolderResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListRevisionsResult(res *ApiRpcResponse) (r *files.ListRevisionsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRelocationBatchJobStatus(res *ApiRpcResponse) (r *files.RelocationBatchJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRelocationBatchLaunch(res *ApiRpcResponse) (r *files.RelocationBatchLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRelocationResult(res *ApiRpcResponse) (r *files.RelocationResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSaveCopyReferenceResult(res *ApiRpcResponse) (r *files.SaveCopyReferenceResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSaveUrlJobStatus(res *ApiRpcResponse) (r *files.SaveUrlJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSaveUrlResult(res *ApiRpcResponse) (r *files.SaveUrlResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSearchResult(res *ApiRpcResponse) (r *files.SearchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseUploadSessionFinishBatchJobStatus(res *ApiRpcResponse) (r *files.UploadSessionFinishBatchJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseUploadSessionFinishBatchLaunch(res *ApiRpcResponse) (r *files.UploadSessionFinishBatchLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseUploadSessionStartResult(res *ApiRpcResponse) (r *files.UploadSessionStartResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFoldersContainingPaperDoc(res *ApiRpcResponse) (r *paper.FoldersContainingPaperDoc, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListPaperDocsResponse(res *ApiRpcResponse) (r *paper.ListPaperDocsResponse, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListUsersOnFolderResponse(res *ApiRpcResponse) (r *paper.ListUsersOnFolderResponse, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListUsersOnPaperDocResponse(res *ApiRpcResponse) (r *paper.ListUsersOnPaperDocResponse, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parsePaperDocCreateUpdateResult(res *ApiRpcResponse) (r *paper.PaperDocCreateUpdateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parsePaperDocExportResult(res *ApiRpcResponse) (r *paper.PaperDocExportResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSharingPolicy(res *ApiRpcResponse) (r *paper.SharingPolicy, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFileMemberActionIndividualResult(res *ApiRpcResponse) (r *sharing.FileMemberActionIndividualResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFileMemberActionResult(res *ApiRpcResponse) (r *sharing.FileMemberActionResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFileMemberRemoveActionResult(res *ApiRpcResponse) (r *sharing.FileMemberRemoveActionResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetSharedLinksResult(res *ApiRpcResponse) (r *sharing.GetSharedLinksResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseJobStatus(res *ApiRpcResponse) (r *sharing.JobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFilesResult(res *ApiRpcResponse) (r *sharing.ListFilesResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListFoldersResult(res *ApiRpcResponse) (r *sharing.ListFoldersResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListSharedLinksResult(res *ApiRpcResponse) (r *sharing.ListSharedLinksResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseMemberAccessLevelResult(res *ApiRpcResponse) (r *sharing.MemberAccessLevelResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parsePathLinkMetadata(res *ApiRpcResponse) (r *sharing.PathLinkMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRemoveMemberJobStatus(res *ApiRpcResponse) (r *sharing.RemoveMemberJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSharedFileMembers(res *ApiRpcResponse) (r *sharing.SharedFileMembers, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSharedFileMetadata(res *ApiRpcResponse) (r *sharing.SharedFileMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSharedFolderMembers(res *ApiRpcResponse) (r *sharing.SharedFolderMembers, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSharedFolderMetadata(res *ApiRpcResponse) (r *sharing.SharedFolderMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseShareFolderJobStatus(res *ApiRpcResponse) (r *sharing.ShareFolderJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseShareFolderLaunch(res *ApiRpcResponse) (r *sharing.ShareFolderLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseExcludedUsersListResult(res *ApiRpcResponse) (r *team.ExcludedUsersListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseExcludedUsersUpdateResult(res *ApiRpcResponse) (r *team.ExcludedUsersUpdateResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFeaturesGetValuesBatchResult(res *ApiRpcResponse) (r *team.FeaturesGetValuesBatchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetActivityReport(res *ApiRpcResponse) (r *team.GetActivityReport, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetDevicesReport(res *ApiRpcResponse) (r *team.GetDevicesReport, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetMembershipReport(res *ApiRpcResponse) (r *team.GetMembershipReport, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetStorageReport(res *ApiRpcResponse) (r *team.GetStorageReport, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGroupFullInfo(res *ApiRpcResponse) (r *team.GroupFullInfo, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGroupMembersChangeResult(res *ApiRpcResponse) (r *team.GroupMembersChangeResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGroupsListResult(res *ApiRpcResponse) (r *team.GroupsListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGroupsMembersListResult(res *ApiRpcResponse) (r *team.GroupsMembersListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListMemberAppsResult(res *ApiRpcResponse) (r *team.ListMemberAppsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListMemberDevicesResult(res *ApiRpcResponse) (r *team.ListMemberDevicesResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListMembersAppsResult(res *ApiRpcResponse) (r *team.ListMembersAppsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListMembersDevicesResult(res *ApiRpcResponse) (r *team.ListMembersDevicesResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListTeamAppsResult(res *ApiRpcResponse) (r *team.ListTeamAppsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseListTeamDevicesResult(res *ApiRpcResponse) (r *team.ListTeamDevicesResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseMembersAddJobStatus(res *ApiRpcResponse) (r *team.MembersAddJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseMembersAddLaunch(res *ApiRpcResponse) (r *team.MembersAddLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseMembersListResult(res *ApiRpcResponse) (r *team.MembersListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseMembersSetPermissionsResult(res *ApiRpcResponse) (r *team.MembersSetPermissionsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRevokeDeviceSessionBatchResult(res *ApiRpcResponse) (r *team.RevokeDeviceSessionBatchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseRevokeLinkedAppBatchResult(res *ApiRpcResponse) (r *team.RevokeLinkedAppBatchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamFolderArchiveJobStatus(res *ApiRpcResponse) (r *team.TeamFolderArchiveJobStatus, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamFolderArchiveLaunch(res *ApiRpcResponse) (r *team.TeamFolderArchiveLaunch, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamFolderListResult(res *ApiRpcResponse) (r *team.TeamFolderListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamFolderMetadata(res *ApiRpcResponse) (r *team.TeamFolderMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamGetInfoResult(res *ApiRpcResponse) (r *team.TeamGetInfoResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamMemberInfo(res *ApiRpcResponse) (r *team.TeamMemberInfo, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTeamNamespacesListResult(res *ApiRpcResponse) (r *team.TeamNamespacesListResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseTokenGetAuthenticatedAdminResult(res *ApiRpcResponse) (r *team.TokenGetAuthenticatedAdminResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseGetTeamEventsResult(res *ApiRpcResponse) (r *team_log.GetTeamEventsResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseBasicAccount(res *ApiRpcResponse) (r *users.BasicAccount, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseFullAccount(res *ApiRpcResponse) (r *users.FullAccount, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseSpaceUsage(res *ApiRpcResponse) (r *users.SpaceUsage, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseIsMetadata(res *ApiRpcResponse) (r files.IsMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
func parseIsSharedLinkMetadata(res *ApiRpcResponse) (r sharing.IsSharedLinkMetadata, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}
