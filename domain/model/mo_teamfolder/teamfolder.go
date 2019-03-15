package mo_teamfolder

import "encoding/json"

type TeamFolder struct {
	Raw                 json.RawMessage
	TeamFolderId        string `path:"team_folder_id"`
	Name                string `path:"name"`
	Status              string `path:"status.\\.tag"`
	IsTeamSharedDropbox bool   `path:"is_team_shared_dropbox"`
	SyncSetting         string `path:"sync_setting.\\.tag"`
}
