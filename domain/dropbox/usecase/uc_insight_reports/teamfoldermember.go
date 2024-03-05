package uc_insight_reports

type TeamFolderMember struct {
	TeamFolderId   string `json:"team_folder_id"`
	TeamFolderName string `json:"team_folder_name"`
	PathDisplay    string `json:"path_display"`

	AccessType  string `json:"access_type"`
	IsInherited bool   `json:"is_inherited"`

	// MemberType user, group, invitee
	MemberType string `json:"member_type"`
	// SameTeam yes, no, unknown
	SameTeam string `json:"same_team"`

	// group
	GroupId          string `json:"group_id"`
	GroupName        string `json:"group_name"`
	GroupType        string `json:"group_type"`
	GroupMemberCount uint64 `json:"group_member_count"`

	// invitee
	InviteeEmail string `json:"invitee_email"`

	// user
	UserTeamMemberId string `json:"user_team_member_id"`
	UserEmail        string `json:"user_email"`
	UserDisplayName  string `json:"user_display_name"`
	UserAccountId    string `json:"user_account_id"`
}
