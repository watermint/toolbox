package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"gorm.io/gorm"
)

type Namespace struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// attributes
	Name          string `path:"name"`
	NamespaceType string `path:"namespace_type.\\.tag"`
	TeamMemberId  string `path:"team_member_id" gorm:"index"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

type NamespaceError struct {
	Dummy             string `gorm:"primaryKey"`
	ScanMemberFolders bool   `json:"scanMemberFolders" path:"scan_member_folders"`
	Error             string `path:"error_summary"`
}

func (z NamespaceError) ToParam() interface{} {
	return &NamespaceParam{
		ScanMemberFolders: z.ScanMemberFolders,
		IsRetry:           true,
	}
}

type NamespaceParam struct {
	ScanMemberFolders bool `json:"scanMemberFolders" path:"scan_member_folders"`
	IsRetry           bool `path:"is_retry" json:"is_retry" json:"is_retry,omitempty"`
}

func NewNamespaceFromJson(data es_json.Json) (ns *Namespace, err error) {
	ns = &Namespace{}
	if err = data.Model(ns); err != nil {
		return nil, err
	}
	return ns, nil
}

func (z tsImpl) scanNamespaces(param *NamespaceParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log()
	qne := stage.Get(teamScanQueueNamespaceEntry)
	qnd := stage.Get(teamScanQueueNamespaceDetail)
	qnm := stage.Get(teamScanQueueNamespaceMember)

	var lastErr error
	opErr := sv_namespace.New(z.client).ListEach(func(namespace *mo_namespace.Namespace) bool {
		ll := l.With(esl.String("namespaceId", namespace.NamespaceId))
		ns, err := NewNamespaceFromJson(es_json.MustParse(namespace.Raw))
		if err != nil {
			lastErr = err
			ll.Debug("unable to parse namespace", esl.Error(err))
			return false
		}
		if !param.ScanMemberFolders {
			switch ns.NamespaceType {
			case "team_member_folder", "app_folder", "team_member_root", "shared_folder":
				ll.Debug("skip member folder")
				return true
			}
		}
		z.adb.Save(ns)
		if z.adb.Error != nil {
			lastErr = z.adb.Error
			ll.Debug("unable to save namespace", esl.Error(z.adb.Error))
			return false
		}

		qne.Enqueue(&NamespaceEntryParam{
			NamespaceId: ns.NamespaceId,
			FolderId:    "",
		})
		if ns.NamespaceType != "team_member_folder" && ns.NamespaceType != "app_folder" && ns.NamespaceType != "team_member_root" {
			qnd.Enqueue(&NamespaceDetailParam{
				NamespaceId: ns.NamespaceId,
			})
			qnm.Enqueue(&NamespaceMemberParam{
				NamespaceId: ns.NamespaceId,
			})
		}
		if ns.NamespaceType == "team_member_folder" || ns.NamespaceType == "app_folder" {
			meta, err := sv_file.NewFiles(z.client.AsAdminId(admin.TeamMemberId)).Resolve(mo_path.NewDropboxPath("ns:" + namespace.NamespaceId))
			if err != nil {
				lastErr = err
				ll.Debug("unable to resolve namespace", esl.Error(err))
				return false
			}
			ce := meta.Concrete()
			z.adb.Save(&NamespaceEntry{
				NamespaceId:              namespace.NamespaceId,
				FileId:                   ce.Id,
				ParentFolderId:           "",
				EntryType:                "folder",
				Name:                     ce.Name,
				Size:                     0,
				Rev:                      "",
				IsDownloadable:           false,
				HasExplicitSharedMembers: false,
				ClientModified:           "",
				ServerModified:           "",
				ContentHash:              "",
				PathLower:                ce.PathLower,
				PathDisplay:              ce.PathDisplay,
				EntryNamespaceId:         namespace.NamespaceId,
				ParentNamespaceId:        ce.ParentSharedFolderId,
				Updated:                  0,
				Raw:                      nil,
			})
		}

		return true
	})

	err = es_lang.NewMultiErrorOrNull(opErr, lastErr)
	if err != nil {
		l.Debug("Operation error", esl.Error(err))
		z.adb.Save(&NamespaceError{
			Dummy: "dummy",
			Error: err.Error(),
		})
		return opErr
	}
	if param.IsRetry {
		z.adb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&NamespaceError{})
	}
	return nil
}
