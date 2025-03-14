package mo_team

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
)

type Info struct {
	Raw                      json.RawMessage
	Name                     string `path:"name" json:"name"`
	TeamId                   string `path:"team_id" json:"team_id"`
	NumLicensedUsers         int    `path:"num_licensed_users" json:"num_licensed_users"`
	NumProvisionedUsers      int    `path:"num_provisioned_users" json:"num_provisioned_users"`
	PolicySharedFolderMember string `path:"policies.sharing.shared_folder_member_policy.\\.tag" json:"policy_shared_folder_member"`
	PolicySharedFolderJoin   string `path:"policies.sharing.shared_folder_join_policy.\\.tag" json:"policy_shared_folder_join"`
	PolicySharedLinkCreate   string `path:"policies.sharing.shared_link_create_policy.\\.tag" json:"policy_shared_link_create"`
	PolicyEmmState           string `path:"policies.emm_state.\\.tag" json:"policy_emm_state"`
	PolicyOfficeAddIn        string `path:"policies.office_addin.\\.tag" json:"policy_office_add_in"`
}

type Feature struct {
	Raw                     json.RawMessage
	UploadApiRateLimit      string `path:"upload_api_rate_limit.\\.tag" json:"upload_api_rate_limit"`
	UploadApiRateLimitCount int    `path:"upload_api_rate_limit.limit" json:"upload_api_rate_limit_count"`
	HasTeamSharedDropbox    bool   `path:"has_team_shared_dropbox.has_team_shared_dropbox" json:"has_team_shared_dropbox"`
	HasTeamFileEvents       bool   `path:"has_team_file_events.enabled" json:"has_team_file_events"`
	HasTeamSelectiveSync    bool   `path:"has_team_selective_sync.has_team_selective_sync" json:"has_team_selective_sync"`
	HasDistinctMemberHomes  bool   `path:"has_distinct_member_homes.has_distinct_member_homes" json:"has_distinct_member_homes"`
}

func (z Feature) FileSystemType() dbx_filesystem.TeamFileSystemType {
	return IdentifyFileSystemType(&z)
}

func IdentifyFileSystemType(f *Feature) dbx_filesystem.TeamFileSystemType {
	return dbx_filesystem.IdentifyFileSystemTypeByParam(f.HasDistinctMemberHomes, f.HasTeamSharedDropbox)
}
