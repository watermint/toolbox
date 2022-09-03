package mo_sharedfolder_member

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestMember(t *testing.T) {
	ju := `{
            "access_type": {
                ".tag": "owner"
            },
            "user": {
                "account_id": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc",
                "email": "bob@example.com",
                "display_name": "Robert Smith",
                "same_team": true,
                "team_member_id": "dbmid:abcd1234"
            },
            "permissions": [],
            "is_inherited": false,
            "time_last_seen": "2016-01-20T00:00:00Z",
            "platform_type": {
                ".tag": "unknown"
            }
        }`
	jg := `{
            "access_type": {
                ".tag": "editor"
            },
            "group": {
                "group_name": "Test group",
                "group_id": "g:e2db7665347abcd600000000001a2b3c",
                "group_management_type": {
                    ".tag": "user_managed"
                },
                "group_type": {
                    ".tag": "user_managed"
                },
                "is_member": false,
                "is_owner": false,
                "same_team": true,
                "member_count": 10
            },
            "permissions": [],
            "is_inherited": false
        }`
	ji := `{
            "access_type": {
                ".tag": "viewer"
            },
            "invitee": {
                ".tag": "email",
                "email": "jessica@example.com"
            },
            "permissions": [],
            "is_inherited": false
        }`

	{
		m := &Metadata{}
		mu := &User{}

		if err := api_parser.ParseModelString(m, ju); err != nil {
			t.Error(err)
		}
		if err := api_parser.ParseModelString(mu, ju); err != nil {
			t.Error(err)
		}
		if m.MemberType() != mu.MemberType() || m.MemberType() != MemberTypeUser {
			t.Error("invalid")
		}
		if _, e := m.User(); !e {
			t.Error("invalid")
		}
		if _, e := mu.User(); !e {
			t.Error("invalid")
		}
		if _, e := m.Group(); e {
			t.Error("invalid")
		}
		if _, e := mu.Group(); e {
			t.Error("invalid")
		}
		if _, e := m.Invitee(); e {
			t.Error("invalid")
		}
		if _, e := mu.Invitee(); e {
			t.Error("invalid")
		}
	}

	{
		m := &Metadata{}
		mg := &Group{}

		if err := api_parser.ParseModelString(m, jg); err != nil {
			t.Error(err)
		}
		if err := api_parser.ParseModelString(mg, jg); err != nil {
			t.Error(err)
		}
		if m.MemberType() != mg.MemberType() || m.MemberType() != MemberTypeGroup {
			t.Error("invalid")
		}
		if _, e := m.User(); e {
			t.Error("invalid")
		}
		if _, e := mg.User(); e {
			t.Error("invalid")
		}
		if _, e := m.Group(); !e {
			t.Error("invalid")
		}
		if _, e := mg.Group(); !e {
			t.Error("invalid")
		}
		if _, e := m.Invitee(); e {
			t.Error("invalid")
		}
		if _, e := mg.Invitee(); e {
			t.Error("invalid")
		}
	}

	{
		m := &Metadata{}
		mi := &Invitee{}

		if err := api_parser.ParseModelString(m, ji); err != nil {
			t.Error(err)
		}
		if err := api_parser.ParseModelString(mi, ji); err != nil {
			t.Error(err)
		}
		if m.MemberType() != mi.MemberType() || m.MemberType() != MemberTypeInvitee {
			t.Error("invalid")
		}
		if _, e := m.User(); e {
			t.Error("invalid")
		}
		if _, e := mi.User(); e {
			t.Error("invalid")
		}
		if _, e := m.Group(); e {
			t.Error("invalid")
		}
		if _, e := mi.Group(); e {
			t.Error("invalid")
		}
		if _, e := m.Invitee(); !e {
			t.Error("invalid")
		}
		if _, e := mi.Invitee(); !e {
			t.Error("invalid")
		}
	}
}
