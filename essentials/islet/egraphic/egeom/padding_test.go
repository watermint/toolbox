package egeom

import (
	"testing"
)

func TestNewPaddingNone(t *testing.T) {
	p := NewPaddingNone()
	if d := p.Define(NewRectangle(ZeroPoint, 100, 100)); d.X() != 0 || d.Y() != 0 {
		t.Error(d)
	}
}

func TestNewPaddingFixed(t *testing.T) {
	p := NewPaddingFixed(12, 34)
	if d := p.Define(NewRectangle(ZeroPoint, 100, 100)); d.X() != 12 || d.Y() != 34 {
		t.Error(d)
	}
}

func TestNewPaddingRatio(t *testing.T) {
	p := NewPaddingRatio(0.5, 1.0)
	if d := p.Define(NewRectangle(ZeroPoint, 100, 100)); d.X() != 50 || d.Y() != 100 {
		t.Error(d)
	}
}
