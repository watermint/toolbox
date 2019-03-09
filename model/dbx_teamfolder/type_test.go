package dbx_teamfolder

import (
	"encoding/json"
	"github.com/watermint/toolbox/model/dbx_api"
	"testing"
)

func TestModelTeamFolder(t *testing.T) {
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

	err := dbx_api.ParseModelJsonForTest(tf, json.RawMessage(j))
	if err != nil {
		t.Error("failed parse", err)
	}
	if tf.Name != "Marketing" {
		t.Error("invalid")
	}
	if tf.TeamFolderId != "123456789" {
		t.Error("invalid")
	}
	if tf.SyncSetting != "default" {
		t.Error("invalid")
	}
	if tf.Status != "active" {
		t.Error("invalid")
	}
}
