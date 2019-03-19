package mo_group_member

import (
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"testing"
)

func TestGroup(t *testing.T) {
	j := `{
  "profile": {
    "team_member_id": "dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx",
    "external_id": "xxxxx x+xxxxxxx.xxx-xxxxx@xxxxxxxxx.xxx",
    "account_id": "dbid:xxxxxxxxxxxxxxxxxxxxxxx-xxxxxxxxxxx",
    "email": "xxx+xxx@xxxxxxxxx.xxx",
    "email_verified": true,
    "status": {
      ".tag": "active"
    },
    "name": {
      "given_name": "xx",
      "surname": "xxxxxxx",
      "familiar_name": "xx",
      "display_name": "xx xxxxxxx",
      "abbreviated_name": "xx"
    },
    "membership_type": {
      ".tag": "full"
    },
    "joined_on": "2016-01-15T05:42:49Z"
  },
  "access_type": {
    ".tag": "member"
  }
}`
	m := &Member{}
	if err := api_parser.ParseModelString(m, j); err != nil {
		t.Error(err)
	}
	if m.TeamMemberId != "dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx" || m.AccessType != "member" {
		t.Error("invalid")
	}
}
