package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type SharedLink struct {
	// primary keys
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	Url          string `path:"url" gorm:"primaryKey"`

	// attributes
	Name      string `path:"name"`
	PathLower string `path:"path_lower"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewSharedLink(data es_json.Json) (sl *SharedLink, err error) {
	sl = &SharedLink{}
	if err = data.Model(sl); err != nil {
		return nil, err
	}
	return sl, nil
}

func NewSharedLinkWithTeamMemberId(teamMemberId string, data es_json.Json) (sl *SharedLink, err error) {
	sl = &SharedLink{}
	if err = data.Model(sl); err != nil {
		return nil, err
	}
	sl.TeamMemberId = teamMemberId
	return sl, nil
}
