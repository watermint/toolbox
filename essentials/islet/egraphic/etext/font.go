package etext

import (
	"github.com/golang/freetype/truetype"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"
	"golang.org/x/image/font"
)

const (
	DefaultFontSize = 12
)

type Font interface {
	WithSize(size int) Font

	BoundString(text string) (bound egeom.Rectangle, advance int)

	// Size returns vertical font size in pixel.
	Size() int

	Face() font.Face
}

func MustNewTrueTypeParse(fontData []byte) Font {
	f, oc := NewTrueTypeParse(fontData)
	if oc.IsError() {
		panic(oc.String())
	}
	return f
}

func NewTrueTypeParse(fontData []byte) (f Font, oc eoutcome.ParseOutcome) {
	ttf, err := truetype.Parse(fontData)
	if err != nil {
		return nil, eoutcome.NewParseInvalidFormat(err.Error())
	}
	return NewTrueType(ttf), eoutcome.NewParseSuccess()
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

func (z ttfImpl) BoundString(text string) (bound egeom.Rectangle, advance int) {
	d := &font.Drawer{
		Face: z.Face(),
	}
	b26, a26 := d.BoundString(text)
	return egeom.NewRectangleFixed26(b26), a26.Round()
}

func (z ttfImpl) Face() font.Face {
	return truetype.NewFace(
		z.ttf,
		&truetype.Options{
			Size: float64(z.size),
		},
	)
}
