package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type NamespaceMember struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// attributes
	AccessMember

	Updated uint64 `gorm:"autoUpdateTime"`
}

type NamespaceMemberError struct {
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	Error       string `path:"error_summary"`
}

func (z NamespaceMemberError) ToParam() interface{} {
	return &NamespaceMemberParam{
		NamespaceId: z.NamespaceId,
		IsRetry:     true,
	}
}

type NamespaceMemberParam struct {
	NamespaceId string `path:"namespace_id" json:"namespace_id"`
	IsRetry     bool   `path:"is_retry" json:"is_retry"`
}

func NewNamespaceMember(namespaceId string, data mo_sharedfolder_member.Member) (ns *NamespaceMember) {
	ns = &NamespaceMember{}
	ns.NamespaceId = namespaceId
	ns.AccessMember = ParseAccessMember(data)
	return ns
}

func (z tsImpl) scanNamespaceMember(param *NamespaceMemberParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log().With(esl.String("namespaceId", param.NamespaceId))
	members, err := sv_sharedfolder_member.NewBySharedFolderId(z.client.AsAdminId(admin.TeamMemberId), param.NamespaceId).List()
	if err != nil {
		l.Debug("Unable to retrieve members", esl.Error(err))
		z.adb.Save(&NamespaceMemberError{
			NamespaceId: param.NamespaceId,
			Error:       err.Error(),
		})
		return err
	}
	for _, member := range members {
		m := NewNamespaceMember(param.NamespaceId, member)
		z.saveIfExternalGroup(member)
		z.adb.Save(m)
		if z.adb.Error != nil {
			return z.adb.Error
		}
	}

	if param.IsRetry {
		z.adb.Delete(&NamespaceMemberError{}, "namespace_id = ?", param.NamespaceId)
	}

	return nil
}
