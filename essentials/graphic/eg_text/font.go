package eg_text

import (
	"github.com/golang/freetype/truetype"
	"github.com/watermint/toolbox/essentials/go/es_errors"
	"github.com/watermint/toolbox/essentials/graphic/eg_geom"
	"golang.org/x/image/font"
)

const (
	DefaultFontSize = 12
)

type Font interface {
	WithSize(size int) Font

	BoundString(text string) (bound eg_geom.Rectangle, advance int)

	// Size returns vertical font size in pixel.
	Size() int

	Face() font.Face
}

func MustNewTrueTypeParse(fontData []byte) Font {
	f, err := NewTrueTypeParse(fontData)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func NewTrueTypeParse(fontData []byte) (f Font, err error) {
	ttf, err := truetype.Parse(fontData)
	if err != nil {
		return nil, es_errors.NewInvalidFormatError("invalid font format: %s", err.Error())
	}
	return NewTrueType(ttf), nil
}

func NewTrueType(f *truetype.Font) Font {
	return &ttfImpl{
		ttf:  f,
		size: DefaultFontSize,
	}
}

type ttfImpl struct {
	ttf  *truetype.Font
	size int
}

func (z ttfImpl) Size() int {
	return z.size
}

func (z ttfImpl) WithSize(size int) Font {
	z.size = size
	return z
}

func (z ttfImpl) BoundString(text string) (bound eg_geom.Rectangle, advance int) {
	d := &font.Drawer{
		Face: z.Face(),
	}
	b26, a26 := d.BoundString(text)
	return eg_geom.NewRectangleFixed26(b26), a26.Round()
}

func (z ttfImpl) Face() font.Face {
	return truetype.NewFace(
		z.ttf,
		&truetype.Options{
			Size: float64(z.size),
		},
	)
}
