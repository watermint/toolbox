package dbx_teamfolder

import (
	"encoding/json"
)

const (
	StatusActive            = "active"
	StatusArchived          = "archived"
	StatusArchiveInProgress = "archive_in_progress"
)

type TeamFolder struct {
	Raw json.RawMessage `json:"-"`

	TeamFolderId string `path:"team_folder_id" json:"team_folder_id"`
	Name         string `path:"name" json:"name"`
	Status       string `path:"status.\\.tag" json:"status"`
	SyncSetting  string `path:"sync_setting.\\.tag" json:"sync_setting"`
}
