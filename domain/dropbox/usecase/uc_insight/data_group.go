package uc_insight

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"gorm.io/gorm"
)

type Group struct {
	// primary keys
	GroupId string `path:"group_id" gorm:"primaryKey"`

	// attributes
	GroupName           string `path:"group_name"`
	GroupManagementType string `path:"group_management_type.\\.tag"`
	GroupExternalId     string `path:"group_external_id"`
	MemberCount         uint64 `path:"member_count"`

	SameTeam string `path:"same_team"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

type GroupError struct {
	Dummy string `path:"dummy" gorm:"primaryKey"`
	Error string `path:"error_summary"`
}

func (z GroupError) ToParam() interface{} {
	return &GroupParam{
		IsRetry: true,
	}
}

type GroupParam struct {
	IsRetry bool `path:"is_retry" json:"is_retry"`
}

func NewGroupFromJson(data es_json.Json) (g *Group, err error) {
	g = &Group{}
	if err = data.Model(g); err != nil {
		return nil, err
	}
	return g, nil
}

func NewGroupFromMember(member mo_sharedfolder_member.Member) (g *Group, err error) {
	g = &Group{}
	if mg, ok := member.Group(); ok {
		g.GroupId = mg.GroupId
		g.GroupName = mg.GroupName
		g.GroupManagementType = mg.GroupManagementType
		g.GroupExternalId = mg.GroupExternalId
		g.SameTeam = ConvertSameTeam(member.SameTeam())
		return g, nil
	}

	return nil, errors.New("not a group")
}

func (z tsImpl) scanGroup(param *GroupParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log()
	gmq := stage.Get(teamScanQueueGroupMember)

	groups, err := sv_group.New(z.client).List()
	if err != nil {
		l.Debug("Unable to retrieve groups", esl.Error(err))
		z.adb.Save(&GroupError{
			Dummy: "dummy",
			Error: err.Error(),
		})
		return err
	}
	for _, group := range groups {
		g, err := NewGroupFromJson(es_json.MustParse(group.Raw))
		if err != nil {
			return err
		}
		z.adb.Save(g)
		if z.adb.Error != nil {
			return z.adb.Error
		}
		gmq.Enqueue(&GroupMemberParam{
			GroupId: g.GroupId,
		})
	}

	if param.IsRetry {
		z.adb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&GroupError{})
	}
	return nil
}
