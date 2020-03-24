package mo_group_member

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_parser"
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

func TestGroupMember_Group(t *testing.T) {
	j := `{
  "group": {
    "group_name": "営業部",
    "group_id": "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "group_external_id": "xxxxx x:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "member_count": 1,
    "group_management_type": {
      ".tag": "company_managed"
    }
  },
  "member": {
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
        "display_name": "xx xx",
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
  }
}
`

	gm := &GroupMember{}
	if err := api_parser.ParseModelString(gm, j); err != nil {
		t.Error(err)
	}
	if gm.GroupId != gm.Group().GroupId || gm.GroupId == "" {
		t.Error("invalid")
	}
	if gm.TeamMemberId != gm.Member().TeamMemberId || gm.TeamMemberId == "" {
		t.Error("invalid")
	}
	if gm.GroupName != gm.Group().GroupName || gm.GroupName != "営業部" {
		t.Error("invalid")
	}

	// round trip test 1 : bytes
	{
		jb1, err := json.Marshal(gm.Raw)
		if err != nil {
			t.Error(err)
		}
		gm2 := &GroupMember{}
		if err := api_parser.ParseModelRaw(gm2, jb1); err != nil {
			t.Error(err)
		}
		if gm2.GroupId != gm2.Group().GroupId || gm2.GroupId == "" {
			t.Error("invalid")
		}
		if gm2.TeamMemberId != gm2.Member().TeamMemberId || gm2.TeamMemberId == "" {
			t.Error("invalid")
		}
		if gm2.GroupName != gm2.Group().GroupName || gm2.GroupName != "営業部" {
			t.Error("invalid")
		}
	}

	//// round trip test 2 : report
	//{
	//	ec := app2.NewExecContextForTest()
	//	reportPath := filepath.Join(ec.JobsPath(), "report")
	//	report := app_report_legacy.Factory{}
	//	report.ExecContext = ec
	//	report.Path = reportPath
	//	if err := report.Init(ec); err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	if err := report.Report(gm); err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	report.Close()
	//
	//	rp := filepath.Join(reportPath, "GroupMember.json")
	//	rf, err := os.Open(rp)
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	defer rf.Close()
	//	jb1, err := ioutil.ReadAll(rf)
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//
	//	gm2 := &GroupMember{}
	//	if err := api_parser.ParseModelRaw(gm2, jb1); err != nil {
	//		t.Error(err)
	//	}
	//	if gm2.GroupId != gm2.Group().GroupId || gm2.GroupId == "" {
	//		t.Error("invalid")
	//	}
	//	if gm2.TeamMemberId != gm2.Member().TeamMemberId || gm2.TeamMemberId == "" {
	//		t.Error("invalid")
	//	}
	//	if gm2.Name != gm2.Group().Name || gm2.Name != "営業部" {
	//		t.Error("invalid")
	//	}
	//}

}
