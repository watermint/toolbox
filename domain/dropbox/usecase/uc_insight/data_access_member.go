package uc_insight

import "github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"

type AccessMember struct {
	// attributes
	AccessType  string `path:"access_type.\\.tag"`
	IsInherited bool   `path:"is_inherited"`

	// MemberType user, group, invitee
	MemberType string
	// SameTeam yes, no, unknown
	SameTeam string

	// group
	GroupId          string `path:"group.group_id" gorm:"index"`
	GroupName        string `path:"group.group_name"`
	GroupType        string `path:"group.group_management_type.\\.tag"`
	GroupMemberCount uint64 `path:"group.member_count"`

	// invitee
	InviteeEmail string `path:"invitee.email"`

	// user
	UserTeamMemberId string `path:"user.team_member_id" gorm:"index"`
	UserEmail        string `path:"user.email"`
	UserDisplayName  string `path:"user.display_name"`
	UserAccountId    string `path:"user.account_id"`
}

func ParseAccessMember(m mo_sharedfolder_member.Member) (am AccessMember) {
	am.AccessType = m.AccessType()
	am.IsInherited = m.IsInherited()
	am.MemberType = m.MemberType()
	am.SameTeam = ConvertSameTeam(m.SameTeam())

	if user, ok := m.User(); ok {
		am.UserTeamMemberId = user.TeamMemberId
		am.UserEmail = user.Email
		am.UserDisplayName = user.DisplayName
		am.UserAccountId = user.AccountId
	}
	if group, ok := m.Group(); ok {
		am.GroupId = group.GroupId
		am.GroupName = group.GroupName
		am.GroupType = group.GroupType
		am.GroupMemberCount = uint64(group.MemberCount)
	}
	if invitee, ok := m.Invitee(); ok {
		am.InviteeEmail = invitee.InviteeEmail
	}
	return am
}
