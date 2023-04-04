package sv_file

import (
	"strings"
	"testing"
)

func TestVerifyFileKey(t *testing.T) {
	if r, _ := VerifyFileKey("1234"); r != VerifyFileKeyLooksOkay {
		t.Error(r)
	}
	if r, _ := VerifyFileKey("@#$%"); r != VerifyFileKeyInvalidChar {
		t.Error(r)
	}
	if r, _ := VerifyFileKey(strings.Repeat("x", FileKeyMaxLength+1)); r != VerifyFileKeyTooLong {
		t.Error(r)
	}
}
