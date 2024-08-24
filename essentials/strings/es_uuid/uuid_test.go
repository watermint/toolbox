package es_uuid

import (
	"testing"
	"time"
)

func TestNewV4(t *testing.T) {
	uv4 := NewV4()
	if x := uv4.String(); !IsUUID(x) {
		t.Error(x)
	}
}

func TestIsUUID(t *testing.T) {
	if x := IsUUID("00010203-0405-0607-0809-0a0b0c0d0e0f"); !x {
		t.Error(x)
	}
	if x := IsUUID("00010203=0405=0607=0809=0a0b0c0d0e0f"); x {
		t.Error(x)
	}
}

func TestParse(t *testing.T) {
	{
		v := "00010203-0405-0607-0809-0a0b0c0d0e0f"
		if x, out := Parse(v); x == nil || x.String() != v || out.IsError() {
			t.Error(x, out)
		}
	}

	{
		for i := 0; i < 10; i++ {
			v := NewV4()
			if x, out := Parse(v.String()); x == nil || x.String() != v.String() || out.IsError() {
				t.Error(x, out)
			}
		}
	}

	// invalid format
	{
		v := "00010203-0405-06070809-0a0b0c0d0e0f"
		if x, out := Parse(v); x != nil || !out.IsInvalidFormat() {
			t.Error(x, out)
		}
	}
}

func TestUuidData_IsNil(t *testing.T) {
	un := &uuidData{u: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
	if x := un.IsNil(); !x {
		t.Error(x)
	}

	uv4 := NewV4()
	if x := uv4.IsNil(); x {
		t.Error(x)
	}

	u := &uuidData{u: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}}
	if x := u.IsNil(); x {
		t.Error(x)
	}
}

func TestUuidData_Equals(t *testing.T) {
	uv4 := NewV4()
	uv1, out := Parse("81451528-00e5-11ec-9a03-0242ac130003")
	if out.IsError() {
		t.Error(out)
		return
	}

	if x := uv1.Equals(uv1); !x {
		t.Error(uv1, x)
	}
	if x := uv4.Equals(uv4); !x {
		t.Error(uv4, x)
	}
	if x := uv1.Equals(uv4); x {
		t.Error(uv1, uv4, x)
	}
	if x := uv4.Equals(uv1); x {
		t.Error(uv1, uv4, x)
	}
}

func TestUuidData_Version(t *testing.T) {
	{
		uv4 := NewV4()
		if x := uv4.Version(); x != Version4 {
			t.Error(x)
		}
	}

	{
		uv1, out := Parse("81451528-00e5-11ec-9a03-0242ac130003")
		if out.IsError() {
			t.Error(out)
		}
		if x := uv1.Version(); x != Version1 {
			t.Error(x)
		}
	}
}

func TestUuidData_Variant(t *testing.T) {
	{
		uv4 := NewV4()
		if x := uv4.Variant(); x != Variant1 {
			t.Error(x)
		}
	}
}

func TestUuidV4_String(t *testing.T) {
	{
		u := &uuidData{u: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}}
		if x := u.String(); x != "00010203-0405-0607-0809-0a0b0c0d0e0f" {
			t.Error(x)
		}
	}
	{
		u := &uuidData{u: []byte{0x71, 0x84, 0xc6, 0x68, 0xb4, 0xe7, 0x4e, 0xcd, 0xb0, 0xbb, 0xf7, 0xd9, 0x87, 0x9f, 0xd9, 0x13}}
		if x := u.String(); x != "7184c668-b4e7-4ecd-b0bb-f7d9879fd913" {
			t.Error(x)
		}
	}
}

func TestUuidV4_Urn(t *testing.T) {
	{
		u := &uuidData{u: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}}
		if x := u.Urn(); x != "urn:uuid:00010203-0405-0607-0809-0a0b0c0d0e0f" {
			t.Error(x)
		}
	}
	{
		u := &uuidData{u: []byte{0x71, 0x84, 0xc6, 0x68, 0xb4, 0xe7, 0x4e, 0xcd, 0xb0, 0xbb, 0xf7, 0xd9, 0x87, 0x9f, 0xd9, 0x13}}
		if x := u.Urn(); x != "urn:uuid:7184c668-b4e7-4ecd-b0bb-f7d9879fd913" {
			t.Error(x)
		}
	}
}

func TestUuidV7(t *testing.T) {
	u := NewV7()
	if x := u.String(); !IsUUID(x) {
		t.Error(x)
	} else {
		t.Log(x)
	}
	if u.Version() != Version7 {
		t.Error(u.Version())
	}
	if u.Variant() != Variant1 {
		t.Error(u.Variant())
	}
}

func TestUuidV7Timestamp(t *testing.T) {
	ts := time.Now()
	u := NewV7WithTimestamp(time.Now())
	if x := u.String(); !IsUUID(x) {
		t.Error(x)
	} else {
		t.Log(x)
	}
	if u.Version() != Version7 {
		t.Error(u.Version())
	}
	if u.Variant() != Variant1 {
		t.Error(u.Variant())
	}
	if x, err := TimestampFromUUIDV7(u); err != nil {
		t.Error(err)
	} else if x.UnixNano() != (ts.UnixNano()/1_000_000)*1_000_000 {
		t.Error(x.UnixNano(), ts.UnixNano())
	} else {
		t.Log(x)
	}
}
