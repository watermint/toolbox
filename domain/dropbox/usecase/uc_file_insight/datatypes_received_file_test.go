package uc_file_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

const (
	mountReceivedFile = `{
            "access_type": {
                ".tag": "viewer"
            },
            "id": "id:3kmLmQFnf1AAAAAAAAAAAw",
            "name": "file.txt",
            "owner_display_names": [
                "Jane Doe"
            ],
            "owner_team": {
                "id": "dbtid:AAFdgehTzw7WlXhZJsbGCLePe8RvQGYDr-I",
                "name": "Acme, Inc."
            },
            "path_display": "/dir/file.txt",
            "path_lower": "/dir/file.txt",
            "permissions": [],
            "policy": {
                "acl_update_policy": {
                    ".tag": "owner"
                },
                "member_policy": {
                    ".tag": "anyone"
                },
                "resolved_member_policy": {
                    ".tag": "team"
                },
                "shared_link_policy": {
                    ".tag": "anyone"
                }
            },
            "preview_url": "https://www.dropbox.com/scl/fi/fir9vjelf",
            "time_invited": "2016-01-20T00:00:00Z"
        }`
)

func TestNewReceivedFileFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		err = db.AutoMigrate(&ReceivedFile{})
		if err != nil {
			return
		}

		rf, err := NewReceivedFileFromJson(es_json.MustParseString(mountReceivedFile))
		if err != nil {
			t.Error(err)
			return
		}
		if err = db.Create(rf).Error; err != nil {
			t.Error(err)
			return
		}
		rf1 := &ReceivedFile{}
		db.First(rf1)
		if !reflect.DeepEqual(rf, rf1) {
			t.Error("Not equal", rf, rf1)
			return
		}
		db.Delete(rf1)
	})
}

func TestNewReceivedFileFromJsonWithTeamMemberId(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		err = db.AutoMigrate(&ReceivedFile{})
		if err != nil {
			return
		}

		rf, err := NewReceivedFileFromJsonWithTeamMemberId("dbmid:12345", es_json.MustParseString(mountReceivedFile))
		if err != nil {
			t.Error(err)
			return
		}
		if err = db.Create(rf).Error; err != nil {
			t.Error(err)
			return
		}
		rf1 := &ReceivedFile{}
		db.First(rf1)
		if !reflect.DeepEqual(rf, rf1) {
			t.Error("Not equal", rf, rf1)
			return
		}
		db.Delete(rf1)
	})
}
