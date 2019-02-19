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
	linkJson = `{".tag": "file", "url": "https://www.dropbox.com/s/4gdqs29ysesj7m2/%E8%A9%B3%E7%B4%B0%E3%83%AD%E3%82%B0%202017-07-03%2C2.csv?dl=0", "id": "id:CgFrhSgVOXAAAAAAAAA8vA", "name": "\u8a73\u7d30\u30ed\u30b0 2017-07-03,2.csv", "expires": "2020-02-19T05:24:36Z", "path_lower": "/dropbox business \u30ec\u30dd\u30fc\u30c8/\u8a73\u7d30\u30ed\u30b0 2017-07-03,2.csv", "link_permissions": {"resolved_visibility": {".tag": "public"}, "requested_visibility": {".tag": "public"}, "can_revoke": true, "visibility_policies": [{"policy": {".tag": "public"}, "resolved_policy": {".tag": "public"}, "allowed": true}, {"policy": {".tag": "team_only"}, "resolved_policy": {".tag": "team_only"}, "allowed": true}, {"policy": {".tag": "password"}, "resolved_policy": {".tag": "password"}, "allowed": true}], "can_set_expiry": true, "can_remove_expiry": true, "allow_download": false, "can_allow_download": true, "can_disallow_download": true, "allow_comments": true, "team_restricts_comments": false}, "team_member_info": {"team_info": {"id": "dbtid:AAAbdpIK3rMwt8tzhYbrA4VzZ1Y1zflXfK0", "name": "DeveloperInc"}, "display_name": "Dropbox\u30c7\u30e2", "member_id": "dbmid:AACeBMtPhHPCQ7-n44DZYq4-FjXM-xeO2ds"}, "preview_type": "excel", "client_modified": "2017-07-03T01:35:39Z", "server_modified": "2019-02-19T05:20:45Z", "rev": "18d053f457393", "size": 727}`

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
