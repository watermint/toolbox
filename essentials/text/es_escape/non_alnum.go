package es_escape

import "regexp"

var (
	nonAlNum = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func ReplaceNonAlNum(s string, alt string) string {
	return nonAlNum.ReplaceAllString(s, alt)
}
