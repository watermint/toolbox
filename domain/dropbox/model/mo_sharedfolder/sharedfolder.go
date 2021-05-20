package mo_sharedfolder

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
)

type SharedFolder struct {
	Raw                  json.RawMessage
	SharedFolderId       string `path:"shared_folder_id" json:"shared_folder_id"`
	ParentSharedFolderId string `path:"parent_shared_folder_id" json:"parent_shared_folder_id"`
	Name                 string `path:"name" json:"name"`
	AccessType           string `path:"access_type.\\.tag" json:"access_type"`
	PathLower            string `path:"path_lower" json:"path_lower"`
	IsInsideTeamFolder   bool   `path:"is_inside_team_folder" json:"is_inside_team_folder"`
	IsTeamFolder         bool   `path:"is_team_folder" json:"is_team_folder"`
	PolicyManageAccess   string `path:"policy.acl_update_policy.\\.tag" json:"policy_manage_access"`
	PolicySharedLink     string `path:"policy.shared_link_policy.\\.tag" json:"policy_shared_link"`
	PolicyMember         string `path:"policy.resolved_member_policy.\\.tag" json:"policy_member"`
	PolicyViewerInfo     string `path:"policy.viewer_info_policy.\\.tag" json:"policy_viewer_info"`
	OwnerTeamId          string `path:"owner_team.id" json:"owner_team_id"`
	OwnerTeamName        string `path:"owner_team.name" json:"owner_team_name"`
	AccessInheritance    string `path:"access_inheritance.\\.tag" json:"access_inheritance"`
}

func (z SharedFolder) IsNoInherit() bool {
	return z.AccessInheritance == "no_inherit"
}

type MemberMount struct {
	TeamMemberId          string `json:"team_member_id"`
	TeamMemberDisplayName string `json:"team_member_display_name"`
	TeamMemberEmail       string `json:"team_member_email"`
	NamespaceId           string `json:"namespace_id"`
	NamespaceName         string `json:"namespace_name"`
	AccessType            string `json:"access_type"`
	MountPath             string `json:"mount_path"`
	IsInsideTeamFolder    bool   `path:"is_inside_team_folder" json:"is_inside_team_folder"`
	IsTeamFolder          bool   `path:"is_team_folder" json:"is_team_folder"`
	PolicyManageAccess    string `path:"policy.acl_update_policy.\\.tag" json:"policy_manage_access"`
	PolicySharedLink      string `path:"policy.shared_link_policy.\\.tag" json:"policy_shared_link"`
	PolicyMember          string `path:"policy.resolved_member_policy.\\.tag" json:"policy_member"`
	PolicyViewerInfo      string `path:"policy.viewer_info_policy.\\.tag" json:"policy_viewer_info"`
	OwnerTeamId           string `path:"owner_team.id" json:"owner_team_id"`
	OwnerTeamName         string `path:"owner_team.name" json:"owner_team_name"`
}

func NewMemberMount(member *mo_member.Member, sf *SharedFolder) *MemberMount {
	return &MemberMount{
		TeamMemberId:          member.TeamMemberId,
		TeamMemberDisplayName: member.DisplayName,
		TeamMemberEmail:       member.Email,
		NamespaceId:           sf.SharedFolderId,
		NamespaceName:         sf.Name,
		AccessType:            sf.AccessType,
		MountPath:             sf.PathLower,
		IsInsideTeamFolder:    sf.IsInsideTeamFolder,
		IsTeamFolder:          sf.IsTeamFolder,
		PolicyManageAccess:    sf.PolicyManageAccess,
		PolicySharedLink:      sf.PolicySharedLink,
		PolicyMember:          sf.PolicyMember,
		PolicyViewerInfo:      sf.PolicyViewerInfo,
		OwnerTeamId:           sf.OwnerTeamId,
		OwnerTeamName:         sf.OwnerTeamName,
	}
}
