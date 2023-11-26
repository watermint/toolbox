package etext

import (
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/gofont/gomediumitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"golang.org/x/image/font/gofont/gosmallcapsitalic"
	"testing"
)

var (
	GoFontRegular         = MustNewTrueTypeParse(goregular.TTF)
	GoFontItalic          = MustNewTrueTypeParse(goitalic.TTF)
	GoFontBold            = MustNewTrueTypeParse(gobold.TTF)
	GoFontBoldItalic      = MustNewTrueTypeParse(gobolditalic.TTF)
	GoFontMedium          = MustNewTrueTypeParse(gomedium.TTF)
	GoFontMediumItalic    = MustNewTrueTypeParse(gomediumitalic.TTF)
	GoFontMono            = MustNewTrueTypeParse(gomono.TTF)
	GoFontMonoItalic      = MustNewTrueTypeParse(gomonoitalic.TTF)
	GoFontMonoBold        = MustNewTrueTypeParse(gomonobold.TTF)
	GoFontMonoBoldItalic  = MustNewTrueTypeParse(gomonobolditalic.TTF)
	GoFontSmallCaps       = MustNewTrueTypeParse(gosmallcaps.TTF)
	GoFontSmallCapsItalic = MustNewTrueTypeParse(gosmallcapsitalic.TTF)

	DefaultFonts = []Font{
		GoFontRegular,
		GoFontItalic,
		GoFontBold,
		GoFontBoldItalic,
		GoFontMedium,
		GoFontMediumItalic,
		GoFontMono,
		GoFontMonoItalic,
		GoFontMonoBold,
		GoFontMonoBoldItalic,
		GoFontSmallCaps,
		GoFontSmallCapsItalic,
	}
)

func TestNewTrueType(t *testing.T) {
	for _, df := range DefaultFonts {
		b, a := df.BoundString("Hello")
		if b.Width() < 1 || b.Height() < 1 || a < 1 {
			t.Error(b, a)
		}
	}
}
