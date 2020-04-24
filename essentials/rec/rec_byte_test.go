package rec

import "testing"

func TestByteDigest(t *testing.T) {
	b := ByteDigest([]byte("github.com/watermint/toolbox"))
	if b["len"] != 28 {
		t.Error(b)
	}
}
