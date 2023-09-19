package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

func TestNewNamespaceMember(t *testing.T) {
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

	mu := &mo_sharedfolder_member.Metadata{}
	mg := &mo_sharedfolder_member.Metadata{}
	mi := &mo_sharedfolder_member.Metadata{}
	if err := api_parser.ParseModelString(mu, ju); err != nil {
		t.Error(err)
	}
	if err := api_parser.ParseModelString(mg, jg); err != nil {
		t.Error(err)
	}
	if err := api_parser.ParseModelString(mi, ji); err != nil {
		t.Error(err)
	}

	nsId := "ns:1234"
	nu := NewNamespaceMember(nsId, mu)
	ng := NewNamespaceMember(nsId, mg)
	ni := NewNamespaceMember(nsId, mi)

	{
		if nu.NamespaceId != nsId {
			t.Error(nu)
		}
		if nu.MemberType != "user" {
			t.Error(nu)
		}
		if nu.UserTeamMemberId != "dbmid:abcd1234" {
			t.Error(nu)
		}
		if nu.UserEmail != "bob@example.com" {
			t.Error(nu)
		}
		if nu.UserDisplayName != "Robert Smith" {
			t.Error(nu)
		}
		if nu.UserAccountId != "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc" {
			t.Error(nu)
		}
		if nu.AccessType != "owner" {
			t.Error(nu)
		}
		if nu.IsInherited != false {
			t.Error(nu)
		}
		if nu.SameTeam != "yes" {
			t.Error(nu)
		}
	}

	{
		if ng.NamespaceId != nsId {
			t.Error(ng)
		}
		if ng.MemberType != "group" {
			t.Error(ng)
		}
		if ng.GroupName != "Test group" {
			t.Error(ng)
		}
		if ng.GroupId != "g:e2db7665347abcd600000000001a2b3c" {
			t.Error(ng)
		}
		if ng.GroupType != "user_managed" {
			t.Error(ng)
		}
		if ng.AccessType != "editor" {
			t.Error(ng)
		}
		if ng.IsInherited != false {
			t.Error(ng)
		}
		if ng.SameTeam != "yes" {
			t.Error(ng)
		}
	}

	{
		if ni.NamespaceId != nsId {
			t.Error(ni)
		}
		if ni.MemberType != "invitee" {
			t.Error(ni)
		}
		if ni.InviteeEmail != "jessica@example.com" {
			t.Error(ni)
		}
		if ni.AccessType != "viewer" {
			t.Error(ni)
		}
		if ni.IsInherited != false {
			t.Error(ni)
		}
		if ni.SameTeam != "unknown" {
			t.Error(ni)
		}
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}

		if err := db.AutoMigrate(&NamespaceMember{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(nu).Error; err != nil {
			t.Error(err)
			return
		}
		nu1 := &NamespaceMember{}
		db.First(nu1)
		if !reflect.DeepEqual(nu, nu1) {
			t.Error(nu1)
		}
		db.Delete(nu1)

		if err := db.Create(ng).Error; err != nil {
			t.Error(err)
			return
		}
		ng1 := &NamespaceMember{}
		db.First(ng1)
		if !reflect.DeepEqual(ng, ng1) {
			t.Error(ng1)
		}
		db.Delete(ng1)

		if err := db.Create(ni).Error; err != nil {
			t.Error(err)
			return
		}
		ni1 := &NamespaceMember{}
		db.First(ni1)
		if !reflect.DeepEqual(ni, ni1) {
			t.Error(ni1)
		}
	})
}
