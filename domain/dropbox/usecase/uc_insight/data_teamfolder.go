package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"gorm.io/gorm"
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

type TeamFolderError struct {
	Dummy string `path:"dummy" gorm:"primaryKey"`
	ApiError
}

func (z TeamFolderError) ToParam() interface{} {
	return &TeamFolderParam{
		IsRetry: true,
	}
}

type TeamFolderParam struct {
	IsRetry bool `path:"is_retry" json:"is_retry"`
}

func NewTeamFolder(data es_json.Json) (tf *TeamFolder, err error) {
	tf = &TeamFolder{}
	if err = data.Model(tf); err != nil {
		return nil, err
	}
	return tf, nil
}

func (z tsImpl) scanTeamFolder(param *TeamFolderParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	l := z.client.Log()
	folders, err := sv_teamfolder.New(z.client).List()
	if err != nil {
		l.Debug("Unable to retrieve team folders", esl.Error(err))
		z.adb.Save(&TeamFolderError{
			Dummy:    "dummy",
			ApiError: ApiErrorFromError(err),
		})
		return err
	}

	for _, folder := range folders {
		f, err := NewTeamFolder(es_json.MustParse(folder.Raw))
		if err != nil {
			return err
		}
		z.adb.Save(f)
	}

	if param.IsRetry {
		z.adb.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&TeamFolderError{})
	}
	return nil
}
