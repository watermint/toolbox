package mo_namespace

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestNamespace(t *testing.T) {
	j1 := `{
            "name": "Marketing",
            "namespace_id": "123456789",
            "namespace_type": {
                ".tag": "shared_folder"
            }
        }`
	j2 := `{
            "name": "Franz Ferdinand",
            "namespace_id": "123456789",
            "namespace_type": {
                ".tag": "team_member_folder"
            },
            "team_member_id": "dbmid:1234567"
        }`

	n1 := &Namespace{}
	n2 := &Namespace{}

	if err := api_parser.ParseModelString(n1, j1); err != nil {
		t.Error(err)
	}
	if err := api_parser.ParseModelString(n2, j2); err != nil {
		t.Error(err)
	}
	if n1.Name != "Marketing" ||
		n1.NamespaceId != "123456789" ||
		n1.NamespaceType != "shared_folder" ||
		n1.TeamMemberId != "" {
		t.Error("invalid")
	}
	if n2.Name != "Franz Ferdinand" ||
		n2.NamespaceId != "123456789" ||
		n2.NamespaceType != "team_member_folder" ||
		n2.TeamMemberId != "dbmid:1234567" {
		t.Error("invalid")
	}
}
