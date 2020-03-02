package sv_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"strings"
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

func TestEndToEndMemberImpl_ResolveByEmail(t *testing.T) {
	qt_api.DoTestBusinessInfo(func(ctx api_context.Context) {
		svm := New(ctx)
		members, err := svm.List()
		if err != nil {
			t.Error(err)
		}

		for i, member := range members {
			if i > 10 {
				break
			}
			m1, err := svm.Resolve(member.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
			m2, err := svm.ResolveByEmail(member.Email)
			if err != nil {
				t.Error(err)
			}

			if m1.TeamMemberId != member.TeamMemberId ||
				m2.TeamMemberId != member.TeamMemberId {
				t.Error("invalid")
			}
		}

		_, err = svm.Resolve("dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx")
		if err == nil {
			t.Error("invalid")
		}

		_, err = svm.ResolveByEmail("non_existent@example.com")
		if err == nil {
			t.Error("invalid")
		}
	})
}

func TestEndToEndMemberImpl_ListResolve(t *testing.T) {
	qt_api.DoTestBusinessInfo(func(ctx api_context.Context) {
		ls := newTest(ctx)
		members, err := ls.List()
		if err != nil {
			t.Error(err)
			return
		}
		if len(members) < 1 {
			t.Error("invalid")
		}
		if !strings.Contains(members[0].Email, "@") {
			t.Error("invalid")
		}

		m, err := ls.Resolve(members[0].TeamMemberId)
		if err != nil {
			t.Error("failed fetch")
		}
		if m.TeamMemberId != members[0].TeamMemberId {
			t.Error("invalid")
		}
		if m.Email != members[0].Email {
			t.Error("invalid")
		}
	})
}

// -- mock test impl

func TestMemberImpl_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Add("test@example.com",
			AddWithGivenName("test"),
			AddWithSurname("example"),
			AddWithExternalId("ADSYNC test@test"),
			AddWithoutSendWelcomeEmail(),
			AddWithRole("user_admin"),
			AddWithDirectoryRestricted(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
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
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_ResolveByEmail(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.ResolveByEmail("test@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Update(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Update(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

// -- Test cached

func TestCachedMember_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.Add("test@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		err := sv.Remove(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.Resolve("test")
		if err != ErrorMemberNotFoundForTeamMemberId {
			t.Error(err)
		}
	})
}

func TestCachedMember_ResolveByEmail(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.ResolveByEmail("test@example.com")
		if err != ErrorMemberNotFoundForEmail {
			t.Error(err)
		}
	})
}

func TestCachedMember_Update(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Update(&mo_member.Member{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
