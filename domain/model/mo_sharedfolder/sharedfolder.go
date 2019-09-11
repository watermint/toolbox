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
	PolicyMember         string `path:"policy.member_policy.\\.tag" json:"policy_member"`
}
