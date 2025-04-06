package es_demo

import (
	"bytes"
	"image"
	"image/gif"
	"io"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

// Image terminal emulator.
type Terminal interface {
	io.Writer
}

type termImpl struct {
	width    int
	height   int
	fontSize int
}

func (z *termImpl) Write(p []byte) (n int, err error) {
	mono, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(mono, &truetype.Options{
		Size:              14,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	})

	plane := image.NewRGBA(image.Rect(0, 0, 640, 480))
	d := &font.Drawer{
		Dst:  plane,
		Src:  image.White,
		Face: face,
		Dot:  fixed.Point26_6{},
	}

	d.Dot.X = fixed.I(0)
	d.Dot.Y = fixed.I(80)

	d.DrawString(string(p))
	boundary, advance := d.BoundBytes(p)

	l := esl.Default()
	l.Info("Boundary",
		esl.Int("length", len(p)),
		esl.Any("minX", boundary.Min.X),
		esl.Any("minY", boundary.Min.Y),
		esl.Any("maxX", boundary.Max.X),
		esl.Any("maxY", boundary.Max.Y),
		esl.Any("w", boundary.Max.X-boundary.Min.X),
		esl.Any("h", boundary.Max.Y-boundary.Min.Y),
		esl.Any("advance", advance))

	buf := &bytes.Buffer{}
	err = gif.Encode(buf, plane, &gif.Options{
		NumColors: 16,
	})
	if err != nil {
		panic(err)
	}

	dest, err := os.CreateTemp("", "dummy")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = dest.Close()
	}()
	if _, err := io.Copy(dest, buf); err != nil {
		panic(err)
	}

	return len(p), nil
}
