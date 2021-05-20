package uc_team_content

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
)

const (
	MaximumTeamFolderMemberCount              = 1000
	MaximumTeamFolderMemberCountWithNoInherit = 250
	MaximumSharedFolderMemberCount            = 1000
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

type MemberCount struct {
	NamespaceId         string          `json:"namespace_id"`
	NamespaceName       string          `json:"namespace_name"`
	ParentNamespaceId   string          `json:"parent_namespace_id"`
	Path                string          `json:"path"`
	FolderType          string          `json:"folder_type"`
	OwnerTeamId         string          `json:"owner_team_id"`
	OwnerTeamName       string          `json:"owner_team_name"`
	HasNoInherit        bool            `json:"has_no_inherit"`
	IsNoInherit         bool            `json:"is_no_inherit"`
	Capacity            *int            `json:"capacity"`
	CountTotal          int             `json:"count_total"`
	CountExternalGroups int             `json:"count_external_groups"`
	MemberEmails        map[string]bool `json:"member_emails"`
}

func (z MemberCount) Merge(other MemberCount) MemberCount {
	if other.HasNoInherit {
		z.HasNoInherit = true
	}
	emails := make(map[string]bool)
	for email := range z.MemberEmails {
		emails[email] = true
	}
	for email := range other.MemberEmails {
		emails[email] = true
	}

	z.MemberEmails = emails
	z.CountTotal = len(emails)
	z.CountExternalGroups += other.CountExternalGroups
	if z.ParentNamespaceId == "" {
		z.Capacity = MemberCapacity(z.HasNoInherit, z.CountExternalGroups, z.CountTotal)
	}

	return z
}

func MemberCapacity(hasNoInherit bool, countExternalGroups, countTotal int) *int {
	// unknown
	if 0 < countExternalGroups {
		return nil
	}

	var capacity int
	if hasNoInherit {
		capacity = MaximumTeamFolderMemberCountWithNoInherit - countTotal
	} else {
		capacity = MaximumTeamFolderMemberCount - countTotal
	}
	return &capacity
}

func MemberCountFromSharedFolder(path string, sf *mo_sharedfolder.SharedFolder) MemberCount {
	return MemberCount{
		NamespaceId:         sf.SharedFolderId,
		NamespaceName:       sf.Name,
		ParentNamespaceId:   sf.ParentSharedFolderId,
		Path:                folderPath(sf, path),
		FolderType:          folderType(sf),
		OwnerTeamId:         sf.OwnerTeamId,
		OwnerTeamName:       sf.OwnerTeamName,
		HasNoInherit:        sf.IsNoInherit(),
		IsNoInherit:         sf.IsNoInherit(),
		Capacity:            MemberCapacity(sf.IsNoInherit(), 0, 0),
		CountTotal:          0,
		CountExternalGroups: 0,
		MemberEmails:        make(map[string]bool),
	}
}

func MemberCountFromMemberEntry(path string, sf *mo_sharedfolder.SharedFolder, member mo_sharedfolder_member.Member, gd sv_group_member.GroupDirectory) (MemberCount, error) {
	mc := MemberCountFromSharedFolder(path, sf)
	defer func() {
		mc.Capacity = MemberCapacity(mc.HasNoInherit, mc.CountExternalGroups, mc.CountTotal)
	}()
	if g, ok := member.Group(); ok {
		if g.IsSameTeam {
			groupMembers, err := gd.List(g.GroupId)
			if err != nil {
				return MemberCount{}, err
			}
			for _, gm := range groupMembers {
				mc.MemberEmails[gm.Email] = true
			}
			mc.CountTotal = len(mc.MemberEmails)

		} else {
			mc.CountExternalGroups++
		}
		return mc, nil
	} else if u, ok := member.User(); ok {
		mc.MemberEmails[u.Email] = true
		mc.CountTotal = len(mc.MemberEmails)
		return mc, nil
	} else if v, ok := member.Invitee(); ok {
		mc.MemberEmails[v.InviteeEmail] = true
		mc.CountTotal = len(mc.MemberEmails)
		return mc, nil
	} else {
		return MemberCount{}, errors.New("unknown member type")
	}
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

type MemberToFolder struct {
	MemberName    string `json:"member_name"`
	TeamMemberId  string `json:"team_member_id"`
	MemberEmail   string `json:"member_email"`
	AccessType    string `json:"access_type"`
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	FolderType    string `json:"folder_type"`
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
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

func NewMemberToFolderByUser(u *mo_sharedfolder_member.User, sf *mo_sharedfolder.SharedFolder, path string, member *mo_sharedfolder_member.Metadata) *MemberToFolder {
	return &MemberToFolder{
		MemberName:    u.DisplayName,
		TeamMemberId:  u.TeamMemberId,
		MemberEmail:   u.Email,
		AccessType:    member.AccessType(),
		NamespaceId:   sf.SharedFolderId,
		NamespaceName: sf.Name,
		Path:          folderPath(sf, path),
		FolderType:    folderType(sf),
		OwnerTeamId:   sf.OwnerTeamId,
		OwnerTeamName: sf.OwnerTeamName,
	}
}

func NewMemberToFolderByInvitee(u *mo_sharedfolder_member.Invitee, sf *mo_sharedfolder.SharedFolder, path string, member *mo_sharedfolder_member.Metadata) *MemberToFolder {
	return &MemberToFolder{
		MemberName:    "",
		TeamMemberId:  "",
		MemberEmail:   u.InviteeEmail,
		AccessType:    member.AccessType(),
		NamespaceId:   sf.SharedFolderId,
		NamespaceName: sf.Name,
		Path:          folderPath(sf, path),
		FolderType:    folderType(sf),
		OwnerTeamId:   sf.OwnerTeamId,
		OwnerTeamName: sf.OwnerTeamName,
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
