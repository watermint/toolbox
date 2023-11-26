package eg_geom

import (
	"testing"
)

func TestParseLocatePosition(t *testing.T) {
	artboard := NewRectangle(ZeroPoint, 640, 400)
	obj := NewRectangle(ZeroPoint, 32, 20)
	pad := NewPaddingFixed(1, 3)

	{
		pp, oc := ParsePosition("top-left")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("top-center")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("top-right")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("center-left")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("center")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("center-right")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("bottom-left")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("bottom-center")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("bottom-right")
		if oc.IsError() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, oc := ParsePosition("undefined")
		if oc.IsOk() {
			t.Error(oc)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 0 || p.Y() != 0 {
			t.Error(p)
		}
	}
}
