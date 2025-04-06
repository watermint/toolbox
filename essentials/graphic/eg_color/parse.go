package eg_color

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/watermint/toolbox/essentials/go/es_errors"
	"github.com/watermint/toolbox/essentials/strings/es_hex"
)

var (
	regexHexRgb           = regexp.MustCompile(`^[0-9a-f]{3}$`)
	regexHexRgba          = regexp.MustCompile(`^[0-9a-f]{4}$`)
	regexHexRrGgBb        = regexp.MustCompile(`^[0-9a-f]{6}$`)
	regexHexRrGgBbAa      = regexp.MustCompile(`^[0-9a-f]{8}$`)
	regexSharpHexRgb      = regexp.MustCompile(`^#[0-9a-f]{3}$`)
	regexSharpHexRgba     = regexp.MustCompile(`^#[0-9a-f]{4}$`)
	regexSharpHexRrGgBb   = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	regexSharpHexRrGgBbAa = regexp.MustCompile(`^#[0-9a-f]{8}$`)
	regexPaletteMarker    = regexp.MustCompile(`^marker\(\s*([a-z]{0,3}[0-9]{1,4})\s*\)$`)
	regexPaletteX11       = regexp.MustCompile(`^x11\(\s*([a-z][\sa-z0-9]+)\s*\)`)
	regexDecRgb           = regexp.MustCompile(`^rgb\(\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*,\s*([0-9]{1,3})\s*\)`)
)

func parseHexRgb(c string) Color {
	return NewRgba(
		es_hex.ParseSingleHex(rune(c[0]))*16+es_hex.ParseSingleHex(rune(c[0])),
		es_hex.ParseSingleHex(rune(c[1]))*16+es_hex.ParseSingleHex(rune(c[1])),
		es_hex.ParseSingleHex(rune(c[2]))*16+es_hex.ParseSingleHex(rune(c[2])),
		255,
	)
}

func parseHexRgba(c string) Color {
	return NewRgba(
		es_hex.ParseSingleHex(rune(c[0]))*16+es_hex.ParseSingleHex(rune(c[0])),
		es_hex.ParseSingleHex(rune(c[1]))*16+es_hex.ParseSingleHex(rune(c[1])),
		es_hex.ParseSingleHex(rune(c[2]))*16+es_hex.ParseSingleHex(rune(c[2])),
		es_hex.ParseSingleHex(rune(c[3]))*16+es_hex.ParseSingleHex(rune(c[3])),
	)
}

func parseHexRrGgBb(c string) Color {
	return NewRgba(
		es_hex.ParseSingleHex(rune(c[0]))*16+es_hex.ParseSingleHex(rune(c[1])),
		es_hex.ParseSingleHex(rune(c[2]))*16+es_hex.ParseSingleHex(rune(c[3])),
		es_hex.ParseSingleHex(rune(c[4]))*16+es_hex.ParseSingleHex(rune(c[5])),
		255,
	)
}

func parseHexRrGgBbAa(c string) Color {
	return NewRgba(
		es_hex.ParseSingleHex(rune(c[0]))*16+es_hex.ParseSingleHex(rune(c[1])),
		es_hex.ParseSingleHex(rune(c[2]))*16+es_hex.ParseSingleHex(rune(c[3])),
		es_hex.ParseSingleHex(rune(c[4]))*16+es_hex.ParseSingleHex(rune(c[5])),
		es_hex.ParseSingleHex(rune(c[6]))*16+es_hex.ParseSingleHex(rune(c[7])),
	)
}

func parseDecRgb(c string) (Color, error) {
	rgb := regexDecRgb.FindStringSubmatch(c)
	if rgb == nil || len(rgb) != 4 {
		return nil, es_errors.NewInvalidFormatError("invalid rgb format")
	}
	r, err := strconv.ParseInt(rgb[1], 10, 32)
	if err != nil || r < 0 || 255 < r {
		return nil, es_errors.NewOutOfRangeError("red: invalid color range. must specify between 0-255")
	}
	g, err := strconv.ParseInt(rgb[2], 10, 32)
	if err != nil || g < 0 || 255 < g {
		return nil, es_errors.NewOutOfRangeError("green: invalid color range. must specify between 0-255")
	}
	b, err := strconv.ParseInt(rgb[3], 10, 32)
	if err != nil || b < 0 || 255 < b {
		return nil, es_errors.NewOutOfRangeError("blue: invalid color range. must specify between 0-255")
	}
	return NewRgba(uint8(r), uint8(g), uint8(b), 255), nil
}

func parsePaletteMarker(c string) (Color, error) {
	name := regexPaletteMarker.FindStringSubmatch(c)
	if name == nil || len(name) != 2 {
		return nil, es_errors.NewInvalidFormatError("invalid marker color name pattern")
	}
	rgb, ok := markerColors[name[1]]
	if !ok {
		return nil, es_errors.NewInvalidFormatError("invalid marker color name")
	}
	return parseHexRrGgBb(rgb), nil
}

func parsePaletteX11(c string) (Color, error) {
	name := regexPaletteX11.FindStringSubmatch(c)
	if name == nil || len(name) != 2 {
		return nil, es_errors.NewInvalidFormatError("invalid x11 color name pattern")
	}
	rgb, ok := xColor[strings.TrimSpace(name[1])]
	if !ok {
		return nil, es_errors.NewInvalidFormatError("invalid x11 color name")
	}
	return rgb, nil
}

func ParseColor(c string) (Color, error) {
	cl := strings.ToLower(strings.TrimSpace(c))
	switch {
	case regexHexRgb.MatchString(cl):
		return parseHexRgb(cl), nil
	case regexHexRgba.MatchString(cl):
		return parseHexRgba(cl), nil
	case regexHexRrGgBb.MatchString(cl):
		return parseHexRrGgBb(cl), nil
	case regexHexRrGgBbAa.MatchString(cl):
		return parseHexRrGgBbAa(cl), nil

	case regexSharpHexRgb.MatchString(cl):
		return parseHexRgb(cl[1:]), nil
	case regexSharpHexRgba.MatchString(cl):
		return parseHexRgba(cl[1:]), nil
	case regexSharpHexRrGgBb.MatchString(cl):
		return parseHexRrGgBb(cl[1:]), nil
	case regexSharpHexRrGgBbAa.MatchString(cl):
		return parseHexRrGgBbAa(cl[1:]), nil

	case regexDecRgb.MatchString(cl):
		return parseDecRgb(cl)

	case regexPaletteMarker.MatchString(cl):
		return parsePaletteMarker(cl)
	case regexPaletteX11.MatchString(cl):
		return parsePaletteX11(cl)
	}

	if rgb, ok := cssColors[cl]; ok {
		return rgb, nil
	}

	return nil, es_errors.NewInvalidFormatError("unsupported color format")
}
