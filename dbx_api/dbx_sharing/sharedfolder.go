package dbx_sharing

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Membership struct {
	AccessType  string          `json:"access_type"`
	Permissions json.RawMessage `json:"permissions,omitempty"`
	IsInherited bool            `json:"is_inherited"`
}

type User struct {
	UserAccountId string `json:"user_account_id,omitempty"`
	Email         string `json:"email"`
	DisplayName   string `json:"display_name,omitempty"`
	SameTeam      bool   `json:"same_team"`
	TeamMemberId  string `json:"team_member_id,omitempty"`
}

type MembershipUser struct {
	Membership *Membership
	User       *User
}

type Group struct {
	GroupName           string `json:"group_name,omitempty"`
	GroupId             string `json:"group_id"`
	GroupManagementType string `json:"group_management_type,omitempty"`
	GroupType           string `json:"group_type,omitempty"`
	IsMember            bool   `json:"is_member"`
	IsOwner             bool   `json:"is_owner"`
	SameTeam            bool   `json:"same_team"`
	MemberCount         int64  `json:"member_count,omitempty"`
}

type MembershipGroup struct {
	Membership *Membership
	Group      *Group
}

type Invitee struct {
	Email string `json:"email"`
}

type MembershipInvitee struct {
	Membership *Membership
	Invitee    *Invitee
	User       *User
}

func ParseMembership(r gjson.Result) (m *Membership) {
	m = &Membership{}

	at := r.Get("access_type." + dbx_api.ResJsonDotTag)
	if !at.Exists() {
		return nil
	}
	m.AccessType = at.String()
	m.IsInherited = r.Get("is_inherited").Bool()
	m.Permissions = json.RawMessage(r.Get("permissions").Raw)

	return
}

func ParseMembershipUser(r gjson.Result) (u *MembershipUser) {
	user := &User{}
	resUser := r.Get("user")
	if !resUser.Exists() {
		return nil
	}
	err := json.Unmarshal([]byte(resUser.Raw), user)
	if err != nil {
		seelog.Warnf("Parse error[%s] body[%s]", err, resUser.Raw)
		return nil
	}
	m := ParseMembership(r)
	if m == nil {
		return nil
	}
	u = &MembershipUser{
		Membership: m,
		User:       user,
	}
	return
}

func ParseMembershipGroup(r gjson.Result) (g *MembershipGroup) {
	resGroup := r.Get("group")
	if !resGroup.Exists() {
		return nil
	}

	group := &Group{
		GroupName:           resGroup.Get("group_name").String(),
		GroupId:             resGroup.Get("group_id").String(),
		GroupManagementType: resGroup.Get("group_management_type." + dbx_api.ResJsonDotTag).String(),
		GroupType:           resGroup.Get("group_type." + dbx_api.ResJsonDotTag).String(),
		IsMember:            resGroup.Get("is_member").Bool(),
		IsOwner:             resGroup.Get("is_owner").Bool(),
		SameTeam:            resGroup.Get("same_team").Bool(),
		MemberCount:         resGroup.Get("member_count").Int(),
	}
	m := ParseMembership(r)
	if m == nil {
		return nil
	}
	g = &MembershipGroup{
		Membership: m,
		Group:      group,
	}
	return g
}

type SharedFolderMembers struct {
	AsMemberId string
	AsAdminId  string
	OnError    func(annotation dbx_api.ErrorAnnotation) bool
	OnUser     func(user *MembershipUser) bool
	OnGroup    func(group *MembershipGroup) bool
	OnInvitee  func(invitee *MembershipInvitee) bool
}

func (s *SharedFolderMembers) List(c *dbx_api.Context, sharedFolderId string) bool {
	list := dbx_rpc.RpcList{
		AsAdminId:            s.AsAdminId,
		AsMemberId:           s.AsMemberId,
		EndpointList:         "sharing/list_folder_members",
		EndpointListContinue: "sharing/list_folder_members/continue",
		UseHasMore:           false,
		OnError:              s.OnError,
		OnResponse: func(body string) bool {
			users := gjson.Get(body, "users")
			if s.OnUser != nil && users.Exists() && users.IsArray() {
				for _, u := range users.Array() {
					user := ParseMembershipUser(u)
					if user == nil {
						continue
					}
					if !s.OnUser(user) {
						return false
					}
				}
			}
			groups := gjson.Get(body, "groups")
			if s.OnGroup != nil && groups.Exists() && groups.IsArray() {
				for _, g := range groups.Array() {
					group := ParseMembershipGroup(g)
					if group == nil {
						continue
					}
					if !s.OnGroup(group) {
						return false
					}
				}
			}
			return true
		},
	}
	type ListParam struct {
		SharedFolderId string `json:"shared_folder_id"`
	}
	lp := &ListParam{
		SharedFolderId: sharedFolderId,
	}

	return list.List(c, lp)
}
