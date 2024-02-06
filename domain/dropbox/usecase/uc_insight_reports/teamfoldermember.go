package uc_insight_reports

type TeamFolderMember struct {
	TeamFolderId   string `json:"team_folder_id,omitempty"`
	TeamFolderName string `json:"team_folder_name,omitempty"`
	PathDisplay    string `json:"path_display,omitempty"`

	AccessType  string `json:"access_type,omitempty"`
	IsInherited bool   `json:"is_inherited,omitempty"`

	// MemberType user, group, invitee
	MemberType string `json:"member_type,omitempty"`
	// SameTeam yes, no, unknown
	SameTeam string `json:"same_team,omitempty"`

	// group
	GroupId          string `json:"group_id,omitempty"`
	GroupName        string `json:"group_name,omitempty"`
	GroupType        string `json:"group_type,omitempty"`
	GroupMemberCount uint64 `json:"group_member_count,omitempty"`

	// invitee
	InviteeEmail string `json:"invitee_email,omitempty"`

	// user
	UserTeamMemberId string `json:"user_team_member_id,omitempty"`
	UserEmail        string `json:"user_email,omitempty"`
	UserDisplayName  string `json:"user_display_name,omitempty"`
	UserAccountId    string `json:"user_account_id,omitempty"`
}
