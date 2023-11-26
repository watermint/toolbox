package eg_geom

import (
	"golang.org/x/image/math/fixed"
	"image"
)

type Rectangle interface {
	X() int
	Y() int
	Width() int
	Height() int

	Equals(q Rectangle) bool

	TopLeft() Point
	TopCenter() Point
	TopRight() Point
	CenterLeft() Point
	Center() Point
	CenterRight() Point
	BottomLeft() Point
	BottomCenter() Point
	BottomRight() Point

	Fixed26() fixed.Rectangle26_6
	Fixed52() fixed.Rectangle52_12
	ImageRect() image.Rectangle
}

func NewRectangle(location Point, width, height int) Rectangle {
	return rectImpl{
		location: location,
		w:        width,
		h:        height,
	}
}
func NewRectangleImage(r image.Rectangle) Rectangle {
	return rectImpl{
		location: NewPointImage(r.Min),
		w:        r.Dx(),
		h:        r.Dy(),
	}
}
func NewRectangleFixed26(q fixed.Rectangle26_6) Rectangle {
	dim := q.Max.Sub(q.Min)
	return rectImpl{
		location: NewPointFixed26(q.Min),
		w:        dim.X.Round(),
		h:        dim.Y.Round(),
	}
}
func NewRectangleFixed52(q fixed.Rectangle52_12) Rectangle {
	dim := q.Max.Sub(q.Min)
	return rectImpl{
		location: NewPointFixed52(q.Min),
		w:        dim.X.Round(),
		h:        dim.Y.Round(),
	}
}

type rectImpl struct {
	location Point
	w, h     int
}

func (z rectImpl) Equals(q Rectangle) bool {
	if q == nil {
		return false
	}
	return z.location.Equals(q.TopLeft()) && z.w == q.Width() && z.h == q.Height()
}

func (z rectImpl) TopLeft() Point {
	return z.location
}

func (z rectImpl) TopCenter() Point {
	return z.location.Add(z.w/2, 0)
}

func (z rectImpl) TopRight() Point {
	return z.location.Add(z.w, 0)
}

func (z rectImpl) CenterLeft() Point {
	return z.location.Add(0, z.h/2)
}

func (z rectImpl) Center() Point {
	return z.location.Add(z.w/2, z.h/2)
}

func (z rectImpl) CenterRight() Point {
	return z.location.Add(z.w, z.h/2)
}

func (z rectImpl) BottomLeft() Point {
	return z.location.Add(0, z.h)
}

func (z rectImpl) BottomCenter() Point {
	return z.location.Add(z.w/2, z.h)
}

func (z rectImpl) BottomRight() Point {
	return z.location.Add(z.w, z.h)
}

func (z rectImpl) X() int {
	return z.location.X()
}

func (z rectImpl) Y() int {
	return z.location.Y()
}

func (z rectImpl) Width() int {
	return z.w
}

func (z rectImpl) Height() int {
	return z.h
}

func (z rectImpl) Fixed26() fixed.Rectangle26_6 {
	return fixed.Rectangle26_6{
		Min: z.TopLeft().Fixed26(),
		Max: z.BottomRight().Fixed26(),
	}
}

func (z rectImpl) Fixed52() fixed.Rectangle52_12 {
	return fixed.Rectangle52_12{
		Min: z.TopLeft().Fixed52(),
		Max: z.BottomRight().Fixed52(),
	}
}

func (z rectImpl) ImageRect() image.Rectangle {
	br := z.BottomRight()
	return image.Rect(z.X(), z.Y(), br.X(), br.Y())
}
