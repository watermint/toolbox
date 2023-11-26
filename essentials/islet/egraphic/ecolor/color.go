package ecolor

import "image/color"

type Color interface {
	color.Color

	Equals(other Color) bool
}

func NewColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return NewRgba(uint8(r), uint8(g), uint8(b), uint8(a))
}

func NewRgb(r, g, b uint8) Color {
	return NewRgba(r, g, b, 255)
}

func NewRgba(r, g, b, a uint8) Color {
	return &rgbaImpl{
		rgba: color.RGBA{
			R: r,
			G: g,
			B: b,
			A: a,
		},
	}
}

type rgbaImpl struct {
	rgba color.RGBA
}

func (z rgbaImpl) RGBA() (r, g, b, a uint32) {
	return z.rgba.RGBA()
}

func (z rgbaImpl) Equals(other Color) bool {
	if other == nil {
		return false
	}
	zr, zg, zb, za := z.RGBA()
	or, og, ob, oa := other.RGBA()
	return zr == or && zg == og && zb == ob && za == oa
}
