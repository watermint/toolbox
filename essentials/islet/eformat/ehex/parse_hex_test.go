package ehex

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	if x, out := Parse("0001020304"); !reflect.DeepEqual(x, []byte{0x00, 0x01, 0x02, 0x03, 0x04}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}
	if x, out := Parse("123456789abcdef0"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}
	if x, out := Parse("123456789ABCDEF0"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}

	if x, out := Parse("０００１０２０３０４"); !reflect.DeepEqual(x, []byte{0x00, 0x01, 0x02, 0x03, 0x04}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}
	if x, out := Parse("１２３４５６７８９ａｂｃｄｅｆ０"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}
	if x, out := Parse("１２３４５６７８９ＡＢＣＤＥＦ０"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}

	if x, out := Parse(""); len(x) != 0 || out.IsError() {
		t.Error(x, ToHexString(x), out)
	}

	// odd number of chars
	if x, out := Parse("0"); x != nil || !out.IsInvalidFormat() {
		t.Error(x, ToHexString(x), out)
	}
	if x, out := Parse("0ab"); x != nil || !out.IsInvalidFormat() {
		t.Error(x, ToHexString(x), out)
	}

	// invalid char
	if x, out := Parse("0-ab"); x != nil || !out.IsInvalidChar() {
		t.Error(x, ToHexString(x), out)
	}
}
