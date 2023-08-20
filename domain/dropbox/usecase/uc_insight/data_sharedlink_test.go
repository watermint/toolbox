package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleSharedLink = `{
            ".tag": "file",
            "client_modified": "2015-05-12T15:50:38Z",
            "id": "id:a4ayc_80_OEAAAAAAAAAXw",
            "link_permissions": {
                "allow_comments": true,
                "allow_download": true,
                "audience_options": [
                    {
                        "allowed": true,
                        "audience": {
                            ".tag": "public"
                        }
                    },
                    {
                        "allowed": false,
                        "audience": {
                            ".tag": "team"
                        }
                    },
                    {
                        "allowed": true,
                        "audience": {
                            ".tag": "no_one"
                        }
                    }
                ],
                "can_allow_download": true,
                "can_disallow_download": false,
                "can_remove_expiry": false,
                "can_remove_password": true,
                "can_revoke": false,
                "can_set_expiry": false,
                "can_set_password": true,
                "can_use_extended_sharing_controls": false,
                "require_password": false,
                "resolved_visibility": {
                    ".tag": "public"
                },
                "revoke_failure_reason": {
                    ".tag": "owner_only"
                },
                "team_restricts_comments": true,
                "visibility_policies": [
                    {
                        "allowed": true,
                        "policy": {
                            ".tag": "public"
                        },
                        "resolved_policy": {
                            ".tag": "public"
                        }
                    },
                    {
                        "allowed": true,
                        "policy": {
                            ".tag": "password"
                        },
                        "resolved_policy": {
                            ".tag": "password"
                        }
                    }
                ]
            },
            "name": "Prime_Numbers.txt",
            "path_lower": "/homework/math/prime_numbers.txt",
            "rev": "a1c10ce0dd78",
            "server_modified": "2015-05-12T15:50:38Z",
            "size": 7212,
            "team_member_info": {
                "display_name": "Roger Rabbit",
                "member_id": "dbmid:abcd1234",
                "team_info": {
                    "id": "dbtid:AAFdgehTzw7WlXhZJsbGCLePe8RvQGYDr-I",
                    "name": "Acme, Inc."
                }
            },
            "url": "https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0"
        }`
)

func TestNewSharedLink(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		sl, err := NewSharedLink(es_json.MustParseString(sampleSharedLink))
		if err != nil {
			t.Error(err)
		}
		if sl.TeamMemberId != "" {
			t.Error(sl.TeamMemberId)
		}
		if sl.Url != "https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0" {
			t.Error(sl.Url)
		}
		if sl.Name != "Prime_Numbers.txt" {
			t.Error(sl.Name)
		}
		if sl.PathLower != "/homework/math/prime_numbers.txt" {
			t.Error(sl.PathLower)
		}

		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&SharedLink{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(sl).Error; err != nil {
			return
		}

		sl1 := &SharedLink{}
		db.First(&sl1)
		if !reflect.DeepEqual(sl, sl1) {
			t.Error(sl1)
		}
	})
}

func TestNewSharedLinkWithTeamMemberId(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		sl, err := NewSharedLinkWithTeamMemberId("dbmid:abcd1234", es_json.MustParseString(sampleSharedLink))
		if err != nil {
			t.Error(err)
		}
		if sl.TeamMemberId != "dbmid:abcd1234" {
			t.Error(sl.TeamMemberId)
		}
		if sl.Url != "https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0" {
			t.Error(sl.Url)
		}
		if sl.Name != "Prime_Numbers.txt" {
			t.Error(sl.Name)
		}
		if sl.PathLower != "/homework/math/prime_numbers.txt" {
			t.Error(sl.PathLower)
		}

		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&SharedLink{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(sl).Error; err != nil {
			return
		}

		sl1 := &SharedLink{}
		db.First(&sl1)
		if !reflect.DeepEqual(sl, sl1) {
			t.Error(sl1)
		}
	})
}
