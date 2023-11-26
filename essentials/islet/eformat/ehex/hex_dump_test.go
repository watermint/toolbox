package ehex

import "testing"

func TestToHexString(t *testing.T) {
	if x := ToHexString([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}); x != "0123456789abcdef" {
		t.Error(x)
	}
	if x := ToHexString([]byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}); x != "fedcba9876543210" {
		t.Error(x)
	}
	if x := ToHexString([]byte{}); x != "" {
		t.Error(x)
	}
	if x := ToHexString(nil); x != "" {
		t.Error(x)
	}
}

func TestToUpperHexString(t *testing.T) {
	if x := ToUpperHexString([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}); x != "0123456789ABCDEF" {
		t.Error(x)
	}
	if x := ToUpperHexString([]byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}); x != "FEDCBA9876543210" {
		t.Error(x)
	}
	if x := ToUpperHexString([]byte{}); x != "" {
		t.Error(x)
	}
	if x := ToUpperHexString(nil); x != "" {
		t.Error(x)
	}
}
