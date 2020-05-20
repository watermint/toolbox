package mo_sharedfolder_member

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"testing"
)

func TestExternalOpt_Accept(t *testing.T) {
	eo := NewExternalOpt()
	if x := eo.Enabled(); x {
		t.Error(x)
	}

	iob := eo.Bind().(*bool)
	*iob = true
	if x := eo.Enabled(); !x {
		t.Error(x)
	}
	if x := eo.NameSuffix(); x != "External" {
		t.Error(x)
	}
	if x := eo.Desc(); x.Key() != MExternalOpt.Desc.Key() {
		t.Error(x)
	}

	// should not accept
	if x := eo.Accept(123); x {
		t.Error(x)
	}
	if x := eo.Accept(&User{IsSameTeam: true}); x {
		t.Error(x)
	}
	if x := eo.Accept(&Group{IsSameTeam: true}); x {
		t.Error(x)
	}
	if x := eo.Accept(&Invitee{InviteeEmail: "external@example.com"}); !x {
		t.Error(x)
	}

	member := mo_member.Member{
		Email: "internal@example.com",
		Raw:   json.RawMessage(`{"profile":{"status":{".tag":"active"}}}`),
	}
	members := []*mo_member.Member{&member}

	eo.SetMembers(members)

	if x := eo.Accept(&Invitee{InviteeEmail: "internal@example.com"}); x {
		t.Error(x)
	}
}
