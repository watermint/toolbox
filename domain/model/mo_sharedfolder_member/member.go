package mo_sharedfolder_member

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

const (
	AccessTypeOwner           = "owner"
	AccessTypeEditor          = "editor"
	AccessTypeViewer          = "viewer"
	AccessTypeViewerNoComment = "viewer_no_comment"
	MemberTypeUser            = "user"
	MemberTypeGroup           = "group"
	MemberTypeInvitee         = "invitee"
)

type Member interface {
	AccessType() string
	IsInherited() bool
	MemberType() string
	User() (user *User, e bool)
	Group() (group *Group, e bool)
	Invitee() (invitee *Invitee, e bool)
	EntryRaw() json.RawMessage
}

type Metadata struct {
	Raw              json.RawMessage
	EntryAccessType  string `path:"access_type.\\.tag"`
	EntryIsInherited bool   `path:"is_inherited"`
}

func (z *Metadata) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Metadata) AccessType() string {
	return z.EntryAccessType
}

func (z *Metadata) IsInherited() bool {
	return z.EntryIsInherited
}

func (z *Metadata) MemberType() string {
	j := gjson.ParseBytes(z.Raw)
	if j.Get("user").Exists() {
		return MemberTypeUser
	}
	if j.Get("group").Exists() {
		return MemberTypeGroup
	}
	if j.Get("invitee").Exists() {
		return MemberTypeInvitee
	}
	app_root.Log().Warn("Unknown member type", zap.String("entry", string(z.Raw)))
	return ""
}

func (z *Metadata) User() (user *User, e bool) {
	if z.MemberType() == "user" {
		user = &User{}
		if err := api_parser.ParseModelRaw(user, z.Raw); err != nil {
			return nil, false
		}
		return user, true
	}
	return nil, false
}

func (z *Metadata) Group() (group *Group, e bool) {
	if z.MemberType() == "group" {
		group = &Group{}
		if err := api_parser.ParseModelRaw(group, z.Raw); err != nil {
			return nil, false
		}
		return group, true
	}
	return nil, false
}

func (z *Metadata) Invitee() (invitee *Invitee, e bool) {
	if z.MemberType() == "invitee" {
		invitee = &Invitee{}
		if err := api_parser.ParseModelRaw(invitee, z.Raw); err != nil {
			return nil, false
		}
		return invitee, true
	}
	return nil, false
}

type User struct {
	Raw              json.RawMessage
	EntryAccessType  string `path:"access_type.\\.tag"`
	EntryIsInherited bool   `path:"is_inherited"`
	AccountId        string `path:"user.account_id"`
	Email            string `path:"user.email"`
	DisplayName      string `path:"user.display_name"`
	SameTeam         bool   `path:"user.same_team"`
	TeamMemberId     string `path:"user.team_member_id"`
}

func (z *User) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *User) AccessType() string {
	return z.EntryAccessType
}

func (z *User) IsInherited() bool {
	return z.EntryIsInherited
}

func (z *User) MemberType() string {
	return MemberTypeUser
}

func (z *User) User() (user *User, e bool) {
	return z, true
}

func (z *User) Group() (group *Group, e bool) {
	return nil, false
}

func (z *User) Invitee() (invitee *Invitee, e bool) {
	return nil, false
}

type Group struct {
	Raw                 json.RawMessage
	EntryAccessType     string `path:"access_type.\\.tag"`
	EntryIsInherited    bool   `path:"is_inherited"`
	GroupName           string `path:"group.group_name"`
	GroupId             string `path:"group.group_id"`
	GroupManagementType string `path:"group.group_management_type.\\.tag"`
	GroupType           string `path:"group.group_type.\\.tag"`
	IsMember            bool   `path:"group.is_member"`
	IsOwner             bool   `path:"group.is_owner"`
	SameTeam            bool   `path:"group.same_team"`
	GroupExternalId     string `path:"group.group_external_id"`
	MemberCount         int    `path:"group.member_count"`
}

func (z *Group) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Group) AccessType() string {
	return z.EntryAccessType
}

func (z *Group) IsInherited() bool {
	return z.EntryIsInherited
}

func (z *Group) MemberType() string {
	return MemberTypeGroup
}

func (z *Group) User() (user *User, e bool) {
	return nil, false
}

func (z *Group) Group() (group *Group, e bool) {
	return z, true
}

func (z *Group) Invitee() (invitee *Invitee, e bool) {
	return nil, false
}

type Invitee struct {
	Raw              json.RawMessage
	EntryAccessType  string `path:"access_type.\\.tag"`
	EntryIsInherited bool   `path:"is_inherited"`
	InviteeEmail     string `path:"invitee.email"`
}

func (z *Invitee) EntryRaw() json.RawMessage {
	return z.Raw
}

func (z *Invitee) AccessType() string {
	return z.EntryAccessType
}

func (z *Invitee) IsInherited() bool {
	return z.EntryIsInherited
}

func (z *Invitee) MemberType() string {
	return MemberTypeInvitee
}

func (z *Invitee) User() (user *User, e bool) {
	return nil, false
}

func (z *Invitee) Group() (group *Group, e bool) {
	return nil, false
}

func (z *Invitee) Invitee() (invitee *Invitee, e bool) {
	return z, true
}

type SharedFolderMember struct {
	Raw                  json.RawMessage
	SharedFolderId       string `path:"sharedfolder.shared_folder_id"`
	ParentSharedFolderId string `path:"sharedfolder.parent_shared_folder_id"`
	Name                 string `path:"sharedfolder.name"`
	AccessType           string `path:"sharedfolder.access_type.\\.tag"`
	PathLower            string `path:"sharedfolder.path_lower"`
	IsInsideTeamFolder   bool   `path:"sharedfolder.is_inside_team_folder"`
	IsTeamFolder         bool   `path:"sharedfolder.is_team_folder"`
	EntryAccessType      string `path:"member.access_type.\\.tag"`
	EntryIsInherited     bool   `path:"member.is_inherited"`
	AccountId            string `path:"member.user.account_id"`
	Email                string `path:"member.user.email"`
	DisplayName          string `path:"member.user.display_name"`
	GroupName            string `path:"member.group.group_name"`
	GroupId              string `path:"member.group.group_id"`
	InviteeEmail         string `path:"member.invitee.email"`
}

func NewSharedFolderMember(sf *mo_sharedfolder.SharedFolder, m Member) *SharedFolderMember {
	raws := make(map[string]json.RawMessage)
	raws["sharedfolder"] = sf.Raw
	raws["member"] = m.EntryRaw()
	raw := api_parser.CombineRaw(raws)

	sfm := &SharedFolderMember{}
	if err := api_parser.ParseModelRaw(sfm, raw); err != nil {
		app_root.Log().Error("unable to parse", zap.Error(err))
	}
	return sfm
}
