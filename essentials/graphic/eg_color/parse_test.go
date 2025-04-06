package eg_color

import (
	"testing"

	"github.com/watermint/toolbox/essentials/go/es_errors"
)

func TestParseColor(t *testing.T) {
	c0 := NewRgba(0, 0, 0, 255)
	c1 := NewRgba(255, 255, 255, 255)
	c2 := NewRgba(16+1, 32+2, 48+3, 255)
	c3 := NewRgba(16+1, 32+2, 48+3, 64+4)
	c4 := NewRgba(1, 2, 3, 4)

	// R-G-B
	{
		c, err := ParseColor("000")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("fff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("123")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}

	// R-G-B-A
	{
		c, err := ParseColor("000f")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("ffff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("123f")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("1234")
		if err != nil || !c.Equals(c3) {
			t.Error(c, err)
		}
	}

	// RR-GG-BB
	{
		c, err := ParseColor("000")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("fff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("123")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}

	// RR-GG-BB-AA
	{
		c, err := ParseColor("000000ff")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("ffffffff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("112233ff")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("11223344")
		if err != nil || !c.Equals(c3) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("01020304")
		if err != nil || !c.Equals(c4) {
			t.Error(c, err)
		}
	}

	// ------ With Sharp like #fed, #ffeedd

	// R-G-B
	{
		c, err := ParseColor("#000")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#fff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#123")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}

	// R-G-B-A
	{
		c, err := ParseColor("#000f")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#ffff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#123f")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#1234")
		if err != nil || !c.Equals(c3) {
			t.Error(c, err)
		}
	}

	// RR-GG-BB
	{
		c, err := ParseColor("#000")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#fff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#123")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}

	// RR-GG-BB-AA
	{
		c, err := ParseColor("#000000ff")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#ffffffff")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#112233ff")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#11223344")
		if err != nil || !c.Equals(c3) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("#01020304")
		if err != nil || !c.Equals(c4) {
			t.Error(c, err)
		}
	}

	// rgb(R, G, B)
	{
		c, err := ParseColor("rgb(0, 0, 0)")
		if err != nil || !c.Equals(c0) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("rgb( 255, 255, 255 )")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("rgb(  17, 34,51)")
		if err != nil || !c.Equals(c2) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("rgb(1, 2, 256)")
		if err == nil || !es_errors.IsOutOfRangeError(err) {
			t.Error(c, err)
		}
	}

	// ----- CSS Colors
	{
		for n, cc := range cssColors {
			c, err := ParseColor(n)
			if err != nil || !c.Equals(cc) {
				t.Error(n, cc, c, err)
			}
		}
	}

	// ----- Marker colors

	{
		c, err := ParseColor("marker(0)")
		if err != nil || !c.Equals(c1) {
			t.Error(c, err)
		}
	}
	{
		c, err := ParseColor("marker(bv01)")
		bv01, _ := ParseColor("C5C9E6")
		if err != nil || !c.Equals(bv01) {
			t.Error(c, err, bv01)
		}
	}
	{
		c, err := ParseColor("marker(z01)")
		if err == nil || !es_errors.IsInvalidFormatError(err) {
			t.Error(c, err)
		}
	}
	{
		for n, cc := range markerColors {
			cl := "marker( " + n + " )"
			c, err := ParseColor(cl)
			expect, err2 := ParseColor(cc)
			if err2 != nil {
				t.Error(cc, err2)
			}
			if err != nil || !c.Equals(expect) {
				t.Error(cl, c, err, expect)
			}
		}
	}

	// ------ x11 color
	{
		for n, cc := range xColor {
			c, err := ParseColor("  x11( " + n + "  ) ")
			if err != nil || !c.Equals(cc) {
				t.Error(n, c, cc, err)
			}
		}
	}

	// invalid format
	{
		c, err := ParseColor("xxx")
		if err == nil {
			t.Error(c, err)
		}
	}
}
