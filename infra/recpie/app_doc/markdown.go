package app_doc

import (
	"fmt"
	"github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/cmd"
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

func LegacyCommands(c cmd.Commandlet, ec *app.ExecContext) map[string]string {
	return parseLegacyCmd(c, []string{}, ec)
}

func parseLegacyCmd(c cmd.Commandlet, line []string, ec *app.ExecContext) map[string]string {
	docs := make(map[string]string)
	if c.IsHidden() {
		return docs
	}
	w := strings.Join(line, " ")
	switch x := c.(type) {
	case *cmd.CommandletGroup:
		q := len(line)
		sl := make([]string, q+1)
		copy(sl, line)
		for _, y := range x.SubCommands {
			sl[q] = y.Name()
			d := parseLegacyCmd(y, sl, ec)
			for k, v := range d {
				docs[k] = v
			}
		}

	default:
		docs[w] = ec.Msg(x.Desc()).T()
	}
	return docs
}
