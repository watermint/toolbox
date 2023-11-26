package edraw

import (
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/egraphic/eimage"
	"github.com/watermint/toolbox/essentials/islet/egraphic/etext"
	"golang.org/x/image/font/gofont/gomedium"
	"testing"
)

func TestNewImageDrawer(t *testing.T) {
	text := "Hello\nyour go World"
	fill, _ := ecolor.ParseColor("marker(b18)")
	textFill, _ := ecolor.ParseColor("marker(w00)")

	img := eimage.NewRgba(640, 400)
	d := NewImageDrawer(img)
	d.FillRectangle(img.Bounds(), fill)

	font := etext.MustNewTrueTypeParse(gomedium.TTF)
	s := etext.NewStyle(font, textFill)
	pos := egeom.PositionTopRight.Locate(img.Bounds(), s.Bound(text), egeom.NewPaddingFixed(20, 20))

	d.DrawString(pos, text, s)
	img.ExportTo(eimage.FormatPng, "/tmp/test_image_draw.png")
}
