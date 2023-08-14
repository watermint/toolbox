package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_member"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type FileMember struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FileId      string `path:"file_id" gorm:"primaryKey"`

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

func NewFileMember(namespaceId string, fileId string, data mo_sharedfolder_member.Member) (ns *FileMember) {
	ns = &FileMember{}
	ns.NamespaceId = namespaceId
	ns.FileId = fileId
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

func (z tsImpl) scanFileMember(entry *FileMemberParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(entry.NamespaceId))

	members, err := sv_file_member.New(client).List(entry.FileId, false)
	if err != nil {
		return err
	}

	for _, member := range members {
		m := NewFileMember(entry.NamespaceId, entry.FileId, member)
		z.saveIfExternalGroup(member)
		z.db.Save(m)
	}

	return nil
}