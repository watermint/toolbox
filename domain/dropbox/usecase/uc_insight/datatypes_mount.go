package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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
