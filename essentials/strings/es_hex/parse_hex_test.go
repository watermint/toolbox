package es_hex

import (
	"reflect"
	"testing"

	"github.com/watermint/toolbox/essentials/go/es_errors"
)

func TestParse(t *testing.T) {
	if x, err := Parse("0001020304"); !reflect.DeepEqual(x, []byte{0x00, 0x01, 0x02, 0x03, 0x04}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}
	if x, err := Parse("123456789abcdef0"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}
	if x, err := Parse("123456789ABCDEF0"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}

	if x, err := Parse("０００１０２０３０４"); !reflect.DeepEqual(x, []byte{0x00, 0x01, 0x02, 0x03, 0x04}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}
	if x, err := Parse("１２３４５６７８９ａｂｃｄｅｆ０"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}
	if x, err := Parse("１２３４５６７８９ＡＢＣＤＥＦ０"); !reflect.DeepEqual(x, []byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}) || err != nil {
		t.Error(x, ToHexString(x), err)
	}

	if x, err := Parse(""); len(x) != 0 || err != nil {
		t.Error(x, ToHexString(x), err)
	}

	// odd number of chars
	if x, err := Parse("0"); x != nil || !es_errors.IsInvalidFormatError(err) {
		t.Error(x, ToHexString(x), err)
	}
	if x, err := Parse("0ab"); x != nil || !es_errors.IsInvalidFormatError(err) {
		t.Error(x, ToHexString(x), err)
	}

	// invalid char
	if x, err := Parse("0-ab"); x != nil || !es_errors.IsInvalidFormatError(err) {
		t.Error(x, ToHexString(x), err)
	}
}
