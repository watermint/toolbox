package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Group struct {
	// primary keys
	GroupId string `path:"group_id" gorm:"primaryKey"`

	// attributes
	GroupName           string `path:"group_name"`
	GroupManagementType string `path:"group_management_type.\\.tag"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewGroupFromJson(data es_json.Json) (g *Group, err error) {
	g = &Group{}
	if err = data.Model(g); err != nil {
		return nil, err
	}
	return g, nil
}
