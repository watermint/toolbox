package dbx_sharing

import (
	"encoding/json"
	"github.com/watermint/toolbox/model/dbx_api"
	"testing"
)

func TestSharedLink(t *testing.T) {
	linkJson := `{
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

	link := SharedLink{}

	if dbx_api.ParseModelJsonForTest(&link, json.RawMessage(linkJson)) != nil {
		t.Error("failed to parse")
	}

	if link.Url != "https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0" {
		t.Error("invalid url")
	}
	if link.Name != "Prime_Numbers.txt" {
		t.Error("invalid name")
	}
	if link.SharedLinkId != "id:a4ayc_80_OEAAAAAAAAAXw" {
		t.Error("invalid id")
	}
	if link.PermissionResolvedVisibility != "public" {
		t.Error("invalid permission")
	}

}
