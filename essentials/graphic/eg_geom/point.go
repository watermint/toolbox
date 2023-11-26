package eg_geom

import (
	"golang.org/x/image/math/fixed"
	"image"
)

type Point interface {
	X() int
	Y() int

	Add(x, y int) Point
	AddPoint(q Point) Point
	Sub(x, y int) Point
	SubPoint(q Point) Point

	Equals(q Point) bool

	Fixed26() fixed.Point26_6
	Fixed52() fixed.Point52_12
	ImagePoint() image.Point
}

var (
	ZeroPoint = NewPoint(0, 0)
)

func NewPoint(x, y int) Point {
	return &pointImpl{
		x: x,
		y: y,
	}
}
func NewPointImage(q image.Point) Point {
	return &pointImpl{
		x: q.X,
		y: q.Y,
	}
}
func NewPointFixed26(q fixed.Point26_6) Point {
	return &pointImpl{
		x: q.X.Round(),
		y: q.Y.Round(),
	}
}
func NewPointFixed52(q fixed.Point52_12) Point {
	return &pointImpl{
		x: q.X.Round(),
		y: q.Y.Round(),
	}
}

type pointImpl struct {
	x, y int
}

func (z pointImpl) Equals(q Point) bool {
	if q == nil {
		return false
	}
	return z.x == q.X() && z.y == q.Y()
}

func (z pointImpl) X() int {
	return z.x
}

func (z pointImpl) Y() int {
	return z.y
}

func (z pointImpl) Add(x, y int) Point {
	z.x += x
	z.y += y
	return z
}

func (z pointImpl) AddPoint(q Point) Point {
	z.x += q.X()
	z.y += q.Y()
	return z
}

func (z pointImpl) Sub(x, y int) Point {
	z.x -= x
	z.y -= y
	return z
}

func (z pointImpl) SubPoint(q Point) Point {
	z.x -= q.X()
	z.y -= q.Y()
	return z
}

func (z pointImpl) Fixed26() fixed.Point26_6 {
	return fixed.P(z.x, z.y)
}

func (z pointImpl) Fixed52() fixed.Point52_12 {
	return fixed.Point52_12{
		X: fixed.Int52_12(z.x),
		Y: fixed.Int52_12(z.y),
	}
}

func (z pointImpl) ImagePoint() image.Point {
	return image.Point{
		X: z.x,
		Y: z.y,
	}
}
