package mo_sharedfolder

import "encoding/json"

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
}
