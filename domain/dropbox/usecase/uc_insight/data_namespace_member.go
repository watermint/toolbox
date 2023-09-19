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
	AccessMember

	Updated uint64 `gorm:"autoUpdateTime"`
}

func NewNamespaceMember(namespaceId string, data mo_sharedfolder_member.Member) (ns *NamespaceMember) {
	ns = &NamespaceMember{}
	ns.NamespaceId = namespaceId
	ns.AccessMember = ParseAccessMember(data)
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
