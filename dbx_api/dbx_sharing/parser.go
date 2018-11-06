package dbx_sharing

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"go.uber.org/zap"
)

func ParseMembership(r gjson.Result, log *zap.Logger) (m *Membership) {
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

func ParseMembershipUser(r gjson.Result, log *zap.Logger) (u *MembershipUser) {
	user := &User{}
	resUser := r.Get("user")
	if !resUser.Exists() {
		return nil
	}
	err := json.Unmarshal([]byte(resUser.Raw), user)
	if err != nil {
		log.Warn(
			"parse error",
			zap.Error(err),
			zap.String("body", resUser.Str),
		)
		return nil
	}
	m := ParseMembership(r, log)
	if m == nil {
		return nil
	}
	u = &MembershipUser{
		Membership: m,
		User:       user,
	}
	return
}

func ParseMembershipGroup(r gjson.Result, log *zap.Logger) (g *MembershipGroup) {
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
	m := ParseMembership(r, log)
	if m == nil {
		return nil
	}
	g = &MembershipGroup{
		Membership: m,
		Group:      group,
	}
	return g
}

func ParseMembershipInvitee(r gjson.Result, log *zap.Logger) (g *MembershipInvitee) {
	m := ParseMembership(r, log)
	if m == nil {
		return nil
	}

	email := r.Get("invitee.email")
	if !email.Exists() {
		return nil
	}

	g = &MembershipInvitee{
		Membership: m,
		Invitee: &Invitee{
			Email: email.String(),
		},
	}
	return
}
