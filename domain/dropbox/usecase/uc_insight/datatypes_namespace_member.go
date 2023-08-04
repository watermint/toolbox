package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
)

type NamespaceMember struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// attributes
	AccessType  string `path:"access_type.\\.tag"`
	IsInherited bool   `path:"is_inherited"`

	// MemberType user, group, invitee
	MemberType string
	// SameTeam yes, no, unknown
	SameTeam string

	// group
	GroupId          string `path:"group.group_id"`
	GroupName        string `path:"group.group_name"`
	GroupType        string `path:"group.group_type.\\.tag"`
	GroupMemberCount uint64 `path:"group.member_count"`

	// invitee
	InviteeEmail string `path:"invitee.email"`

	// user
	UserTeamMemberId string `path:"user.team_member_id"`
	UserEmail        string `path:"user.email"`
	UserDisplayName  string `path:"user.display_name"`
	UserAccountId    string `path:"user.account_id"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func NewNamespaceMember(namespaceId string, data mo_sharedfolder_member.Member) (ns *NamespaceMember) {
	ns = &NamespaceMember{}
	ns.NamespaceId = namespaceId
	ns.AccessType = data.AccessType()
	ns.IsInherited = data.IsInherited()
	ns.MemberType = data.MemberType()
	switch data.SameTeam() {
	case "true":
		ns.SameTeam = "yes"
	case "false":
		ns.SameTeam = "no"
	default:
		ns.SameTeam = "unknown"
	}

	if user, ok := data.User(); ok {
		ns.UserTeamMemberId = user.TeamMemberId
		ns.UserEmail = user.Email
		ns.UserDisplayName = user.DisplayName
		ns.UserAccountId = user.AccountId
	}
	if group, ok := data.Group(); ok {
		ns.GroupId = group.GroupId
		ns.GroupName = group.GroupName
		ns.GroupType = group.GroupType
		ns.GroupMemberCount = uint64(group.MemberCount)
	}
	if invitee, ok := data.Invitee(); ok {
		ns.InviteeEmail = invitee.InviteeEmail
	}

	return ns
}
