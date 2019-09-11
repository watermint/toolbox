package mo_member

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

type Member struct {
	Raw             json.RawMessage
	TeamMemberId    string `path:"profile.team_member_id" json:"team_member_id"`
	Email           string `path:"profile.email" json:"email"`
	EmailVerified   bool   `path:"profile.email_verified" json:"email_verified"`
	Status          string `path:"profile.status.\\.tag" json:"status"`
	GivenName       string `path:"profile.name.given_name" json:"given_name"`
	Surname         string `path:"profile.name.surname" json:"surname"`
	FamiliarName    string `path:"profile.name.familiar_name" json:"familiar_name"`
	DisplayName     string `path:"profile.name.display_name" json:"display_name"`
	AbbreviatedName string `path:"profile.name.abbreviated_name" json:"abbreviated_name"`
	MemberFolderId  string `path:"profile.member_folder_id" json:"member_folder_id"`
	ExternalId      string `path:"profile.external_id" json:"external_id"`
	AccountId       string `path:"profile.account_id" json:"account_id"`
	PersistentId    string `path:"profile.persistent_id" json:"persistent_id"`
	JoinedOn        string `path:"profile.joined_on" json:"joined_on"`
	Role            string `path:"role.\\.tag" json:"role"`
	Tag             string `path:"\\.tag" json:"tag"`
}

func (z *Member) Profile() *mo_profile.Profile {
	p := &mo_profile.Profile{}
	if err := api_parser.ParseModelPathRaw(p, z.Raw, "profile"); err != nil {
		return &mo_profile.Profile{}
	}
	return p
}

func MapByEmail(list []*Member) (members map[string]*Member) {
	members = make(map[string]*Member)
	for _, m := range list {
		members[m.Email] = m
	}
	return members
}

func MapByTeamMemberId(list []*Member) (members map[string]*Member) {
	members = make(map[string]*Member)
	for _, m := range list {
		members[m.TeamMemberId] = m
	}
	return members
}
