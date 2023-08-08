package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_lang"
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
	qne := stage.Get(teamScanQueueNamespaceEntry)
	qnd := stage.Get(teamScanQueueNamespaceDetail)
	qnm := stage.Get(teamScanQueueNamespaceMember)

	var lastErr error
	opErr := sv_namespace.New(z.client).ListEach(func(namespace *mo_namespace.Namespace) bool {
		ns, err := NewNamespaceFromJson(es_json.MustParse(namespace.Raw))
		if err != nil {
			lastErr = err
			return false
		}
		z.db.Save(ns)
		if z.db.Error != nil {
			lastErr = z.db.Error
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

		return true
	})

	return es_lang.NewMultiErrorOrNull(opErr, lastErr)
}
