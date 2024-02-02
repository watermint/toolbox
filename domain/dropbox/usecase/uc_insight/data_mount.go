package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type Mount struct {
	// primary keys
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	NamespaceId  string `path:"shared_folder_id" gorm:"primaryKey"`

	// attributes
	Name               string `path:"name"`
	Path               string `path:"path_lower"`
	AccessType         string `path:"access_type.\\.tag"`
	TimeInvited        string `path:"time_invited"`
	IsTeamFolder       bool   `path:"is_team_folder"`
	IsInsideTeamFolder bool   `path:"is_inside_team_folder"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

type MountParam struct {
	TeamMemberId string `path:"team_member_id" json:"team_member_id"`
	IsRetry      bool   `path:"is_retry" json:"is_retry"`
}

type MountError struct {
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	Error        string `path:"error_summary"`
}

func (z MountError) ToParam() interface{} {
	return &MountParam{
		TeamMemberId: z.TeamMemberId,
		IsRetry:      true,
	}
}

func NewMountFromJson(data es_json.Json) (m *Mount, err error) {
	m = &Mount{}
	if err = data.Model(m); err != nil {
		return nil, err
	}
	return m, nil
}

func NewMountFromJsonWithTeamMemberId(teamMemberId string, data es_json.Json) (m *Mount, err error) {
	m = &Mount{}
	if err = data.Model(m); err != nil {
		return nil, err
	}
	m.TeamMemberId = teamMemberId
	return m, nil
}

func (z tsImpl) scanMount(param *MountParam, stage eq_sequence.Stage, admin *mo_profile.Profile, team *mo_team.Info) (err error) {
	l := z.ctl.Log().With(esl.String("teamMemberId", param.TeamMemberId))
	client := z.client.AsMemberId(param.TeamMemberId)
	qnd := stage.Get(teamScanQueueNamespaceDetail)

	mountables, err := sv_sharedfolder_mount.New(client).Mountables()
	if err != nil {
		l.Debug("Unable to retrieve mountables", esl.Error(err))
		z.adb.Save(&MountError{
			TeamMemberId: param.TeamMemberId,
			Error:        err.Error(),
		})
		return err
	}

	for _, mount := range mountables {
		m, err := NewMountFromJsonWithTeamMemberId(param.TeamMemberId, es_json.MustParse(mount.Raw))
		if err != nil {
			return err
		}
		z.adb.Save(m)

		if team.TeamId != mount.OwnerTeamId {
			qnd.Enqueue(&NamespaceDetailParam{
				NamespaceId: mount.SharedFolderId,
			})
		}
	}

	if param.IsRetry {
		z.adb.Delete(&MountError{}, "team_member_id = ?", param.TeamMemberId)
	}

	return nil
}
