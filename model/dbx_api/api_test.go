package dbx_api

import (
	"github.com/tidwall/gjson"
	"testing"
	"time"
)

func TestRebaseTimeForAPI(t *testing.T) {
	jst, err := time.LoadLocation("Japan")
	if err != nil {
		t.Error(err)
	}
	nowUtc := time.Now()
	nowJst := nowUtc.In(jst)
	nowRoundedUtc := nowUtc.Round(time.Second)

	if !RebaseTimeForAPI(nowJst).Equal(nowRoundedUtc) {
		t.Error("Invalid state")
	}
}

func TestParseError(t *testing.T) {
	errResponse := `{"error_summary": "not_a_member/..", "error": {".tag": "not_a_member"}, "user_message": {"locale": "ja", "text": "\u3053\u306e\u5171\u6709\u30d5\u30a9\u30eb\u30c0\u306e\u30e1\u30f3\u30d0\u30fc\u3067\u306f\u3042\u308a\u307e\u305b\u3093\u3002"}}`

	ae := ParseApiError(errResponse)
	if ae.ErrorTag != "not_a_member" {
		t.Error("ParseModel failed")
	}
	if ae.ErrorSummary != "not_a_member/.." {
		t.Error("ParseModel failed")
	}
	if ae.UserMessageLocale != "ja" {
		t.Error("ParseModel failed")
	}

	errResponse = `{"error_summary": "not_a_member/..", "error": {".tag": "not_a_member"}}`
	ae = ParseApiError(errResponse)
	if ae.ErrorTag != "not_a_member" {
		t.Error("ParseModel failed")
	}
	if ae.ErrorSummary != "not_a_member/.." {
		t.Error("ParseModel failed")
	}
	if ae.UserMessageLocale != "" {
		t.Error("ParseModel failed")
	}
	if ae.UserMessage != "" {
		t.Error("ParseModel failed")
	}

	errResponse = `{
    "error_summary": "bad_member/invalid_dropbox_id/...",
    "error": {
        ".tag": "bad_member",
        "bad_member": {
            ".tag": "invalid_dropbox_id",
            "invalid_dropbox_id": "dbid:AAEufNrMPSPe0dMQijRP0N_aZtBJRm26W4Q"
        }
    }
}`
	ae = ParseApiError(errResponse)

	if ae.ErrorTag != "bad_member" {
		t.Error("ParseModel failed")
	}
	if ae.ErrorSummary != "bad_member/invalid_dropbox_id/..." {
		t.Error("ParseModel failed")
	}
	if gjson.Get(string(ae.ErrorBody), "bad_member.invalid_dropbox_id").String() != "dbid:AAEufNrMPSPe0dMQijRP0N_aZtBJRm26W4Q" {
		t.Error("ParseModel failed")
	}

}
