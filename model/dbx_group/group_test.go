package dbx_group

import (
	"encoding/json"
	"github.com/watermint/toolbox/model/dbx_api"
	"testing"
)

func TestModelGroup(t *testing.T) {
	j := `{
            "group_name": "Test group",
            "group_id": "g:e2db7665347abcd600000000001a2b3c",
            "group_management_type": {
                ".tag": "user_managed"
            },
            "member_count": 10
        }`

	g := &Group{}
	err := dbx_api.ParseModelJsonForTest(g, json.RawMessage(j))
	if err != nil {
		t.Error(err)
	}
	if g.GroupId != "g:e2db7665347abcd600000000001a2b3c" {
		t.Error("invalid")
	}
	if g.GroupExternalId != "" {
		t.Error("invalid")
	}
	if g.GroupManagementType != "user_managed" {
		t.Error("invalid")
	}
	if g.MemberCount != 10 {
		t.Error("invalid")
	}
}
