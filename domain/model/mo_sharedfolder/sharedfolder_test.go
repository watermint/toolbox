package mo_sharedfolder

import (
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"testing"
)

func TestSharedFolder(t *testing.T) {
	j := `{
            "access_type": {
                ".tag": "owner"
            },
            "is_inside_team_folder": false,
            "is_team_folder": false,
            "name": "dir",
            "policy": {
                "acl_update_policy": {
                    ".tag": "owner"
                },
                "shared_link_policy": {
                    ".tag": "anyone"
                },
                "member_policy": {
                    ".tag": "anyone"
                },
                "resolved_member_policy": {
                    ".tag": "team"
                }
            },
            "preview_url": "https://www.dropbox.com/scl/fo/fir9vjelf",
            "shared_folder_id": "84528192421",
            "time_invited": "2016-01-20T00:00:00Z",
            "path_lower": "/dir",
            "link_metadata": {
                "audience_options": [
                    {
                        ".tag": "public"
                    },
                    {
                        ".tag": "team"
                    },
                    {
                        ".tag": "members"
                    }
                ],
                "current_audience": {
                    ".tag": "public"
                },
                "link_permissions": [
                    {
                        "action": {
                            ".tag": "change_audience"
                        },
                        "allow": true
                    }
                ],
                "password_protected": false,
                "url": ""
            },
            "permissions": [],
            "access_inheritance": {
                ".tag": "inherit"
            }
        }`

	sf := &SharedFolder{}
	if err := api_parser.ParseModelString(sf, j); err != nil {
		t.Error(err)
	}
	if sf.SharedFolderId != "84528192421" || sf.AccessType != "owner" || sf.Name != "dir" {
		t.Error("invalid")
	}
}
