package eg_image

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
	"image"
)

type ImageFormat int

const (
	FormatPng ImageFormat = iota
	FormatJpeg
)

// Image is a mutable instance
type Image interface {
	Bounds() eg_geom2.Rectangle

	GetPixel(p eg_geom2.Point) eg_color.Color

	// SetPixel update the pixel. Nothing happens if `p` is out of bounds.
	SetPixel(p eg_geom2.Point, c eg_color.Color)

	// GoImageRGBA returns the image instance of image.RGBA
	GoImageRGBA() *image.RGBA

	// Deprecated: ExportTo exports file to the file. This func will change method signature in the future
	ExportTo(format ImageFormat, path string) ExportOutcome
}

type ExportOutcome interface {
	es_idiom_deprecated.Outcome

	IsUnsupportedFormat() bool
	IsEncodeFailure() bool
	IsWriteFailure() bool
}
