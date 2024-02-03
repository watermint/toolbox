package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type FileMember struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FileId      string `path:"file_id" gorm:"primaryKey"`

	AccessMember

	Updated uint64 `gorm:"autoUpdateTime"`
}

type FileMemberError struct {
	FileMemberParam
	Error string `path:"error_summary"`
}

func (z FileMemberError) ToParam() interface{} {
	return &FileMemberParam{
		NamespaceId: z.NamespaceId,
		FileId:      z.FileId,
		IsRetry:     true,
	}
}

func NewFileMember(namespaceId string, fileId string, data mo_sharedfolder_member.Member) (ns *FileMember) {
	ns = &FileMember{}
	ns.NamespaceId = namespaceId
	ns.FileId = fileId
	ns.AccessMember = ParseAccessMember(data)
	return ns
}

func (z tsImpl) scanFileMember(entry *FileMemberParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log().With(esl.String("namespaceId", entry.NamespaceId), esl.String("fileId", entry.FileId))
	client := z.client.AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(entry.NamespaceId))

	members, err := sv_file_member.New(client).List(entry.FileId, false)
	if err != nil {
		l.Debug("Unable to retrieve members", esl.Error(err))
		dbxErr := dbx_error.NewErrors(err)
		if dbxErr != nil && dbxErr.Path().IsNotFound() {
			l.Debug("File not found, maybe removed during the scan", esl.String("fileId", entry.FileId))
			return nil
		}

		z.adb.Save(&FileMemberError{
			FileMemberParam: *entry,
			Error:           err.Error(),
		})
		return err
	}

	for _, member := range members {
		m := NewFileMember(entry.NamespaceId, entry.FileId, member)
		z.saveIfExternalGroup(member)
		z.adb.Save(m)
	}

	if entry.IsRetry {
		z.adb.Delete(&FileMemberError{}, "namespace_id = ? and file_id = ?", entry.NamespaceId, entry.FileId)
	}

	return nil
}
