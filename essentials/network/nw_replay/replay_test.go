package nw_replay

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestResponse_Http(t *testing.T) {
	rec := `{
    "code": 200,
    "proto": "HTTP/2.0",
    "body": "{\"name\": \"xxxxxxxxx xxx\", \"team_id\": \"dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\", \"num_licensed_users\": 10, \"num_provisioned_users\": 8, \"policies\": {\"sharing\": {\"shared_folder_member_policy\": {\".tag\": \"anyone\"}, \"shared_folder_join_policy\": {\".tag\": \"from_anyone\"}, \"shared_link_create_policy\": {\".tag\": \"team_only\"}}, \"emm_state\": {\".tag\": \"disabled\"}, \"office_addin\": {\".tag\": \"disabled\"}}}",
    "headers": {
      "Cache-Control": "no-cache",
      "Content-Type": "application/json",
      "Date": "Wed, 06 May 2020 15:03:47 GMT",
      "Pragma": "no-cache",
      "Server": "nginx",
      "Vary": "Accept-Encoding",
      "X-Content-Type-Options": "nosniff",
      "X-Dropbox-Request-Id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "X-Envoy-Upstream-Service-Time": "74",
      "X-Frame-Options": "SAMEORIGIN",
      "X-Server-Response-Time": "67"
    },
    "content_length": 392
  }`

	res := &Response{}
	if err := json.Unmarshal([]byte(rec), res); err != nil {
		t.Error(err)
	}
	if res.Code != 200 || res.Headers["Date"] != "Wed, 06 May 2020 15:03:47 GMT" {
		t.Error(res)
	}

	hr := res.Http()
	if hr.StatusCode != 200 || hr.Header.Get("Date") != "Wed, 06 May 2020 15:03:47 GMT" {
		t.Error(hr)
	}
}

func TestReplay_Call(t *testing.T) {
	rec := `
[
  {
    "code": 200,
    "proto": "HTTP/2.0",
    "body": "{\"name\": \"xxxxxxxxx xxx\", \"team_id\": \"dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\", \"num_licensed_users\": 10, \"num_provisioned_users\": 8, \"policies\": {\"sharing\": {\"shared_folder_member_policy\": {\".tag\": \"anyone\"}, \"shared_folder_join_policy\": {\".tag\": \"from_anyone\"}, \"shared_link_create_policy\": {\".tag\": \"team_only\"}}, \"emm_state\": {\".tag\": \"disabled\"}, \"office_addin\": {\".tag\": \"disabled\"}}}",
    "headers": {
      "Cache-Control": "no-cache",
      "Content-Type": "application/json",
      "Date": "Wed, 06 May 2020 15:03:47 GMT",
      "Pragma": "no-cache",
      "Server": "nginx",
      "Vary": "Accept-Encoding",
      "X-Content-Type-Options": "nosniff",
      "X-Dropbox-Request-Id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "X-Envoy-Upstream-Service-Time": "74",
      "X-Frame-Options": "SAMEORIGIN",
      "X-Server-Response-Time": "67"
    },
    "content_length": 392
  },
  {
    "code": 409,
    "headers": {
      "Content-Disposition": "attachment; filename='error'",
      "Content-Security-Policy": "sandbox; frame-ancestors 'none'",
      "Content-Type": "application/json",
      "Date": "Wed, 06 May 2020 15:03:02 GMT",
      "Server": "nginx",
      "X-Content-Type-Options": "nosniff",
      "X-Dropbox-Request-Id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "X-Envoy-Upstream-Service-Time": "118",
      "X-Frame-Options": "DENY"
    },
    "content_length": 0
  }
]
`

	var res []Response
	if err := json.Unmarshal([]byte(rec), &res); err != nil {
		t.Error(err)
	}
	rr := NewSequentialReplay(res)
	res1, _, err1 := rr.Call("", "", &http.Request{})
	if err1 != nil {
		t.Error(err1)
	}
	if res1.StatusCode != 200 || res1.Header.Get("Date") != "Wed, 06 May 2020 15:03:47 GMT" {
		t.Error(res1)
	}
	body1, err := io.ReadAll(res1.Body)
	if err != nil {
		t.Error(err)
	}
	if len(body1) != 392 {
		t.Error(err)
	}

	res2, _, err2 := rr.Call("", "", &http.Request{})
	if err2 != nil {
		t.Error(err2)
	}
	if res2.StatusCode != 409 || res2.Header.Get("Date") != "Wed, 06 May 2020 15:03:02 GMT" {
		t.Error(res2)
	}
	body2, err := io.ReadAll(res2.Body)
	if err != nil {
		t.Error(err)
	}
	if len(body2) != 0 {
		t.Error(err)
	}

	res3, _, err3 := rr.Call("", "", &http.Request{})
	if err3 == nil || res3 != nil {
		t.Error(res3, err3)
	}
}
