package eg_placeholder

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
	"image"
	"image/draw"
)

type placeholderOpts struct {
	Fill eg_color.Color
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
	Fill     eg_color.Color
	FontSize int
	Position eg_geom2.Position
	Padding  eg_geom2.Padding
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
	es_idiom_deprecated.Outcome
}
