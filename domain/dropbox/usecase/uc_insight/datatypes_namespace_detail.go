package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
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

func (z tsImpl) scanNamespaceDetail(namespaceId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	ns, err := sv_sharedfolder.New(z.client.AsAdminId(admin.TeamMemberId)).Resolve(namespaceId)
	if err != nil {
		return err
	}
	n, err := NewNamespaceDetail(es_json.MustParse(ns.Raw))
	if err != nil {
		return err
	}
	z.db.Save(n)
	if z.db.Error != nil {
		return z.db.Error
	}
	return nil
}
