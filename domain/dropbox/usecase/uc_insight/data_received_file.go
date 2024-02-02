package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharing"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
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

type ReceivedFileParam struct {
	TeamMemberId string `path:"team_member_id" json:"team_member_id"`
	IsRetry      bool   `path:"is_retry" json:"is_retry"`
}

type ReceivedFileError struct {
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	Error        string `path:"error_summary"`
}

func (z ReceivedFileError) ToParam() interface{} {
	return &ReceivedFileParam{
		TeamMemberId: z.TeamMemberId,
		IsRetry:      true,
	}
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

func (z tsImpl) scanReceivedFile(param *ReceivedFileParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.client.Log().With(esl.String("teamMemberId", param.TeamMemberId))
	client := z.client.AsMemberId(param.TeamMemberId)
	received, err := sv_sharing.NewReceived(client).List()
	dbxErr := dbx_error.NewErrors(err)
	switch {
	case dbxErr == nil:
		// fall through
	case dbxErr.IsEmailUnverified():
		l.Debug("Email unverified, skip")
		return nil
	default:
		l.Debug("List received files", esl.Error(err))

		z.adb.Save(&ReceivedFileError{
			TeamMemberId: param.TeamMemberId,
		})
		return err
	}

	for _, rf := range received {
		r, err := NewReceivedFileFromJsonWithTeamMemberId(param.TeamMemberId, es_json.MustParse(rf.Raw))
		if err != nil {
			return err
		}
		z.adb.Save(r)
	}

	if param.IsRetry {
		z.adb.Delete(&ReceivedFileError{}, "team_member_id = ?", param.TeamMemberId)
	}

	return nil
}
