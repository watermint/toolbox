package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	mountSample = `{
            "access_inheritance": {
                ".tag": "inherit"
            },
            "access_type": {
                ".tag": "owner"
            },
            "is_inside_team_folder": false,
            "is_team_folder": false,
            "link_metadata": {
                "audience_options": [
                    {
                        ".tag": "public"
                    },
                    {
                        ".tag": "team"
                    },
                    {
                        ".tag": "members"
                    }
                ],
                "current_audience": {
                    ".tag": "public"
                },
                "link_permissions": [
                    {
                        "action": {
                            ".tag": "change_audience"
                        },
                        "allow": true
                    }
                ],
                "password_protected": false,
                "url": ""
            },
            "name": "dir",
            "path_lower": "/dir",
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
            "preview_url": "https://www.dropbox.com/scl/fo/fir9vjelf",
            "shared_folder_id": "84528192421",
            "time_invited": "2016-01-20T00:00:00Z"
        }`
)

func TestMount_Exec(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		err = db.AutoMigrate(&Mount{})
		if err != nil {
			t.Error(err)
			return
		}

		{
			m0, err := NewMountFromJson(es_json.MustParseString(mountSample))
			if err != nil {
				t.Error(err)
				return
			}

			if err = db.Create(m0).Error; err != nil {
				t.Error(err)
				return
			}

			m1 := &Mount{}
			db.First(m1)

			if !reflect.DeepEqual(m0, m1) {
				t.Error("Not equal", m0, m1)
				return
			}

			db.Delete(m1)
		}

		{

			m0, err := NewMountFromJsonWithTeamMemberId("dbmid:12345", es_json.MustParseString(mountSample))
			if err != nil {
				t.Error(err)
				return
			}

			if err = db.Create(m0).Error; err != nil {
				t.Error(err)
				return
			}

			m1 := &Mount{}
			db.First(m1)

			if !reflect.DeepEqual(m0, m1) {
				t.Error("Not equal", m0, m1)
				return
			}
		}
	})
}
