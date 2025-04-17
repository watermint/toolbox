package eg_image

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrUnsupportedFormatError = fmt.Errorf("%s", ErrUnsupportedFormat)
	ErrEncodeFailureError     = fmt.Errorf("%s", ErrEncodeFailure)
	ErrWriteFailureError      = fmt.Errorf("%s", ErrWriteFailure)
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

func (z rgbaImpl) ExportTo(format ImageFormat, path string) error {
	l := esl.Default()
	f, err := os.Create(path)
	if err != nil {
		l.Debug("Failed to create file", esl.String("path", path), esl.Error(err))
		return ErrWriteFailureError
	}
	defer func() {
		_ = f.Close()
	}()

	switch format {
	case FormatPng:
		if encErr := png.Encode(f, z.GoImageRGBA()); encErr != nil {
			l.Debug("Failed to encode PNG", esl.String("path", path), esl.Error(encErr))
			return ErrEncodeFailureError
		}
	default:
		l.Debug("Unsupported format", esl.String("path", path), esl.Int("format", int(format)))
		return ErrUnsupportedFormatError
	}
	return nil
}

func (z rgbaImpl) Bounds() eg_geom2.Rectangle {
	return eg_geom2.NewRectangleImage(z.img.Bounds())
}

func (z rgbaImpl) GetPixel(p eg_geom2.Point) eg_color.Color {
	return eg_color.NewColor(z.img.At(p.X(), p.Y()))
}

func (z rgbaImpl) SetPixel(p eg_geom2.Point, c eg_color.Color) {
	z.img.Set(p.X(), p.Y(), c)
}

func (z rgbaImpl) GoImageRGBA() *image.RGBA {
	return z.img
}
