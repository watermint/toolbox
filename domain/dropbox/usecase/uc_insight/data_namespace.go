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

func NewNamespaceFromJson(data es_json.Json) (ns *Namespace, err error) {
	ns = &Namespace{}
	if err = data.Model(ns); err != nil {
		return nil, err
	}
	return ns, nil
}

func (z tsImpl) scanNamespaces(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
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
		if ns.NamespaceType != "team_member_folder" && ns.NamespaceType != "app_folder" {
			qnd.Enqueue(ns.NamespaceId)
			qnm.Enqueue(ns.NamespaceId)
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

	return es_lang.NewMultiErrorOrNull(opErr, lastErr)
}
