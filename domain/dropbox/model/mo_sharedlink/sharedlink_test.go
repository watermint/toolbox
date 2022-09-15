package mo_sharedlink

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestSharedLink(t *testing.T) {
	j1 := `{
            ".tag": "file",
            "url": "https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0",
            "name": "Prime_Numbers.txt",
            "link_permissions": {
                "can_revoke": false,
                "resolved_visibility": {
                    ".tag": "public"
                },
                "revoke_failure_reason": {
                    ".tag": "owner_only"
                }
            },
            "client_modified": "2015-05-12T15:50:38Z",
            "server_modified": "2015-05-12T15:50:38Z",
            "rev": "a1c10ce0dd78",
            "size": 7212,
            "id": "id:a4ayc_80_OEAAAAAAAAAXw",
            "path_lower": "/homework/math/prime_numbers.txt",
            "team_member_info": {
                "team_info": {
                    "id": "dbtid:AAFdgehTzw7WlXhZJsbGCLePe8RvQGYDr-I",
                    "name": "Acme, Inc."
                },
                "display_name": "Roger Rabbit",
                "member_id": "dbmid:abcd1234"
            }
        }`
	j2 := `{
  ".tag": "folder",
  "url": "https://www.dropbox.com/sh/xxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxxxxxxxxxxx?dl=0",
  "id": "id:xxxxxxxxxxxxxxxxxxxxxx",
  "name": "Apps",
  "expires": "2019-10-23T12:41:00Z",
  "path_lower": "/xxxx",
  "link_permissions": {
    "resolved_visibility": {
      ".tag": "public"
    },
    "requested_visibility": {
      ".tag": "public"
    },
    "can_revoke": true,
    "visibility_policies": [
      {
        "policy": {
          ".tag": "public"
        },
        "resolved_policy": {
          ".tag": "public"
        },
        "allowed": true
      },
      {
        "policy": {
          ".tag": "team_only"
        },
        "resolved_policy": {
          ".tag": "team_only"
        },
        "allowed": true
      },
      {
        "policy": {
          ".tag": "password"
        },
        "resolved_policy": {
          ".tag": "password"
        },
        "allowed": true
      }
    ],
    "can_set_expiry": true,
    "can_remove_expiry": true,
    "allow_download": true,
    "can_allow_download": true,
    "can_disallow_download": true,
    "allow_comments": true,
    "team_restricts_comments": false
  },
  "team_member_info": {
    "team_info": {
      "id": "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "name": "xxxxxxxxx xxx"
    },
    "display_name": "xx xxxxxxx",
    "member_id": "dbmid:xxxxxxxxxxxxxx-xxxxxxxx-xxxx-xxxxxx"
  }
}`

	f1 := &Metadata{}
	f2 := &Metadata{}

	if err := api_parser.ParseModelString(f1, j1); err != nil {
		t.Error(err)
	}
	if err := api_parser.ParseModelString(f2, j2); err != nil {
		t.Error(err)
	}
	if f1.SharedLinkId() != "id:a4ayc_80_OEAAAAAAAAAXw" ||
		f1.Id != "id:a4ayc_80_OEAAAAAAAAAXw" ||
		f1.LinkVisibility() != "public" ||
		f1.Visibility != "public" {
		t.Error("invalid")
	}
	g1, ok := f1.File()
	if !ok || g1.Tag != "file" {
		t.Error("invalid")
	}
	if g1.SharedLinkId() != f1.SharedLinkId() {
		t.Error("invalid")
	}
	if f2.SharedLinkId() != "id:xxxxxxxxxxxxxxxxxxxxxx" ||
		f2.Id != "id:xxxxxxxxxxxxxxxxxxxxxx" ||
		f2.LinkVisibility() != "public" ||
		f2.Visibility != "public" {
		t.Error("invalid")
	}
	g2, ok := f2.Folder()
	if !ok || g2.Tag != "folder" {
		t.Error("invalid")
	}
	if g2.SharedLinkId() != f2.SharedLinkId() {
		t.Error("invalid")
	}
}
