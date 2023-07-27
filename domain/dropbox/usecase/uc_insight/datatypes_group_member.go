package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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
