package dbx_device

import (
	"encoding/json"
	"testing"
)

func TestModelDevice(t *testing.T) {
	resListMembersDevices := `{
  "devices": [
    {
      "team_member_id": "dbmid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiQ",
      "web_sessions": [],
      "desktop_clients": [],
      "mobile_clients": []
    },
    {
      "team_member_id": "dbmid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAds",
      "web_sessions": [
        {
          "session_id": "dbwsid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA787",
          "ip_address": "xx.xx.xx.xx",
          "country": "Japan",
          "created": "2019-03-04T07:20:56Z",
          "updated": "2019-03-04T07:20:56Z",
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36",
          "os": "Mac OS X",
          "browser": "Chrome",
          "expires": "2019-04-03T07:20:56Z"
        },
        {
          "session_id": "dbwsid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA425",
          "ip_address": "xx.xx.xx.xx",
          "country": "Japan",
          "created": "2019-03-03T03:46:41Z",
          "updated": "2019-03-03T03:46:41Z",
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36",
          "os": "Mac OS X",
          "browser": "Chrome",
          "expires": "2019-04-02T03:46:41Z"
        }
      ],
      "desktop_clients": [
        {
          "session_id": "dbdsid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAArc",
          "ip_address": "xx.xxx.xx.xxx",
          "country": "United States",
          "created": "2017-12-25T05:28:05Z",
          "updated": "2017-12-27T01:57:31Z",
          "host_name": "xxxxxxxx",
          "client_type": {
            ".tag": "windows"
          },
          "client_version": "40.4.46",
          "platform": "Windows 7",
          "is_delete_on_unlink_supported": true
        },
        {
          "session_id": "dbdsid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAzk",
          "ip_address": "xx.xxx.xx.xxx",
          "country": "United States",
          "created": "2017-04-21T09:13:56Z",
          "updated": "2017-04-25T05:41:54Z",
          "host_name": "xxxxxxx",
          "client_type": {
            ".tag": "windows"
          },
          "client_version": "24.4.16",
          "platform": "Windows 10",
          "is_delete_on_unlink_supported": true
        }
      ],
      "mobile_clients": [
        {
          "session_id": "dbmsid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA-AAAAAAAAAIP",
          "ip_address": "xx.xxx.xx.xxx",
          "country": "Japan",
          "created": "2017-06-15T01:37:04Z",
          "updated": "2017-06-15T01:58:43Z",
          "device_name": "xxxxxxxxx.xxx xxxxxx xx xxxx",
          "client_type": {
            ".tag": "iphone"
          },
          "client_version": "52.2.2",
          "os_version": "10.3.2",
          "last_carrier": "ソフトバンク"
        }
      ]
    }
  ],
  "has_more": false
}`

	d := Devices{}
	err := json.Unmarshal([]byte(resListMembersDevices), &d)
	if err != nil {
		t.Error(err)
	}

	if d.Devices[0].TeamMemberId != "dbmid:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiQ" {
		t.Error("invalid")
	}

	m := d.MobileClients()
	if len(m) != 1 {
		t.Error("invalid")
	}

	w := d.WebSessions()
	if len(w) != 2 {
		t.Error("invalid")
	}

	x := d.DesktopClients()
	if len(x) != 2 {
		t.Error("invalid")
	}
}
