package eimage

import (
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/eidiom"
	"image"
)

type ImageFormat int

const (
	FormatPng ImageFormat = iota
	FormatJpeg
)

// Image is a mutable instance
type Image interface {
	Bounds() egeom.Rectangle

	GetPixel(p egeom.Point) ecolor.Color

	// SetPixel update the pixel. Nothing happens if `p` is out of bounds.
	SetPixel(p egeom.Point, c ecolor.Color)

	// GoImageRGBA returns the image instance of image.RGBA
	GoImageRGBA() *image.RGBA

	// Deprecated: ExportTo exports file to the file. This func will change method signature in the future
	ExportTo(format ImageFormat, path string) ExportOutcome
}

type ExportOutcome interface {
	eidiom.Outcome

	IsUnsupportedFormat() bool
	IsEncodeFailure() bool
	IsWriteFailure() bool
}
