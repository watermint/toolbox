package eg_geom

import (
	"testing"
)

func TestParseLocatePosition(t *testing.T) {
	artboard := NewRectangle(ZeroPoint, 640, 400)
	obj := NewRectangle(ZeroPoint, 32, 20)
	pad := NewPaddingFixed(1, 3)

	{
		pp, err := ParsePosition("top-left")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("top-center")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("top-right")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 3 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("center-left")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("center")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("center-right")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 190 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("bottom-left")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 1 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("bottom-center")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 304 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("bottom-right")
		if err != nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 607 || p.Y() != 377 {
			t.Error(p)
		}
	}

	{
		pp, err := ParsePosition("undefined")
		if err == nil {
			t.Error(err)
		}
		p := pp.Locate(artboard, obj, pad)
		if p.X() != 0 || p.Y() != 0 {
			t.Error(p)
		}
	}
}
