package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
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

type GroupMemberParam struct {
	GroupId string `path:"group_id" json:"group_id"`
	IsRetry bool   `path:"is_retry" json:"is_retry"`
}

type GroupMemberError struct {
	GroupId string `path:"group_id" gorm:"primaryKey"`
	ApiError
}

func (z GroupMemberError) ToParam() interface{} {
	return &GroupMemberParam{
		GroupId: z.GroupId,
		IsRetry: true,
	}
}

func NewGroupMemberFromJson(groupId string, data es_json.Json) (gm *GroupMember, err error) {
	gm = &GroupMember{}
	if err = data.Model(gm); err != nil {
		return nil, err
	}
	gm.GroupId = groupId
	return gm, nil
}

func (z tsImpl) scanGroupMember(param *GroupMemberParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log().With(esl.String("groupId", param.GroupId))
	members, err := sv_group_member.NewByGroupId(z.client, param.GroupId).List()
	if err != nil {
		l.Debug("Unable to retrieve members", esl.Error(err))
		z.adb.Save(&GroupMemberError{
			GroupId:  param.GroupId,
			ApiError: ApiErrorFromError(err),
		})
		return err
	}
	for _, member := range members {
		m, err := NewGroupMemberFromJson(param.GroupId, es_json.MustParse(member.Raw))
		if err != nil {
			return err
		}
		z.adb.Save(m)
		if z.adb.Error != nil {
			return z.adb.Error
		}
	}

	if param.IsRetry {
		z.adb.Delete(&GroupMemberError{}, "group_id = ?", param.GroupId)
	}

	return nil
}
