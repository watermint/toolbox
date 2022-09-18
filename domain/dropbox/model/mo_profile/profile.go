package mo_profile

import "encoding/json"

const (
	StatusActive    = "active"
	StatusInvited   = "invited"
	StatusSuspended = "suspended"
	StatusRemoved   = "removed"
)

type Profile struct {
	Raw                   json.RawMessage
	TeamMemberId          string `path:"team_member_id" json:"team_member_id"`
	Email                 string `path:"email" json:"email"`
	EmailVerified         bool   `path:"email_verified" json:"email_verified"`
	Status                string `path:"status.\\.tag" json:"status"`
	GivenName             string `path:"name.given_name" json:"given_name"`
	Surname               string `path:"name.surname" json:"surname"`
	FamiliarName          string `path:"name.familiar_name" json:"familiar_name"`
	DisplayName           string `path:"name.display_name" json:"display_name"`
	AbbreviatedName       string `path:"name.abbreviated_name" json:"abbreviated_name"`
	MemberFolderId        string `path:"member_folder_id" json:"member_folder_id"`
	RootNamespaceId       string `path:"root_info.root_namespace_id" json:"root_namespace_id"`
	HomeNamespaceId       string `path:"root_info.home_namespace_id" json:"home_namespace_id"`
	HomePath              string `path:"root_info.home_path" json:"home_path"`
	ExternalId            string `path:"external_id" json:"external_id"`
	AccountId             string `path:"account_id" json:"account_id"`
	AccountType           string `path:"account_type.\\.tag" json:"account_type"`
	JoinedOn              string `path:"joined_on" json:"joined_on"`
	IsDirectoryRestricted bool   `path:"is_directory_restricted" json:"is_directory_restricted"`
}

func (z Profile) IsActive() bool {
	return z.Status == StatusActive
}
func (z Profile) IsInvited() bool {
	return z.Status == StatusInvited
}
func (z Profile) IsSuspended() bool {
	return z.Status == StatusSuspended
}
func (z Profile) IsRemoved() bool {
	return z.Status == StatusRemoved
}

// True when the status is active or suspended.
func (z Profile) IsAccepted() bool {
	return z.Status == StatusActive || z.Status == StatusSuspended
}
