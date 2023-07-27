package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleGroupMember = `{
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
		"access_type": {
			".tag": "owner"
		}
	}`
)

func TestNewGroupMemberFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		g, err := NewGroupMemberFromJson("g:1234567890", es_json.MustParseString(sampleGroupMember))
		if err != nil {
			t.Error(err)
		}
		if g.GroupId != "g:1234567890" {
			t.Error(g.GroupId)
		}
		if g.TeamMemberId != "dbmid:FDFSVF-DFSDF" {
			t.Error(g.TeamMemberId)
		}
		if g.Email != "tami@seagull.com" {
			t.Error(g.Email)
		}
		if g.AccessType != "owner" {
			t.Error(g.AccessType)
		}

		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&GroupMember{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(g).Error; err != nil {
			t.Error(err)
			return
		}

		g1 := &GroupMember{}
		db.First(&g1)
		if !reflect.DeepEqual(g, g1) {
			t.Error(g1)
		}
	})
}
