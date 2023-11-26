package eg_draw

import (
	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
	"github.com/watermint/toolbox/essentials/graphic/eg_image"
	"github.com/watermint/toolbox/essentials/graphic/eg_text"
	"golang.org/x/image/font"
	"image"
	"image/draw"
)

type Draw interface {
	// FillRectangle a rectangle with the color
	FillRectangle(rect eg_geom2.Rectangle, color eg_color.Color)

	DrawString(pos eg_geom2.Point, text string, style eg_text.Style)
}

func NewImageDrawer(img eg_image.Image) Draw {
	return &drawImpl{
		img: img,
	}
}

type drawImpl struct {
	img eg_image.Image
}

func (z drawImpl) FillRectangle(rect eg_geom2.Rectangle, color eg_color.Color) {
	draw.Draw(z.img.GoImageRGBA(), rect.ImageRect(), image.NewUniform(color), rect.TopLeft().ImagePoint(), draw.Src)
}

func (z drawImpl) DrawString(pos eg_geom2.Point, text string, style eg_text.Style) {
	dr := &font.Drawer{
		Dst:  z.img.GoImageRGBA(),
		Src:  image.NewUniform(style.Color()),
		Face: style.Font().Face(),
	}
	style.Layout(text, pos, func(s string, p eg_geom2.Point) {
		dr.Dot = p.Add(0, dr.Face.Metrics().Height.Round()).Fixed26()
		dr.DrawString(s)
	})
}
