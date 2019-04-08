package mo_profile

import "encoding/json"

type Profile struct {
	Raw                   json.RawMessage
	TeamMemberId          string `path:"team_member_id"`
	Email                 string `path:"email"`
	EmailVerified         bool   `path:"email_verified"`
	Status                string `path:"status.\\.tag"`
	GivenName             string `path:"name.given_name"`
	Surname               string `path:"name.surname"`
	FamiliarName          string `path:"name.familiar_name"`
	DisplayName           string `path:"name.display_name"`
	AbbreviatedName       string `path:"name.abbreviated_name"`
	MemberFolderId        string `path:"member_folder_id"`
	ExternalId            string `path:"external_id"`
	AccountId             string `path:"account_id"`
	JoinedOn              string `path:"joined_on"`
	IsDirectoryRestricted bool   `path:"is_directory_restricted"`
}
