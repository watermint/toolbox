package eg_image

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/essentials/graphic/eg_color"
	eg_geom2 "github.com/watermint/toolbox/essentials/graphic/eg_geom"
)

func TestNewRgba_Bounds(t *testing.T) {
	img := NewRgba(10, 20)
	b := img.Bounds()
	if b.Width() != 10 || b.Height() != 20 {
		t.Errorf("unexpected bounds: got %dx%d", b.Width(), b.Height())
	}
}

func TestSetGetPixel(t *testing.T) {
	img := NewRgba(5, 5)
	p := eg_geom2.NewPoint(2, 3)
	c := eg_color.NewRgba(10, 20, 30, 40)
	img.SetPixel(p, c)
	got := img.GetPixel(p)
	if !got.Equals(c) {
		t.Errorf("pixel mismatch: got %v, want %v", got, c)
	}
}

func TestGoImageRGBA(t *testing.T) {
	img := NewRgba(3, 4)
	rgba := img.GoImageRGBA()
	if rgba.Bounds().Dx() != 3 || rgba.Bounds().Dy() != 4 {
		t.Errorf("unexpected GoImageRGBA bounds: got %dx%d", rgba.Bounds().Dx(), rgba.Bounds().Dy())
	}
}

func TestExportTo_UnsupportedFormat(t *testing.T) {
	img := NewRgba(1, 1)
	dir := t.TempDir()
	path := dir + "/test.jpg"
	err := img.ExportTo(FormatJpeg, path)
	if err == nil {
		t.Error("expected error for unsupported format")
	}
}

func TestExportTo_Png(t *testing.T) {
	img := NewRgba(1, 1)
	p := eg_geom2.NewPoint(0, 0)
	img.SetPixel(p, eg_color.NewRgba(255, 0, 0, 255))
	dir := t.TempDir()
	path := dir + "/test_export.png"
	err := img.ExportTo(FormatPng, path)
	if err != nil {
		t.Errorf("ExportTo PNG failed: %v", err)
	}
	if _, statErr := os.Stat(path); statErr != nil {
		t.Errorf("PNG file not created: %v", statErr)
	}
}
