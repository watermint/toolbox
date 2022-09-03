package mo_profile

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestProfile(t *testing.T) {
	j := `{
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

	p := &Profile{}
	if err := api_parser.ParseModelString(p, j); err != nil {
		t.Error(err)
	}
	if p.TeamMemberId != "dbmid:FDFSVF-DFSDF" {
		t.Error("invalid")
	}
	if p.Status != "active" {
		t.Error("invalid")
	}
}
