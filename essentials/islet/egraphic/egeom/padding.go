package egeom

type Padding interface {
	Define(obj Rectangle) Point
}

func NewPaddingNone() Padding {
	return &paddingNone{}
}

type paddingNone struct {
}

func (z paddingNone) Define(obj Rectangle) Point {
	return NewPoint(0, 0)
}

func NewPaddingFixed(x, y int) Padding {
	return &paddingFixed{
		x: x,
		y: y,
	}
}

type paddingFixed struct {
	x, y int
}

func (z paddingFixed) Define(obj Rectangle) Point {
	return NewPoint(z.x, z.y)
}

func NewPaddingRatio(rx, ry float64) Padding {
	return &paddingRatio{
		rx: rx,
		ry: ry,
	}
}

type paddingRatio struct {
	rx, ry float64
}

func (z paddingRatio) Define(obj Rectangle) Point {
	return NewPoint(int(float64(obj.Width())*z.rx), int(float64(obj.Height())*z.ry))
}
