package eg_draw

import (
	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
	eg_image2 "github.com/watermint/toolbox/essentials/graphic/eg_image"
	eg_text2 "github.com/watermint/toolbox/essentials/graphic/eg_text"
	"golang.org/x/image/font/gofont/gomedium"
	"testing"
)

func TestNewImageDrawer(t *testing.T) {
	text := "Hello\nyour go World"
	fill, _ := eg_color.ParseColor("marker(b18)")
	textFill, _ := eg_color.ParseColor("marker(w00)")

	img := eg_image2.NewRgba(640, 400)
	d := NewImageDrawer(img)
	d.FillRectangle(img.Bounds(), fill)

	font := eg_text2.MustNewTrueTypeParse(gomedium.TTF)
	s := eg_text2.NewStyle(font, textFill)
	pos := eg_geom2.PositionTopRight.Locate(img.Bounds(), s.Bound(text), eg_geom2.NewPaddingFixed(20, 20))

	d.DrawString(pos, text, s)
	img.ExportTo(eg_image2.FormatPng, "/tmp/test_image_draw.png")
}
