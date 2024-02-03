package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type NamespaceEntry struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FileId      string `path:"id" gorm:"primaryKey"`

	// attributes
	ParentFolderId           string `gorm:"index"`
	EntryType                string `path:"\\.tag"`
	Name                     string `path:"name"`
	Size                     uint64 `path:"size"`
	Rev                      string `path:"rev"`
	IsDownloadable           bool   `path:"is_downloadable"`
	HasExplicitSharedMembers bool   `path:"has_explicit_shared_members"`
	ClientModified           string `path:"client_modified"`
	ServerModified           string `path:"server_modified"`
	ContentHash              string `path:"content_hash"`
	PathLower                string `path:"path_lower"`
	PathDisplay              string `path:"path_display"`
	EntryNamespaceId         string `path:"sharing_info.shared_folder_id" gorm:"index"`
	ParentNamespaceId        string `path:"sharing_info.parent_shared_folder_id" gorm:"index"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

type NamespaceEntryParam struct {
	NamespaceId string `json:"namespaceId" path:"namespace_id"`
	FolderId    string `json:"folderId" path:"folder_id"`
	IsRetry     bool   `json:"isRetry" path:"is_retry"`
}

type NamespaceEntryError struct {
	NamespaceId string `path:"shared_folder_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

	Error string `path:"error_summary"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z NamespaceEntryError) ToParam() interface{} {
	return &NamespaceEntryParam{
		NamespaceId: z.NamespaceId,
		FolderId:    z.FolderId,
		IsRetry:     true,
	}
}

func NewNamespaceEntry(namespaceId string, parentFolderId string, data es_json.Json) (ne *NamespaceEntry, err error) {
	ne = &NamespaceEntry{}
	if err = data.Model(ne); err != nil {
		return nil, err
	}
	ne.NamespaceId = namespaceId
	ne.ParentFolderId = parentFolderId
	return ne, nil
}

func (z tsImpl) scanNamespaceEntry(param *NamespaceEntryParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log().With(esl.String("namespaceId", param.NamespaceId), esl.String("folderId", param.FolderId))
	qne := stage.Get(teamScanQueueNamespaceEntry)
	qfm := stage.Get(teamScanQueueFileMember)

	targetPath := mo_path.NewDropboxPath(param.FolderId)
	if param.FolderId == "" {
		f, err := sv_file.NewFiles(z.client.AsAdminId(admin.TeamMemberId)).Resolve(mo_path.NewDropboxPath("ns:" + param.NamespaceId))
		dbxErr := dbx_error.NewErrors(err)
		switch {
		case dbxErr == nil:
			param.FolderId = f.Concrete().Id
			// fall through

		case dbxErr.Path().IsNotFound():
			l.Debug("Namespace folder not found, the folder removed during the process", esl.String("namespaceId", param.NamespaceId))
			return nil

		default:
			l.Debug("Unable to resolve namespace folder", esl.Error(err))
			return err
		}
	}

	client := z.client.AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(param.NamespaceId))
	err = sv_file.NewFiles(client).ListEach(targetPath,
		func(entry mo_file.Entry) {
			ce := entry.Concrete()
			f, err := NewNamespaceEntry(param.NamespaceId, param.FolderId, es_json.MustParse(ce.Raw))
			if err != nil {
				return
			}
			f.ParentFolderId = param.FolderId
			z.adb.Save(f)

			if ce.IsFolder() {
				if ce.SharedFolderId == "" {
					qne.Enqueue(&NamespaceEntryParam{
						NamespaceId: param.NamespaceId,
						FolderId:    ce.Id,
					})
				} else {
					l.Debug("Skip nested", esl.String("folderId", ce.Id), esl.String("sharedFolderId", ce.SharedFolderId))
				}
			}
			if ce.IsFile() && ce.HasExplicitSharedMembers {
				qfm.Enqueue(&FileMemberParam{
					NamespaceId: param.NamespaceId,
					FileId:      ce.Id,
				})
			}
		},
		sv_file.Recursive(false),
		sv_file.IncludeDeleted(true),
		sv_file.IncludeHasExplicitSharedMembers(true),
	)
	dbxErr := dbx_error.NewErrors(err)
	switch {
	case dbxErr == nil:
		if param.IsRetry {
			z.adb.Delete(&NamespaceEntryError{
				NamespaceId: param.NamespaceId,
				FolderId:    param.FolderId,
			})
		}

	case dbxErr.Path().IsNotFound():
		l.Debug("Namespace folder not found, the folder removed during the process", esl.String("namespaceId", param.NamespaceId))
		return nil

	default:
		dbxErr := dbx_error.NewErrors(err)
		l.Debug("List namespace entry", esl.Error(err), esl.String("dbxErrorSummary", dbxErr.Summary()))

		z.adb.Save(&NamespaceEntryError{
			NamespaceId: param.NamespaceId,
			FolderId:    param.FolderId,
			Error:       err.Error(),
		})
	}
	return err
}
