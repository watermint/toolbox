package mo_teamfolder

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type TeamFolder struct {
	Raw                 json.RawMessage
	TeamFolderId        string `path:"team_folder_id" json:"team_folder_id"`
	Name                string `path:"name" json:"name"`
	Status              string `path:"status.\\.tag" json:"status"`
	IsTeamSharedDropbox bool   `path:"is_team_shared_dropbox" json:"is_team_shared_dropbox"`
	SyncSetting         string `path:"sync_setting.\\.tag" json:"sync_setting"`
}

func (z TeamFolder) ContentSyncSettings() (settings []ContentSyncSetting) {
	settings = make([]ContentSyncSetting, 0)

	j := es_json.MustParse(z.Raw)
	_ = j.FindArrayEach("content_sync_settings", func(e es_json.Json) error {
		folderId, found := e.FindString("id")
		if !found {
			return nil
		}
		syncSetting, found := e.FindString("sync_setting.\\.tag")
		if !found {
			return nil
		}

		settings = append(settings, ContentSyncSetting{
			Id:          folderId,
			SyncSetting: syncSetting,
		})
		return nil
	})
	return
}

type ContentSyncSetting struct {
	Id          string `json:"id"`
	SyncSetting string `json:"sync_setting"`
}
