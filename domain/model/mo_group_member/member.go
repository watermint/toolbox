package mo_group_member

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"go.uber.org/zap"
)

type Member struct {
	Raw             json.RawMessage
	TeamMemberId    string `path:"profile.team_member_id"`
	Email           string `path:"profile.email"`
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
	AccessType      string `path:"access_type.\\.tag"`
}

func (z *Member) Profile() mo_profile.Profile {
	p := mo_profile.Profile{}
	if err := api_parser.ParseModelRaw(&p, z.Raw); err != nil {
		return mo_profile.Profile{}
	}
	return p
}

func NewGroupMember(group *mo_group.Group, member *Member) (gm *GroupMember) {
	raw := struct {
		Group  json.RawMessage `json:"group,string"`
		Member json.RawMessage `json:"member,string"`
	}{
		Group:  group.Raw,
		Member: member.Raw,
	}
	r, err := json.Marshal(raw)
	if err != nil {
		app.Root().Log().Warn("unable to marshal raw JSON", zap.Error(err))
		r = json.RawMessage("{}")
	}

	gm = &GroupMember{
		Raw:                 r,
		GroupId:             group.GroupId,
		GroupName:           group.GroupName,
		GroupManagementType: group.GroupManagementType,
		AccessType:          member.AccessType,
		AccountId:           member.AccountId,
		TeamMemberId:        member.TeamMemberId,
		Email:               member.Email,
		Status:              member.Status,
		Surname:             member.Surname,
		GivenName:           member.GivenName,
	}
	return gm
}

// Group and member information
type GroupMember struct {
	Raw                 json.RawMessage `json:"-"`
	GroupId             string          `path:"group.group_id" json:"group_id"`
	GroupName           string          `path:"group.group_name" json:"group_name"`
	GroupManagementType string          `path:"group.group_management_type.\\.tag" json:"group_management_type"`
	AccessType          string          `path:"member.access_type.\\.tag" json:"access_type"`
	AccountId           string          `path:"member.profile.account_id" json:"account_id"`
	TeamMemberId        string          `path:"member.profile.team_member_id" json:"team_member_id"`
	Email               string          `path:"member.profile.email" json:"email"`
	Status              string          `path:"member.profile.status.\\.tag" json:"status"`
	Surname             string          `path:"member.profile.name.surname" json:"surname"`
	GivenName           string          `path:"member.profile.name.given_name" json:"given_name"`
}

func (z *GroupMember) Group() (group *mo_group.Group) {
	group = &mo_group.Group{}
	j := gjson.ParseBytes(z.Raw)
	jg := j.Get("group")
	if !jg.Exists() {
		app.Root().Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)))
		// return empty
		return group
	}
	if err := api_parser.ParseModel(group, jg); err != nil {
		app.Root().Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)), zap.Error(err))
		// return empty
		return group
	}
	return group
}

func (z *GroupMember) Member() (member *Member) {
	member = &Member{}
	j := gjson.ParseBytes(z.Raw)
	jg := j.Get("member")
	if !jg.Exists() {
		app.Root().Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)))
		// return empty
		return
	}
	if err := api_parser.ParseModel(member, jg); err != nil {
		app.Root().Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)), zap.Error(err))
		// return empty
		return
	}
	return member
}
