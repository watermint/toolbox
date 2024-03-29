package mo_teamfolder

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestTeamFolder(t *testing.T) {
	j := `{
            "team_folder_id": "123456789",
            "name": "Marketing",
            "status": {
                ".tag": "active"
            },
            "is_team_shared_dropbox": false,
            "sync_setting": {
                ".tag": "default"
            },
            "content_sync_settings": [
                {
                    "id": "id:a4ayc_80_OEAAAAAAAAAXw",
                    "sync_setting": {
                        ".tag": "default"
                    }
                }
            ]
        }`

	tf := &TeamFolder{}
	if err := api_parser.ParseModelString(tf, j); err != nil {
		t.Error(err)
	}
	if tf.TeamFolderId != "123456789" ||
		tf.Name != "Marketing" {
		t.Error("invalid")
	}

	css := tf.ContentSyncSettings()
	if lcss := len(css); lcss != 1 {
		t.Error(lcss)
	}
	if css[0].Id != "id:a4ayc_80_OEAAAAAAAAAXw" {
		t.Error(css[0].Id)
	}
	if css[0].SyncSetting != "default" {
		t.Error(css[0].SyncSetting)
	}
}
