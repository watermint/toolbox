package sv_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

const (
	memberGetInfoSuccessRes = `[
  {
    ".tag": "member_info",
    "profile": {
      "team_member_id": "dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx",
      "external_id": "xxxxx x+xxxxxxx.xxx-xxxxx@xxxxxxxxx.xxx",
      "account_id": "xxxx:xxxxxxxxxxxxxxxxxxxxxxx-xxxxxxxxxxx",
      "email": "xxx+xxx@xxxxxxxxx.xxx",
      "email_verified": true,
      "status": {
        ".tag": "active"
      },
      "name": {
        "given_name": "デモ",
        "surname": "Dropbox",
        "familiar_name": "デモ",
        "display_name": "デモ Dropbox",
        "abbreviated_name": "デD"
      },
      "membership_type": {
        ".tag": "full"
      },
      "joined_on": "2016-01-15T05:42:49Z",
      "groups": [
        "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "g:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      ],
      "member_folder_id": "xxxxxxxxxx"
    },
    "role": {
      ".tag": "team_admin"
    }
  }
]`
)

func TestModelMemberImpl_Resolve(t *testing.T) {
	member := &mo_member.Member{}
	j := gjson.Parse(memberGetInfoSuccessRes)
	if !j.IsArray() {
		t.Error("invalid")
	}
	a := j.Array()[0]
	err := api_parser.ParseModel(member, a)
	if err != nil {
		t.Error("invalid")
	}
	if member.TeamMemberId != "dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx" {
		t.Error("invalid")
	}
}

// -- mock test impl

func TestMemberImpl_Add(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Add("test@example.com",
			AddWithGivenName("test"),
			AddWithSurname("example"),
			AddWithExternalId("ADSYNC test@test"),
			AddWithoutSendWelcomeEmail(),
			AddWithRole("user_admin"),
			AddWithDirectoryRestricted(true),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.Remove(&mo_member.Member{},
			Downgrade(),
			RemoveWipeData(),
			RetainTeamShares(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_ResolveByEmail(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.ResolveByEmail("test@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Update(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Update(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

// -- Test cached

func TestCachedMember_Add(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewCached(ctx)
		_, err := sv.Add("test@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_List(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewCached(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewCached(ctx)
		err := sv.Remove(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewCached(ctx)
		_, err := sv.Resolve("test")
		if err != ErrorMemberNotFoundForTeamMemberId && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_ResolveByEmail(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewCached(ctx)
		_, err := sv.ResolveByEmail("test@example.com")
		if err != ErrorMemberNotFoundForEmail && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Update(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Update(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
