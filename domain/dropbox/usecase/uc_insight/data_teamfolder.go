package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
)

type TeamFolder struct {
	// primary keys
	TeamFolderId string `path:"team_folder_id" gorm:"primaryKey"`

	// attributes
	Name        string `path:"name"`
	Status      string `path:"status.\\.tag"`
	SyncSetting string `path:"sync_setting.\\.tag"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewTeamFolder(data es_json.Json) (tf *TeamFolder, err error) {
	tf = &TeamFolder{}
	if err = data.Model(tf); err != nil {
		return nil, err
	}
	return tf, nil
}

func (z tsImpl) scanTeamFolder(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	folders, err := sv_teamfolder.New(z.client).List()
	if err != nil {
		return err
	}

	for _, folder := range folders {
		f, err := NewTeamFolder(es_json.MustParse(folder.Raw))
		if err != nil {
			return err
		}
		z.db.Save(f)
	}
	return nil
}
