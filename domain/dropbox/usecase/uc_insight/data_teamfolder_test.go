package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleTeamFolder = `{
            "content_sync_settings": [
                {
                    "id": "id:a4ayc_80_OEAAAAAAAAAXw",
                    "sync_setting": {
                        ".tag": "default"
                    }
                }
            ],
            "is_team_shared_dropbox": false,
            "name": "Marketing",
            "status": {
                ".tag": "active"
            },
            "sync_setting": {
                ".tag": "default"
            },
            "team_folder_id": "123456789"
        }`
)

func TestNewTeamFolderFromJson(t *testing.T) {
	tf, err := NewTeamFolder(es_json.MustParseString(sampleTeamFolder))
	if err != nil {
		t.Error(err)
	}

	if tf.TeamFolderId != "123456789" {
		t.Error(tf.TeamFolderId)
	}
	if tf.Name != "Marketing" {
		t.Error(tf.Name)
	}
	if tf.Status != "active" {
		t.Error(tf.Status)
	}
	if tf.SyncSetting != "default" {
		t.Error(tf.SyncSetting)
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&TeamFolder{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(tf).Error; err != nil {
			t.Error(err)
			return
		}

		tf1 := &TeamFolder{}
		db.First(&tf1)
		if !reflect.DeepEqual(tf, tf1) {
			t.Error(tf1)
		}
	})
}
