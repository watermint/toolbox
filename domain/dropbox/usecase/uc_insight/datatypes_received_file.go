package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharing"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
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

func (z tsImpl) scanReceivedFile(teamMemberId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsMemberId(teamMemberId)
	received, err := sv_sharing.NewReceived(client).List()
	if err != nil {
		return err
	}
	for _, rf := range received {
		r, err := NewReceivedFileFromJsonWithTeamMemberId(teamMemberId, es_json.MustParse(rf.Raw))
		if err != nil {
			return err
		}
		z.db.Save(r)
	}
	return nil
}