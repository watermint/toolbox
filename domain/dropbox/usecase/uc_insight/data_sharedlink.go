package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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

func (z tsImpl) scanSharedLink(teamMemberId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsMemberId(teamMemberId)

	links, err := sv_sharedlink.New(client).List()
	if err != nil {
		return err
	}

	for _, link := range links {
		l, err := NewSharedLinkWithTeamMemberId(teamMemberId, es_json.MustParse(link.Metadata().Raw))
		if err != nil {
			return err
		}
		z.adb.Save(l)
	}
	return nil
}
