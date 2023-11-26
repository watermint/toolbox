package eimage

import (
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"image"
	"image/png"
	"os"
)

func NewRgba(width, height int) Image {
	return &rgbaImpl{
		img: image.NewRGBA(
			image.Rect(0, 0, width, height),
		),
	}
}

type rgbaImpl struct {
	img *image.RGBA
}

func (z rgbaImpl) ExportTo(format ImageFormat, path string) ExportOutcome {
	f, err := os.Create(path)
	if err != nil {
		return NewExportOutcomeWriteFailure(err)
	}
	defer func() {
		_ = f.Close()
	}()

	switch format {
	case FormatPng:
		if encErr := png.Encode(f, z.GoImageRGBA()); encErr != nil {
			return NewExportOutcomeEncodeFailure(encErr)
		}
	default:
		return NewExportOutcomeUnsupportedFormat(format)

	}
	return NewExportOutcomeSuccess()
}

func (z rgbaImpl) Bounds() egeom.Rectangle {
	return egeom.NewRectangleImage(z.img.Bounds())
}

func (z rgbaImpl) GetPixel(p egeom.Point) ecolor.Color {
	return ecolor.NewColor(z.img.At(p.X(), p.Y()))
}

func (z rgbaImpl) SetPixel(p egeom.Point, c ecolor.Color) {
	z.img.Set(p.X(), p.Y(), c)
}

func (z rgbaImpl) GoImageRGBA() *image.RGBA {
	return z.img
}
