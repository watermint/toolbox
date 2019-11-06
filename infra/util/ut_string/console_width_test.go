package ut_string

import "testing"

func TestWidth(t *testing.T) {
	if Width("ABC") != 3 {
		t.Error("invalid")
	}
	if Width("あいう") != 6 {
		t.Error("invalid")
	}
	if Width("あa") != 3 {
		t.Error("invalid")
	}
	if Width("道具a") != 5 {
		t.Error("invalid")
	}
	if Width("ウォーター・ミント") != 18 {
		t.Error("invalid")
	}
	if Width("バージョン５５") != 14 {
		t.Error("invalid")
	}
	if Width("バージョン55") != 12 {
		t.Error("invalid")
	}
}
