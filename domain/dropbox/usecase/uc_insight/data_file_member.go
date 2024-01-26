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

	AccessMember

	Updated uint64 `gorm:"autoUpdateTime"`
}

func NewFileMember(namespaceId string, fileId string, data mo_sharedfolder_member.Member) (ns *FileMember) {
	ns = &FileMember{}
	ns.NamespaceId = namespaceId
	ns.FileId = fileId
	ns.AccessMember = ParseAccessMember(data)
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
		z.adb.Save(m)
	}

	return nil
}
