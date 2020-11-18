package uc_team_content

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
)

type Membership struct {
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	FolderType    string `json:"folder_type"`
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	AccessType    string `json:"access_type"`
	MemberType    string `json:"member_type"`
	MemberId      string `json:"member_id"`
	MemberName    string `json:"member_name"`
	MemberEmail   string `json:"member_email"`
	SameTeam      string `json:"same_team"`
}

type NoMember struct {
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	FolderType    string `json:"folder_type"`
}

type FolderPolicy struct {
	NamespaceId        string `json:"namespace_id"`
	NamespaceName      string `json:"namespace_name"`
	Path               string `json:"path"`
	IsTeamFolder       bool   `json:"is_team_folder"`
	OwnerTeamId        string `json:"owner_team_id"`
	OwnerTeamName      string `json:"owner_team_name"`
	PolicyManageAccess string `json:"policy_manage_access"`
	PolicySharedLink   string `json:"policy_shared_link"`
	PolicyMember       string `json:"policy_member"`
	PolicyViewerInfo   string `json:"policy_viewer_info"`
}

type GroupToFolder struct {
	GroupName       string `json:"group_name"`
	GroupType       string `json:"group_type"`
	GroupIsSameTeam bool   `json:"group_is_same_team"`
	AccessType      string `json:"access_type"`
	NamespaceId     string `json:"namespace_id"`
	NamespaceName   string `json:"namespace_name"`
	Path            string `json:"path"`
	FolderType      string `json:"folder_type"`
	OwnerTeamId     string `json:"owner_team_id"`
	OwnerTeamName   string `json:"owner_team_name"`
}

func NewMembership(sf *mo_sharedfolder.SharedFolder, path string, member *mo_sharedfolder_member.Metadata) *Membership {
	var memberId, memberName, memberEmail string
	if u, ok := member.User(); ok {
		memberId = u.AccountId
		memberName = u.DisplayName
		memberEmail = u.Email
	}
	if g, ok := member.Group(); ok {
		memberId = g.GroupId
		memberName = g.GroupName
	}
	if e, ok := member.Invitee(); ok {
		memberEmail = e.InviteeEmail
	}

	return &Membership{
		OwnerTeamId:   sf.OwnerTeamId,
		OwnerTeamName: sf.OwnerTeamName,
		NamespaceId:   sf.SharedFolderId,
		NamespaceName: sf.Name,
		Path:          folderPath(sf, path),
		FolderType:    folderType(sf),
		AccessType:    member.AccessType(),
		MemberType:    member.MemberType(),
		MemberId:      memberId,
		MemberName:    memberName,
		MemberEmail:   memberEmail,
		SameTeam:      member.SameTeam(),
	}
}

func NewNoMember(sf *mo_sharedfolder.SharedFolder, path string) *NoMember {
	return &NoMember{
		OwnerTeamId:   sf.OwnerTeamId,
		OwnerTeamName: sf.OwnerTeamName,
		NamespaceId:   sf.SharedFolderId,
		NamespaceName: sf.Name,
		Path:          folderPath(sf, path),
		FolderType:    folderType(sf),
	}
}

func NewFolderPolicy(sf *mo_sharedfolder.SharedFolder, path string) *FolderPolicy {
	return &FolderPolicy{
		NamespaceId:        sf.SharedFolderId,
		NamespaceName:      sf.Name,
		Path:               folderPath(sf, path),
		IsTeamFolder:       sf.IsTeamFolder || sf.IsInsideTeamFolder,
		OwnerTeamId:        sf.OwnerTeamId,
		OwnerTeamName:      sf.OwnerTeamName,
		PolicyManageAccess: sf.PolicyManageAccess,
		PolicySharedLink:   sf.PolicySharedLink,
		PolicyMember:       sf.PolicyMember,
		PolicyViewerInfo:   sf.PolicyViewerInfo,
	}
}

func NewGroupToFolder(g *mo_sharedfolder_member.Group, sf *mo_sharedfolder.SharedFolder, path string, member *mo_sharedfolder_member.Metadata) *GroupToFolder {
	return &GroupToFolder{
		GroupName:       g.GroupName,
		GroupType:       g.GroupManagementType,
		GroupIsSameTeam: g.IsSameTeam,
		AccessType:      member.AccessType(),
		NamespaceId:     sf.SharedFolderId,
		NamespaceName:   sf.Name,
		Path:            folderPath(sf, path),
		FolderType:      folderType(sf),
		OwnerTeamId:     sf.OwnerTeamId,
		OwnerTeamName:   sf.OwnerTeamName,
	}
}

func folderType(m *mo_sharedfolder.SharedFolder) string {
	switch {
	case m.IsTeamFolder, m.IsInsideTeamFolder:
		return "team_folder"
	default:
		return "shared_folder"
	}
}

func folderPath(m *mo_sharedfolder.SharedFolder, path string) string {
	if path == "" {
		return m.Name
	} else {
		return path
	}
}
