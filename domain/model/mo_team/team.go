package mo_team

import "encoding/json"

type Info struct {
	Raw                      json.RawMessage
	Name                     string `path:"name"`
	TeamId                   string `path:"team_id"`
	NumLicensedUsers         int    `path:"num_licensed_users"`
	NumProvisionedUsers      int    `path:"num_provisioned_users"`
	PolicySharedFolderMember string `path:"policies.sharing.shared_folder_member_policy.\\.tag"`
	PolicySharedFolderJoin   string `path:"policies.sharing.shared_folder_join_policy.\\.tag"`
	PolicySharedLinkCreate   string `path:"policies.sharing.shared_link_create_policy.\\.tag"`
	PolicyEmmState           string `path:"policies.emm_state.\\.tag"`
	PolicyOfficeAddIn        string `path:"policies.office_addin.\\.tag"`
}

type Feature struct {
	Raw                     json.RawMessage
	UploadApiRateLimit      string `path:"upload_api_rate_limit.\\.tag"`
	UploadApiRateLimitCount int    `path:"upload_api_rate_limit.limit"`
	HasTeamSharedDropbox    bool   `path:"has_team_shared_dropbox.has_team_shared_dropbox"`
	HasTeamFileEvents       bool   `path:"has_team_file_events.enabled"`
	HasTeamSelectiveSync    bool   `path:"has_team_selective_sync.has_team_selective_sync"`
}
