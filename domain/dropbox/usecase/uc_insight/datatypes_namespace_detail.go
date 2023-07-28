package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type NamespaceDetail struct {
	// primary keys
	NamespaceId string `path:"shared_folder_id" gorm:"primaryKey"`

	// attributes
	Name               string `path:"name"`
	PolicyAclUpdate    string `path:"policy.acl_update_policy.\\.tag"`
	PolicyMemberPolicy string `path:"policy.member_policy.\\.tag"`
	PolicyResolvedAcl  string `path:"policy.resolved_member_policy.\\.tag"`
	PolicySharedLink   string `path:"policy.shared_link_policy.\\.tag"`
	AccessInheritance  string `path:"access_inheritance.\\.tag"`
	IsInsideTeamFolder bool   `path:"is_inside_team_folder"`
	IsTeamFolder       bool   `path:"is_team_folder"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewNamespaceDetail(data es_json.Json) (ns *NamespaceDetail, err error) {
	ns = &NamespaceDetail{}
	if err = data.Model(ns); err != nil {
		return nil, err
	}
	return ns, nil
}
