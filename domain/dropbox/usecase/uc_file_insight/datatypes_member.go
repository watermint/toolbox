package uc_file_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Member struct {
	// primary keys
	TeamMemberId string `path:"profile.team_member_id" gorm:"primaryKey"`

	// attributes
	Email          string `path:"profile.email"`
	DisplayName    string `path:"profile.name.display_name"`
	MemberFolderId string `path:"profile.member_folder_id"`
	Status         string `path:"profile.status.\\.tag"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewMemberFromJson(data es_json.Json) (m *Member, err error) {
	m = &Member{}
	if err = data.Model(m); err != nil {
		return nil, err
	}
	return m, nil
}
