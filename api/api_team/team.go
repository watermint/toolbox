package api_team

type ArgMemberAdd struct {
	MemberEmail           string `json:"member_email"`
	MemberGivenName       string `json:"member_given_name,omitempty"`
	MemberSurname         string `json:"member_surname,omitempty"`
	MemberExternalId      string `json:"member_external_id,omitempty"`
	MemberPersistentId    string `json:"member_persistent_id,omitempty"`
	SendWelcomeEmail      bool   `json:"send_welcome_email"`
	Role                  string `json:"role,omitempty"`
	IsDirectoryRestricted bool   `json:"is_directory_restricted,omitempty"`
}

type ArgMembersAdd struct {
	NewMembers []ArgMemberAdd `json:"new_members"`
	ForceAsync bool           `json:"force_async"`
}

const (
	ADMIN_TIER_TEAM_ADMIN            = "team_admin"
	ADMIN_TIER_USER_MANAGEMENT_ADMIN = "user_management_admin"
	ADMIN_TIER_SUPPORT_ADMIN         = "support_admin"
	ADMIN_TIER_MEMBER_ONLY           = "member_only"
)
