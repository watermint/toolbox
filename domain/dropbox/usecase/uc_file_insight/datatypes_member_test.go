package uc_file_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleMember = `{
            "profile": {
                "account_id": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc",
                "email": "tami@seagull.com",
                "email_verified": false,
                "external_id": "244423",
                "groups": [
                    "g:e2db7665347abcd600000000001a2b3c"
                ],
                "joined_on": "2015-05-12T15:50:38Z",
                "member_folder_id": "20",
                "membership_type": {
                    ".tag": "full"
                },
                "name": {
                    "abbreviated_name": "FF",
                    "display_name": "Franz Ferdinand (Personal)",
                    "familiar_name": "Franz",
                    "given_name": "Franz",
                    "surname": "Ferdinand"
                },
                "profile_photo_url": "https://dl-web.dropbox.com/account_photo/get/dbaphid%3AAAHWGmIXV3sUuOmBfTz0wPsiqHUpBWvv3ZA?vers=1556069330102&size=128x128",
                "secondary_emails": [
                    {
                        "email": "grape@strawberry.com",
                        "is_verified": false
                    },
                    {
                        "email": "apple@orange.com",
                        "is_verified": true
                    }
                ],
                "status": {
                    ".tag": "active"
                },
                "team_member_id": "dbmid:FDFSVF-DFSDF"
            },
            "roles": [
                {
                    "description": "Add, remove, and manage member accounts.",
                    "name": "User management admin",
                    "role_id": "pid_dbtmr:3456"
                }
            ]
        }`
)

func TestNewMemberFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		m, err := NewMemberFromJson(es_json.MustParseString(sampleMember))
		if err != nil {
			t.Error(err)
			return
		}
		if m.TeamMemberId != "dbmid:FDFSVF-DFSDF" {
			t.Error(m.TeamMemberId)
		}
		if m.Email != "tami@seagull.com" {
			t.Error(m.Email)
		}

		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		err = db.AutoMigrate(&Member{})
		if err != nil {
			t.Error(err)
			return
		}
		db.Create(m)

		m1 := &Member{}
		db.First(m1)
		if !reflect.DeepEqual(m, m1) {
			t.Error(m1)
		}
	})
}
