package mo_member_quota

import (
	"github.com/watermint/toolbox/infra/api/api_parser"
	"testing"
)

func TestQuota(t *testing.T) {
	{
		j := `{
    ".tag": "success",
    "user": {
      ".tag": "team_member_id",
      "team_member_id": "dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    },
    "quota_gb": 15
  }`

		q := &Quota{}
		if err := api_parser.ParseModelString(q, j); err != nil {
			t.Error(err)
		}
		if q.TeamMemberId != "dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error("invalid")
		}
		if q.Quota != 15 {
			t.Error("invalid")
		}
		if q.IsUnlimited() {
			t.Error("invalid")
		}
	}

	{
		j := `{
    ".tag": "success",
    "user": {
      ".tag": "team_member_id",
      "team_member_id": "dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    }
  }`

		q := &Quota{}
		if err := api_parser.ParseModelString(q, j); err != nil {
			t.Error(err)
		}
		if q.TeamMemberId != "dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error("invalid")
		}
		if q.Quota != 0 {
			t.Error("invalid")
		}
		if !q.IsUnlimited() {
			t.Error("invalid")
		}
	}
}
