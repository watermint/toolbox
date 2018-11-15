package model

import (
	"testing"
)

func TestProfile_Email(t *testing.T) {
	rawJson := `{
                "team_member_id": "dbmid:FDFSVF-DFSDF",
                "email": "tami@seagull.com",
                "email_verified": false,
                "status": {
                    ".tag": "active"
                },
                "name": {
                    "given_name": "Franz",
                    "surname": "Ferdinand",
                    "familiar_name": "Franz",
                    "display_name": "Franz Ferdinand (Personal)",
                    "abbreviated_name": "FF"
                },
                "membership_type": {
                    ".tag": "full"
                },
                "groups": [
                    "g:e2db7665347abcd600000000001a2b3c"
                ],
                "member_folder_id": "20",
                "external_id": "244423",
                "account_id": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc",
                "joined_on": "2015-05-12T15:50:38Z"
            }`

	prof := Profile{
		Raw: Raw{Json: []byte(rawJson)},
	}

	if x, ok := prof.Email(); ok {
		if x != "tami@seagull.com" {
			t.Error("email not match")
		}
	} else {
		t.Error("email not found")
	}

	if x, ok := prof.TeamMemberId(); ok {
		if x != "dbmid:FDFSVF-DFSDF" {
			t.Error("team member id not match")
		}
	} else {
		t.Error("team member id not found")
	}

	if x, ok := prof.Status(); ok {
		if x != "active" {
			t.Error("status not match")
		}
	} else {
		t.Error("status not found")
	}

	ColumnHeader(&prof)
}
