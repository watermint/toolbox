package eg_color

import (
	"testing"
)

func TestParseColor(t *testing.T) {
	c0 := NewRgba(0, 0, 0, 255)
	c1 := NewRgba(255, 255, 255, 255)
	c2 := NewRgba(16+1, 32+2, 48+3, 255)
	c3 := NewRgba(16+1, 32+2, 48+3, 64+4)
	c4 := NewRgba(1, 2, 3, 4)

	// R-G-B
	{
		c, oc := ParseColor("000")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("fff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("123")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}

	// R-G-B-A
	{
		c, oc := ParseColor("000f")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("ffff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("123f")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("1234")
		if oc.IsError() || !c.Equals(c3) {
			t.Error(c, oc)
		}
	}

	// RR-GG-BB
	{
		c, oc := ParseColor("000")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("fff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("123")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}

	// RR-GG-BB-AA
	{
		c, oc := ParseColor("000000ff")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("ffffffff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("112233ff")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("11223344")
		if oc.IsError() || !c.Equals(c3) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("01020304")
		if oc.IsError() || !c.Equals(c4) {
			t.Error(c, oc)
		}
	}

	// ------ With Sharp like #fed, #ffeedd

	// R-G-B
	{
		c, oc := ParseColor("#000")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#fff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#123")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}

	// R-G-B-A
	{
		c, oc := ParseColor("#000f")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#ffff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#123f")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#1234")
		if oc.IsError() || !c.Equals(c3) {
			t.Error(c, oc)
		}
	}

	// RR-GG-BB
	{
		c, oc := ParseColor("#000")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#fff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#123")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}

	// RR-GG-BB-AA
	{
		c, oc := ParseColor("#000000ff")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#ffffffff")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#112233ff")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#11223344")
		if oc.IsError() || !c.Equals(c3) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("#01020304")
		if oc.IsError() || !c.Equals(c4) {
			t.Error(c, oc)
		}
	}

	// rgb(R, G, B)
	{
		c, oc := ParseColor("rgb(0, 0, 0)")
		if oc.IsError() || !c.Equals(c0) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("rgb( 255, 255, 255 )")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("rgb(  17, 34,51)")
		if oc.IsError() || !c.Equals(c2) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("rgb(1, 2, 256)")
		if oc.IsOk() || !oc.IsInvalidFormat() {
			t.Error(c, oc)
		}
	}

	// ----- CSS Colors
	{
		for n, cc := range cssColors {
			c, oc := ParseColor(n)
			if oc.IsError() || !c.Equals(cc) {
				t.Error(n, cc, c, oc)
			}
		}
	}

	// ----- Marker colors

	{
		c, oc := ParseColor("marker(0)")
		if oc.IsError() || !c.Equals(c1) {
			t.Error(c, oc)
		}
	}
	{
		c, oc := ParseColor("marker(bv01)")
		bv01, _ := ParseColor("C5C9E6")
		if oc.IsError() || !c.Equals(bv01) {
			t.Error(c, oc, bv01)
		}
	}
	{
		c, oc := ParseColor("marker(z01)")
		if oc.IsOk() || !oc.IsInvalidFormat() {
			t.Error(c, oc)
		}
	}
	{
		for n, cc := range markerColors {
			cl := "marker( " + n + " )"
			c, oc := ParseColor(cl)
			expect, oc2 := ParseColor(cc)
			if oc2.IsError() {
				t.Error(cc, oc2)
			}
			if oc.IsError() || !c.Equals(expect) {
				t.Error(cl, c, oc, expect)
			}
		}
	}

	// ------ x11 color
	{
		for n, cc := range xColor {
			c, oc := ParseColor("  x11( " + n + "  ) ")
			if oc.IsError() || !c.Equals(cc) {
				t.Error(n, c, cc, oc)
			}
		}
	}

	// invalid format
	{
		c, oc := ParseColor("xxx")
		if oc.IsOk() {
			t.Error(c, oc)
		}
	}
}
