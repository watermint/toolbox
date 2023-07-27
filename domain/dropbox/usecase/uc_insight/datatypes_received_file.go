package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type ReceivedFile struct {
	// primary keys
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	FileId       string `path:"id" gorm:"primaryKey"`

	// attributes
	Name          string `path:"name"`
	AccessType    string `path:"access_type.\\.tag"`
	OwnerTeamId   string `path:"owner_team.id"`
	OwnerTeamName string `path:"owner_team.name"`
	PathDisplay   string `path:"path_display"`
	PathLower     string `path:"path_lower"`
	TimeInvited   string `path:"time_invited"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewReceivedFileFromJson(data es_json.Json) (rf *ReceivedFile, err error) {
	rf = &ReceivedFile{}
	if err = data.Model(rf); err != nil {
		return nil, err
	}
	return rf, nil
}

func NewReceivedFileFromJsonWithTeamMemberId(teamMemberId string, data es_json.Json) (rf *ReceivedFile, err error) {
	rf = &ReceivedFile{}
	if err = data.Model(rf); err != nil {
		return nil, err
	}
	rf.TeamMemberId = teamMemberId
	return rf, nil
}
