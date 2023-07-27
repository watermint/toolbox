package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Namespace struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// attributes
	Name          string `path:"name"`
	NamespaceType string `path:"namespace_type.\\.tag"`
	TeamMemberId  string `path:"team_member_id"`

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
