package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type GroupMember struct {
	// primary keys
	GroupId      string `path:"group_id" gorm:"primaryKey"`
	TeamMemberId string `path:"profile.team_member_id" gorm:"primaryKey"`

	// attributes
	Email       string `path:"profile.email"`
	DisplayName string `path:"profile.name.display_name"`
	AccessType  string `path:"access_type.\\.tag"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewGroupMemberFromJson(groupId string, data es_json.Json) (gm *GroupMember, err error) {
	gm = &GroupMember{}
	if err = data.Model(gm); err != nil {
		return nil, err
	}
	gm.GroupId = groupId
	return gm, nil
}

func (z tsImpl) scanGroupMember(groupId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	members, err := sv_group_member.NewByGroupId(z.client, groupId).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		m, err := NewGroupMemberFromJson(groupId, es_json.MustParse(member.Raw))
		if err != nil {
			return err
		}
		z.db.Save(m)
		if z.db.Error != nil {
			return z.db.Error
		}
	}
	return nil
}
