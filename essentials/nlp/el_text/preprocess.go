package el_text

import "strings"

func IgnoreLineBreak(content []byte, ignoreLineBreak bool) string {
	inContent := string(content)
	if ignoreLineBreak {
		inContent = strings.ReplaceAll(inContent, "\r\n", " ")
		inContent = strings.ReplaceAll(inContent, "\n", " ")
	}
	return inContent
}
