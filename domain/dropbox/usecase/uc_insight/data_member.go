package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type Member struct {
	// primary keys
	TeamMemberId string `path:"profile.team_member_id" gorm:"primaryKey"`

	// attributes
	Email          string `path:"profile.email"`
	DisplayName    string `path:"profile.name.display_name"`
	MemberFolderId string `path:"profile.member_folder_id" gorm:"index"`
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

func (z tsImpl) scanMembers(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	var lastErr error
	opErr := sv_member.New(z.client).ListEach(func(member *mo_member.Member) bool {
		m, err := NewMemberFromJson(es_json.MustParse(member.Raw))
		if err != nil {
			lastErr = err
			return false
		}
		z.adb.Save(m)
		if z.adb.Error != nil {
			lastErr = z.adb.Error
			return false
		}
		if err = z.dispatchMember(member, stage, admin); err != nil {
			lastErr = err
			return false
		}

		return true
	}, sv_member.IncludeDeleted(true))

	return es_lang.NewMultiErrorOrNull(opErr, lastErr)
}
