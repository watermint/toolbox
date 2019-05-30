package mo_member

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/model/mo_profile"
)

type Member struct {
	Raw             json.RawMessage
	TeamMemberId    string `path:"profile.team_member_id"`
	Email           string `path:"profile.email"`
	EmailVerified   bool   `path:"profile.email_verified"`
	Status          string `path:"profile.status.\\.tag"`
	GivenName       string `path:"profile.name.given_name"`
	Surname         string `path:"profile.name.surname"`
	FamiliarName    string `path:"profile.name.familiar_name"`
	DisplayName     string `path:"profile.name.display_name"`
	AbbreviatedName string `path:"profile.name.abbreviated_name"`
	MemberFolderId  string `path:"profile.member_folder_id"`
	ExternalId      string `path:"profile.external_id"`
	AccountId       string `path:"profile.account_id"`
	PersistentId    string `path:"profile.persistent_id"`
	JoinedOn        string `path:"profile.joined_on"`
	Role            string `path:"role.\\.tag"`
	Tag             string `path:"\\.tag"`
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
