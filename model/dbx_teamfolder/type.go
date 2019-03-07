package dbx_teamfolder

import (
	"encoding/json"
)

type TeamFolder struct {
	Raw json.RawMessage `json:"-"`

	TeamFolderId string `path:"team_folder_id" json:"team_folder_id"`
	Name         string `path:"name" json:"name"`
	SyncSetting  string `path:"sync_setting.\\.tag" json:"sync_setting"`
}
