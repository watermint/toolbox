package etext

import (
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/emath"
	"strings"
)

type Alignment int

const (
	AlignLeft = iota
	AlignCenter
	AlignRight
)

const (
	defaultLineHeight = 1.2
	defaultAlign      = AlignLeft
)

type Style interface {
	WithLineHeight(h float64) Style
	WithAlignment(a Alignment) Style
	WithFont(f Font) Style
	WithColor(c ecolor.Color) Style

	// Alignment is for specify text alignment
	Alignment() Alignment

	// LineHeight is a multiplier for font height. If you set 2.0 that means one empty line between the lines.
	LineHeight() float64

	// Font returns font setting for this style
	Font() Font

	// Color returns color setting for this style
	Color() ecolor.Color

	Layout(text string, start egeom.Point, f func(s string, p egeom.Point)) (bound egeom.Rectangle)

	Bound(text string) (bound egeom.Rectangle)
}

func NewStyle(f Font, c ecolor.Color) Style {
	return &styleImpl{
		font:       f,
		color:      c,
		lineHeight: defaultLineHeight,
		align:      defaultAlign,
	}
}

type styleImpl struct {
	align      Alignment
	lineHeight float64
	font       Font
	color      ecolor.Color
}

func (z styleImpl) WithColor(c ecolor.Color) Style {
	z.color = c
	return z
}

func (z styleImpl) Color() ecolor.Color {
	return z.color
}

func (z styleImpl) Font() Font {
	return z.font
}

func (z styleImpl) WithLineHeight(h float64) Style {
	z.lineHeight = h
	return z
}

func (z styleImpl) WithAlignment(a Alignment) Style {
	z.align = a
	return z
}

func (z styleImpl) WithFont(f Font) Style {
	z.font = f
	return z
}

func (z styleImpl) Alignment() Alignment {
	return z.align
}

func (z styleImpl) LineHeight() float64 {
	return z.lineHeight
}

func (z styleImpl) Layout(text string, start egeom.Point, f func(s string, p egeom.Point)) (bound egeom.Rectangle) {
	lines := strings.Split(text, "\n")
	numLines := len(lines)
	lineWidths := make([]int, numLines)

	for i := 0; i < numLines; i++ {
		bounds, _ := z.font.BoundString(lines[i])
		lineWidths[i] = bounds.Width()
	}
	lineHeight := int(z.lineHeight * float64(z.font.Size()))
	maxWidth := emath.MaxInt(lineWidths...)

	switch z.align {
	case AlignCenter:
		for i := 0; i < numLines; i++ {
			f(lines[i], egeom.NewPoint(start.X()+(maxWidth-lineWidths[i])/2, start.Y()+lineHeight*i))
		}

	case AlignRight:
		for i := 0; i < numLines; i++ {
			f(lines[i], egeom.NewPoint(start.X()+maxWidth-lineWidths[i], start.Y()+lineHeight*i))
		}

	default:
		// align left if unknown Alignment
		for i := 0; i < numLines; i++ {
			f(lines[i], egeom.NewPoint(start.X(), start.Y()+lineHeight*i))
		}
	}

	return egeom.NewRectangle(
		start,
		maxWidth,
		lineHeight*numLines,
	)
}

func (z styleImpl) Bound(text string) (bound egeom.Rectangle) {
	return z.Layout(text, egeom.NewPoint(0, 0), func(s string, p egeom.Point) {})
}
