package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type SharedLink struct {
	// primary keys
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	Url          string `path:"url" gorm:"primaryKey"`

	// attributes
	FileId    string `path:"id" gorm:"index"`
	Name      string `path:"name"`
	PathLower string `path:"path_lower"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

type SharedLinkParam struct {
	TeamMemberId string `path:"team_member_id" json:"team_member_id"`
	IsRetry      bool   `path:"is_retry" json:"is_retry"`
}

type SharedLinkError struct {
	TeamMemberId string `path:"team_member_id" gorm:"primaryKey"`
	ApiError
}

func (z SharedLinkError) ToParam() interface{} {
	return &SharedLinkParam{
		TeamMemberId: z.TeamMemberId,
		IsRetry:      true,
	}
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

func (z tsImpl) scanSharedLink(param *SharedLinkParam, stage eq_sequence.Stage, admin *mo_profile.Profile, baseNamespace dbx_filesystem.BaseNamespaceType) (err error) {
	l := z.client.Log().With(esl.String("teamMemberId", param.TeamMemberId))
	client := z.client.AsMemberId(param.TeamMemberId, baseNamespace)

	links, err := sv_sharedlink.New(client).List()
	if err != nil {
		l.Debug("Unable to retrieve links", esl.Error(err))
		z.db.Save(&SharedLinkError{
			TeamMemberId: param.TeamMemberId,
			ApiError:     ApiErrorFromError(err),
		})
		return err
	}

	for _, link := range links {
		l, err := NewSharedLinkWithTeamMemberId(param.TeamMemberId, es_json.MustParse(link.Metadata().Raw))
		if err != nil {
			return err
		}
		z.db.Save(l)
	}

	if param.IsRetry {
		z.db.Delete(&SharedLinkError{}, "team_member_id = ?", param.TeamMemberId)
	}
	return nil
}
