package mo_sharedfolder_member

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"testing"
)

func TestInternalOpt_Accept(t *testing.T) {
	io := NewInternalOpt()
	if x := io.Enabled(); x {
		t.Error(x)
	}

	iob := io.Bind().(*bool)
	*iob = true
	if x := io.Enabled(); !x {
		t.Error(x)
	}
	if x := io.NameSuffix(); x != "Internal" {
		t.Error(x)
	}
	if x := io.Desc(); x.Key() != MInternalOpt.Desc.Key() {
		t.Error(x)
	}

	// should not accept
	if x := io.Accept(123); x {
		t.Error(x)
	}
	if x := io.Accept(&User{IsSameTeam: true}); !x {
		t.Error(x)
	}
	if x := io.Accept(&Group{IsSameTeam: true}); !x {
		t.Error(x)
	}
	if x := io.Accept(&Invitee{InviteeEmail: "external@example.com"}); x {
		t.Error(x)
	}

	member := mo_member.Member{
		Email: "internal@example.com",
		Raw:   json.RawMessage(`{"profile":{"status":{".tag":"active"}}}`),
	}
	members := []*mo_member.Member{&member}

	io.SetMembers(members)

	if x := io.Accept(&Invitee{InviteeEmail: "internal@example.com"}); !x {
		t.Error(x)
	}
}
