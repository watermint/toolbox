package mo_teamfolder

import (
	"github.com/watermint/toolbox/infra/api/api_parser"
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
}
