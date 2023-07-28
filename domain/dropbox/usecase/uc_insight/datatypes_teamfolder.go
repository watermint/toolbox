package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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
