package ut_doc

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Print markdown table of key, value pair
func PrintMarkdown(out io.Writer, headKey, headValue string, values map[string]string) {
	lenKey := 0
	lenVal := 0
	for k, v := range values {
		if len(k) > lenKey {
			lenKey = len(k)
		}
		if len(v) > lenVal {
			lenVal = len(v)
		}
	}
	fCols := "| %-" + fmt.Sprintf("%d", lenKey+2) + "s | %-" + fmt.Sprintf("%d", lenVal) + "s |"
	fBorder := "|" + strings.Repeat("-", lenKey+4) + "|" + strings.Repeat("-", lenVal+2) + "|"

	fmt.Fprintf(out, fCols, headKey, headValue)
	fmt.Fprintln(out)
	fmt.Fprintln(out, fBorder)
	keys := make([]string, 0)
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if d, e := values[k]; e {
			fmt.Fprintf(out, fCols, "`"+k+"`", d)
			fmt.Fprintln(out)
		}
	}
}
