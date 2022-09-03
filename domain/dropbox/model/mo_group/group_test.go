package mo_group

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestGroup(t *testing.T) {
	j := `{
            "group_name": "Test group",
            "group_id": "g:e2db7665347abcd600000000001a2b3c",
            "group_management_type": {
                ".tag": "user_managed"
            },
            "member_count": 10
        }`

	g := Group{}
	if err := api_parser.ParseModelString(&g, j); err != nil {
		t.Error(err)
	}
	if g.GroupId != "g:e2db7665347abcd600000000001a2b3c" {
		t.Error("invalid")
	}
	if g.GroupName != "Test group" {
		t.Error("invalid")
	}
	if g.GroupManagementType != "user_managed" {
		t.Error("invalid")
	}
	if g.MemberCount != 10 {
		t.Error("invalid")
	}
}
