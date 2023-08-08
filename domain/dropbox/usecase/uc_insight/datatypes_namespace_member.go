package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
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

	Updated uint64 `gorm:"autoUpdateTime"`
}

func NewNamespaceMember(namespaceId string, data mo_sharedfolder_member.Member) (ns *NamespaceMember) {
	ns = &NamespaceMember{}
	ns.NamespaceId = namespaceId
	ns.AccessType = data.AccessType()
	ns.IsInherited = data.IsInherited()
	ns.MemberType = data.MemberType()
	ns.SameTeam = ConvertSameTeam(data.SameTeam())

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

func (z tsImpl) scanNamespaceMember(namespaceId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	members, err := sv_sharedfolder_member.NewBySharedFolderId(z.client.AsAdminId(admin.TeamMemberId), namespaceId).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		m := NewNamespaceMember(namespaceId, member)
		z.saveIfExternalGroup(member)
		z.db.Save(m)
		if z.db.Error != nil {
			return z.db.Error
		}
	}
	return nil
}
