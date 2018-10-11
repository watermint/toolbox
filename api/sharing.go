package api

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"io"
)

type ApiSharing struct {
	Context *ApiContext
}

func (a *ApiSharing) Compat() sharing.Client {
	return sharing.New(a.Context.compatConfig())
}

func (a *ApiSharing) AddFileMember(arg *sharing.AddFileMemberArgs) (res []*sharing.FileMemberActionResult, err error) {
	return a.Compat().AddFileMember(arg)
}
func (a *ApiSharing) AddFolderMember(arg *sharing.AddFolderMemberArg) (err error) {
	return a.Compat().AddFolderMember(arg)
}
func (a *ApiSharing) ChangeFileMemberAccess(arg *sharing.ChangeFileMemberAccessArgs) (res *sharing.FileMemberActionResult, err error) {
	return a.Compat().ChangeFileMemberAccess(arg)
}
func (a *ApiSharing) CheckJobStatus(arg *async.PollArg) (res *sharing.JobStatus, err error) {
	return a.Compat().CheckJobStatus(arg)
}
func (a *ApiSharing) CheckRemoveMemberJobStatus(arg *async.PollArg) (res *sharing.RemoveMemberJobStatus, err error) {
	return a.Compat().CheckRemoveMemberJobStatus(arg)
}
func (a *ApiSharing) CheckShareJobStatus(arg *async.PollArg) (res *sharing.ShareFolderJobStatus, err error) {
	return a.Compat().CheckShareJobStatus(arg)
}
func (a *ApiSharing) CreateSharedLink(arg *sharing.CreateSharedLinkArg) (res *sharing.PathLinkMetadata, err error) {
	return a.Compat().CreateSharedLink(arg)
}
func (a *ApiSharing) CreateSharedLinkWithSettings(arg *sharing.CreateSharedLinkWithSettingsArg) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().CreateSharedLinkWithSettings(arg)
}
func (a *ApiSharing) GetFileMetadata(arg *sharing.GetFileMetadataArg) (res *sharing.SharedFileMetadata, err error) {
	return a.Compat().GetFileMetadata(arg)
}
func (a *ApiSharing) GetFileMetadataBatch(arg *sharing.GetFileMetadataBatchArg) (res []*sharing.GetFileMetadataBatchResult, err error) {
	return a.Compat().GetFileMetadataBatch(arg)
}
func (a *ApiSharing) GetFolderMetadata(arg *sharing.GetMetadataArgs) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().GetFolderMetadata(arg)
}
func (a *ApiSharing) GetSharedLinkFile(arg *sharing.GetSharedLinkMetadataArg) (res sharing.IsSharedLinkMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetSharedLinkFile(arg)
}
func (a *ApiSharing) GetSharedLinkMetadata(arg *sharing.GetSharedLinkMetadataArg) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().GetSharedLinkMetadata(arg)
}
func (a *ApiSharing) GetSharedLinks(arg *sharing.GetSharedLinksArg) (res *sharing.GetSharedLinksResult, err error) {
	return a.Compat().GetSharedLinks(arg)
}
func (a *ApiSharing) ListFileMembers(arg *sharing.ListFileMembersArg) (res *sharing.SharedFileMembers, err error) {
	return a.Compat().ListFileMembers(arg)
}
func (a *ApiSharing) ListFileMembersBatch(arg *sharing.ListFileMembersBatchArg) (res []*sharing.ListFileMembersBatchResult, err error) {
	return a.Compat().ListFileMembersBatch(arg)
}
func (a *ApiSharing) ListFileMembersContinue(arg *sharing.ListFileMembersContinueArg) (res *sharing.SharedFileMembers, err error) {
	return a.Compat().ListFileMembersContinue(arg)
}
func (a *ApiSharing) ListFolderMembers(arg *sharing.ListFolderMembersArgs) (res *sharing.SharedFolderMembers, err error) {
	return a.Compat().ListFolderMembers(arg)
}
func (a *ApiSharing) ListFolderMembersContinue(arg *sharing.ListFolderMembersContinueArg) (res *sharing.SharedFolderMembers, err error) {
	return a.Compat().ListFolderMembersContinue(arg)
}
func (a *ApiSharing) ListFolders(arg *sharing.ListFoldersArgs) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListFolders(arg)
}
func (a *ApiSharing) ListFoldersContinue(arg *sharing.ListFoldersContinueArg) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListFoldersContinue(arg)
}
func (a *ApiSharing) ListMountableFolders(arg *sharing.ListFoldersArgs) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListMountableFolders(arg)
}
func (a *ApiSharing) ListMountableFoldersContinue(arg *sharing.ListFoldersContinueArg) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListMountableFoldersContinue(arg)
}
func (a *ApiSharing) ListReceivedFiles(arg *sharing.ListFilesArg) (res *sharing.ListFilesResult, err error) {
	return a.Compat().ListReceivedFiles(arg)
}
func (a *ApiSharing) ListReceivedFilesContinue(arg *sharing.ListFilesContinueArg) (res *sharing.ListFilesResult, err error) {
	return a.Compat().ListReceivedFilesContinue(arg)
}
func (a *ApiSharing) ListSharedLinks(arg *sharing.ListSharedLinksArg) (res *sharing.ListSharedLinksResult, err error) {
	return a.Compat().ListSharedLinks(arg)
}
func (a *ApiSharing) ModifySharedLinkSettings(arg *sharing.ModifySharedLinkSettingsArgs) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().ModifySharedLinkSettings(arg)
}
func (a *ApiSharing) MountFolder(arg *sharing.MountFolderArg) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().MountFolder(arg)
}
func (a *ApiSharing) RelinquishFileMembership(arg *sharing.RelinquishFileMembershipArg) (err error) {
	return a.Compat().RelinquishFileMembership(arg)
}
func (a *ApiSharing) RelinquishFolderMembership(arg *sharing.RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().RelinquishFolderMembership(arg)
}
func (a *ApiSharing) RemoveFileMember(arg *sharing.RemoveFileMemberArg) (res *sharing.FileMemberActionIndividualResult, err error) {
	return a.Compat().RemoveFileMember(arg)
}
func (a *ApiSharing) RemoveFileMember2(arg *sharing.RemoveFileMemberArg) (res *sharing.FileMemberRemoveActionResult, err error) {
	return a.Compat().RemoveFileMember2(arg)
}
func (a *ApiSharing) RemoveFolderMember(arg *sharing.RemoveFolderMemberArg) (res *async.LaunchResultBase, err error) {
	return a.Compat().RemoveFolderMember(arg)
}
func (a *ApiSharing) RevokeSharedLink(arg *sharing.RevokeSharedLinkArg) (err error) {
	return a.Compat().RevokeSharedLink(arg)
}
func (a *ApiSharing) SetAccessInheritance(arg *sharing.SetAccessInheritanceArg) (res *sharing.ShareFolderLaunch, err error) {
	return a.Compat().SetAccessInheritance(arg)
}
func (a *ApiSharing) ShareFolder(arg *sharing.ShareFolderArg) (res *sharing.ShareFolderLaunch, err error) {
	return a.Compat().ShareFolder(arg)
}
func (a *ApiSharing) TransferFolder(arg *sharing.TransferFolderArg) (err error) {
	return a.Compat().TransferFolder(arg)
}
func (a *ApiSharing) UnmountFolder(arg *sharing.UnmountFolderArg) (err error) {
	return a.Compat().UnmountFolder(arg)
}
func (a *ApiSharing) UnshareFile(arg *sharing.UnshareFileArg) (err error) {
	return a.Compat().UnshareFile(arg)
}
func (a *ApiSharing) UnshareFolder(arg *sharing.UnshareFolderArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().UnshareFolder(arg)
}
func (a *ApiSharing) UpdateFileMember(arg *sharing.UpdateFileMemberArgs) (res *sharing.MemberAccessLevelResult, err error) {
	return a.Compat().UpdateFileMember(arg)
}
func (a *ApiSharing) UpdateFolderMember(arg *sharing.UpdateFolderMemberArg) (res *sharing.MemberAccessLevelResult, err error) {
	return a.Compat().UpdateFolderMember(arg)
}
func (a *ApiSharing) UpdateFolderPolicy(arg *sharing.UpdateFolderPolicyArg) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().UpdateFolderPolicy(arg)
}
