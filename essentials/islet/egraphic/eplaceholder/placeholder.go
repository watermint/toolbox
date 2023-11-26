package eplaceholder

import (
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/eidiom"
	"image"
	"image/draw"
)

type placeholderOpts struct {
	Fill ecolor.Color
}

func (z placeholderOpts) Apply(opts []PlaceholderOpt) placeholderOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type placeholderText struct {
	Text     string
	Fill     ecolor.Color
	FontSize int
	Position egeom.Position
	Padding  egeom.Padding
}

type PlaceholderOpt func(o placeholderOpts) placeholderOpts

func NewPlaceholder(width, height int, opts ...PlaceholderOpt) {
	po := placeholderOpts{}.Apply(opts)
	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)

	if po.Fill != nil {
		draw.Draw(img, img.Bounds(), image.NewUniform(po.Fill), image.Point{}, draw.Src)
	}

}

type PlaceholderOutcome interface {
	eidiom.Outcome
}
