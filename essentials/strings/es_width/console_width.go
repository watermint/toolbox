package es_width

import (
	"golang.org/x/text/width"
)

func Width(s string) int {
	w := 0
	for _, r := range s {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2

		default:
			w++
		}
	}
	return w
}
