package es_width

import (
	"golang.org/x/text/width"
	"unicode"
)

func Width(s string) int {
	w := 0
	for _, r := range s {
		if unicode.IsControl(r) {
			continue
		}
		switch width.LookupRune(r).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2

		default:
			w++
		}
	}
	return w
}
