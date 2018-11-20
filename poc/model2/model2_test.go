package model2

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
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

	prof := Profile{}

	if !Parse(&prof, json.RawMessage(rawJson)) {
		t.Error("failed to parse")
	}

	if prof.Email != "tami@seagull.com" {
		t.Error("invalid email")
	}

	if Parse(&prof, json.RawMessage("{}")) {
		t.Error("required field must be validated")
	}
}
