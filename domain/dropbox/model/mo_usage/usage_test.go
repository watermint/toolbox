package mo_usage

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

const (
	individual = `{
  "used": 1433469,
  "allocation": {
    ".tag": "individual",
    "allocated": 2147483648
  }
}`
	team = `{
  "used": 34767814042,
  "allocation": {
    ".tag": "team",
    "used": 86112460529,
    "allocated": 10995116277760,
    "user_within_team_space_allocated": 0,
    "user_within_team_space_limit_type": {
      ".tag": "off"
    },
    "user_within_team_space_used_cached": 34760310091
  }
}`
)

func TestUsage(t *testing.T) {
	{
		u := &Usage{}
		if err := api_parser.ParseModelString(u, individual); err != nil {
			t.Error(err)
		}
		if u.Used != 1433469 {
			t.Error("invalid")
		}
		if u.Allocation != "individual" {
			t.Error("invalid")
		}
		if u.Allocated != 2147483648 {
			t.Error("invalid")
		}
	}
	{
		u := &Usage{}
		if err := api_parser.ParseModelString(u, team); err != nil {
			t.Error(err)
		}
		if u.Used != 34767814042 {
			t.Error("invalid")
		}
		if u.Allocation != "team" {
			t.Error("invalid")
		}
		if u.Allocated != 10995116277760 {
			t.Error("invalid")
		}
		if u.TeamUserWithinTeamSpaceAllocated != 0 {
			t.Error("invalid")
		}
		if u.TeamUserWithinTeamSpaceLimitType != "off" {
			t.Error("invalid")
		}
		if u.TeamUserWithinTeamSpaceUsedCached != 34760310091 {
			t.Error("invalid")
		}
	}
}
