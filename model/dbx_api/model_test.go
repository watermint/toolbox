package dbx_api

import (
	"bytes"
	"encoding/json"
	"testing"
)

type Profile struct {
	Raw json.RawMessage

	// for processing, csv export, etc
	Email        string `path:"email,required"`
	TeamMemberId string `path:"team_member_id"`
	Status       string `path:"status.\\.tag"`
}

func TestParseModel(t *testing.T) {
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

	if ParseModelJsonForTest(&prof, json.RawMessage(rawJson)) != nil {
		t.Error("failed to parse")
	}

	if prof.Email != "tami@seagull.com" {
		t.Error("invalid email")
	}

	if bytes.Compare(prof.Raw, json.RawMessage(rawJson)) != 0 {
		t.Error("invalid raw data")
	}

	if ParseModelJsonForTest(&prof, json.RawMessage("{}")) == nil {
		t.Error("required field must be validated")
	}
}
